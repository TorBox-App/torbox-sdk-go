package torboxapi

import (
	"time"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/general"
	"torbox-sdk-go/pkg/integrations"
	"torbox-sdk-go/pkg/notifications"
	"torbox-sdk-go/pkg/queued"
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
	Queued             *queued.QueuedService
	manager            *configmanager.ConfigManager
}

func NewTorboxApi(config torboxapiconfig.Config) *TorboxApi {
	torrents := torrents.NewTorrentsService()
	usenet := usenet.NewUsenetService()
	webDownloadsDebrid := webdownloadsdebrid.NewWebDownloadsDebridService()
	general := general.NewGeneralService()
	notifications := notifications.NewNotificationsService()
	user := user.NewUserService()
	rssFeeds := rssfeeds.NewRssFeedsService()
	integrations := integrations.NewIntegrationsService()
	queued := queued.NewQueuedService()

	manager := configmanager.NewConfigManager(config)
	torrents.WithConfigManager(manager)
	usenet.WithConfigManager(manager)
	webDownloadsDebrid.WithConfigManager(manager)
	general.WithConfigManager(manager)
	notifications.WithConfigManager(manager)
	user.WithConfigManager(manager)
	rssFeeds.WithConfigManager(manager)
	integrations.WithConfigManager(manager)
	queued.WithConfigManager(manager)

	return &TorboxApi{
		Torrents:           torrents,
		Usenet:             usenet,
		WebDownloadsDebrid: webDownloadsDebrid,
		General:            general,
		Notifications:      notifications,
		User:               user,
		RssFeeds:           rssFeeds,
		Integrations:       integrations,
		Queued:             queued,
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
