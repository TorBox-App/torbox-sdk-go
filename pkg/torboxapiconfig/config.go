package torboxapiconfig

import "time"

type Config struct {
	BaseUrl     *string
	Timeout     *time.Duration
	AccessToken *string
	HookParams  map[string]string
}

func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	timeout := time.Second * 10
	newConfig := Config{
		BaseUrl:    &baseUrl,
		Timeout:    &timeout,
		HookParams: make(map[string]string),
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) GetBaseUrl() string {
	return *c.BaseUrl
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = &timeout
}

func (c *Config) GetTimeout() time.Duration {
	return *c.Timeout
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = &accessToken
}

func (c *Config) GetAccessToken() string {
	return *c.AccessToken
}
