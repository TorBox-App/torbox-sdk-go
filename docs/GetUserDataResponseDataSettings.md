# GetUserDataResponseDataSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailNotifications** | Pointer to **bool** |  | [optional] 
**WebNotifications** | Pointer to **bool** |  | [optional] 
**MobileNotifications** | Pointer to **bool** |  | [optional] 
**RssNotifications** | Pointer to **bool** |  | [optional] 
**DownloadSpeedInTab** | Pointer to **bool** |  | [optional] 
**ShowTrackerInTorrent** | Pointer to **bool** |  | [optional] 
**StremioQuality** | Pointer to **[]float32** |  | [optional] 
**StremioResolution** | Pointer to **[]float32** |  | [optional] 
**StremioLanguage** | Pointer to **[]float32** |  | [optional] 
**StremioCache** | Pointer to **[]float32** |  | [optional] 
**StremioSizeLower** | Pointer to **int64** |  | [optional] 
**StremioSizeUpper** | Pointer to **int64** |  | [optional] 
**GoogleDriveFolderId** | Pointer to **string** |  | [optional] 
**OnedriveSavePath** | Pointer to **string** |  | [optional] 
**DiscordId** | Pointer to **string** |  | [optional] 
**DiscordNotifications** | Pointer to **bool** |  | [optional] 
**StremioAllowAdult** | Pointer to **bool** |  | [optional] 
**WebdavFlatten** | Pointer to **bool** |  | [optional] 
**StremioSeedTorrents** | Pointer to **float32** |  | [optional] 
**SeedTorrents** | Pointer to **float32** |  | [optional] 
**AllowZipped** | Pointer to **bool** |  | [optional] 
**StremioAllowZipped** | Pointer to **bool** |  | [optional] 
**GofileFolderId** | Pointer to **string** |  | [optional] 
**OnefichierFolderId** | Pointer to **string** |  | [optional] 
**JdownloaderNotifications** | Pointer to **bool** |  | [optional] 
**WebhookNotifications** | Pointer to **bool** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**TelegramId** | Pointer to **string** |  | [optional] 
**TelegramNotifications** | Pointer to **bool** |  | [optional] 
**MegaEmail** | Pointer to **string** |  | [optional] 
**MegaPassword** | Pointer to **string** |  | [optional] 
**StremioSort** | Pointer to **string** |  | [optional] 
**DashboardFilter** | Pointer to [**GetUserDataResponseDataSettingsDashboardFilter**](GetUserDataResponseDataSettingsDashboardFilter.md) |  | [optional] 
**DashboardSort** | Pointer to **string** |  | [optional] 
**CdnSelection** | Pointer to **string** |  | [optional] 
**StremioUseCustomSearchEngines** | Pointer to **bool** |  | [optional] 
**StremioResultSort** | Pointer to **string** |  | [optional] 
**StremioLegacyYourMedia** | Pointer to **bool** |  | [optional] 
**WebdavUseLocalFiles** | Pointer to **bool** |  | [optional] 
**StremioDisableYourMediaStreams** | Pointer to **bool** |  | [optional] 
**StremioOnlyYourMediaStreams** | Pointer to **bool** |  | [optional] 
**WebdavUseFolderView** | Pointer to **bool** |  | [optional] 
**StremioLimitPerResolutionTorrent** | Pointer to **int64** |  | [optional] 
**StremioLimitPerResolutionUsenet** | Pointer to **int64** |  | [optional] 
**StremioTorrentSeedersCutoff** | Pointer to **float32** |  | [optional] 
**PixeldrainApiKey** | Pointer to **string** |  | [optional] 
**OnefichierApiKey** | Pointer to **string** |  | [optional] 
**GofileApiKey** | Pointer to **string** |  | [optional] 
**PatreonId** | Pointer to **string** |  | [optional] 
**StremioWaitForDownloadUsenet** | Pointer to **bool** |  | [optional] 
**StremioWaitForDownloadTorrent** | Pointer to **bool** |  | [optional] 
**StremioDisableFilteredNote** | Pointer to **bool** |  | [optional] 
**StremioEmojiInDescription** | Pointer to **bool** |  | [optional] 
**WebPlayerAlwaysTranscode** | Pointer to **bool** |  | [optional] 
**WebPlayerAlwaysSkipIntro** | Pointer to **bool** |  | [optional] 
**WebPlayerAudioPreferredLanguage** | Pointer to **string** |  | [optional] 
**WebPlayerSubtitlePreferredLanguage** | Pointer to **string** |  | [optional] 
**WebPlayerDisablePrestreamSelector** | Pointer to **bool** |  | [optional] 
**WebPlayerDisableNextUpDialogue** | Pointer to **bool** |  | [optional] 
**WebPlayerEnableScrobbling** | Pointer to **bool** |  | [optional] 
**AppendFilenameToLinks** | Pointer to **bool** |  | [optional] 
**DisableAi** | Pointer to **bool** |  | [optional] 
**SubscribedEmailSegments** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetUserDataResponseDataSettings

`func NewGetUserDataResponseDataSettings() *GetUserDataResponseDataSettings`

NewGetUserDataResponseDataSettings instantiates a new GetUserDataResponseDataSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDataResponseDataSettingsWithDefaults

`func NewGetUserDataResponseDataSettingsWithDefaults() *GetUserDataResponseDataSettings`

NewGetUserDataResponseDataSettingsWithDefaults instantiates a new GetUserDataResponseDataSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailNotifications

`func (o *GetUserDataResponseDataSettings) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *GetUserDataResponseDataSettings) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *GetUserDataResponseDataSettings) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetWebNotifications

`func (o *GetUserDataResponseDataSettings) GetWebNotifications() bool`

GetWebNotifications returns the WebNotifications field if non-nil, zero value otherwise.

### GetWebNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetWebNotificationsOk() (*bool, bool)`

GetWebNotificationsOk returns a tuple with the WebNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebNotifications

`func (o *GetUserDataResponseDataSettings) SetWebNotifications(v bool)`

SetWebNotifications sets WebNotifications field to given value.

### HasWebNotifications

`func (o *GetUserDataResponseDataSettings) HasWebNotifications() bool`

HasWebNotifications returns a boolean if a field has been set.

### GetMobileNotifications

`func (o *GetUserDataResponseDataSettings) GetMobileNotifications() bool`

GetMobileNotifications returns the MobileNotifications field if non-nil, zero value otherwise.

### GetMobileNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetMobileNotificationsOk() (*bool, bool)`

GetMobileNotificationsOk returns a tuple with the MobileNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNotifications

`func (o *GetUserDataResponseDataSettings) SetMobileNotifications(v bool)`

SetMobileNotifications sets MobileNotifications field to given value.

### HasMobileNotifications

`func (o *GetUserDataResponseDataSettings) HasMobileNotifications() bool`

HasMobileNotifications returns a boolean if a field has been set.

### GetRssNotifications

`func (o *GetUserDataResponseDataSettings) GetRssNotifications() bool`

GetRssNotifications returns the RssNotifications field if non-nil, zero value otherwise.

### GetRssNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetRssNotificationsOk() (*bool, bool)`

GetRssNotificationsOk returns a tuple with the RssNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssNotifications

`func (o *GetUserDataResponseDataSettings) SetRssNotifications(v bool)`

SetRssNotifications sets RssNotifications field to given value.

### HasRssNotifications

`func (o *GetUserDataResponseDataSettings) HasRssNotifications() bool`

HasRssNotifications returns a boolean if a field has been set.

### GetDownloadSpeedInTab

`func (o *GetUserDataResponseDataSettings) GetDownloadSpeedInTab() bool`

GetDownloadSpeedInTab returns the DownloadSpeedInTab field if non-nil, zero value otherwise.

### GetDownloadSpeedInTabOk

`func (o *GetUserDataResponseDataSettings) GetDownloadSpeedInTabOk() (*bool, bool)`

GetDownloadSpeedInTabOk returns a tuple with the DownloadSpeedInTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeedInTab

`func (o *GetUserDataResponseDataSettings) SetDownloadSpeedInTab(v bool)`

SetDownloadSpeedInTab sets DownloadSpeedInTab field to given value.

### HasDownloadSpeedInTab

`func (o *GetUserDataResponseDataSettings) HasDownloadSpeedInTab() bool`

HasDownloadSpeedInTab returns a boolean if a field has been set.

### GetShowTrackerInTorrent

`func (o *GetUserDataResponseDataSettings) GetShowTrackerInTorrent() bool`

GetShowTrackerInTorrent returns the ShowTrackerInTorrent field if non-nil, zero value otherwise.

### GetShowTrackerInTorrentOk

`func (o *GetUserDataResponseDataSettings) GetShowTrackerInTorrentOk() (*bool, bool)`

GetShowTrackerInTorrentOk returns a tuple with the ShowTrackerInTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTrackerInTorrent

`func (o *GetUserDataResponseDataSettings) SetShowTrackerInTorrent(v bool)`

SetShowTrackerInTorrent sets ShowTrackerInTorrent field to given value.

### HasShowTrackerInTorrent

`func (o *GetUserDataResponseDataSettings) HasShowTrackerInTorrent() bool`

HasShowTrackerInTorrent returns a boolean if a field has been set.

### GetStremioQuality

`func (o *GetUserDataResponseDataSettings) GetStremioQuality() []float32`

GetStremioQuality returns the StremioQuality field if non-nil, zero value otherwise.

### GetStremioQualityOk

`func (o *GetUserDataResponseDataSettings) GetStremioQualityOk() (*[]float32, bool)`

GetStremioQualityOk returns a tuple with the StremioQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioQuality

`func (o *GetUserDataResponseDataSettings) SetStremioQuality(v []float32)`

SetStremioQuality sets StremioQuality field to given value.

### HasStremioQuality

`func (o *GetUserDataResponseDataSettings) HasStremioQuality() bool`

HasStremioQuality returns a boolean if a field has been set.

### GetStremioResolution

`func (o *GetUserDataResponseDataSettings) GetStremioResolution() []float32`

GetStremioResolution returns the StremioResolution field if non-nil, zero value otherwise.

### GetStremioResolutionOk

`func (o *GetUserDataResponseDataSettings) GetStremioResolutionOk() (*[]float32, bool)`

GetStremioResolutionOk returns a tuple with the StremioResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioResolution

`func (o *GetUserDataResponseDataSettings) SetStremioResolution(v []float32)`

SetStremioResolution sets StremioResolution field to given value.

### HasStremioResolution

`func (o *GetUserDataResponseDataSettings) HasStremioResolution() bool`

HasStremioResolution returns a boolean if a field has been set.

### GetStremioLanguage

`func (o *GetUserDataResponseDataSettings) GetStremioLanguage() []float32`

GetStremioLanguage returns the StremioLanguage field if non-nil, zero value otherwise.

### GetStremioLanguageOk

`func (o *GetUserDataResponseDataSettings) GetStremioLanguageOk() (*[]float32, bool)`

GetStremioLanguageOk returns a tuple with the StremioLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLanguage

`func (o *GetUserDataResponseDataSettings) SetStremioLanguage(v []float32)`

SetStremioLanguage sets StremioLanguage field to given value.

### HasStremioLanguage

`func (o *GetUserDataResponseDataSettings) HasStremioLanguage() bool`

HasStremioLanguage returns a boolean if a field has been set.

### GetStremioCache

`func (o *GetUserDataResponseDataSettings) GetStremioCache() []float32`

GetStremioCache returns the StremioCache field if non-nil, zero value otherwise.

### GetStremioCacheOk

`func (o *GetUserDataResponseDataSettings) GetStremioCacheOk() (*[]float32, bool)`

GetStremioCacheOk returns a tuple with the StremioCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioCache

`func (o *GetUserDataResponseDataSettings) SetStremioCache(v []float32)`

SetStremioCache sets StremioCache field to given value.

### HasStremioCache

`func (o *GetUserDataResponseDataSettings) HasStremioCache() bool`

HasStremioCache returns a boolean if a field has been set.

### GetStremioSizeLower

`func (o *GetUserDataResponseDataSettings) GetStremioSizeLower() int64`

GetStremioSizeLower returns the StremioSizeLower field if non-nil, zero value otherwise.

### GetStremioSizeLowerOk

`func (o *GetUserDataResponseDataSettings) GetStremioSizeLowerOk() (*int64, bool)`

GetStremioSizeLowerOk returns a tuple with the StremioSizeLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSizeLower

`func (o *GetUserDataResponseDataSettings) SetStremioSizeLower(v int64)`

SetStremioSizeLower sets StremioSizeLower field to given value.

### HasStremioSizeLower

`func (o *GetUserDataResponseDataSettings) HasStremioSizeLower() bool`

HasStremioSizeLower returns a boolean if a field has been set.

### GetStremioSizeUpper

`func (o *GetUserDataResponseDataSettings) GetStremioSizeUpper() int64`

GetStremioSizeUpper returns the StremioSizeUpper field if non-nil, zero value otherwise.

### GetStremioSizeUpperOk

`func (o *GetUserDataResponseDataSettings) GetStremioSizeUpperOk() (*int64, bool)`

GetStremioSizeUpperOk returns a tuple with the StremioSizeUpper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSizeUpper

`func (o *GetUserDataResponseDataSettings) SetStremioSizeUpper(v int64)`

SetStremioSizeUpper sets StremioSizeUpper field to given value.

### HasStremioSizeUpper

`func (o *GetUserDataResponseDataSettings) HasStremioSizeUpper() bool`

HasStremioSizeUpper returns a boolean if a field has been set.

### GetGoogleDriveFolderId

`func (o *GetUserDataResponseDataSettings) GetGoogleDriveFolderId() string`

GetGoogleDriveFolderId returns the GoogleDriveFolderId field if non-nil, zero value otherwise.

### GetGoogleDriveFolderIdOk

`func (o *GetUserDataResponseDataSettings) GetGoogleDriveFolderIdOk() (*string, bool)`

