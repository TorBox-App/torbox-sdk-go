package handlers

import (
	"errors"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"example.com/main-api-sdk/internal/clients/rest/httptransport"
)

// rng is seeded at package init to ensure non-deterministic jitter across Go versions.
// The global math/rand source is only auto-seeded from Go 1.20 onward.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// calculateDelay returns the backoff duration before the next retry attempt.
// Uses exponential backoff capped at maxDelay, plus optional random jitter.
func calculateDelay(attempt int, base, max, jitter time.Duration, factor float64) time.Duration {
	delay := float64(base) * math.Pow(factor, float64(attempt))
	if delay > float64(max) {
		delay = float64(max)
	}
	if jitter > 0 {
		delay += float64(rng.Int63n(int64(jitter)))
	}
	return time.Duration(delay)
}

var retryAfterDeltaRe = regexp.MustCompile(`^\d+(\.\d+)?$`)

// retryAfterDelay returns the server-directed retry delay from rate-limit response headers,
// honoring Retry-After (delta-seconds or HTTP-date) and, when absent, X-RateLimit-Reset
// (epoch seconds), clamped to maxCap. The bool is false when no usable header is present so
// the caller falls back to the computed exponential backoff.
func retryAfterDelay(retryAfterMs, retryAfter, reset string, maxCap time.Duration) (time.Duration, bool) {
	if maxCap <= 0 {
		return 0, false
	}
	// retry-after-ms (milliseconds) is a non-standard but finer-grained hint some APIs send
	// (e.g. OpenAI); it takes precedence over the whole-second Retry-After. Uses the same
	// strict delta format as Retry-After so every SDK honors the same set of values.
	if retryAfterMs = strings.TrimSpace(retryAfterMs); retryAfterDeltaRe.MatchString(retryAfterMs) {
		if ms, err := strconv.ParseFloat(retryAfterMs, 64); err == nil {
			ns := ms * float64(time.Millisecond)
			// Guard the float->Duration conversion: a huge value clamps to the cap
			// (matching the other headers and SDKs) rather than falling through.
			if ns >= float64(math.MaxInt64) {
				return maxCap, true
			}
			return clampDelay(time.Duration(ns), maxCap), true
		}
	}
	if d, ok := parseRetryAfter(retryAfter); ok {
		return clampDelay(d, maxCap), true
	}
	// X-RateLimit-Reset (epoch seconds) is only consulted when Retry-After is absent.
	// An already-elapsed reset window is treated as stale (fall back to backoff), whereas
	// an elapsed Retry-After above resolves to 0 ("retry now") — this mirrors the Ruby SDK.
	if reset = strings.TrimSpace(reset); reset != "" {
		if epoch, err := strconv.ParseInt(reset, 10, 64); err == nil {
			if d := time.Until(time.Unix(epoch, 0)); d > 0 {
				return clampDelay(d, maxCap), true
			}
		}
	}
	return 0, false
}

// parseRetryAfter parses a Retry-After header value: integer/float delta-seconds, or an
// HTTP-date (a past date yields 0). The bool is false when the value is empty or unparseable.
func parseRetryAfter(value string) (time.Duration, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	if retryAfterDeltaRe.MatchString(value) {
		secs, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, false
		}
		ns := secs * float64(time.Second)
		// Guard the float->Duration conversion: an out-of-range value is
		// implementation-defined (negative on amd64), which would defeat clampDelay.
		if ns >= float64(math.MaxInt64) {
			return time.Duration(math.MaxInt64), true
		}
		return time.Duration(ns), true
	}
	if t, err := http.ParseTime(value); err == nil {
		if d := time.Until(t); d > 0 {
			return d, true
		}
		return 0, true
	}
	return 0, false
}

// clampDelay bounds d to [0, maxCap].
func clampDelay(d, maxCap time.Duration) time.Duration {
	if d < 0 {
		return 0
	}
	if d > maxCap {
		return maxCap
	}
	return d
}

// RetryHandler automatically retries failed requests with configurable backoff.
// Retries HTTP 5xx errors and 408/429 by default; specific status codes can be configured.
// Non-HTTP errors (network, serialization) are not retried. T is the response type, E is the error type.
type RetryHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewRetryHandler creates a new retry handler. Retry configuration is read from the
// request's Config on each call, allowing per-call overrides via RequestOption.
func NewRetryHandler[T any, E any]() *RetryHandler[T, E] {
	return &RetryHandler[T, E]{}
}

