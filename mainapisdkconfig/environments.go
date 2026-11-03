package mainapisdkconfig

// Environment type defines the available API environments.
type Environment string

// Environment constants define the available base URLs for different deployment environments.
// Use these constants when configuring the SDK client.
const (
	DefaultEnvironment Environment = "https://api_base"
	ApiBaseEnvironment Environment = "https://api_base"
)
