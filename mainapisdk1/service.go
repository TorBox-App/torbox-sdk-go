package mainapisdk1

import (
	"context"
	restClient "example.com/main-api-sdk/internal/clients/rest"
	"example.com/main-api-sdk/internal/clients/rest/hooks"
	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"example.com/main-api-sdk/internal/configmanager"
	"example.com/main-api-sdk/mainapisdkconfig"
	"example.com/main-api-sdk/sharedmodels"
	"time"
)

// Service provides methods to interact with MainAPISDK1-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                                     *configmanager.ConfigManager
	hook                                        hooks.Hook
	createTorrentConfig                         []mainapisdkconfig.RequestOption
	asyncCreateTorrentConfig                    []mainapisdkconfig.RequestOption
	controlTorrentConfig                        []mainapisdkconfig.RequestOption
	requestDownloadLinkConfig                   []mainapisdkconfig.RequestOption
	requestDownloadLinksGetConfig               []mainapisdkconfig.RequestOption
	requestDownloadLinksPostConfig              []mainapisdkconfig.RequestOption
	getTorrentListConfig                        []mainapisdkconfig.RequestOption
	getTorrentCachedAvailabilityConfig          []mainapisdkconfig.RequestOption
	getTorrentCachedAvailabilityBatchConfig     []mainapisdkconfig.RequestOption
	exportTorrentDataConfig                     []mainapisdkconfig.RequestOption
	getTorrentInfoConfig                        []mainapisdkconfig.RequestOption
	getTorrentInfo1Config                       []mainapisdkconfig.RequestOption
	editTorrentItemConfig                       []mainapisdkconfig.RequestOption
	magnetToFileConfig                          []mainapisdkconfig.RequestOption
	createUsenetDownloadConfig                  []mainapisdkconfig.RequestOption
	controlUsenetDownloadConfig                 []mainapisdkconfig.RequestOption
	requestDownloadLink1Config                  []mainapisdkconfig.RequestOption
	requestDownloadLinksGet1Config              []mainapisdkconfig.RequestOption
	requestDownloadLinksPost1Config             []mainapisdkconfig.RequestOption
	getUsenetListConfig                         []mainapisdkconfig.RequestOption
	getUsenetCachedAvailabilityConfig           []mainapisdkconfig.RequestOption
	getUsenetCachedAvailabilityBatchConfig      []mainapisdkconfig.RequestOption
	editUsenetItemConfig                        []mainapisdkconfig.RequestOption
	getNewsServerConnectionInformationConfig    []mainapisdkconfig.RequestOption
	getNewsServerAccountConfig                  []mainapisdkconfig.RequestOption
	resetNewsServerAccountPasswordConfig        []mainapisdkconfig.RequestOption
	createWebDownloadConfig                     []mainapisdkconfig.RequestOption
	controlWebDownloadConfig                    []mainapisdkconfig.RequestOption
	requestDownloadLink2Config                  []mainapisdkconfig.RequestOption
	requestDownloadLinksGet2Config              []mainapisdkconfig.RequestOption
	requestDownloadLinksPost2Config             []mainapisdkconfig.RequestOption
	getWebDownloadListConfig                    []mainapisdkconfig.RequestOption
	getWebDownloadCachedAvailabilityConfig      []mainapisdkconfig.RequestOption
	getWebDownloadCachedAvailabilityBatchConfig []mainapisdkconfig.RequestOption
	getHosterListConfig                         []mainapisdkconfig.RequestOption
	editWebDownloadItemConfig                   []mainapisdkconfig.RequestOption
	getUpStatusConfig                           []mainapisdkconfig.RequestOption
	getStatsConfig                              []mainapisdkconfig.RequestOption
	getChangelogsRssFeedConfig                  []mainapisdkconfig.RequestOption
	getChangelogsJSONConfig                     []mainapisdkconfig.RequestOption
	getSpeedtestFilesConfig                     []mainapisdkconfig.RequestOption
	getRssNotificationFeedConfig                []mainapisdkconfig.RequestOption
	getNotificationFeedConfig                   []mainapisdkconfig.RequestOption
	clearAllNotificationsConfig                 []mainapisdkconfig.RequestOption
	clearSingleNotificationConfig               []mainapisdkconfig.RequestOption
	sendTestNotificationConfig                  []mainapisdkconfig.RequestOption
	refreshAPITokenConfig                       []mainapisdkconfig.RequestOption
	getUserDataConfig                           []mainapisdkconfig.RequestOption
	addReferralToAccountConfig                  []mainapisdkconfig.RequestOption
	getConfirmationCodeConfig                   []mainapisdkconfig.RequestOption
	startDeviceCodeAuthorizationProcessConfig   []mainapisdkconfig.RequestOption
	getTokenFromDeviceCodeProcessConfig         []mainapisdkconfig.RequestOption
	getReferralDataConfig                       []mainapisdkconfig.RequestOption
	getSubscriptionsConfig                      []mainapisdkconfig.RequestOption
	getTransactionsConfig                       []mainapisdkconfig.RequestOption
	getTransactionPdfConfig                     []mainapisdkconfig.RequestOption
	editUserSettingsConfig                      []mainapisdkconfig.RequestOption
	getUserStatsConfig                          []mainapisdkconfig.RequestOption
	addRssFeedConfig                            []mainapisdkconfig.RequestOption
	controlRssFeedConfig                        []mainapisdkconfig.RequestOption
	editRssFeedConfig                           []mainapisdkconfig.RequestOption
	getUserRssFeedsConfig                       []mainapisdkconfig.RequestOption
	getRssFeedItemsConfig                       []mainapisdkconfig.RequestOption
	getAllJobsConfig                            []mainapisdkconfig.RequestOption
	getSpecificJobConfig                        []mainapisdkconfig.RequestOption
	getAllJobsByHashConfig                      []mainapisdkconfig.RequestOption
	cancelSpecificJobConfig                     []mainapisdkconfig.RequestOption
	getQueuedDownloadsConfig                    []mainapisdkconfig.RequestOption
	controlQueuedDownloadsConfig                []mainapisdkconfig.RequestOption
	createStreamConfig                          []mainapisdkconfig.RequestOption
	getStreamDataConfig                         []mainapisdkconfig.RequestOption
}

func NewService() *Service {
	return &Service{
		manager: configmanager.NewConfigManager(mainapisdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *Service) WithConfigManager(manager *configmanager.ConfigManager) *Service {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *Service) WithHook(hook hooks.Hook) *Service {
	api.hook = hook
	return api
}

func (api *Service) config() *mainapisdkconfig.Config {
	return api.manager.GetMainAPISDK1()
}

func (api *Service) getHook() hooks.Hook {
	return api.hook
}

func (api *Service) SetBaseURL(baseURL string) {
	config := api.config()
	config.SetBaseURL(baseURL)
}

func (api *Service) SetTimeout(timeout time.Duration) {
	config := api.config()
	config.SetTimeout(timeout)
}

func (api *Service) SetAccessToken(accessToken string) {
	config := api.config()
	config.SetAccessToken(accessToken)
}

// SetCreateTorrentConfig sets method-level configuration for CreateTorrent.
// Options are applied to every future call to CreateTorrent and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateTorrentConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.createTorrentConfig = opts
	return api
}

// SetAsyncCreateTorrentConfig sets method-level configuration for AsyncCreateTorrent.
// Options are applied to every future call to AsyncCreateTorrent and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAsyncCreateTorrentConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.asyncCreateTorrentConfig = opts
	return api
}