GetGoogleDriveFolderIdOk returns a tuple with the GoogleDriveFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleDriveFolderId

`func (o *GetUserDataResponseDataSettings) SetGoogleDriveFolderId(v string)`

SetGoogleDriveFolderId sets GoogleDriveFolderId field to given value.

### HasGoogleDriveFolderId

`func (o *GetUserDataResponseDataSettings) HasGoogleDriveFolderId() bool`

HasGoogleDriveFolderId returns a boolean if a field has been set.

### GetOnedriveSavePath

`func (o *GetUserDataResponseDataSettings) GetOnedriveSavePath() string`

GetOnedriveSavePath returns the OnedriveSavePath field if non-nil, zero value otherwise.

### GetOnedriveSavePathOk

`func (o *GetUserDataResponseDataSettings) GetOnedriveSavePathOk() (*string, bool)`

GetOnedriveSavePathOk returns a tuple with the OnedriveSavePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnedriveSavePath

`func (o *GetUserDataResponseDataSettings) SetOnedriveSavePath(v string)`

SetOnedriveSavePath sets OnedriveSavePath field to given value.

### HasOnedriveSavePath

`func (o *GetUserDataResponseDataSettings) HasOnedriveSavePath() bool`

HasOnedriveSavePath returns a boolean if a field has been set.

### GetDiscordId

`func (o *GetUserDataResponseDataSettings) GetDiscordId() string`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *GetUserDataResponseDataSettings) GetDiscordIdOk() (*string, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *GetUserDataResponseDataSettings) SetDiscordId(v string)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *GetUserDataResponseDataSettings) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### GetDiscordNotifications

`func (o *GetUserDataResponseDataSettings) GetDiscordNotifications() bool`

GetDiscordNotifications returns the DiscordNotifications field if non-nil, zero value otherwise.

### GetDiscordNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetDiscordNotificationsOk() (*bool, bool)`

GetDiscordNotificationsOk returns a tuple with the DiscordNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordNotifications

`func (o *GetUserDataResponseDataSettings) SetDiscordNotifications(v bool)`

SetDiscordNotifications sets DiscordNotifications field to given value.

### HasDiscordNotifications

`func (o *GetUserDataResponseDataSettings) HasDiscordNotifications() bool`

HasDiscordNotifications returns a boolean if a field has been set.

### GetStremioAllowAdult

`func (o *GetUserDataResponseDataSettings) GetStremioAllowAdult() bool`

GetStremioAllowAdult returns the StremioAllowAdult field if non-nil, zero value otherwise.

### GetStremioAllowAdultOk

`func (o *GetUserDataResponseDataSettings) GetStremioAllowAdultOk() (*bool, bool)`

GetStremioAllowAdultOk returns a tuple with the StremioAllowAdult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioAllowAdult

`func (o *GetUserDataResponseDataSettings) SetStremioAllowAdult(v bool)`

SetStremioAllowAdult sets StremioAllowAdult field to given value.

### HasStremioAllowAdult

`func (o *GetUserDataResponseDataSettings) HasStremioAllowAdult() bool`

HasStremioAllowAdult returns a boolean if a field has been set.

### GetWebdavFlatten

`func (o *GetUserDataResponseDataSettings) GetWebdavFlatten() bool`

GetWebdavFlatten returns the WebdavFlatten field if non-nil, zero value otherwise.

### GetWebdavFlattenOk

`func (o *GetUserDataResponseDataSettings) GetWebdavFlattenOk() (*bool, bool)`

GetWebdavFlattenOk returns a tuple with the WebdavFlatten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavFlatten

`func (o *GetUserDataResponseDataSettings) SetWebdavFlatten(v bool)`

SetWebdavFlatten sets WebdavFlatten field to given value.

### HasWebdavFlatten

`func (o *GetUserDataResponseDataSettings) HasWebdavFlatten() bool`

HasWebdavFlatten returns a boolean if a field has been set.

### GetStremioSeedTorrents

`func (o *GetUserDataResponseDataSettings) GetStremioSeedTorrents() float32`

GetStremioSeedTorrents returns the StremioSeedTorrents field if non-nil, zero value otherwise.

### GetStremioSeedTorrentsOk

`func (o *GetUserDataResponseDataSettings) GetStremioSeedTorrentsOk() (*float32, bool)`

GetStremioSeedTorrentsOk returns a tuple with the StremioSeedTorrents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSeedTorrents

`func (o *GetUserDataResponseDataSettings) SetStremioSeedTorrents(v float32)`

SetStremioSeedTorrents sets StremioSeedTorrents field to given value.

### HasStremioSeedTorrents

`func (o *GetUserDataResponseDataSettings) HasStremioSeedTorrents() bool`

HasStremioSeedTorrents returns a boolean if a field has been set.

### GetSeedTorrents

`func (o *GetUserDataResponseDataSettings) GetSeedTorrents() float32`

GetSeedTorrents returns the SeedTorrents field if non-nil, zero value otherwise.

### GetSeedTorrentsOk

`func (o *GetUserDataResponseDataSettings) GetSeedTorrentsOk() (*float32, bool)`

GetSeedTorrentsOk returns a tuple with the SeedTorrents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedTorrents

`func (o *GetUserDataResponseDataSettings) SetSeedTorrents(v float32)`

SetSeedTorrents sets SeedTorrents field to given value.

### HasSeedTorrents

`func (o *GetUserDataResponseDataSettings) HasSeedTorrents() bool`

HasSeedTorrents returns a boolean if a field has been set.

### GetAllowZipped

`func (o *GetUserDataResponseDataSettings) GetAllowZipped() bool`

GetAllowZipped returns the AllowZipped field if non-nil, zero value otherwise.

### GetAllowZippedOk

`func (o *GetUserDataResponseDataSettings) GetAllowZippedOk() (*bool, bool)`

GetAllowZippedOk returns a tuple with the AllowZipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZipped

`func (o *GetUserDataResponseDataSettings) SetAllowZipped(v bool)`

SetAllowZipped sets AllowZipped field to given value.

### HasAllowZipped

`func (o *GetUserDataResponseDataSettings) HasAllowZipped() bool`

HasAllowZipped returns a boolean if a field has been set.

### GetStremioAllowZipped

`func (o *GetUserDataResponseDataSettings) GetStremioAllowZipped() bool`

GetStremioAllowZipped returns the StremioAllowZipped field if non-nil, zero value otherwise.

### GetStremioAllowZippedOk

`func (o *GetUserDataResponseDataSettings) GetStremioAllowZippedOk() (*bool, bool)`

GetStremioAllowZippedOk returns a tuple with the StremioAllowZipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioAllowZipped

`func (o *GetUserDataResponseDataSettings) SetStremioAllowZipped(v bool)`

SetStremioAllowZipped sets StremioAllowZipped field to given value.

### HasStremioAllowZipped

`func (o *GetUserDataResponseDataSettings) HasStremioAllowZipped() bool`

HasStremioAllowZipped returns a boolean if a field has been set.

