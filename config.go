package mainapisdk

import (
	"example.com/main-api-sdk/mainapisdkconfig"
	"example.com/main-api-sdk/param"
)

// The type aliases below let consumers use a single import path for the entire SDK.
// Internally the concrete types live in mainapisdkconfig and param.

// Config holds all configuration parameters for the SDK client.
type Config = mainapisdkconfig.Config

// RequestOption is a function that configures a single request.
type RequestOption = mainapisdkconfig.RequestOption

// Environment defines the available API base URLs.
type Environment = mainapisdkconfig.Environment

// RetryConfig holds all runtime-configurable retry parameters.
type RetryConfig = mainapisdkconfig.RetryConfig

// NewConfig creates a Config with spec-derived defaults.
var NewConfig = mainapisdkconfig.NewConfig

// NewRetryConfig returns a RetryConfig initialized with spec-derived defaults.
var NewRetryConfig = mainapisdkconfig.NewRetryConfig

// WithBaseURL returns a RequestOption that overrides BaseURL for a single request.
var WithBaseURL = mainapisdkconfig.WithBaseURL

// WithTimeout returns a RequestOption that overrides Timeout for a single request.
var WithTimeout = mainapisdkconfig.WithTimeout

// WithAccessToken returns a RequestOption that overrides AccessToken for a single request.
var WithAccessToken = mainapisdkconfig.WithAccessToken

// WithRetryConfig returns a RequestOption that overrides the RetryConfig for a single request.
var WithRetryConfig = mainapisdkconfig.WithRetryConfig

// Nullable returns a *param.Nullable[T] set to v — use for nullable fields with a value.
func Nullable[T any](v T) *param.Nullable[T] { return &param.Nullable[T]{Value: v} }

// Null returns a *param.Nullable[T] with IsNull set to true, signalling an explicit JSON null.
func Null[T any]() *param.Nullable[T] { return param.Null[T]() }

// Ptr returns a pointer to v — use when no type-specific helper exists.
func Ptr[T any](v T) *T { return param.Ptr(v) }

// Environment constants for the available API base URLs.
const (
	DefaultEnvironment Environment = mainapisdkconfig.DefaultEnvironment
	ApiBaseEnvironment Environment = mainapisdkconfig.ApiBaseEnvironment
)