// SetControlTorrentConfig sets method-level configuration for ControlTorrent.
// Options are applied to every future call to ControlTorrent and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetControlTorrentConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.controlTorrentConfig = opts
	return api
}

// SetRequestDownloadLinkConfig sets method-level configuration for RequestDownloadLink.
// Options are applied to every future call to RequestDownloadLink and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinkConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinkConfig = opts
	return api
}

// SetRequestDownloadLinksGetConfig sets method-level configuration for RequestDownloadLinksGet.
// Options are applied to every future call to RequestDownloadLinksGet and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinksGetConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinksGetConfig = opts
	return api
}

// SetRequestDownloadLinksPostConfig sets method-level configuration for RequestDownloadLinksPost.
// Options are applied to every future call to RequestDownloadLinksPost and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinksPostConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinksPostConfig = opts
	return api
}

// SetGetTorrentListConfig sets method-level configuration for GetTorrentList.
// Options are applied to every future call to GetTorrentList and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTorrentListConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTorrentListConfig = opts
	return api
}

// SetGetTorrentCachedAvailabilityConfig sets method-level configuration for GetTorrentCachedAvailability.
// Options are applied to every future call to GetTorrentCachedAvailability and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTorrentCachedAvailabilityConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTorrentCachedAvailabilityConfig = opts
	return api
}

// SetGetTorrentCachedAvailabilityBatchConfig sets method-level configuration for GetTorrentCachedAvailabilityBatch.
// Options are applied to every future call to GetTorrentCachedAvailabilityBatch and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTorrentCachedAvailabilityBatchConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTorrentCachedAvailabilityBatchConfig = opts
	return api
}

// SetExportTorrentDataConfig sets method-level configuration for ExportTorrentData.
// Options are applied to every future call to ExportTorrentData and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetExportTorrentDataConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.exportTorrentDataConfig = opts
	return api
}

// SetGetTorrentInfoConfig sets method-level configuration for GetTorrentInfo.
// Options are applied to every future call to GetTorrentInfo and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTorrentInfoConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTorrentInfoConfig = opts
	return api
}

// SetGetTorrentInfo1Config sets method-level configuration for GetTorrentInfo1.
// Options are applied to every future call to GetTorrentInfo1 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTorrentInfo1Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTorrentInfo1Config = opts
	return api
}

// SetEditTorrentItemConfig sets method-level configuration for EditTorrentItem.
// Options are applied to every future call to EditTorrentItem and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditTorrentItemConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.editTorrentItemConfig = opts
	return api
}

// SetMagnetToFileConfig sets method-level configuration for MagnetToFile.
// Options are applied to every future call to MagnetToFile and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetMagnetToFileConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.magnetToFileConfig = opts
	return api
}

// SetCreateUsenetDownloadConfig sets method-level configuration for CreateUsenetDownload.
// Options are applied to every future call to CreateUsenetDownload and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateUsenetDownloadConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.createUsenetDownloadConfig = opts
	return api
}

// SetControlUsenetDownloadConfig sets method-level configuration for ControlUsenetDownload.
// Options are applied to every future call to ControlUsenetDownload and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetControlUsenetDownloadConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.controlUsenetDownloadConfig = opts
	return api
}

// SetRequestDownloadLink1Config sets method-level configuration for RequestDownloadLink1.
// Options are applied to every future call to RequestDownloadLink1 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLink1Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLink1Config = opts
	return api
}

// SetRequestDownloadLinksGet1Config sets method-level configuration for RequestDownloadLinksGet1.
// Options are applied to every future call to RequestDownloadLinksGet1 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinksGet1Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinksGet1Config = opts
	return api
}

// SetRequestDownloadLinksPost1Config sets method-level configuration for RequestDownloadLinksPost1.
// Options are applied to every future call to RequestDownloadLinksPost1 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinksPost1Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinksPost1Config = opts
	return api
}

// SetGetUsenetListConfig sets method-level configuration for GetUsenetList.
// Options are applied to every future call to GetUsenetList and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUsenetListConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUsenetListConfig = opts
	return api
}

// SetGetUsenetCachedAvailabilityConfig sets method-level configuration for GetUsenetCachedAvailability.
// Options are applied to every future call to GetUsenetCachedAvailability and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUsenetCachedAvailabilityConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUsenetCachedAvailabilityConfig = opts
	return api
}

// SetGetUsenetCachedAvailabilityBatchConfig sets method-level configuration for GetUsenetCachedAvailabilityBatch.
// Options are applied to every future call to GetUsenetCachedAvailabilityBatch and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUsenetCachedAvailabilityBatchConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUsenetCachedAvailabilityBatchConfig = opts
	return api
}

// SetEditUsenetItemConfig sets method-level configuration for EditUsenetItem.
// Options are applied to every future call to EditUsenetItem and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditUsenetItemConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.editUsenetItemConfig = opts
	return api
}

// SetGetNewsServerConnectionInformationConfig sets method-level configuration for GetNewsServerConnectionInformation.
// Options are applied to every future call to GetNewsServerConnectionInformation and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetNewsServerConnectionInformationConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getNewsServerConnectionInformationConfig = opts
	return api
}

// SetGetNewsServerAccountConfig sets method-level configuration for GetNewsServerAccount.
// Options are applied to every future call to GetNewsServerAccount and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetNewsServerAccountConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getNewsServerAccountConfig = opts
	return api
}

// SetResetNewsServerAccountPasswordConfig sets method-level configuration for ResetNewsServerAccountPassword.
// Options are applied to every future call to ResetNewsServerAccountPassword and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetResetNewsServerAccountPasswordConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.resetNewsServerAccountPasswordConfig = opts
	return api
}

// SetCreateWebDownloadConfig sets method-level configuration for CreateWebDownload.
// Options are applied to every future call to CreateWebDownload and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateWebDownloadConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.createWebDownloadConfig = opts
	return api
}

// SetControlWebDownloadConfig sets method-level configuration for ControlWebDownload.
// Options are applied to every future call to ControlWebDownload and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetControlWebDownloadConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.controlWebDownloadConfig = opts
	return api
}

// SetRequestDownloadLink2Config sets method-level configuration for RequestDownloadLink2.
// Options are applied to every future call to RequestDownloadLink2 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLink2Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLink2Config = opts
	return api
}

// SetRequestDownloadLinksGet2Config sets method-level configuration for RequestDownloadLinksGet2.
// Options are applied to every future call to RequestDownloadLinksGet2 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinksGet2Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinksGet2Config = opts
	return api
}

// SetRequestDownloadLinksPost2Config sets method-level configuration for RequestDownloadLinksPost2.
// Options are applied to every future call to RequestDownloadLinksPost2 and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRequestDownloadLinksPost2Config(opts ...mainapisdkconfig.RequestOption) *Service {
	api.requestDownloadLinksPost2Config = opts
	return api
}

// SetGetWebDownloadListConfig sets method-level configuration for GetWebDownloadList.
// Options are applied to every future call to GetWebDownloadList and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetWebDownloadListConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getWebDownloadListConfig = opts
	return api
}