### GetGofileFolderId

`func (o *GetUserDataResponseDataSettings) GetGofileFolderId() string`

GetGofileFolderId returns the GofileFolderId field if non-nil, zero value otherwise.

### GetGofileFolderIdOk

`func (o *GetUserDataResponseDataSettings) GetGofileFolderIdOk() (*string, bool)`

GetGofileFolderIdOk returns a tuple with the GofileFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGofileFolderId

`func (o *GetUserDataResponseDataSettings) SetGofileFolderId(v string)`

SetGofileFolderId sets GofileFolderId field to given value.

### HasGofileFolderId

`func (o *GetUserDataResponseDataSettings) HasGofileFolderId() bool`

HasGofileFolderId returns a boolean if a field has been set.

### GetOnefichierFolderId

`func (o *GetUserDataResponseDataSettings) GetOnefichierFolderId() string`

GetOnefichierFolderId returns the OnefichierFolderId field if non-nil, zero value otherwise.

### GetOnefichierFolderIdOk

`func (o *GetUserDataResponseDataSettings) GetOnefichierFolderIdOk() (*string, bool)`

GetOnefichierFolderIdOk returns a tuple with the OnefichierFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnefichierFolderId

`func (o *GetUserDataResponseDataSettings) SetOnefichierFolderId(v string)`

SetOnefichierFolderId sets OnefichierFolderId field to given value.

### HasOnefichierFolderId

`func (o *GetUserDataResponseDataSettings) HasOnefichierFolderId() bool`

HasOnefichierFolderId returns a boolean if a field has been set.

### GetJdownloaderNotifications

`func (o *GetUserDataResponseDataSettings) GetJdownloaderNotifications() bool`

GetJdownloaderNotifications returns the JdownloaderNotifications field if non-nil, zero value otherwise.

### GetJdownloaderNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetJdownloaderNotificationsOk() (*bool, bool)`

GetJdownloaderNotificationsOk returns a tuple with the JdownloaderNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdownloaderNotifications

`func (o *GetUserDataResponseDataSettings) SetJdownloaderNotifications(v bool)`

SetJdownloaderNotifications sets JdownloaderNotifications field to given value.

### HasJdownloaderNotifications

`func (o *GetUserDataResponseDataSettings) HasJdownloaderNotifications() bool`

HasJdownloaderNotifications returns a boolean if a field has been set.

### GetWebhookNotifications

`func (o *GetUserDataResponseDataSettings) GetWebhookNotifications() bool`

GetWebhookNotifications returns the WebhookNotifications field if non-nil, zero value otherwise.

### GetWebhookNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetWebhookNotificationsOk() (*bool, bool)`

GetWebhookNotificationsOk returns a tuple with the WebhookNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookNotifications

`func (o *GetUserDataResponseDataSettings) SetWebhookNotifications(v bool)`

SetWebhookNotifications sets WebhookNotifications field to given value.

### HasWebhookNotifications

`func (o *GetUserDataResponseDataSettings) HasWebhookNotifications() bool`

HasWebhookNotifications returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *GetUserDataResponseDataSettings) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *GetUserDataResponseDataSettings) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *GetUserDataResponseDataSettings) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *GetUserDataResponseDataSettings) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetTelegramId

`func (o *GetUserDataResponseDataSettings) GetTelegramId() string`

GetTelegramId returns the TelegramId field if non-nil, zero value otherwise.

### GetTelegramIdOk

`func (o *GetUserDataResponseDataSettings) GetTelegramIdOk() (*string, bool)`

GetTelegramIdOk returns a tuple with the TelegramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramId

`func (o *GetUserDataResponseDataSettings) SetTelegramId(v string)`

SetTelegramId sets TelegramId field to given value.

### HasTelegramId

`func (o *GetUserDataResponseDataSettings) HasTelegramId() bool`

HasTelegramId returns a boolean if a field has been set.

### GetTelegramNotifications

`func (o *GetUserDataResponseDataSettings) GetTelegramNotifications() bool`

GetTelegramNotifications returns the TelegramNotifications field if non-nil, zero value otherwise.

### GetTelegramNotificationsOk

`func (o *GetUserDataResponseDataSettings) GetTelegramNotificationsOk() (*bool, bool)`

GetTelegramNotificationsOk returns a tuple with the TelegramNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramNotifications

`func (o *GetUserDataResponseDataSettings) SetTelegramNotifications(v bool)`

SetTelegramNotifications sets TelegramNotifications field to given value.

### HasTelegramNotifications

`func (o *GetUserDataResponseDataSettings) HasTelegramNotifications() bool`

HasTelegramNotifications returns a boolean if a field has been set.

### GetMegaEmail

`func (o *GetUserDataResponseDataSettings) GetMegaEmail() string`

GetMegaEmail returns the MegaEmail field if non-nil, zero value otherwise.

### GetMegaEmailOk

`func (o *GetUserDataResponseDataSettings) GetMegaEmailOk() (*string, bool)`

GetMegaEmailOk returns a tuple with the MegaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMegaEmail

`func (o *GetUserDataResponseDataSettings) SetMegaEmail(v string)`

SetMegaEmail sets MegaEmail field to given value.

### HasMegaEmail

`func (o *GetUserDataResponseDataSettings) HasMegaEmail() bool`

HasMegaEmail returns a boolean if a field has been set.

### GetMegaPassword

`func (o *GetUserDataResponseDataSettings) GetMegaPassword() string`

GetMegaPassword returns the MegaPassword field if non-nil, zero value otherwise.

### GetMegaPasswordOk

`func (o *GetUserDataResponseDataSettings) GetMegaPasswordOk() (*string, bool)`

GetMegaPasswordOk returns a tuple with the MegaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMegaPassword

`func (o *GetUserDataResponseDataSettings) SetMegaPassword(v string)`

SetMegaPassword sets MegaPassword field to given value.

### HasMegaPassword

`func (o *GetUserDataResponseDataSettings) HasMegaPassword() bool`

HasMegaPassword returns a boolean if a field has been set.

### GetStremioSort

`func (o *GetUserDataResponseDataSettings) GetStremioSort() string`

GetStremioSort returns the StremioSort field if non-nil, zero value otherwise.

### GetStremioSortOk

`func (o *GetUserDataResponseDataSettings) GetStremioSortOk() (*string, bool)`

GetStremioSortOk returns a tuple with the StremioSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSort

`func (o *GetUserDataResponseDataSettings) SetStremioSort(v string)`

SetStremioSort sets StremioSort field to given value.

### HasStremioSort

`func (o *GetUserDataResponseDataSettings) HasStremioSort() bool`

HasStremioSort returns a boolean if a field has been set.

### GetDashboardFilter

`func (o *GetUserDataResponseDataSettings) GetDashboardFilter() GetUserDataResponseDataSettingsDashboardFilter`

GetDashboardFilter returns the DashboardFilter field if non-nil, zero value otherwise.

### GetDashboardFilterOk

