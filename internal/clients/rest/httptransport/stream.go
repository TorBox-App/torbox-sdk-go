package httptransport

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// ErrStreamClosed is returned when trying to read from a closed stream.
var ErrStreamClosed = errors.New("stream is closed")

// StreamChunk represents a single chunk from a streaming response.
// Contains the deserialized data, raw bytes, and metadata about the response.
type StreamChunk[T any] struct {
	Data     T
	Raw      []byte
	Metadata StreamChunkMetadata
}

// StreamChunkMetadata contains metadata about a stream chunk.
// Includes HTTP headers and status code from the streaming response.
type StreamChunkMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// Stream represents a streaming response from the API using Server-Sent Events (SSE) protocol.
// It handles chunk-by-chunk processing with context cancellation and thread-safe closure.
type Stream[T any] struct {
	response   *http.Response
	scanner    *bufio.Scanner
	ctx        context.Context
	statusCode int
	headers    map[string]string

	mu     sync.Mutex
	closed bool
}

// NewStream creates a new stream wrapper from an HTTP response with context for cancellation.
// Converts response headers to a map and initializes the stream scanner for reading chunks.
func NewStream[T any](resp *http.Response, ctx context.Context) *Stream[T] {
	// Convert http.Header to map[string]string
	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &Stream[T]{
		response:   resp,
		scanner:    bufio.NewScanner(resp.Body),
		ctx:        ctx,
		statusCode: resp.StatusCode,
		headers:    headers,
		closed:     false,
	}
}

// Next returns the next chunk from the stream by reading and parsing SSE data.
// Returns ErrStreamClosed when the stream has ended normally or receives [DONE] signal.
// Returns context errors if the context is cancelled or times out. Thread-safe for concurrent access.
func (s *Stream[T]) Next() (*StreamChunk[T], error) {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return nil, ErrStreamClosed
	}
	s.mu.Unlock()

	// Check if context is cancelled
	select {
	case <-s.ctx.Done():
		s.Close()
		return nil, s.ctx.Err()
	default:
	}

	for s.scanner.Scan() {
		line := s.scanner.Text()
		data, ok := parseSSELine(line)
		if !ok {
			// Skip non-data lines or stream termination
			if data == "" && line == "" {
				continue
			}
			// [DONE] signal received
			s.Close()
			return nil, ErrStreamClosed
		}

		var chunk T
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			return nil, err
		}

		return &StreamChunk[T]{
			Data: chunk,
			Raw:  []byte(data),
			Metadata: StreamChunkMetadata{
				StatusCode: s.statusCode,
				Headers:    s.headers,
			},
		}, nil
	}

	if err := s.scanner.Err(); err != nil {
		s.Close()
		return nil, err
	}

	s.Close()
	return nil, ErrStreamClosed
}

// Close closes the stream and releases resources including the response body.
// Idempotent - calling Close multiple times is safe. Thread-safe for concurrent access.
func (s *Stream[T]) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}
	s.closed = true
	return s.response.Body.Close()
}

// IsClosed returns whether the stream has been closed.
// Thread-safe for concurrent access.
func (s *Stream[T]) IsClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed
}

// parseSSELine parses a Server-Sent Events line according to SSE protocol.
// SSE format: "data: {json content}\n\n". Returns the data content and true if valid,
// or empty string and false if the line should be skipped or signals stream termination ([DONE]).
func parseSSELine(line string) (string, bool) {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "data: ") {
		data := strings.TrimPrefix(line, "data: ")
		// Check for stream termination signal
		if data == "[DONE]" {
			return "", false
		}
		return data, true
	}
	return "", false
}

// HandleStream creates a streaming HTTP request and returns a Stream for consuming chunks.
// Executes the HTTP request, validates the status code, and wraps the response in a Stream.
func HandleStream[T any](request Request) (*Stream[T], *ErrorResponse[T]) {
	requestClone := request.Clone()

	client := http.Client{}
	if requestClone.Config.Timeout != 0 {
		client.Timeout = requestClone.Config.Timeout
	}

	req, err := requestClone.CreateHTTPRequest()
	if err != nil {
		return nil, NewErrorResponse[T](err, nil)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, NewErrorResponse[T](err, nil)
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()
		err := fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
		return nil, NewErrorResponse[T](err, nil)
	}

	// Get context from request, or use background if not available
	ctx := request.Context
	if ctx == nil {
		ctx = context.Background()
	}

	stream := NewStream[T](resp, ctx)
	return stream, nil
}