// SetGetWebDownloadCachedAvailabilityConfig sets method-level configuration for GetWebDownloadCachedAvailability.
// Options are applied to every future call to GetWebDownloadCachedAvailability and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetWebDownloadCachedAvailabilityConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getWebDownloadCachedAvailabilityConfig = opts
	return api
}

// SetGetWebDownloadCachedAvailabilityBatchConfig sets method-level configuration for GetWebDownloadCachedAvailabilityBatch.
// Options are applied to every future call to GetWebDownloadCachedAvailabilityBatch and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetWebDownloadCachedAvailabilityBatchConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getWebDownloadCachedAvailabilityBatchConfig = opts
	return api
}

// SetGetHosterListConfig sets method-level configuration for GetHosterList.
// Options are applied to every future call to GetHosterList and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetHosterListConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getHosterListConfig = opts
	return api
}

// SetEditWebDownloadItemConfig sets method-level configuration for EditWebDownloadItem.
// Options are applied to every future call to EditWebDownloadItem and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditWebDownloadItemConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.editWebDownloadItemConfig = opts
	return api
}

// SetGetUpStatusConfig sets method-level configuration for GetUpStatus.
// Options are applied to every future call to GetUpStatus and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUpStatusConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUpStatusConfig = opts
	return api
}

// SetGetStatsConfig sets method-level configuration for GetStats.
// Options are applied to every future call to GetStats and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetStatsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getStatsConfig = opts
	return api
}

// SetGetChangelogsRssFeedConfig sets method-level configuration for GetChangelogsRssFeed.
// Options are applied to every future call to GetChangelogsRssFeed and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetChangelogsRssFeedConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getChangelogsRssFeedConfig = opts
	return api
}

// SetGetChangelogsJSONConfig sets method-level configuration for GetChangelogsJSON.
// Options are applied to every future call to GetChangelogsJSON and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetChangelogsJSONConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getChangelogsJSONConfig = opts
	return api
}

// SetGetSpeedtestFilesConfig sets method-level configuration for GetSpeedtestFiles.
// Options are applied to every future call to GetSpeedtestFiles and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetSpeedtestFilesConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getSpeedtestFilesConfig = opts
	return api
}

// SetGetRssNotificationFeedConfig sets method-level configuration for GetRssNotificationFeed.
// Options are applied to every future call to GetRssNotificationFeed and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetRssNotificationFeedConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getRssNotificationFeedConfig = opts
	return api
}

// SetGetNotificationFeedConfig sets method-level configuration for GetNotificationFeed.
// Options are applied to every future call to GetNotificationFeed and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetNotificationFeedConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getNotificationFeedConfig = opts
	return api
}

// SetClearAllNotificationsConfig sets method-level configuration for ClearAllNotifications.
// Options are applied to every future call to ClearAllNotifications and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetClearAllNotificationsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.clearAllNotificationsConfig = opts
	return api
}

// SetClearSingleNotificationConfig sets method-level configuration for ClearSingleNotification.
// Options are applied to every future call to ClearSingleNotification and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetClearSingleNotificationConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.clearSingleNotificationConfig = opts
	return api
}

// SetSendTestNotificationConfig sets method-level configuration for SendTestNotification.
// Options are applied to every future call to SendTestNotification and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSendTestNotificationConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.sendTestNotificationConfig = opts
	return api
}

// SetRefreshAPITokenConfig sets method-level configuration for RefreshAPIToken.
// Options are applied to every future call to RefreshAPIToken and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRefreshAPITokenConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.refreshAPITokenConfig = opts
	return api
}

// SetGetUserDataConfig sets method-level configuration for GetUserData.
// Options are applied to every future call to GetUserData and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUserDataConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUserDataConfig = opts
	return api
}

// SetAddReferralToAccountConfig sets method-level configuration for AddReferralToAccount.
// Options are applied to every future call to AddReferralToAccount and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddReferralToAccountConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.addReferralToAccountConfig = opts
	return api
}

// SetGetConfirmationCodeConfig sets method-level configuration for GetConfirmationCode.
// Options are applied to every future call to GetConfirmationCode and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetConfirmationCodeConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getConfirmationCodeConfig = opts
	return api
}

// SetStartDeviceCodeAuthorizationProcessConfig sets method-level configuration for StartDeviceCodeAuthorizationProcess.
// Options are applied to every future call to StartDeviceCodeAuthorizationProcess and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetStartDeviceCodeAuthorizationProcessConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.startDeviceCodeAuthorizationProcessConfig = opts
	return api
}

// SetGetTokenFromDeviceCodeProcessConfig sets method-level configuration for GetTokenFromDeviceCodeProcess.
// Options are applied to every future call to GetTokenFromDeviceCodeProcess and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTokenFromDeviceCodeProcessConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTokenFromDeviceCodeProcessConfig = opts
	return api
}

// SetGetReferralDataConfig sets method-level configuration for GetReferralData.
// Options are applied to every future call to GetReferralData and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetReferralDataConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getReferralDataConfig = opts
	return api
}

// SetGetSubscriptionsConfig sets method-level configuration for GetSubscriptions.
// Options are applied to every future call to GetSubscriptions and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetSubscriptionsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getSubscriptionsConfig = opts
	return api
}

// SetGetTransactionsConfig sets method-level configuration for GetTransactions.
// Options are applied to every future call to GetTransactions and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTransactionsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTransactionsConfig = opts
	return api
}

// SetGetTransactionPdfConfig sets method-level configuration for GetTransactionPdf.
// Options are applied to every future call to GetTransactionPdf and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTransactionPdfConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getTransactionPdfConfig = opts
	return api
}

// SetEditUserSettingsConfig sets method-level configuration for EditUserSettings.
// Options are applied to every future call to EditUserSettings and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditUserSettingsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.editUserSettingsConfig = opts
	return api
}

// SetGetUserStatsConfig sets method-level configuration for GetUserStats.
// Options are applied to every future call to GetUserStats and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUserStatsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUserStatsConfig = opts
	return api
}

// SetAddRssFeedConfig sets method-level configuration for AddRssFeed.
// Options are applied to every future call to AddRssFeed and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddRssFeedConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.addRssFeedConfig = opts
	return api
}

// SetControlRssFeedConfig sets method-level configuration for ControlRssFeed.
// Options are applied to every future call to ControlRssFeed and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetControlRssFeedConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.controlRssFeedConfig = opts
	return api
}

// SetEditRssFeedConfig sets method-level configuration for EditRssFeed.
// Options are applied to every future call to EditRssFeed and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetEditRssFeedConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.editRssFeedConfig = opts
	return api
}

// SetGetUserRssFeedsConfig sets method-level configuration for GetUserRssFeeds.
// Options are applied to every future call to GetUserRssFeeds and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetUserRssFeedsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getUserRssFeedsConfig = opts
	return api
}

// SetGetRssFeedItemsConfig sets method-level configuration for GetRssFeedItems.
// Options are applied to every future call to GetRssFeedItems and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetRssFeedItemsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getRssFeedItemsConfig = opts
	return api
}

// SetGetAllJobsConfig sets method-level configuration for GetAllJobs.
// Options are applied to every future call to GetAllJobs and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetAllJobsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getAllJobsConfig = opts
	return api
}