// shouldRetry determines if a failed request should be retried.
// Only HTTP errors on retryable methods with retryable status codes trigger a retry.
// Non-HTTP errors (network failures, serialization errors) are not retried.
func (h *RetryHandler[T, E]) shouldRetry(errResp *httptransport.ErrorResponse[E], method string, httpMethods []string, statusCodes []int) bool {
	if !errResp.IsHTTPError {
		return false
	}

	methodAllowed := false
	for _, m := range httpMethods {
		if m == method {
			methodAllowed = true
			break
		}
	}
	if !methodAllowed {
		return false
	}

	if len(statusCodes) > 0 {
		for _, code := range statusCodes {
			if code == errResp.StatusCode {
				return true
			}
		}
		return false
	}
	return errResp.StatusCode >= 500 || errResp.StatusCode == 408 || errResp.StatusCode == 429
}

// Handle processes a request with automatic retry logic on retryable HTTP failures.
// Uses exponential backoff with jitter between attempts. Returns the first successful response
// or the final error after all attempts are exhausted.
func (h *RetryHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		return nil, httptransport.NewErrorResponse[E](errors.New("Handler chain terminated without terminating handler"), nil)
	}

	retryConfig := request.Config.Retry
	if retryConfig.MaxAttempts <= 0 {
		retryConfig.MaxAttempts = 1
	}
	for attempt := 0; attempt < retryConfig.MaxAttempts; attempt++ {
		resp, err := h.nextHandler.Handle(request.Clone())
		if err == nil {
			return resp, nil
		}

		if !h.shouldRetry(err, request.Method, retryConfig.HTTPMethodsToRetry, retryConfig.HTTPCodesToRetry) {
			return nil, err
		}

		if attempt == retryConfig.MaxAttempts-1 {
			return nil, err
		}

		if d, ok := retryAfterDelay(err.GetHeader(http.CanonicalHeaderKey("Retry-After-Ms")), err.GetHeader(http.CanonicalHeaderKey("Retry-After")), err.GetHeader(http.CanonicalHeaderKey("X-RateLimit-Reset")), retryConfig.MaxRetryAfterDelay); ok {
			time.Sleep(d)
		} else {
			time.Sleep(calculateDelay(attempt, retryConfig.RetryDelay, retryConfig.MaxDelay, retryConfig.RetryDelayJitter, retryConfig.BackOffFactor))
		}
	}

	return nil, httptransport.NewErrorResponse[E](errors.New("max retries exceeded"), nil)
}

// HandleStream processes a streaming request with automatic retry logic on retryable HTTP failures.
// Retries failed stream connections with exponential backoff.
func (h *RetryHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		return nil, httptransport.NewErrorResponse[E](errors.New("Handler chain terminated without terminating handler"), nil)
	}

	retryConfig := request.Config.Retry
	if retryConfig.MaxAttempts <= 0 {
		retryConfig.MaxAttempts = 1
	}
	for attempt := 0; attempt < retryConfig.MaxAttempts; attempt++ {
		stream, err := h.nextHandler.HandleStream(request.Clone())
		if err == nil {
			return stream, nil
		}

		if !h.shouldRetry(err, request.Method, retryConfig.HTTPMethodsToRetry, retryConfig.HTTPCodesToRetry) {
			return nil, err
		}

		if attempt == retryConfig.MaxAttempts-1 {
			return nil, err
		}

		if d, ok := retryAfterDelay(err.GetHeader(http.CanonicalHeaderKey("Retry-After-Ms")), err.GetHeader(http.CanonicalHeaderKey("Retry-After")), err.GetHeader(http.CanonicalHeaderKey("X-RateLimit-Reset")), retryConfig.MaxRetryAfterDelay); ok {
			time.Sleep(d)
		} else {
			time.Sleep(calculateDelay(attempt, retryConfig.RetryDelay, retryConfig.MaxDelay, retryConfig.RetryDelayJitter, retryConfig.BackOffFactor))
		}
	}

	return nil, httptransport.NewErrorResponse[E](errors.New("max retries exceeded"), nil)
}

// SetNext sets the next handler in the chain.
func (h *RetryHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
