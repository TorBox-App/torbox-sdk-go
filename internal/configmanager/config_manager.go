package configmanager

import (
	"time"

	"torbox-sdk-go/pkg/torboxapiconfig"
)

type ConfigManager struct {
	Torrents           torboxapiconfig.Config
	Usenet             torboxapiconfig.Config
	WebDownloadsDebrid torboxapiconfig.Config
	General            torboxapiconfig.Config
	Notifications      torboxapiconfig.Config
	User               torboxapiconfig.Config
	RssFeeds           torboxapiconfig.Config
	Integrations       torboxapiconfig.Config
}

func NewConfigManager(config torboxapiconfig.Config) *ConfigManager {
	return &ConfigManager{
		Torrents:           config,
		Usenet:             config,
		WebDownloadsDebrid: config,
		General:            config,
		Notifications:      config,
		User:               config,
		RssFeeds:           config,
		Integrations:       config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Torrents.SetBaseUrl(baseUrl)
	c.Usenet.SetBaseUrl(baseUrl)
	c.WebDownloadsDebrid.SetBaseUrl(baseUrl)
	c.General.SetBaseUrl(baseUrl)
	c.Notifications.SetBaseUrl(baseUrl)
	c.User.SetBaseUrl(baseUrl)
	c.RssFeeds.SetBaseUrl(baseUrl)
	c.Integrations.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.Torrents.SetTimeout(timeout)
	c.Usenet.SetTimeout(timeout)
	c.WebDownloadsDebrid.SetTimeout(timeout)
	c.General.SetTimeout(timeout)
	c.Notifications.SetTimeout(timeout)
	c.User.SetTimeout(timeout)
	c.RssFeeds.SetTimeout(timeout)
	c.Integrations.SetTimeout(timeout)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.Torrents.SetAccessToken(accessToken)
	c.Usenet.SetAccessToken(accessToken)
	c.WebDownloadsDebrid.SetAccessToken(accessToken)
	c.General.SetAccessToken(accessToken)
	c.Notifications.SetAccessToken(accessToken)
	c.User.SetAccessToken(accessToken)
	c.RssFeeds.SetAccessToken(accessToken)
	c.Integrations.SetAccessToken(accessToken)
}

func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {

	if c.Torrents.AccessToken != nil && *c.Torrents.AccessToken == originalValue {
		c.Torrents.SetAccessToken(newValue)
	}

	if c.Usenet.AccessToken != nil && *c.Usenet.AccessToken == originalValue {
		c.Usenet.SetAccessToken(newValue)
	}

	if c.WebDownloadsDebrid.AccessToken != nil && *c.WebDownloadsDebrid.AccessToken == originalValue {
		c.WebDownloadsDebrid.SetAccessToken(newValue)
	}

	if c.General.AccessToken != nil && *c.General.AccessToken == originalValue {
		c.General.SetAccessToken(newValue)
	}

	if c.Notifications.AccessToken != nil && *c.Notifications.AccessToken == originalValue {
		c.Notifications.SetAccessToken(newValue)
	}

	if c.User.AccessToken != nil && *c.User.AccessToken == originalValue {
		c.User.SetAccessToken(newValue)
	}

	if c.RssFeeds.AccessToken != nil && *c.RssFeeds.AccessToken == originalValue {
		c.RssFeeds.SetAccessToken(newValue)
	}

	if c.Integrations.AccessToken != nil && *c.Integrations.AccessToken == originalValue {
		c.Integrations.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetTorrents() *torboxapiconfig.Config {
	return &c.Torrents
}
func (c *ConfigManager) GetUsenet() *torboxapiconfig.Config {
	return &c.Usenet
}
func (c *ConfigManager) GetWebDownloadsDebrid() *torboxapiconfig.Config {
	return &c.WebDownloadsDebrid
}
func (c *ConfigManager) GetGeneral() *torboxapiconfig.Config {
	return &c.General
}
func (c *ConfigManager) GetNotifications() *torboxapiconfig.Config {
	return &c.Notifications
}
func (c *ConfigManager) GetUser() *torboxapiconfig.Config {
	return &c.User
}
func (c *ConfigManager) GetRssFeeds() *torboxapiconfig.Config {
	return &c.RssFeeds
}
func (c *ConfigManager) GetIntegrations() *torboxapiconfig.Config {
	return &c.Integrations
}