// SetGetSpecificJobConfig sets method-level configuration for GetSpecificJob.
// Options are applied to every future call to GetSpecificJob and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetSpecificJobConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getSpecificJobConfig = opts
	return api
}

// SetGetAllJobsByHashConfig sets method-level configuration for GetAllJobsByHash.
// Options are applied to every future call to GetAllJobsByHash and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetAllJobsByHashConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getAllJobsByHashConfig = opts
	return api
}

// SetCancelSpecificJobConfig sets method-level configuration for CancelSpecificJob.
// Options are applied to every future call to CancelSpecificJob and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCancelSpecificJobConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.cancelSpecificJobConfig = opts
	return api
}

// SetGetQueuedDownloadsConfig sets method-level configuration for GetQueuedDownloads.
// Options are applied to every future call to GetQueuedDownloads and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetQueuedDownloadsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getQueuedDownloadsConfig = opts
	return api
}

// SetControlQueuedDownloadsConfig sets method-level configuration for ControlQueuedDownloads.
// Options are applied to every future call to ControlQueuedDownloads and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetControlQueuedDownloadsConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.controlQueuedDownloadsConfig = opts
	return api
}

// SetCreateStreamConfig sets method-level configuration for CreateStream.
// Options are applied to every future call to CreateStream and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateStreamConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.createStreamConfig = opts
	return api
}

// SetGetStreamDataConfig sets method-level configuration for GetStreamData.
// Options are applied to every future call to GetStreamData and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetStreamDataConfig(opts ...mainapisdkconfig.RequestOption) *Service {
	api.getStreamDataConfig = opts
	return api
}