`func (o *GetUserDataResponseDataSettings) GetDashboardFilterOk() (*GetUserDataResponseDataSettingsDashboardFilter, bool)`

GetDashboardFilterOk returns a tuple with the DashboardFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilter

`func (o *GetUserDataResponseDataSettings) SetDashboardFilter(v GetUserDataResponseDataSettingsDashboardFilter)`

SetDashboardFilter sets DashboardFilter field to given value.

### HasDashboardFilter

`func (o *GetUserDataResponseDataSettings) HasDashboardFilter() bool`

HasDashboardFilter returns a boolean if a field has been set.

### GetDashboardSort

`func (o *GetUserDataResponseDataSettings) GetDashboardSort() string`

GetDashboardSort returns the DashboardSort field if non-nil, zero value otherwise.

### GetDashboardSortOk

`func (o *GetUserDataResponseDataSettings) GetDashboardSortOk() (*string, bool)`

GetDashboardSortOk returns a tuple with the DashboardSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSort

`func (o *GetUserDataResponseDataSettings) SetDashboardSort(v string)`

SetDashboardSort sets DashboardSort field to given value.

### HasDashboardSort

`func (o *GetUserDataResponseDataSettings) HasDashboardSort() bool`

HasDashboardSort returns a boolean if a field has been set.

### GetCdnSelection

`func (o *GetUserDataResponseDataSettings) GetCdnSelection() string`

GetCdnSelection returns the CdnSelection field if non-nil, zero value otherwise.

### GetCdnSelectionOk

`func (o *GetUserDataResponseDataSettings) GetCdnSelectionOk() (*string, bool)`

GetCdnSelectionOk returns a tuple with the CdnSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnSelection

`func (o *GetUserDataResponseDataSettings) SetCdnSelection(v string)`

SetCdnSelection sets CdnSelection field to given value.

### HasCdnSelection

`func (o *GetUserDataResponseDataSettings) HasCdnSelection() bool`

HasCdnSelection returns a boolean if a field has been set.

### GetStremioUseCustomSearchEngines

`func (o *GetUserDataResponseDataSettings) GetStremioUseCustomSearchEngines() bool`

GetStremioUseCustomSearchEngines returns the StremioUseCustomSearchEngines field if non-nil, zero value otherwise.

### GetStremioUseCustomSearchEnginesOk

`func (o *GetUserDataResponseDataSettings) GetStremioUseCustomSearchEnginesOk() (*bool, bool)`

GetStremioUseCustomSearchEnginesOk returns a tuple with the StremioUseCustomSearchEngines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioUseCustomSearchEngines

`func (o *GetUserDataResponseDataSettings) SetStremioUseCustomSearchEngines(v bool)`

SetStremioUseCustomSearchEngines sets StremioUseCustomSearchEngines field to given value.

### HasStremioUseCustomSearchEngines

`func (o *GetUserDataResponseDataSettings) HasStremioUseCustomSearchEngines() bool`

HasStremioUseCustomSearchEngines returns a boolean if a field has been set.

### GetStremioResultSort

`func (o *GetUserDataResponseDataSettings) GetStremioResultSort() string`

GetStremioResultSort returns the StremioResultSort field if non-nil, zero value otherwise.

### GetStremioResultSortOk

`func (o *GetUserDataResponseDataSettings) GetStremioResultSortOk() (*string, bool)`

GetStremioResultSortOk returns a tuple with the StremioResultSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioResultSort

`func (o *GetUserDataResponseDataSettings) SetStremioResultSort(v string)`

SetStremioResultSort sets StremioResultSort field to given value.

### HasStremioResultSort

`func (o *GetUserDataResponseDataSettings) HasStremioResultSort() bool`

HasStremioResultSort returns a boolean if a field has been set.

### GetStremioLegacyYourMedia

`func (o *GetUserDataResponseDataSettings) GetStremioLegacyYourMedia() bool`

GetStremioLegacyYourMedia returns the StremioLegacyYourMedia field if non-nil, zero value otherwise.

### GetStremioLegacyYourMediaOk

`func (o *GetUserDataResponseDataSettings) GetStremioLegacyYourMediaOk() (*bool, bool)`

GetStremioLegacyYourMediaOk returns a tuple with the StremioLegacyYourMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLegacyYourMedia

`func (o *GetUserDataResponseDataSettings) SetStremioLegacyYourMedia(v bool)`

SetStremioLegacyYourMedia sets StremioLegacyYourMedia field to given value.

### HasStremioLegacyYourMedia

`func (o *GetUserDataResponseDataSettings) HasStremioLegacyYourMedia() bool`

HasStremioLegacyYourMedia returns a boolean if a field has been set.

### GetWebdavUseLocalFiles

`func (o *GetUserDataResponseDataSettings) GetWebdavUseLocalFiles() bool`

GetWebdavUseLocalFiles returns the WebdavUseLocalFiles field if non-nil, zero value otherwise.

### GetWebdavUseLocalFilesOk

`func (o *GetUserDataResponseDataSettings) GetWebdavUseLocalFilesOk() (*bool, bool)`

GetWebdavUseLocalFilesOk returns a tuple with the WebdavUseLocalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavUseLocalFiles

`func (o *GetUserDataResponseDataSettings) SetWebdavUseLocalFiles(v bool)`

SetWebdavUseLocalFiles sets WebdavUseLocalFiles field to given value.

### HasWebdavUseLocalFiles

`func (o *GetUserDataResponseDataSettings) HasWebdavUseLocalFiles() bool`

HasWebdavUseLocalFiles returns a boolean if a field has been set.

### GetStremioDisableYourMediaStreams

`func (o *GetUserDataResponseDataSettings) GetStremioDisableYourMediaStreams() bool`

GetStremioDisableYourMediaStreams returns the StremioDisableYourMediaStreams field if non-nil, zero value otherwise.

### GetStremioDisableYourMediaStreamsOk

`func (o *GetUserDataResponseDataSettings) GetStremioDisableYourMediaStreamsOk() (*bool, bool)`

GetStremioDisableYourMediaStreamsOk returns a tuple with the StremioDisableYourMediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioDisableYourMediaStreams

`func (o *GetUserDataResponseDataSettings) SetStremioDisableYourMediaStreams(v bool)`

SetStremioDisableYourMediaStreams sets StremioDisableYourMediaStreams field to given value.

### HasStremioDisableYourMediaStreams

`func (o *GetUserDataResponseDataSettings) HasStremioDisableYourMediaStreams() bool`

HasStremioDisableYourMediaStreams returns a boolean if a field has been set.

### GetStremioOnlyYourMediaStreams

`func (o *GetUserDataResponseDataSettings) GetStremioOnlyYourMediaStreams() bool`

GetStremioOnlyYourMediaStreams returns the StremioOnlyYourMediaStreams field if non-nil, zero value otherwise.

### GetStremioOnlyYourMediaStreamsOk

`func (o *GetUserDataResponseDataSettings) GetStremioOnlyYourMediaStreamsOk() (*bool, bool)`

GetStremioOnlyYourMediaStreamsOk returns a tuple with the StremioOnlyYourMediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioOnlyYourMediaStreams

`func (o *GetUserDataResponseDataSettings) SetStremioOnlyYourMediaStreams(v bool)`

SetStremioOnlyYourMediaStreams sets StremioOnlyYourMediaStreams field to given value.

### HasStremioOnlyYourMediaStreams

`func (o *GetUserDataResponseDataSettings) HasStremioOnlyYourMediaStreams() bool`

HasStremioOnlyYourMediaStreams returns a boolean if a field has been set.

### GetWebdavUseFolderView

`func (o *GetUserDataResponseDataSettings) GetWebdavUseFolderView() bool`

GetWebdavUseFolderView returns the WebdavUseFolderView field if non-nil, zero value otherwise.

### GetWebdavUseFolderViewOk

`func (o *GetUserDataResponseDataSettings) GetWebdavUseFolderViewOk() (*bool, bool)`

GetWebdavUseFolderViewOk returns a tuple with the WebdavUseFolderView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavUseFolderView

`func (o *GetUserDataResponseDataSettings) SetWebdavUseFolderView(v bool)`

SetWebdavUseFolderView sets WebdavUseFolderView field to given value.

### HasWebdavUseFolderView

`func (o *GetUserDataResponseDataSettings) HasWebdavUseFolderView() bool`

HasWebdavUseFolderView returns a boolean if a field has been set.

### GetStremioLimitPerResolutionTorrent

`func (o *GetUserDataResponseDataSettings) GetStremioLimitPerResolutionTorrent() int64`

GetStremioLimitPerResolutionTorrent returns the StremioLimitPerResolutionTorrent field if non-nil, zero value otherwise.

### GetStremioLimitPerResolutionTorrentOk

`func (o *GetUserDataResponseDataSettings) GetStremioLimitPerResolutionTorrentOk() (*int64, bool)`

GetStremioLimitPerResolutionTorrentOk returns a tuple with the StremioLimitPerResolutionTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLimitPerResolutionTorrent

`func (o *GetUserDataResponseDataSettings) SetStremioLimitPerResolutionTorrent(v int64)`

SetStremioLimitPerResolutionTorrent sets StremioLimitPerResolutionTorrent field to given value.

### HasStremioLimitPerResolutionTorrent

`func (o *GetUserDataResponseDataSettings) HasStremioLimitPerResolutionTorrent() bool`

HasStremioLimitPerResolutionTorrent returns a boolean if a field has been set.

### GetStremioLimitPerResolutionUsenet

`func (o *GetUserDataResponseDataSettings) GetStremioLimitPerResolutionUsenet() int64`

GetStremioLimitPerResolutionUsenet returns the StremioLimitPerResolutionUsenet field if non-nil, zero value otherwise.

### GetStremioLimitPerResolutionUsenetOk

`func (o *GetUserDataResponseDataSettings) GetStremioLimitPerResolutionUsenetOk() (*int64, bool)`

GetStremioLimitPerResolutionUsenetOk returns a tuple with the StremioLimitPerResolutionUsenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLimitPerResolutionUsenet

`func (o *GetUserDataResponseDataSettings) SetStremioLimitPerResolutionUsenet(v int64)`

SetStremioLimitPerResolutionUsenet sets StremioLimitPerResolutionUsenet field to given value.

### HasStremioLimitPerResolutionUsenet

`func (o *GetUserDataResponseDataSettings) HasStremioLimitPerResolutionUsenet() bool`

HasStremioLimitPerResolutionUsenet returns a boolean if a field has been set.

### GetStremioTorrentSeedersCutoff

`func (o *GetUserDataResponseDataSettings) GetStremioTorrentSeedersCutoff() float32`

GetStremioTorrentSeedersCutoff returns the StremioTorrentSeedersCutoff field if non-nil, zero value otherwise.

### GetStremioTorrentSeedersCutoffOk

`func (o *GetUserDataResponseDataSettings) GetStremioTorrentSeedersCutoffOk() (*float32, bool)`

GetStremioTorrentSeedersCutoffOk returns a tuple with the StremioTorrentSeedersCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioTorrentSeedersCutoff

`func (o *GetUserDataResponseDataSettings) SetStremioTorrentSeedersCutoff(v float32)`

SetStremioTorrentSeedersCutoff sets StremioTorrentSeedersCutoff field to given value.

### HasStremioTorrentSeedersCutoff

`func (o *GetUserDataResponseDataSettings) HasStremioTorrentSeedersCutoff() bool`

HasStremioTorrentSeedersCutoff returns a boolean if a field has been set.

### GetPixeldrainApiKey

`func (o *GetUserDataResponseDataSettings) GetPixeldrainApiKey() string`

GetPixeldrainApiKey returns the PixeldrainApiKey field if non-nil, zero value otherwise.

### GetPixeldrainApiKeyOk

`func (o *GetUserDataResponseDataSettings) GetPixeldrainApiKeyOk() (*string, bool)`

GetPixeldrainApiKeyOk returns a tuple with the PixeldrainApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixeldrainApiKey

`func (o *GetUserDataResponseDataSettings) SetPixeldrainApiKey(v string)`

SetPixeldrainApiKey sets PixeldrainApiKey field to given value.

### HasPixeldrainApiKey

`func (o *GetUserDataResponseDataSettings) HasPixeldrainApiKey() bool`

HasPixeldrainApiKey returns a boolean if a field has been set.

### GetOnefichierApiKey

`func (o *GetUserDataResponseDataSettings) GetOnefichierApiKey() string`

GetOnefichierApiKey returns the OnefichierApiKey field if non-nil, zero value otherwise.

### GetOnefichierApiKeyOk

`func (o *GetUserDataResponseDataSettings) GetOnefichierApiKeyOk() (*string, bool)`

GetOnefichierApiKeyOk returns a tuple with the OnefichierApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnefichierApiKey

`func (o *GetUserDataResponseDataSettings) SetOnefichierApiKey(v string)`

SetOnefichierApiKey sets OnefichierApiKey field to given value.

### HasOnefichierApiKey

`func (o *GetUserDataResponseDataSettings) HasOnefichierApiKey() bool`

HasOnefichierApiKey returns a boolean if a field has been set.

### GetGofileApiKey

`func (o *GetUserDataResponseDataSettings) GetGofileApiKey() string`

GetGofileApiKey returns the GofileApiKey field if non-nil, zero value otherwise.

### GetGofileApiKeyOk

`func (o *GetUserDataResponseDataSettings) GetGofileApiKeyOk() (*string, bool)`

GetGofileApiKeyOk returns a tuple with the GofileApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGofileApiKey

`func (o *GetUserDataResponseDataSettings) SetGofileApiKey(v string)`

SetGofileApiKey sets GofileApiKey field to given value.

### HasGofileApiKey

`func (o *GetUserDataResponseDataSettings) HasGofileApiKey() bool`

HasGofileApiKey returns a boolean if a field has been set.

### GetPatreonId

`func (o *GetUserDataResponseDataSettings) GetPatreonId() string`

