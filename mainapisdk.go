package mainapisdk

import (
	"example.com/main-api-sdk/internal/clients/rest/hooks"
	"example.com/main-api-sdk/internal/configmanager"
	"example.com/main-api-sdk/mainapisdk1"
	"time"
)

// MainAPISDK is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type MainAPISDK struct {
	MainAPISDK1 *mainapisdk1.Service
	manager     *configmanager.ConfigManager
}

func NewMainAPISDK(config Config) *MainAPISDK {
	mainApisdk1 := mainapisdk1.NewService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	mainApisdk1.WithConfigManager(manager)
	mainApisdk1.WithHook(hook)

	return &MainAPISDK{
		MainAPISDK1: mainApisdk1,
		manager:     manager,
	}
}

func (m *MainAPISDK) SetBaseURL(baseURL string) {
	m.manager.SetBaseURL(baseURL)
}

func (m *MainAPISDK) SetTimeout(timeout time.Duration) {
	m.manager.SetTimeout(timeout)
}

func (m *MainAPISDK) SetAccessToken(accessToken string) {
	m.manager.SetAccessToken(accessToken)
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
func (m *MainAPISDK) SetEnvironment(environment Environment) {
	m.manager.SetBaseURL(string(environment))
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