// ### Overview
//
// Creates a torrent under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large.
//
// ### Rate Limits
//
// This endpoint has two limits that apply at the same time:
//
// - **300 requests per minute:** a shared ceiling that _every_ request counts against, whether the torrent is cached or uncached.
//
// - **60 uncached torrents per hour:** an additional, stricter limit that only uncached torrents count against.
//
// The important thing is that cached and uncached requests both draw from the same 300/min pool, but uncached requests _also_ draw from their own separate 60/hour pool. An uncached request therefore counts twice: once against the per-minute limit, and once against the per-hour uncached limit.
//
// **Example:** In a single minute you add 250 cached + 50 uncached torrents. That's 300 total, so you've hit the per-minute limit for that minute. But those 50 uncached torrents also came out of your 60/hour uncached budget, leaving only **10 uncached torrents** for the remaining 59 minutes of that hour. (You could still add plenty more _cached_ torrents during that hour, since cached requests aren't capped hourly: only by the 300/min ceiling.)
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) CreateTorrent(ctx context.Context, apiVersion string, createTorrentRequest CreateTorrentRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.createTorrentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/createtorrent").
		WithConfig(config).
		WithBody(createTorrentRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Creates a torrent asynchrously under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large.
//
// The request will return instantly, not returning any data other than a success message. This is the same for failures, as they will return the error in the background, rather than immediately.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) AsyncCreateTorrent(ctx context.Context, apiVersion string, asyncCreateTorrentRequest AsyncCreateTorrentRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.asyncCreateTorrentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/asynccreatetorrent").
		WithConfig(config).
		WithBody(asyncCreateTorrentRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Controls a torrent. By sending the torrent's ID and the type of operation you want to perform, it will send that request to the torrent client.
//
// Operations are either:
//
// - **Reannounce** `reannounces the torrent to get new peers`
//
// - **Delete** `deletes the torrent from the client and your account permanently`
//
// - **Resume** `resumes a paused torrent`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ControlTorrent(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.controlTorrentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/controltorrent").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Permalinks
//
// Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLink(ctx context.Context, apiVersion string, params RequestDownloadLinkRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinkConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/requestdl").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLinksGet(ctx context.Context, apiVersion string, params RequestDownloadLinksGetRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinksGetConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/requestdl/batch").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLinksPost(ctx context.Context, apiVersion string, mainAPISDKRequestDownloadLinksPostRequest1 MainAPISDKRequestDownloadLinksPostRequest1, params RequestDownloadLinksPostRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinksPostConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/requestdl/batch").
		WithConfig(config).
		WithBody(mainAPISDKRequestDownloadLinksPostRequest1).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the user's torrent list. This gives you the needed information to perform other torrent actions. This information only gets updated every 600 seconds, or when the _Request Update On Torrent_ request is sent to the _relay API_.
//
// #### Download States:
//
// - "downloading" -> The torrent is currently downloading.
//
// - "uploading" -> The torrent is currently seeding.
//
// - "stalled (no seeds)" -> The torrent is trying to download, but there are no seeds connected to download from.
//
// - "paused" -> The torrent is paused.
//
// - "completed" -> The torrent is completely downloaded. Do not use this for download completion status.
//
// - "cached" -> The torrent is cached from the server.
//
// - "metaDL" -> The torrent is downloading metadata from the hoard.
//
// - "checkingResumeData" -> The torrent is checking resumable data.
//
// All other statuses are basic qBittorrent states. [Check out here for the full list](https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#torrent-management).
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetTorrentList(ctx context.Context, apiVersion string, params GetTorrentListRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTorrentListConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/mylist").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Takes in a list of comma separated torrent hashes and checks if the torrent is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as:
// `?hash=XXXX&hash=XXXX&hash=XXXX`
//
// or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX`
// and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetTorrentCachedAvailability(ctx context.Context, apiVersion string, params GetTorrentCachedAvailabilityRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTorrentCachedAvailabilityConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/checkcached").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Takes in a list of comma separated torrent hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters.
//
// Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetTorrentCachedAvailabilityBatch(ctx context.Context, apiVersion string, getTorrentCachedAvailabilityBatchRequest GetTorrentCachedAvailabilityBatchRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTorrentCachedAvailabilityBatchConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/checkcached").
		WithConfig(config).
		WithBody(getTorrentCachedAvailabilityBatchRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Exports the magnet or torrent file. Requires a type to be passed. If type is **magnet**, it will return a JSON response with the magnet as a string in the _data_ key. If type is **file**, it will return a bittorrent file as a download. Not compatible with cached downloads.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ExportTorrentData(ctx context.Context, apiVersion string, params ExportTorrentDataRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.exportTorrentDataConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/exportdata").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant.
//
// **It is not recommended to use this for checking TorBox's cache. Please use the GET /checkcache endpoint instead.**
//
// ### Authorization
//
// None required.
func (api *Service) GetTorrentInfo(ctx context.Context, apiVersion string, params GetTorrentInfoRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTorrentInfoConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/torrentinfo").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Same as the GET route, but allows posting data such as magnet, and torrent files.
//
// Hashes will have precedence over magnets, and magnets will have precedence over torrent files.
//
// Only proper torrent files are accepted.
//
// At least one of hash, magnet, or torrent file is required.
//
// A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant.
//
// **It is not recommended to use this for checking TorBox's cache. Please use the GET /checkcache endpoint instead.**
//
// ### Authorization
//
// None required.
func (api *Service) GetTorrentInfo1(ctx context.Context, apiVersion string, getTorrentInfoRequest GetTorrentInfoRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTorrentInfo1Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/torrentinfo").
		WithConfig(config).
		WithBody(getTorrentInfoRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Updates a download item based on the item's ID.
// Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important.
//
// ### Restrictions
//
// - Does not affect the cached database (or any other user's item).
//
// - Name must pass validation:
//
// - Above 1 character long.
//
// - Below 200 characters long.
//
// - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically).
//
// - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically).
//
// - Tags follow the same rules as the name.
//
// - Alternative hashes must be of MD5, SHA1, or SHA256.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) EditTorrentItem(ctx context.Context, apiVersion string, editTorrentItemRequest EditTorrentItemRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.editTorrentItemConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/{api_version}/api/torrents/edittorrent").
		WithConfig(config).
		WithBody(editTorrentItemRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Converts any magnet to a torrent file. Does not require the magnet to be on your TorBox account or in the TorBox cache. Free endpoint for public use.
//
// ### Authorization
//
// None required.
func (api *Service) MagnetToFile(ctx context.Context, apiVersion string, magnetToFileRequest MagnetToFileRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.magnetToFileConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/magnettofile").
		WithConfig(config).
		WithBody(magnetToFileRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Creates a usenet download under your account. Simply send **either** a link, or an nzb file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large.
//
// #### Post Processing Options:
//
// All post processing options that the Usenet client will perform before TorBox's own processing to make the files available. It is recommended you either don't send this parameter, or keep it at `-1` for default, which will give only the wanted files.
//
// - `-1`, Default. This runs repairs, and extractions as well as deletes the source files leaving only the wanted downloaded files.
//
// - `0`, None. No post-processing at all. Just download all the files (including all PAR2). TorBox will still do its normal processing to make the download available, but no repairs, or extraction will take place.
//
// - `1`, Repair. Download files and do a PAR2 verification. If the verification fails, download more PAR2 files and attempt to repair the files.
//
// - `2`, Repair and Unpack. Download all files, do a PAR2 verification and unpack the files. The final folder will also include the RAR and ZIP files.
//
// - `3`, Repair, Unpack and Delete. Download all files, do a PAR2 verification, unpack the files to the final folder and delete the source files.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) CreateUsenetDownload(ctx context.Context, apiVersion string, createUsenetDownloadRequest CreateUsenetDownloadRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.createUsenetDownloadConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/usenet/createusenetdownload").
		WithConfig(config).
		WithBody(createUsenetDownloadRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Controls a usenet download. By sending the usenet download's ID and the type of operation you want to perform, it will send that request to the usenet client.
//
// Operations are either:
//
// - **Delete** `deletes the nzb from the client and your account permanently`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ControlUsenetDownload(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.controlUsenetDownloadConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/usenet/controlusenetdownload").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Permalinks
//
// Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLink1(ctx context.Context, apiVersion string, params RequestDownloadLink1RequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLink1Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/usenet/requestdl").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLinksGet1(ctx context.Context, apiVersion string, params RequestDownloadLinksGet1RequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinksGet1Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/usenet/requestdl/batch").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLinksPost1(ctx context.Context, apiVersion string, mainAPISDKRequestDownloadLinksPostRequest2 MainAPISDKRequestDownloadLinksPostRequest2, params RequestDownloadLinksPost1RequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinksPost1Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/usenet/requestdl/batch").
		WithConfig(config).
		WithBody(mainAPISDKRequestDownloadLinksPostRequest2).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the user's usenet download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live usenet downloads.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetUsenetList(ctx context.Context, apiVersion string, params GetUsenetListRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUsenetListConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/usenet/mylist").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Takes in a list of comma separated usenet hashes and checks if the usenet download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as:
// `?hash=XXXX&hash=XXXX&hash=XXXX`
//
// or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX`
// and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// To get the hash of a usenet download, you can read [this document](https://support.torbox.app/en/articles/13681109-technical-getting-hashes-for-searches) for all of the strategies.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetUsenetCachedAvailability(ctx context.Context, apiVersion string, params GetUsenetCachedAvailabilityRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUsenetCachedAvailabilityConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/usenet/checkcached").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Takes in a list of comma separated hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters.
//
// Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetUsenetCachedAvailabilityBatch(ctx context.Context, apiVersion string, getUsenetCachedAvailabilityBatchRequest GetUsenetCachedAvailabilityBatchRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUsenetCachedAvailabilityBatchConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/usenet/checkcached").
		WithConfig(config).
		WithBody(getUsenetCachedAvailabilityBatchRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Updates a download item based on the item's ID.
// Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important.
//
// ### Restrictions
//
// - Does not affect the cached database (or any other user's item).
//
// - Name must pass validation:
//
// - Above 1 character long.
//
// - Below 200 characters long.
//
// - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically).
//
// - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically).
//
// - Tags follow the same rules as the name.
//
// - Alternative hashes must be of MD5, SHA1, or SHA256.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) EditUsenetItem(ctx context.Context, apiVersion string, editUsenetItemRequest EditUsenetItemRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.editUsenetItemConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/{api_version}/api/usenet/editusenetdownload").
		WithConfig(config).
		WithBody(editUsenetItemRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the TorBox News Server Connection information. This is a static path, and is not dynamic.
//
// ### Authorization
//
// No authentication required.
func (api *Service) GetNewsServerConnectionInformation(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getNewsServerConnectionInformationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/usenet/provider/connection").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the user's TorBox News Server Account with authentication details. Does not show password if not creating the account. Account creation happens on the first request. If you need to get the password, or reset the password, use the `/usenet/provider/account/resetpw` `POST` endpoint.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetNewsServerAccount(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getNewsServerAccountConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/usenet/provider/account").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Resets the user's TorBox News Server account password. This password is automatically generated, and is not saved, it is only shown once.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ResetNewsServerAccountPassword(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.resetNewsServerAccountPasswordConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/usenet/provider/account/resetpw").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Creates a web download under your account. Simply send a link to any file on the internet. Once it has been checked, it will begin downloading assuming your account has available active download slots, and they aren't too large.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) CreateWebDownload(ctx context.Context, apiVersion string, createWebDownloadRequest CreateWebDownloadRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.createWebDownloadConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/webdl/createwebdownload").
		WithConfig(config).
		WithBody(createWebDownloadRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Controls a web download. By sending the web download's ID and the type of operation you want to perform, it will send that request to the debrid client.
//
// Operations are either:
//
// - **Delete** `deletes the download from the client and your account permanently`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ControlWebDownload(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.controlWebDownloadConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/webdl/controlwebdownload").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Permalinks
//
// Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLink2(ctx context.Context, apiVersion string, params RequestDownloadLink2RequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLink2Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/requestdl").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLinksGet2(ctx context.Context, apiVersion string, params RequestDownloadLinksGet2RequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinksGet2Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/requestdl/batch").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *Service) RequestDownloadLinksPost2(ctx context.Context, apiVersion string, mainAPISDKRequestDownloadLinksPostRequest3 MainAPISDKRequestDownloadLinksPostRequest3, params RequestDownloadLinksPost2RequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.requestDownloadLinksPost2Config {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/webdl/requestdl/batch").
		WithConfig(config).
		WithBody(mainAPISDKRequestDownloadLinksPostRequest3).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the user's web download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live web downloads.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetWebDownloadList(ctx context.Context, apiVersion string, params GetWebDownloadListRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getWebDownloadListConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/mylist").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Takes in a list of comma separated usenet hashes and checks if the web download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as:
// `?hash=XXXX&hash=XXXX&hash=XXXX`
//
// or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX`
// and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// To get the hash of a web download you can read [this document](https://support.torbox.app/en/articles/13681109-technical-getting-hashes-for-searches) for all of the strategies.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetWebDownloadCachedAvailability(ctx context.Context, apiVersion string, params GetWebDownloadCachedAvailabilityRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getWebDownloadCachedAvailabilityConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/checkcached").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Takes in a list of comma separated hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters.
//
// Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetWebDownloadCachedAvailabilityBatch(ctx context.Context, apiVersion string, getWebDownloadCachedAvailabilityBatchRequest GetWebDownloadCachedAvailabilityBatchRequest, params GetWebDownloadCachedAvailabilityBatchRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getWebDownloadCachedAvailabilityBatchConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/webdl/checkcached").
		WithConfig(config).
		WithBody(getWebDownloadCachedAvailabilityBatchRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// A dynamic list of hosters that TorBox is capable of downloading through its paid service.
//
// - **Name** - a clean name for display use, the well known name of the service, should be recognizable to users.
//
// - **Domains** - an array of known domains that the hoster uses. While each may serve a different purpose it is still included.
//
// - **URL** - the main url of the service. This should take you to the home page or a service page of the hoster.
//
// - **Icon** - a square image, usually a favicon or logo, that represents the service, should be recognizable as the hoster's icon.
//
// - **Status** - whether this hoster can be used on TorBox or not at the current time. It is usually a good idea to check this value before submitting a download to TorBox's servers for download.
//
// - **Type** - values are either "hoster" or "stream". Both do the same thing, but is good to differentiate services used for different things.
//
// - **Note** - a string value (or null) that may give helpful information about the current status or state of a hoster. This can and should be shown to end users.
//
// - **Daily Link Limit** - the number of downloads a user can use per day. As a user submits links, once they hit this number, the API will deny them from adding anymore of this type of link. A zero value means that it is unlimited.
//
// - **Daily Link Used** - the number of downloads a user has already used. Usually zero unless you send authentication to the endpoint. This will return accurate values.
//
// - **Daily Bandwidth Limit** - the value in bytes that a user is allowed to download from this hoster. A zero value means that it is unlimited. It is recommended to use the Daily Link Limit instead.
//
// - **Daily Bandwidth Used** - the value in bytes that a user has already used to download from this hoster. Usually zero unless you send authentication to the endpoint. This will return accurate values.
//
// ### Authorization
//
// Optional authorization. Authorization is not required in this endpoint unless you want to get the user's live data. Requires an API key using the Authorization Bearer Header to get the live and accurate data for **Daily Link Used** and **Daily Bandwidth Used**.
func (api *Service) GetHosterList(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getHosterListConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/hosters").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Updates a download item based on the item's ID.
// Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important.
//
// ### Restrictions
//
// - Does not affect the cached database (or any other user's item).
//
// - Name must pass validation:
//
// - Above 1 character long.
//
// - Below 200 characters long.
//
// - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically).
//
// - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically).
//
// - Tags follow the same rules as the name.
//
// - Alternative hashes must be of MD5, SHA1, or SHA256.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) EditWebDownloadItem(ctx context.Context, apiVersion string, editWebDownloadItemRequest EditWebDownloadItemRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.editWebDownloadItemConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/{api_version}/api/webdl/editwebdownload").
		WithConfig(config).
		WithBody(editWebDownloadItemRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Simply gets if the TorBox API is available for use or not. Also might include information about downtimes.
//
// ### Authorization
//
// None needed.
func (api *Service) GetUpStatus(ctx context.Context, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUpStatusConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Simply gets general stats about the TorBox service.
//
// ### Authorization
//
// None needed.
func (api *Service) GetStats(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getStatsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/stats").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for people who want updates on the TorBox changelog. Includes all the necessary items to make this compatible with RSS feed viewers for notifications, and proper HTML viewing.
//
// ### Authorization
//
// None needed.
func (api *Service) GetChangelogsRssFeed(ctx context.Context, apiVersion string, params GetChangelogsRssFeedRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getChangelogsRssFeedConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/changelogs/rss").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeXML).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets most recent changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for developers who want to integrate the changelog into their apps for their users to see. Includes content in HTML for developers to easily render the text in a fancy yet simple way.
//
// ### Authorization
//
// None needed.
func (api *Service) GetChangelogsJSON(ctx context.Context, apiVersion string, params GetChangelogsJSONRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getChangelogsJSONConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/changelogs/json").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets CDN speedtest files. This can be used for speedtesting TorBox for users or other usages, such as checking download speeds for verification. Provides all necessary data such as region, server name, and even coordinates. Uses the requesting IP to determine if the server is the closest to the user.
//
// You also have the ability to choose between long tests or short tests via the "test_length" parameter. You may also force the region selection by passing the "region" as a specific region.
//
// ### Authorization
//
// None needed.
func (api *Service) GetSpeedtestFiles(ctx context.Context, apiVersion string, params GetSpeedtestFilesRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getSpeedtestFilesConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/speedtest").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets your notifications in an RSS Feed which allows you to use them with RSS Feed readers or notification services that can take RSS Feeds and listen to updates. As soon as a notification goes to your account, it will be added to your feed.
//
// ### Authorization
//
// Requires an API key using as a query parameter using the `token` key.
func (api *Service) GetRssNotificationFeed(ctx context.Context, apiVersion string, params GetRssNotificationFeedRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getRssNotificationFeedConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/notifications/rss").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeXML).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets your notifications in a JSON object that is easily parsable compared to the RSS Feed. Gives all the same data as the RSS Feed.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetNotificationFeed(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getNotificationFeedConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/notifications/mynotifications").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Marks all of your notifications as read and deletes them permanently.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ClearAllNotifications(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.clearAllNotificationsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/notifications/clear").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Marks a single notification as read and permanently deletes it from your notifications. Requires a `notification_id` which can be found by getting your notification feed.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ClearSingleNotification(ctx context.Context, apiVersion string, notificationID string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.clearSingleNotificationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/notifications/clear/{notification_id}").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		AddPathParam("notification_id", notificationID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Sends a test notification to all enabled notification types. This can be useful for validating setups. No need for any body in this request.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) SendTestNotification(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.sendTestNotificationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/notifications/test").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// If you want a new API token, or your old one has been compromised, you may request a new one. If you happen to forget the token, go the [TorBox settings page ](https://torbox.app/settings) and copy the current one. If this still doesn't work, you may contact us at our support email for a new one.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header as well as passing the `session_token` from the website to be passed in the body. You can find the `session_token` in the localStorage of your browser on any TorBox.app page under the key `torbox_session_token`. This is a temporary token that only lasts for 1 hour, which is why it is used here to verify the identity of a user as well as their API token.
func (api *Service) RefreshAPIToken(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.refreshAPITokenConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/user/refreshtoken").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets a users account data and information.
//
// ### Plans
//
// `0` is Free plan
//
// `1` is Essential plan (_$3 plan_)
//
// `2` is Pro plan (_$10 plan_)
//
// `3` is Standard plan (_$5 plan_)
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetUserData(ctx context.Context, apiVersion string, params GetUserDataRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUserDataConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/me").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Automatically adds a referral code to the user's account. This can be used for developers to automatically add their referral to user's accounts who use their service.
//
// This will not override any referral a user already has. If they already have one, then it cannot be changed using this endpoint. It can only be done by the user on the [website](https://torbox.app/subscription).
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header. Use the user's API key, not your own.
func (api *Service) AddReferralToAccount(ctx context.Context, apiVersion string, params AddReferralToAccountRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.addReferralToAccountConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/user/addreferral").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Requests a 6 digit code to be sent to the user's email for verification. Used to verify a user actually wants to perform a potentially dangerous action.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetConfirmationCode(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getConfirmationCodeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/getconfirmation").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Starts the device code authorization process. Returns all the necessary information to complete the rest of the flow. This includes the following:
//
// - `device_code` - The code your app will use for consecutive requests for getting the user token.
//
// - `code` - This is the 6 digit code that you show to the user to be entered at the following verification_url.
//
// - `verification_url` - This url which you will show as a QR code.
//
// - `friendly_verification_url` - The friendly short url which you can show as plain text to user if they cannot use the QR code method.
//
// You may also submit the `app` parameter to this request to show your app name on the `verification_url` page. This concretes user trust in authorizing your application to use their TorBox account. This is not required, but is a good idea. Default is "Third Party App". This name is also only allowed to be alpha-numeric characters (_A-Z,0-9_). Anything else will be filtered our when being shown to the user. The `friendly_verification_url` will always show "Third Party App" to keep the URL short.
//
// The device_code is only valid for 10 minutes. Make sure the user is aware of this time limit so they submit and login before this time limit is over.
//
// ### Authorization
//
// None required.
func (api *Service) StartDeviceCodeAuthorizationProcess(ctx context.Context, apiVersion string, params StartDeviceCodeAuthorizationProcessRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.startDeviceCodeAuthorizationProcessConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/auth/device/start").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the user's API token from the device code authorization process. This endpoint is expected to be polled in interval until it returns success, or until the expiry has passed. All data you need for this request is in the [original request](https://www.postman.com/wamy-dev/workspace/torbox/request/29572726-744bc274-88b7-4faf-a833-e30573c8c1e0?action=share&creator=29572726&ctx=documentation), which includes the interval to be polled by, the expiry time to stop at, and `device_code` to be sent to the API to get the correct data.
//
// ### Authorization
//
// None required.
func (api *Service) GetTokenFromDeviceCodeProcess(ctx context.Context, apiVersion string, getTokenFromDeviceCodeProcessRequest GetTokenFromDeviceCodeProcessRequest, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTokenFromDeviceCodeProcessConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/user/auth/device/token").
		WithConfig(config).
		WithBody(getTokenFromDeviceCodeProcessRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Use this endpoint to get the user's referral data. This includes information such as how many referrals they have made, and their referral code.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetReferralData(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getReferralDataConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/referraldata").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets a users subscriptions. This includes active and inactive subscriptions for all payment processors.
//
// It should be noted that "plan_code" is not normalized to the standard plan codes such as 1, 2, or 3. These "plan_code"s are payment processor specific and used internally. If you want to display these subscriptions, use the "plan_name".
//
// For displaying information to users, use the "status", "gateway", and "plan_name". No other information is relevant to the user, and the rest is used for TorBox internally.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetSubscriptions(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getSubscriptionsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/subscriptions").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Use this endpoint to get the user's list of transactions. This data isn't exactly relevant to a third party developer, unless you want to build full parity with TorBox's dashboard, or simply want to show this information.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetTransactions(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTransactionsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/transactions").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Generates a PDF of the transaction information. Requires the transaction ID from the "Get Transactions" request.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetTransactionPdf(ctx context.Context, apiVersion string, params GetTransactionPdfRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTransactionPdfConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/transaction/pdf").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeBinary).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Edits a user's TorBox settings. These values take effect immediately. Please view the body to read about allowed options and parameters.
//
// To get the user settings, simply use the "[Get User Data](https://www.postman.com/torbox/workspace/torbox-api/request/29572726-a3775d9b-f3ee-49cd-98a4-fc909ed2923b?action=share&source=copy-link&creator=29572726)" endpoint, while sending the "?settings=true" parameter
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) EditUserSettings(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.editUserSettingsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/{api_version}/api/user/settings/editsettings").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets a users stats. This data is currently split into two sections, "general" and "bandwidth". General shows general data, such as "Total Data Downloaded", "Total Data Egress", "Overall Ratio", and "Total Downloaded". Bandwidth shows the amount of bandwidth a user has downloaded within the past 30 days. We don't store this data for longer than 30 days, and it can be made more or less granular. Makes cool graphs!
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetUserStats(ctx context.Context, apiVersion string, params GetUserStatsRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUserStatsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/stats").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Allows adding an RSS feed to your account. This API gives you a lot of customization options, but is best used with the UI on the website. This endpoint only works for Pro users (plan: 2).
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) AddRssFeed(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.addRssFeedConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/rss/addrss").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Controls an RSS Feed. By sending the RSS feed's ID and the type of operation you want to perform, it will perform said action on the RSS feed checker.
//
// Operations are either:
//
// - **Update** `forces an update on the rss feed`
// - **Delete** `deletes the rss feed from your account permanently. This also deletes all associated RSS feed items. This does not remove queued items spawned from the RSS feed.`
// - **Pause** `pauses checking the rss feed on the scan interval`
// - **Resume** `resumes a paused rss feed`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ControlRssFeed(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.controlRssFeedConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/rss/controlrss").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Allows you to edit the RSS feed based on the RSS feed ID passed. Allows you to change everything about an RSS feed, except for the URL.
//
// When updating the "rss_type" or the "torrent_seeding" options, it will update all of the RSS items associated with the RSS feed, including all previous items.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) EditRssFeed(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.editRssFeedConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/rss/modifyrss").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets all of a user's RSS feeds located on their account. Can only be accessed on a Pro (plan: 2) account. Also allows you to get a specific by passing an "id" in the parameters.
//
// ##### Type Key
//
// The "type" key in the response is used for showing what type of downloads are spawned from this RSS feed. This affects how the item is downloaded, as well as what is scraped, so make sure it is correct.
//
// - **"**torrent**"** `uses .torrent files and magnet links and adds as torrent when downloaded.`
// - "usenet" `uses .nzb files and nzb links and adds as a usenet download when downloaded.`
// - "webdl" `uses generic links and adds as a web download when downloaded.`
//
// ##### Status Key
//
// The "status" key tells the user selected status of the RSS feed. These statuses are **not** the same as the status key for the RSS items associated with the RSS feed. Each feed also includes a "status_message" key which gives a user friendly message about anything about the RSS feed, such as errors and explainations.
//
// - "active" `the RSS feed is set to be scraped.`
// - "paused" `the RSS feed will not be scraped as it is paused by the user.`
//
// - "error" `the RSS feed had an error while scraping, and cannot continue.`
//
// ##### State Key
//
// The "state" key tells you exactly in real time, what the RSS feed is doing. This is useful for realtime status updates.
//
// - "idle" `the RSS feed is idle, and is currently not doing anything in the current moment.`
// - "scraping" `the RSS feed is currently scraping new RSS items.`
// - "downloading" `the RSS feed is currently downloading the newly scraped RSS items.`
//
// ##### Torrent Seeding Key
//
// The "torrent_seeding" key is returned with every response, regardless of the type of the RSS feed. This controls what seeding setting the RSS feed is using when it adds new torrents. These are the same options as when adding a torrent using the /torrent/createtorrent endpoint.
//
// > Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings.
//
// - 1 `auto, the torrent will seed based on what the user has in their settings.`
// - 2 `always, TorBox will be instructed to seed this torrent regardless of any other settings.`
// - 3 `never, TorBox will not seed this torrent.`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetUserRssFeeds(ctx context.Context, apiVersion string, params GetUserRssFeedsRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getUserRssFeedsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/rss/getfeeds").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets the first 10,000 RSS feed items associated with an RSS feed. RSS feed items are scraped items from the RSS feed. They contain a direct download link as well as the content name, and additional information.
//
// ##### Type Key
//
// The "type" key in the response is used for showing what type of downloads this RSS feed item is. This affects how the item is downloaded, as well as what is scraped, so make sure it is correct. When the parent RSS feed's type is changed, it changes all of the items types as well.
//
// - **"**torrent**"** `uses .torrent files and magnet links and adds as torrent when downloaded.`
// - "usenet" `uses .nzb files and nzb links and adds as a usenet download when downloaded.`
// - "webdl" `uses generic links and adds as a web download when downloaded.`
//
// ##### Status Key
//
// The "status" key tells the status of the RSS feed item. This is useful for seeing what the item is doing, and what it will do next. Each item also includes a "status_message" key which gives a user friendly message about anything about the RSS feed item, such as errors and explainations.
//
// - "new" `the RSS feed item is set to be downloaded.`
// - "ignored" `the RSS feed item has been ignored due to the parents RSS feed settings, as it doesn't match the chosen settings. The item will not be downloaded.`
// - "error" `the RSS feed item had an error while attempting to download. The item will be retried on the next scrape.`
// - "downloaded" `the RSS feed item has already been downloaded to the user's account. It will not be downloaded again.`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetRssFeedItems(ctx context.Context, apiVersion string, params GetRssFeedItemsRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getRssFeedItemsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/rss/getfeeditems").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets all the jobs attached to a user account. This is good for an overall view of the jobs, such as on a dashboard, or something similar.
//
// ### Statuses
//
// - "pending" -> Upload is still waiting in the queue. Waiting for spot to upload.
// - "uploading" -> Upload is uploading to the proper remote. Progress will be updated as upload continues.
// - "completed" -> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated.
//
// - "failed" -> The upload has failed. Check the Detail key for information.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetAllJobs(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getAllJobsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/integration/jobs").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets a specifc job using the Job's ID. To get the ID, you will have to Get All Jobs, and get the ID you want.
//
// ### Statuses
//
// - "pending" -> Upload is still waiting in the queue. Waiting for spot to upload.
// - "uploading" -> Upload is uploading to the proper remote. Progress will be updated as upload continues.
// - "completed" -> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated.
// - "failed" -> The upload has failed. Check the Detail key for information.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetSpecificJob(ctx context.Context, apiVersion string, jobID string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getSpecificJobConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/integration/job/{job_id}").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		AddPathParam("job_id", jobID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets all jobs that match a specific hash. Good for checking on specific downloads such as a download page, that could contain a lot of jobs.
//
// ### Statuses
//
// - "pending" -> Upload is still waiting in the queue. Waiting for spot to upload.
// - "uploading" -> Upload is uploading to the proper remote. Progress will be updated as upload continues.
// - "completed" -> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated.
// - "failed" -> The upload has failed. Check the Detail key for information.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetAllJobsByHash(ctx context.Context, apiVersion string, hash string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getAllJobsByHashConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/integration/jobs/{hash}").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		AddPathParam("hash", hash).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Cancels a job or deletes the job. Cancels while in progess (pending, uploading), or deletes the job any other time. It will delete it from the database completely.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) CancelSpecificJob(ctx context.Context, apiVersion string, jobID string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.cancelSpecificJobConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/{api_version}/api/integration/job/{job_id}").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		AddPathParam("job_id", jobID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Retrieves all of a user's queued downloads by type. If you want to get all 3 types, "torrent", "usenet" and "webdl" then you will need to run this request 3 times, each with the different type.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetQueuedDownloads(ctx context.Context, apiVersion string, params GetQueuedDownloadsRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getQueuedDownloadsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/queued/getqueued").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Controls a queued torrent. By sending the queued torrent's ID and the type of operation you want to perform, it will perform that action on the queued torrent.
//
// Operations are either:
//
// - **Delete** `deletes the queued download from your account`
//
// - **Start** `starts a queued download, cannot be used with the "all" parameter`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) ControlQueuedDownloads(ctx context.Context, apiVersion string, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.controlQueuedDownloadsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/queued/controlqueued").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Creates a new stream based off of the data provided.
//
// #### Selecting Subtitles and Audios
//
// We use indexes to determine which audio and subtitle you would like to select. To view all available indexes, you must request this endpoint intially with "null" or 0, for the chosen indexes. This will return a metadata dictionary in the response, which will contain the list of audios and subtitles available for this stream. When choosing a subtitle or audio stream, use the **relative index**, starting from 0, rather than the "index" value. This is important for selecting the correct index. Once your user gets a selection of audios and subtitles available, you can submit this to the endpoint again to get the chosen stream.
//
// - **null** - Selecting no subtitle. Not available for audio as one must always be selected.
//
// - **0** - Selecting first subtitle or audio
//
// #### Selecting Resolution
//
// Similar to the subtitles and audios, we use indexes to determine the resolution chosen. Unlike the above, these are predetermined. We do not upscale video, so if for example, you pass a 1080p video, and request it with resolution, 7, it will stay at 1080p. Please see the below for options.
//
// - **null** - Original video resolution. No downscaling will take place.
//
// - **0** - 144p
//
// - **1** - 240p
//
// - **2** - 360p
//
// - **3** - 480p
//
// - **4** - 720p
//
// - **5** - 1080p
//
// - **6** - 1440p
//
// - **7** - 2160p
//
// #### Types
//
// This endpoint allows you to handle torrents, usenet and web downloads all in one. This is why the "type" parameter is available. It defaults to "torrent".
//
// - **torrent** - Torrents
//
// - **usenet** - Usenet downloads
//
// - **webdownload** - Web Downloads
//
// #### Important Notes
//
// - Only text based subtitles are supported. If a non-text based subtitle is selected, it will be ignored.
//
// - If a subtitle or audio index does not exist on the stream, it will be ignored.
//
// - There are a lot of edge-case videos, if something doesn't work, don't hesitate to report the download hash and file_id of the failing video file.
//
// - This is a **Pro** only feature.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) CreateStream(ctx context.Context, apiVersion string, params CreateStreamRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.createStreamConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/stream/createstream").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}

// ### Overview
//
// Gets stream data from an already created stream. Requires the presigned token and user API token from the original request.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *Service) GetStreamData(ctx context.Context, apiVersion string, params GetStreamDataRequestParams, opts ...mainapisdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getStreamDataConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/stream/getstreamdata").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewMainAPISDKError[[]byte](err)
	}

	return resp.Data, nil
}