GetPatreonId returns the PatreonId field if non-nil, zero value otherwise.

### GetPatreonIdOk

`func (o *GetUserDataResponseDataSettings) GetPatreonIdOk() (*string, bool)`

GetPatreonIdOk returns a tuple with the PatreonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatreonId

`func (o *GetUserDataResponseDataSettings) SetPatreonId(v string)`

SetPatreonId sets PatreonId field to given value.

### HasPatreonId

`func (o *GetUserDataResponseDataSettings) HasPatreonId() bool`

HasPatreonId returns a boolean if a field has been set.

### GetStremioWaitForDownloadUsenet

`func (o *GetUserDataResponseDataSettings) GetStremioWaitForDownloadUsenet() bool`

GetStremioWaitForDownloadUsenet returns the StremioWaitForDownloadUsenet field if non-nil, zero value otherwise.

### GetStremioWaitForDownloadUsenetOk

`func (o *GetUserDataResponseDataSettings) GetStremioWaitForDownloadUsenetOk() (*bool, bool)`

GetStremioWaitForDownloadUsenetOk returns a tuple with the StremioWaitForDownloadUsenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioWaitForDownloadUsenet

`func (o *GetUserDataResponseDataSettings) SetStremioWaitForDownloadUsenet(v bool)`

SetStremioWaitForDownloadUsenet sets StremioWaitForDownloadUsenet field to given value.

### HasStremioWaitForDownloadUsenet

`func (o *GetUserDataResponseDataSettings) HasStremioWaitForDownloadUsenet() bool`

HasStremioWaitForDownloadUsenet returns a boolean if a field has been set.

### GetStremioWaitForDownloadTorrent

`func (o *GetUserDataResponseDataSettings) GetStremioWaitForDownloadTorrent() bool`

GetStremioWaitForDownloadTorrent returns the StremioWaitForDownloadTorrent field if non-nil, zero value otherwise.

### GetStremioWaitForDownloadTorrentOk

`func (o *GetUserDataResponseDataSettings) GetStremioWaitForDownloadTorrentOk() (*bool, bool)`

GetStremioWaitForDownloadTorrentOk returns a tuple with the StremioWaitForDownloadTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioWaitForDownloadTorrent

`func (o *GetUserDataResponseDataSettings) SetStremioWaitForDownloadTorrent(v bool)`

SetStremioWaitForDownloadTorrent sets StremioWaitForDownloadTorrent field to given value.

### HasStremioWaitForDownloadTorrent

`func (o *GetUserDataResponseDataSettings) HasStremioWaitForDownloadTorrent() bool`

HasStremioWaitForDownloadTorrent returns a boolean if a field has been set.

### GetStremioDisableFilteredNote

`func (o *GetUserDataResponseDataSettings) GetStremioDisableFilteredNote() bool`

GetStremioDisableFilteredNote returns the StremioDisableFilteredNote field if non-nil, zero value otherwise.

### GetStremioDisableFilteredNoteOk

`func (o *GetUserDataResponseDataSettings) GetStremioDisableFilteredNoteOk() (*bool, bool)`

GetStremioDisableFilteredNoteOk returns a tuple with the StremioDisableFilteredNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioDisableFilteredNote

`func (o *GetUserDataResponseDataSettings) SetStremioDisableFilteredNote(v bool)`

SetStremioDisableFilteredNote sets StremioDisableFilteredNote field to given value.

### HasStremioDisableFilteredNote

`func (o *GetUserDataResponseDataSettings) HasStremioDisableFilteredNote() bool`

HasStremioDisableFilteredNote returns a boolean if a field has been set.

### GetStremioEmojiInDescription

`func (o *GetUserDataResponseDataSettings) GetStremioEmojiInDescription() bool`

GetStremioEmojiInDescription returns the StremioEmojiInDescription field if non-nil, zero value otherwise.

### GetStremioEmojiInDescriptionOk

`func (o *GetUserDataResponseDataSettings) GetStremioEmojiInDescriptionOk() (*bool, bool)`

GetStremioEmojiInDescriptionOk returns a tuple with the StremioEmojiInDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioEmojiInDescription

`func (o *GetUserDataResponseDataSettings) SetStremioEmojiInDescription(v bool)`

SetStremioEmojiInDescription sets StremioEmojiInDescription field to given value.

### HasStremioEmojiInDescription

`func (o *GetUserDataResponseDataSettings) HasStremioEmojiInDescription() bool`

HasStremioEmojiInDescription returns a boolean if a field has been set.

### GetWebPlayerAlwaysTranscode

`func (o *GetUserDataResponseDataSettings) GetWebPlayerAlwaysTranscode() bool`

GetWebPlayerAlwaysTranscode returns the WebPlayerAlwaysTranscode field if non-nil, zero value otherwise.

### GetWebPlayerAlwaysTranscodeOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerAlwaysTranscodeOk() (*bool, bool)`

GetWebPlayerAlwaysTranscodeOk returns a tuple with the WebPlayerAlwaysTranscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerAlwaysTranscode

`func (o *GetUserDataResponseDataSettings) SetWebPlayerAlwaysTranscode(v bool)`

SetWebPlayerAlwaysTranscode sets WebPlayerAlwaysTranscode field to given value.

### HasWebPlayerAlwaysTranscode

`func (o *GetUserDataResponseDataSettings) HasWebPlayerAlwaysTranscode() bool`

HasWebPlayerAlwaysTranscode returns a boolean if a field has been set.

### GetWebPlayerAlwaysSkipIntro

`func (o *GetUserDataResponseDataSettings) GetWebPlayerAlwaysSkipIntro() bool`

GetWebPlayerAlwaysSkipIntro returns the WebPlayerAlwaysSkipIntro field if non-nil, zero value otherwise.

### GetWebPlayerAlwaysSkipIntroOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerAlwaysSkipIntroOk() (*bool, bool)`

GetWebPlayerAlwaysSkipIntroOk returns a tuple with the WebPlayerAlwaysSkipIntro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerAlwaysSkipIntro

`func (o *GetUserDataResponseDataSettings) SetWebPlayerAlwaysSkipIntro(v bool)`

SetWebPlayerAlwaysSkipIntro sets WebPlayerAlwaysSkipIntro field to given value.

### HasWebPlayerAlwaysSkipIntro

`func (o *GetUserDataResponseDataSettings) HasWebPlayerAlwaysSkipIntro() bool`

HasWebPlayerAlwaysSkipIntro returns a boolean if a field has been set.

### GetWebPlayerAudioPreferredLanguage

`func (o *GetUserDataResponseDataSettings) GetWebPlayerAudioPreferredLanguage() string`

GetWebPlayerAudioPreferredLanguage returns the WebPlayerAudioPreferredLanguage field if non-nil, zero value otherwise.

### GetWebPlayerAudioPreferredLanguageOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerAudioPreferredLanguageOk() (*string, bool)`

GetWebPlayerAudioPreferredLanguageOk returns a tuple with the WebPlayerAudioPreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerAudioPreferredLanguage

