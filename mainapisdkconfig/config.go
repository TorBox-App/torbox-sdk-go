package mainapisdkconfig

import (
	"time"

	"example.com/main-api-sdk/internal/clients/rest/hooks"
)

// RetryConfig holds all runtime-configurable retry parameters for the SDK client.
// Always construct via NewRetryConfig() or indirectly through NewConfig() to get
// spec-derived defaults. Direct struct literals will zero all fields.
type RetryConfig struct {
	MaxAttempts        int
	RetryDelay         time.Duration
	MaxDelay           time.Duration
	RetryDelayJitter   time.Duration
	BackOffFactor      float64
	MaxRetryAfterDelay time.Duration
	HTTPMethodsToRetry []string
	HTTPCodesToRetry   []int
}

// NewRetryConfig returns a RetryConfig initialized with the spec-derived default values.
func NewRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts:        3,
		RetryDelay:         150 * time.Millisecond,
		MaxDelay:           5000 * time.Millisecond,
		RetryDelayJitter:   50 * time.Millisecond,
		BackOffFactor:      2,
		MaxRetryAfterDelay: 60000 * time.Millisecond,
		HTTPMethodsToRetry: []string{"GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
	}
}

// Config holds all configuration parameters for the SDK client.
// It manages base URL, timeout, authentication credentials, and custom hooks.
type Config struct {
	BaseURL     string
	Timeout     time.Duration
	AccessToken *string
	Retry       RetryConfig
	HookParams  map[string]string
	hook        hooks.Hook
}

// NewConfig creates a new Config instance with default values.
// Sets the base URL to the default environment and timeout to 10 seconds.
func NewConfig() Config {
	return Config{
		BaseURL:    string(DefaultEnvironment),
		Timeout:    time.Second * 10,
		Retry:      NewRetryConfig(),
		HookParams: make(map[string]string),
	}
}

func (c *Config) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = &accessToken
}

func (c *Config) SetRetryConfig(retry RetryConfig) {
	c.Retry = retry
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
// It is equivalent to calling SetBaseURL with the environment's URL value.
func (c *Config) SetEnvironment(environment Environment) {
	c.SetBaseURL(string(environment))
}

// RequestOption is a function that configures a single request.
// Options are applied to a copy of the service config and do not affect subsequent calls.
type RequestOption func(*Config)

// WithBaseURL returns a RequestOption that overrides the BaseURL for a single request.
func WithBaseURL(baseURL string) RequestOption {
	return func(c *Config) { c.SetBaseURL(baseURL) }
}

// WithTimeout returns a RequestOption that overrides the Timeout for a single request.
func WithTimeout(timeout time.Duration) RequestOption {
	return func(c *Config) { c.SetTimeout(timeout) }
}

// WithAccessToken returns a RequestOption that overrides the AccessToken for a single request.
func WithAccessToken(accessToken string) RequestOption {
	return func(c *Config) { c.SetAccessToken(accessToken) }
}

// WithRetryConfig returns a RequestOption that overrides the RetryConfig for a single request.
func WithRetryConfig(retry RetryConfig) RequestOption {
	return func(c *Config) { c.SetRetryConfig(retry) }
}
