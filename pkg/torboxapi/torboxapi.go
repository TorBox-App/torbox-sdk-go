package torboxapi

import (
	"time"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/general"
	"torbox-sdk-go/pkg/integrations"
	"torbox-sdk-go/pkg/notifications"
	"torbox-sdk-go/pkg/rssfeeds"
	"torbox-sdk-go/pkg/torboxapiconfig"
	"torbox-sdk-go/pkg/torrents"
	"torbox-sdk-go/pkg/usenet"
	"torbox-sdk-go/pkg/user"
	"torbox-sdk-go/pkg/webdownloadsdebrid"
)

type TorboxApi struct {
	Torrents           *torrents.TorrentsService
	Usenet             *usenet.UsenetService
	WebDownloadsDebrid *webdownloadsdebrid.WebDownloadsDebridService
	General            *general.GeneralService
	Notifications      *notifications.NotificationsService
	User               *user.UserService
	RssFeeds           *rssfeeds.RssFeedsService
	Integrations       *integrations.IntegrationsService
	manager            *configmanager.ConfigManager
}

func NewTorboxApi(config torboxapiconfig.Config) *TorboxApi {
	manager := configmanager.NewConfigManager(config)
	return &TorboxApi{
		Torrents:           torrents.NewTorrentsService(manager),
		Usenet:             usenet.NewUsenetService(manager),
		WebDownloadsDebrid: webdownloadsdebrid.NewWebDownloadsDebridService(manager),
		General:            general.NewGeneralService(manager),
		Notifications:      notifications.NewNotificationsService(manager),
		User:               user.NewUserService(manager),
		RssFeeds:           rssfeeds.NewRssFeedsService(manager),
		Integrations:       integrations.NewIntegrationsService(manager),
		manager:            manager,
	}
}

func (t *TorboxApi) SetBaseUrl(baseUrl string) {
	t.manager.SetBaseUrl(baseUrl)
}

func (t *TorboxApi) SetTimeout(timeout time.Duration) {
	t.manager.SetTimeout(timeout)
}

func (t *TorboxApi) SetAccessToken(accessToken string) {
	t.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