`func (o *GetUserDataResponseDataSettings) SetWebPlayerAudioPreferredLanguage(v string)`

SetWebPlayerAudioPreferredLanguage sets WebPlayerAudioPreferredLanguage field to given value.

### HasWebPlayerAudioPreferredLanguage

`func (o *GetUserDataResponseDataSettings) HasWebPlayerAudioPreferredLanguage() bool`

HasWebPlayerAudioPreferredLanguage returns a boolean if a field has been set.

### GetWebPlayerSubtitlePreferredLanguage

`func (o *GetUserDataResponseDataSettings) GetWebPlayerSubtitlePreferredLanguage() string`

GetWebPlayerSubtitlePreferredLanguage returns the WebPlayerSubtitlePreferredLanguage field if non-nil, zero value otherwise.

### GetWebPlayerSubtitlePreferredLanguageOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerSubtitlePreferredLanguageOk() (*string, bool)`

GetWebPlayerSubtitlePreferredLanguageOk returns a tuple with the WebPlayerSubtitlePreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerSubtitlePreferredLanguage

`func (o *GetUserDataResponseDataSettings) SetWebPlayerSubtitlePreferredLanguage(v string)`

SetWebPlayerSubtitlePreferredLanguage sets WebPlayerSubtitlePreferredLanguage field to given value.

### HasWebPlayerSubtitlePreferredLanguage

`func (o *GetUserDataResponseDataSettings) HasWebPlayerSubtitlePreferredLanguage() bool`

HasWebPlayerSubtitlePreferredLanguage returns a boolean if a field has been set.

### GetWebPlayerDisablePrestreamSelector

`func (o *GetUserDataResponseDataSettings) GetWebPlayerDisablePrestreamSelector() bool`

GetWebPlayerDisablePrestreamSelector returns the WebPlayerDisablePrestreamSelector field if non-nil, zero value otherwise.

### GetWebPlayerDisablePrestreamSelectorOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerDisablePrestreamSelectorOk() (*bool, bool)`

GetWebPlayerDisablePrestreamSelectorOk returns a tuple with the WebPlayerDisablePrestreamSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerDisablePrestreamSelector

`func (o *GetUserDataResponseDataSettings) SetWebPlayerDisablePrestreamSelector(v bool)`

SetWebPlayerDisablePrestreamSelector sets WebPlayerDisablePrestreamSelector field to given value.

### HasWebPlayerDisablePrestreamSelector

`func (o *GetUserDataResponseDataSettings) HasWebPlayerDisablePrestreamSelector() bool`

HasWebPlayerDisablePrestreamSelector returns a boolean if a field has been set.

### GetWebPlayerDisableNextUpDialogue

`func (o *GetUserDataResponseDataSettings) GetWebPlayerDisableNextUpDialogue() bool`

GetWebPlayerDisableNextUpDialogue returns the WebPlayerDisableNextUpDialogue field if non-nil, zero value otherwise.

### GetWebPlayerDisableNextUpDialogueOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerDisableNextUpDialogueOk() (*bool, bool)`

GetWebPlayerDisableNextUpDialogueOk returns a tuple with the WebPlayerDisableNextUpDialogue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerDisableNextUpDialogue

`func (o *GetUserDataResponseDataSettings) SetWebPlayerDisableNextUpDialogue(v bool)`

SetWebPlayerDisableNextUpDialogue sets WebPlayerDisableNextUpDialogue field to given value.

### HasWebPlayerDisableNextUpDialogue

`func (o *GetUserDataResponseDataSettings) HasWebPlayerDisableNextUpDialogue() bool`

HasWebPlayerDisableNextUpDialogue returns a boolean if a field has been set.

### GetWebPlayerEnableScrobbling

`func (o *GetUserDataResponseDataSettings) GetWebPlayerEnableScrobbling() bool`

GetWebPlayerEnableScrobbling returns the WebPlayerEnableScrobbling field if non-nil, zero value otherwise.

### GetWebPlayerEnableScrobblingOk

`func (o *GetUserDataResponseDataSettings) GetWebPlayerEnableScrobblingOk() (*bool, bool)`

GetWebPlayerEnableScrobblingOk returns a tuple with the WebPlayerEnableScrobbling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerEnableScrobbling

`func (o *GetUserDataResponseDataSettings) SetWebPlayerEnableScrobbling(v bool)`

SetWebPlayerEnableScrobbling sets WebPlayerEnableScrobbling field to given value.

### HasWebPlayerEnableScrobbling

`func (o *GetUserDataResponseDataSettings) HasWebPlayerEnableScrobbling() bool`

HasWebPlayerEnableScrobbling returns a boolean if a field has been set.

### GetAppendFilenameToLinks

`func (o *GetUserDataResponseDataSettings) GetAppendFilenameToLinks() bool`

GetAppendFilenameToLinks returns the AppendFilenameToLinks field if non-nil, zero value otherwise.

### GetAppendFilenameToLinksOk

`func (o *GetUserDataResponseDataSettings) GetAppendFilenameToLinksOk() (*bool, bool)`

GetAppendFilenameToLinksOk returns a tuple with the AppendFilenameToLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendFilenameToLinks

`func (o *GetUserDataResponseDataSettings) SetAppendFilenameToLinks(v bool)`

SetAppendFilenameToLinks sets AppendFilenameToLinks field to given value.

### HasAppendFilenameToLinks

`func (o *GetUserDataResponseDataSettings) HasAppendFilenameToLinks() bool`

HasAppendFilenameToLinks returns a boolean if a field has been set.

### GetDisableAi

`func (o *GetUserDataResponseDataSettings) GetDisableAi() bool`

GetDisableAi returns the DisableAi field if non-nil, zero value otherwise.

### GetDisableAiOk

`func (o *GetUserDataResponseDataSettings) GetDisableAiOk() (*bool, bool)`

GetDisableAiOk returns a tuple with the DisableAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAi

`func (o *GetUserDataResponseDataSettings) SetDisableAi(v bool)`

SetDisableAi sets DisableAi field to given value.

### HasDisableAi

`func (o *GetUserDataResponseDataSettings) HasDisableAi() bool`

HasDisableAi returns a boolean if a field has been set.

### GetSubscribedEmailSegments

`func (o *GetUserDataResponseDataSettings) GetSubscribedEmailSegments() []string`

GetSubscribedEmailSegments returns the SubscribedEmailSegments field if non-nil, zero value otherwise.

### GetSubscribedEmailSegmentsOk

`func (o *GetUserDataResponseDataSettings) GetSubscribedEmailSegmentsOk() (*[]string, bool)`

GetSubscribedEmailSegmentsOk returns a tuple with the SubscribedEmailSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEmailSegments

`func (o *GetUserDataResponseDataSettings) SetSubscribedEmailSegments(v []string)`

SetSubscribedEmailSegments sets SubscribedEmailSegments field to given value.

### HasSubscribedEmailSegments

`func (o *GetUserDataResponseDataSettings) HasSubscribedEmailSegments() bool`

HasSubscribedEmailSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


