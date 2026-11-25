# BaseSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailNotifications** | Pointer to **bool** | When the user recieves a notification, send to registered email | [optional] 
**WebNotifications** | Pointer to **bool** | When the user recieves a notification, send to web (if browser is open, otherwise it will be in their notifications panel) | [optional] 
**MobileNotifications** | Pointer to **bool** | When the user recieves a notification, send to TorBox mobile app (deprecated) | [optional] 
**RssNotifications** | Pointer to **bool** | When the user recieves a notification, add to user&#39;s RSS feed | [optional] 
**DownloadSpeedInTab** | Pointer to **bool** | Shows download speed in tab name in browser when using the TorBox dashboard. Third party apps may still use this setting. | [optional] 
**ShowTrackerInTorrent** | Pointer to **bool** | Shows active tracker in torrents on the TorBox dashboard. Third party apps may still use this setting. | [optional] 
**StremioQuality** | Pointer to **[]int64** | Determines what qualities are shown to the user when using the TorBox Stremio addon. It is always recommended to include 0 (Unknown). Third party apps may still use this setting.  - &#x60;0&#x60;: Unknown Quality - &#x60;1&#x60;: Cam Quality - &#x60;2&#x60;: DVD-R Quality - &#x60;3&#x60;: HDTV Quality - &#x60;4&#x60;: HDRip Quality - &#x60;5&#x60;: WEB-DL Quality - &#x60;6&#x60;: WEBRip Quality - &#x60;7&#x60;: Blu-Ray/BD/BDRip Quality - &#x60;8&#x60;: HDR Quality - &#x60;9&#x60;: HDR10 Quality - &#x60;10&#x60;: HDR10+ Quality - &#x60;11&#x60;: Dolby Vision Quality  | [optional] 
**StremioResolution** | Pointer to **[]int64** | Determines what resolutions are shown to the user. It is always recommended to include 0 (Unknown). Third party apps may still use this setting.  - &#x60;0&#x60;: Unknown Resolution - &#x60;1&#x60;: SD Resolution - &#x60;2&#x60;: HD Resolution - &#x60;3&#x60;: UHD Resolution  | [optional] 
**StremioLanguage** | Pointer to **[]int64** | Determines what languages are shown to the user. It is always recommended to include 0 (Unknown). Third party apps may still use this setting.  - &#x60;0&#x60;: Unknown - &#x60;1&#x60;: English - &#x60;2&#x60;: Spanish - &#x60;3&#x60;: French - &#x60;4&#x60;: German - &#x60;5&#x60;: Italian - &#x60;6&#x60;: Portuguese - &#x60;7&#x60;: Russian - &#x60;8&#x60;: Japanese - &#x60;9&#x60;: Korean - &#x60;10&#x60;: Chinese - &#x60;11&#x60;: Arabic - &#x60;12&#x60;: Hindi - &#x60;13&#x60;: Bengali - &#x60;14&#x60;: Ukrainian - &#x60;15&#x60;: Nordic - &#x60;16&#x60;: Czech - &#x60;17&#x60;: Polish - &#x60;18&#x60;: Tamil - &#x60;19&#x60;: Telugu - &#x60;20&#x60;: Kannada - &#x60;21&#x60;: Dutch  | [optional] 
**StremioCache** | Pointer to **[]int64** | Determines what cached status items to show to users. While this is an array, only a single value is accepted at a time to prevent issues. Still submit this as an array to the API. Third party apps may still use this setting.  - &#x60;1&#x60;: Cached and Download - &#x60;2&#x60;: Cached Only - &#x60;3&#x60;: Download Only  | [optional] 
**StremioSizeLower** | Pointer to **int64** | The minimum size of files that will be shown to the user in the TorBox Stremio addon. Third party apps may still use this setting. | [optional] 
**GoogleDriveFolderId** | Pointer to **string** |  | [optional] 
**OnedriveSavePath** | Pointer to **string** |  | [optional] 
**DiscordId** | Pointer to **string** |  | [optional] 
**DiscordNotifications** | Pointer to **bool** | When the user recieves a notification, use the official TorBox Discord Bot to DM the user the notification. Must have a Discord_ID present and must share a server with the Discord Bot. | [optional] 
**StremioAllowAdult** | Pointer to **bool** | Determines whether to show adult content on the TorBox Stremio addon. Third party apps may still use this setting. | [optional] 
**StremioSeedTorrents** | Pointer to **int64** | 1 &#x3D; Auto, 2 &#x3D; Always, 3 &#x3D; Never | [optional] 
**SeedTorrents** | Pointer to **int64** | 1 &#x3D; Auto, 2 &#x3D; Always, 3 &#x3D; Never | [optional] 
**AllowZipped** | Pointer to **bool** |  | [optional] 
**StremioAllowZipped** | Pointer to **bool** |  | [optional] 
**OnefichierFolderId** | Pointer to **string** |  | [optional] 
**GofileFolderId** | Pointer to **string** |  | [optional] 
**WebhookNotifications** | Pointer to **bool** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**TelegramNotifications** | Pointer to **bool** |  | [optional] 
**TelegramId** | Pointer to **string** |  | [optional] 
**MegaEmail** | Pointer to **string** |  | [optional] 
**MegaPassword** | Pointer to **string** |  | [optional] 
**StremioSort** | Pointer to **string** | seeders &#x3D; Seeders/Age (highest seeders first/lowest age first), size &#x3D; Size (highest size first), resolution &#x3D; Resolution (highest resolution first) | [optional] 
**CdnSelection** | Pointer to **string** | any available CDN. Can be found using test file route | [optional] 
**StremioUseCustomSearchEngines** | Pointer to **bool** |  | [optional] 
**StremioResultSort** | Pointer to **string** | torrents &#x3D; Torrents first, usenet &#x3D; Usenet first | [optional] 
**WebdavUseLocalFiles** | Pointer to **bool** |  | [optional] 
**StremioLegacyYourMedia** | Pointer to **bool** |  | [optional] 
**StremioOnlyYourMediaStreams** | Pointer to **bool** | only shows Your Media streams in Stremio addon | [optional] 
**StremioDisableYourMediaStreams** | Pointer to **bool** |  | [optional] 
**WebdavUseFolderView** | Pointer to **bool** |  | [optional] 
**WebdavFlatten** | Pointer to **bool** |  | [optional] 
**StremioLimitPerResolutionTorrent** | Pointer to **int64** |  | [optional] 
**StremioLimitPerResolutionUsenet** | Pointer to **int64** |  | [optional] 
**StremioSizeUpper** | Pointer to **int64** |  | [optional] 
**StremioTorrentSeedersCutoff** | Pointer to **int64** |  | [optional] 
**PixeldrainApiKey** | Pointer to **string** |  | [optional] 
**OnefichierApiKey** | Pointer to **string** |  | [optional] 
**GofileApiKey** | Pointer to **string** |  | [optional] 
**PatreonId** | Pointer to **string** |  | [optional] 
**StremioWaitForDownloadUsenet** | Pointer to **bool** |  | [optional] 
**StremioWaitForDownloadTorrent** | Pointer to **bool** |  | [optional] 
**StremioDisableFilteredNote** | Pointer to **bool** |  | [optional] 
**StremioEmojiInDescription** | Pointer to **bool** |  | [optional] 
**DashboardFilter** | Pointer to **map[string]interface{}** | sorts the user&#39;s dashboard. Does not affect third party dashboards. This entire dict MUST be presented with all keys for any update to be given. | [optional] 
**WebPlayerAlwaysTranscode** | Pointer to **bool** |  | [optional] 
**WebPlayerAlwaysSkipIntro** | Pointer to **bool** |  | [optional] 
**WebPlayerAudioPreferredLanguage** | Pointer to **string** |  | [optional] 
**WebPlayerSubtitlePreferredLanguage** | Pointer to **string** |  | [optional] 
**WebPlayerDisablePrestreamSelector** | Pointer to **bool** |  | [optional] 
**DashboardSort** | Pointer to **string** | default &#x3D; TorBox default sorting, name &#x3D; AZ sorting, size &#x3D; largest size first, added_date &#x3D; newest first, cached_date &#x3D; newest first, last_updated &#x3D; newest first, progress &#x3D; highest progress first, ratio &#x3D; highest ratio first, download_speed &#x3D; highest value first, upload_speed &#x3D; highest value first, eta &#x3D; smallest eta first | [optional] 
**WebPlayerDisableNextUpDialogue** | Pointer to **bool** |  | [optional] 
**WebPlayerEnableScrobbling** | Pointer to **bool** |  | [optional] 
**AppendFilenameToLinks** | Pointer to **bool** |  | [optional] 
**SubscribedEmailSegments** | Pointer to **[]string** |  | [optional] 
**DisableAi** | Pointer to **bool** |  | [optional] 

## Methods

### NewBaseSettingsModel

`func NewBaseSettingsModel() *BaseSettingsModel`

NewBaseSettingsModel instantiates a new BaseSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSettingsModelWithDefaults

`func NewBaseSettingsModelWithDefaults() *BaseSettingsModel`

NewBaseSettingsModelWithDefaults instantiates a new BaseSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailNotifications

`func (o *BaseSettingsModel) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *BaseSettingsModel) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *BaseSettingsModel) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *BaseSettingsModel) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetWebNotifications

`func (o *BaseSettingsModel) GetWebNotifications() bool`

GetWebNotifications returns the WebNotifications field if non-nil, zero value otherwise.

### GetWebNotificationsOk

`func (o *BaseSettingsModel) GetWebNotificationsOk() (*bool, bool)`

GetWebNotificationsOk returns a tuple with the WebNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebNotifications

`func (o *BaseSettingsModel) SetWebNotifications(v bool)`

SetWebNotifications sets WebNotifications field to given value.

### HasWebNotifications

`func (o *BaseSettingsModel) HasWebNotifications() bool`

HasWebNotifications returns a boolean if a field has been set.

### GetMobileNotifications

`func (o *BaseSettingsModel) GetMobileNotifications() bool`

GetMobileNotifications returns the MobileNotifications field if non-nil, zero value otherwise.

### GetMobileNotificationsOk

`func (o *BaseSettingsModel) GetMobileNotificationsOk() (*bool, bool)`

GetMobileNotificationsOk returns a tuple with the MobileNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNotifications

`func (o *BaseSettingsModel) SetMobileNotifications(v bool)`

SetMobileNotifications sets MobileNotifications field to given value.

### HasMobileNotifications

`func (o *BaseSettingsModel) HasMobileNotifications() bool`

HasMobileNotifications returns a boolean if a field has been set.

### GetRssNotifications

`func (o *BaseSettingsModel) GetRssNotifications() bool`

GetRssNotifications returns the RssNotifications field if non-nil, zero value otherwise.

### GetRssNotificationsOk

`func (o *BaseSettingsModel) GetRssNotificationsOk() (*bool, bool)`

GetRssNotificationsOk returns a tuple with the RssNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssNotifications

`func (o *BaseSettingsModel) SetRssNotifications(v bool)`

SetRssNotifications sets RssNotifications field to given value.

### HasRssNotifications

`func (o *BaseSettingsModel) HasRssNotifications() bool`

HasRssNotifications returns a boolean if a field has been set.

### GetDownloadSpeedInTab

`func (o *BaseSettingsModel) GetDownloadSpeedInTab() bool`

GetDownloadSpeedInTab returns the DownloadSpeedInTab field if non-nil, zero value otherwise.

### GetDownloadSpeedInTabOk

`func (o *BaseSettingsModel) GetDownloadSpeedInTabOk() (*bool, bool)`

GetDownloadSpeedInTabOk returns a tuple with the DownloadSpeedInTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeedInTab

`func (o *BaseSettingsModel) SetDownloadSpeedInTab(v bool)`

SetDownloadSpeedInTab sets DownloadSpeedInTab field to given value.

### HasDownloadSpeedInTab

`func (o *BaseSettingsModel) HasDownloadSpeedInTab() bool`

HasDownloadSpeedInTab returns a boolean if a field has been set.

### GetShowTrackerInTorrent

`func (o *BaseSettingsModel) GetShowTrackerInTorrent() bool`

GetShowTrackerInTorrent returns the ShowTrackerInTorrent field if non-nil, zero value otherwise.

### GetShowTrackerInTorrentOk

`func (o *BaseSettingsModel) GetShowTrackerInTorrentOk() (*bool, bool)`

GetShowTrackerInTorrentOk returns a tuple with the ShowTrackerInTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTrackerInTorrent

`func (o *BaseSettingsModel) SetShowTrackerInTorrent(v bool)`

SetShowTrackerInTorrent sets ShowTrackerInTorrent field to given value.

### HasShowTrackerInTorrent

`func (o *BaseSettingsModel) HasShowTrackerInTorrent() bool`

HasShowTrackerInTorrent returns a boolean if a field has been set.

### GetStremioQuality

`func (o *BaseSettingsModel) GetStremioQuality() []int64`

GetStremioQuality returns the StremioQuality field if non-nil, zero value otherwise.

### GetStremioQualityOk

`func (o *BaseSettingsModel) GetStremioQualityOk() (*[]int64, bool)`

GetStremioQualityOk returns a tuple with the StremioQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioQuality

`func (o *BaseSettingsModel) SetStremioQuality(v []int64)`

SetStremioQuality sets StremioQuality field to given value.

### HasStremioQuality

`func (o *BaseSettingsModel) HasStremioQuality() bool`

HasStremioQuality returns a boolean if a field has been set.

### GetStremioResolution

`func (o *BaseSettingsModel) GetStremioResolution() []int64`

GetStremioResolution returns the StremioResolution field if non-nil, zero value otherwise.

### GetStremioResolutionOk

`func (o *BaseSettingsModel) GetStremioResolutionOk() (*[]int64, bool)`

GetStremioResolutionOk returns a tuple with the StremioResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioResolution

`func (o *BaseSettingsModel) SetStremioResolution(v []int64)`

SetStremioResolution sets StremioResolution field to given value.

### HasStremioResolution

`func (o *BaseSettingsModel) HasStremioResolution() bool`

HasStremioResolution returns a boolean if a field has been set.

### GetStremioLanguage

`func (o *BaseSettingsModel) GetStremioLanguage() []int64`

GetStremioLanguage returns the StremioLanguage field if non-nil, zero value otherwise.

### GetStremioLanguageOk

`func (o *BaseSettingsModel) GetStremioLanguageOk() (*[]int64, bool)`

GetStremioLanguageOk returns a tuple with the StremioLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLanguage

`func (o *BaseSettingsModel) SetStremioLanguage(v []int64)`

SetStremioLanguage sets StremioLanguage field to given value.

### HasStremioLanguage

`func (o *BaseSettingsModel) HasStremioLanguage() bool`

HasStremioLanguage returns a boolean if a field has been set.

### GetStremioCache

`func (o *BaseSettingsModel) GetStremioCache() []int64`

GetStremioCache returns the StremioCache field if non-nil, zero value otherwise.

### GetStremioCacheOk

`func (o *BaseSettingsModel) GetStremioCacheOk() (*[]int64, bool)`

GetStremioCacheOk returns a tuple with the StremioCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioCache

`func (o *BaseSettingsModel) SetStremioCache(v []int64)`

SetStremioCache sets StremioCache field to given value.

### HasStremioCache

`func (o *BaseSettingsModel) HasStremioCache() bool`

HasStremioCache returns a boolean if a field has been set.

### GetStremioSizeLower

`func (o *BaseSettingsModel) GetStremioSizeLower() int64`

GetStremioSizeLower returns the StremioSizeLower field if non-nil, zero value otherwise.

### GetStremioSizeLowerOk

`func (o *BaseSettingsModel) GetStremioSizeLowerOk() (*int64, bool)`

GetStremioSizeLowerOk returns a tuple with the StremioSizeLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSizeLower

`func (o *BaseSettingsModel) SetStremioSizeLower(v int64)`

SetStremioSizeLower sets StremioSizeLower field to given value.

### HasStremioSizeLower

`func (o *BaseSettingsModel) HasStremioSizeLower() bool`

HasStremioSizeLower returns a boolean if a field has been set.

### GetGoogleDriveFolderId

`func (o *BaseSettingsModel) GetGoogleDriveFolderId() string`

GetGoogleDriveFolderId returns the GoogleDriveFolderId field if non-nil, zero value otherwise.

### GetGoogleDriveFolderIdOk

`func (o *BaseSettingsModel) GetGoogleDriveFolderIdOk() (*string, bool)`

GetGoogleDriveFolderIdOk returns a tuple with the GoogleDriveFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleDriveFolderId

`func (o *BaseSettingsModel) SetGoogleDriveFolderId(v string)`

SetGoogleDriveFolderId sets GoogleDriveFolderId field to given value.

### HasGoogleDriveFolderId

`func (o *BaseSettingsModel) HasGoogleDriveFolderId() bool`

HasGoogleDriveFolderId returns a boolean if a field has been set.

### GetOnedriveSavePath

`func (o *BaseSettingsModel) GetOnedriveSavePath() string`

GetOnedriveSavePath returns the OnedriveSavePath field if non-nil, zero value otherwise.

### GetOnedriveSavePathOk

`func (o *BaseSettingsModel) GetOnedriveSavePathOk() (*string, bool)`

GetOnedriveSavePathOk returns a tuple with the OnedriveSavePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnedriveSavePath

`func (o *BaseSettingsModel) SetOnedriveSavePath(v string)`

SetOnedriveSavePath sets OnedriveSavePath field to given value.

### HasOnedriveSavePath

`func (o *BaseSettingsModel) HasOnedriveSavePath() bool`

HasOnedriveSavePath returns a boolean if a field has been set.

### GetDiscordId

`func (o *BaseSettingsModel) GetDiscordId() string`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *BaseSettingsModel) GetDiscordIdOk() (*string, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *BaseSettingsModel) SetDiscordId(v string)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *BaseSettingsModel) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### GetDiscordNotifications

`func (o *BaseSettingsModel) GetDiscordNotifications() bool`

GetDiscordNotifications returns the DiscordNotifications field if non-nil, zero value otherwise.

### GetDiscordNotificationsOk

`func (o *BaseSettingsModel) GetDiscordNotificationsOk() (*bool, bool)`

GetDiscordNotificationsOk returns a tuple with the DiscordNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordNotifications

`func (o *BaseSettingsModel) SetDiscordNotifications(v bool)`

SetDiscordNotifications sets DiscordNotifications field to given value.

### HasDiscordNotifications

`func (o *BaseSettingsModel) HasDiscordNotifications() bool`

HasDiscordNotifications returns a boolean if a field has been set.

### GetStremioAllowAdult

`func (o *BaseSettingsModel) GetStremioAllowAdult() bool`

GetStremioAllowAdult returns the StremioAllowAdult field if non-nil, zero value otherwise.

### GetStremioAllowAdultOk

`func (o *BaseSettingsModel) GetStremioAllowAdultOk() (*bool, bool)`

GetStremioAllowAdultOk returns a tuple with the StremioAllowAdult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioAllowAdult

`func (o *BaseSettingsModel) SetStremioAllowAdult(v bool)`

SetStremioAllowAdult sets StremioAllowAdult field to given value.

### HasStremioAllowAdult

`func (o *BaseSettingsModel) HasStremioAllowAdult() bool`

HasStremioAllowAdult returns a boolean if a field has been set.

### GetStremioSeedTorrents

`func (o *BaseSettingsModel) GetStremioSeedTorrents() int64`

GetStremioSeedTorrents returns the StremioSeedTorrents field if non-nil, zero value otherwise.

### GetStremioSeedTorrentsOk

`func (o *BaseSettingsModel) GetStremioSeedTorrentsOk() (*int64, bool)`

GetStremioSeedTorrentsOk returns a tuple with the StremioSeedTorrents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSeedTorrents

`func (o *BaseSettingsModel) SetStremioSeedTorrents(v int64)`

SetStremioSeedTorrents sets StremioSeedTorrents field to given value.

### HasStremioSeedTorrents

`func (o *BaseSettingsModel) HasStremioSeedTorrents() bool`

HasStremioSeedTorrents returns a boolean if a field has been set.

### GetSeedTorrents

`func (o *BaseSettingsModel) GetSeedTorrents() int64`

GetSeedTorrents returns the SeedTorrents field if non-nil, zero value otherwise.

### GetSeedTorrentsOk

`func (o *BaseSettingsModel) GetSeedTorrentsOk() (*int64, bool)`

GetSeedTorrentsOk returns a tuple with the SeedTorrents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedTorrents

`func (o *BaseSettingsModel) SetSeedTorrents(v int64)`

SetSeedTorrents sets SeedTorrents field to given value.

### HasSeedTorrents

`func (o *BaseSettingsModel) HasSeedTorrents() bool`

HasSeedTorrents returns a boolean if a field has been set.

### GetAllowZipped

`func (o *BaseSettingsModel) GetAllowZipped() bool`

GetAllowZipped returns the AllowZipped field if non-nil, zero value otherwise.

### GetAllowZippedOk

`func (o *BaseSettingsModel) GetAllowZippedOk() (*bool, bool)`

GetAllowZippedOk returns a tuple with the AllowZipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZipped

`func (o *BaseSettingsModel) SetAllowZipped(v bool)`

SetAllowZipped sets AllowZipped field to given value.

### HasAllowZipped

`func (o *BaseSettingsModel) HasAllowZipped() bool`

HasAllowZipped returns a boolean if a field has been set.

### GetStremioAllowZipped

`func (o *BaseSettingsModel) GetStremioAllowZipped() bool`

GetStremioAllowZipped returns the StremioAllowZipped field if non-nil, zero value otherwise.

### GetStremioAllowZippedOk

`func (o *BaseSettingsModel) GetStremioAllowZippedOk() (*bool, bool)`

GetStremioAllowZippedOk returns a tuple with the StremioAllowZipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioAllowZipped

`func (o *BaseSettingsModel) SetStremioAllowZipped(v bool)`

SetStremioAllowZipped sets StremioAllowZipped field to given value.

### HasStremioAllowZipped

`func (o *BaseSettingsModel) HasStremioAllowZipped() bool`

HasStremioAllowZipped returns a boolean if a field has been set.

### GetOnefichierFolderId

`func (o *BaseSettingsModel) GetOnefichierFolderId() string`

GetOnefichierFolderId returns the OnefichierFolderId field if non-nil, zero value otherwise.

### GetOnefichierFolderIdOk

`func (o *BaseSettingsModel) GetOnefichierFolderIdOk() (*string, bool)`

GetOnefichierFolderIdOk returns a tuple with the OnefichierFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnefichierFolderId

`func (o *BaseSettingsModel) SetOnefichierFolderId(v string)`

SetOnefichierFolderId sets OnefichierFolderId field to given value.

### HasOnefichierFolderId

`func (o *BaseSettingsModel) HasOnefichierFolderId() bool`

HasOnefichierFolderId returns a boolean if a field has been set.

### GetGofileFolderId

`func (o *BaseSettingsModel) GetGofileFolderId() string`

GetGofileFolderId returns the GofileFolderId field if non-nil, zero value otherwise.

### GetGofileFolderIdOk

`func (o *BaseSettingsModel) GetGofileFolderIdOk() (*string, bool)`

GetGofileFolderIdOk returns a tuple with the GofileFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGofileFolderId

`func (o *BaseSettingsModel) SetGofileFolderId(v string)`

SetGofileFolderId sets GofileFolderId field to given value.

### HasGofileFolderId

`func (o *BaseSettingsModel) HasGofileFolderId() bool`

HasGofileFolderId returns a boolean if a field has been set.

### GetWebhookNotifications

`func (o *BaseSettingsModel) GetWebhookNotifications() bool`

GetWebhookNotifications returns the WebhookNotifications field if non-nil, zero value otherwise.

### GetWebhookNotificationsOk

`func (o *BaseSettingsModel) GetWebhookNotificationsOk() (*bool, bool)`

GetWebhookNotificationsOk returns a tuple with the WebhookNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookNotifications

`func (o *BaseSettingsModel) SetWebhookNotifications(v bool)`

SetWebhookNotifications sets WebhookNotifications field to given value.

### HasWebhookNotifications

`func (o *BaseSettingsModel) HasWebhookNotifications() bool`

HasWebhookNotifications returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *BaseSettingsModel) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *BaseSettingsModel) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *BaseSettingsModel) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *BaseSettingsModel) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetTelegramNotifications

`func (o *BaseSettingsModel) GetTelegramNotifications() bool`

GetTelegramNotifications returns the TelegramNotifications field if non-nil, zero value otherwise.

### GetTelegramNotificationsOk

`func (o *BaseSettingsModel) GetTelegramNotificationsOk() (*bool, bool)`

GetTelegramNotificationsOk returns a tuple with the TelegramNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramNotifications

`func (o *BaseSettingsModel) SetTelegramNotifications(v bool)`

SetTelegramNotifications sets TelegramNotifications field to given value.

### HasTelegramNotifications

`func (o *BaseSettingsModel) HasTelegramNotifications() bool`

HasTelegramNotifications returns a boolean if a field has been set.

### GetTelegramId

`func (o *BaseSettingsModel) GetTelegramId() string`

GetTelegramId returns the TelegramId field if non-nil, zero value otherwise.

### GetTelegramIdOk

`func (o *BaseSettingsModel) GetTelegramIdOk() (*string, bool)`

GetTelegramIdOk returns a tuple with the TelegramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramId

`func (o *BaseSettingsModel) SetTelegramId(v string)`

SetTelegramId sets TelegramId field to given value.

### HasTelegramId

`func (o *BaseSettingsModel) HasTelegramId() bool`

HasTelegramId returns a boolean if a field has been set.

### GetMegaEmail

`func (o *BaseSettingsModel) GetMegaEmail() string`

GetMegaEmail returns the MegaEmail field if non-nil, zero value otherwise.

### GetMegaEmailOk

`func (o *BaseSettingsModel) GetMegaEmailOk() (*string, bool)`

GetMegaEmailOk returns a tuple with the MegaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMegaEmail

`func (o *BaseSettingsModel) SetMegaEmail(v string)`

SetMegaEmail sets MegaEmail field to given value.

### HasMegaEmail

`func (o *BaseSettingsModel) HasMegaEmail() bool`

HasMegaEmail returns a boolean if a field has been set.

### GetMegaPassword

`func (o *BaseSettingsModel) GetMegaPassword() string`

GetMegaPassword returns the MegaPassword field if non-nil, zero value otherwise.

### GetMegaPasswordOk

`func (o *BaseSettingsModel) GetMegaPasswordOk() (*string, bool)`

GetMegaPasswordOk returns a tuple with the MegaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMegaPassword

`func (o *BaseSettingsModel) SetMegaPassword(v string)`

SetMegaPassword sets MegaPassword field to given value.

### HasMegaPassword

`func (o *BaseSettingsModel) HasMegaPassword() bool`

HasMegaPassword returns a boolean if a field has been set.

### GetStremioSort

`func (o *BaseSettingsModel) GetStremioSort() string`

GetStremioSort returns the StremioSort field if non-nil, zero value otherwise.

### GetStremioSortOk

`func (o *BaseSettingsModel) GetStremioSortOk() (*string, bool)`

GetStremioSortOk returns a tuple with the StremioSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSort

`func (o *BaseSettingsModel) SetStremioSort(v string)`

SetStremioSort sets StremioSort field to given value.

### HasStremioSort

`func (o *BaseSettingsModel) HasStremioSort() bool`

HasStremioSort returns a boolean if a field has been set.

### GetCdnSelection

`func (o *BaseSettingsModel) GetCdnSelection() string`

GetCdnSelection returns the CdnSelection field if non-nil, zero value otherwise.

### GetCdnSelectionOk

`func (o *BaseSettingsModel) GetCdnSelectionOk() (*string, bool)`

GetCdnSelectionOk returns a tuple with the CdnSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnSelection

`func (o *BaseSettingsModel) SetCdnSelection(v string)`

SetCdnSelection sets CdnSelection field to given value.

### HasCdnSelection

`func (o *BaseSettingsModel) HasCdnSelection() bool`

HasCdnSelection returns a boolean if a field has been set.

### GetStremioUseCustomSearchEngines

`func (o *BaseSettingsModel) GetStremioUseCustomSearchEngines() bool`

GetStremioUseCustomSearchEngines returns the StremioUseCustomSearchEngines field if non-nil, zero value otherwise.

### GetStremioUseCustomSearchEnginesOk

`func (o *BaseSettingsModel) GetStremioUseCustomSearchEnginesOk() (*bool, bool)`

GetStremioUseCustomSearchEnginesOk returns a tuple with the StremioUseCustomSearchEngines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioUseCustomSearchEngines

`func (o *BaseSettingsModel) SetStremioUseCustomSearchEngines(v bool)`

SetStremioUseCustomSearchEngines sets StremioUseCustomSearchEngines field to given value.

### HasStremioUseCustomSearchEngines

`func (o *BaseSettingsModel) HasStremioUseCustomSearchEngines() bool`

HasStremioUseCustomSearchEngines returns a boolean if a field has been set.

### GetStremioResultSort

`func (o *BaseSettingsModel) GetStremioResultSort() string`

GetStremioResultSort returns the StremioResultSort field if non-nil, zero value otherwise.

### GetStremioResultSortOk

`func (o *BaseSettingsModel) GetStremioResultSortOk() (*string, bool)`

GetStremioResultSortOk returns a tuple with the StremioResultSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioResultSort

`func (o *BaseSettingsModel) SetStremioResultSort(v string)`

SetStremioResultSort sets StremioResultSort field to given value.

### HasStremioResultSort

`func (o *BaseSettingsModel) HasStremioResultSort() bool`

HasStremioResultSort returns a boolean if a field has been set.

### GetWebdavUseLocalFiles

`func (o *BaseSettingsModel) GetWebdavUseLocalFiles() bool`

GetWebdavUseLocalFiles returns the WebdavUseLocalFiles field if non-nil, zero value otherwise.

### GetWebdavUseLocalFilesOk

`func (o *BaseSettingsModel) GetWebdavUseLocalFilesOk() (*bool, bool)`

GetWebdavUseLocalFilesOk returns a tuple with the WebdavUseLocalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavUseLocalFiles

`func (o *BaseSettingsModel) SetWebdavUseLocalFiles(v bool)`

SetWebdavUseLocalFiles sets WebdavUseLocalFiles field to given value.

### HasWebdavUseLocalFiles

`func (o *BaseSettingsModel) HasWebdavUseLocalFiles() bool`

HasWebdavUseLocalFiles returns a boolean if a field has been set.

### GetStremioLegacyYourMedia

`func (o *BaseSettingsModel) GetStremioLegacyYourMedia() bool`

GetStremioLegacyYourMedia returns the StremioLegacyYourMedia field if non-nil, zero value otherwise.

### GetStremioLegacyYourMediaOk

`func (o *BaseSettingsModel) GetStremioLegacyYourMediaOk() (*bool, bool)`

GetStremioLegacyYourMediaOk returns a tuple with the StremioLegacyYourMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLegacyYourMedia

`func (o *BaseSettingsModel) SetStremioLegacyYourMedia(v bool)`

SetStremioLegacyYourMedia sets StremioLegacyYourMedia field to given value.

### HasStremioLegacyYourMedia

`func (o *BaseSettingsModel) HasStremioLegacyYourMedia() bool`

HasStremioLegacyYourMedia returns a boolean if a field has been set.

### GetStremioOnlyYourMediaStreams

`func (o *BaseSettingsModel) GetStremioOnlyYourMediaStreams() bool`

GetStremioOnlyYourMediaStreams returns the StremioOnlyYourMediaStreams field if non-nil, zero value otherwise.

### GetStremioOnlyYourMediaStreamsOk

`func (o *BaseSettingsModel) GetStremioOnlyYourMediaStreamsOk() (*bool, bool)`

GetStremioOnlyYourMediaStreamsOk returns a tuple with the StremioOnlyYourMediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioOnlyYourMediaStreams

`func (o *BaseSettingsModel) SetStremioOnlyYourMediaStreams(v bool)`

SetStremioOnlyYourMediaStreams sets StremioOnlyYourMediaStreams field to given value.

### HasStremioOnlyYourMediaStreams

`func (o *BaseSettingsModel) HasStremioOnlyYourMediaStreams() bool`

HasStremioOnlyYourMediaStreams returns a boolean if a field has been set.

### GetStremioDisableYourMediaStreams

`func (o *BaseSettingsModel) GetStremioDisableYourMediaStreams() bool`

GetStremioDisableYourMediaStreams returns the StremioDisableYourMediaStreams field if non-nil, zero value otherwise.

### GetStremioDisableYourMediaStreamsOk

`func (o *BaseSettingsModel) GetStremioDisableYourMediaStreamsOk() (*bool, bool)`

GetStremioDisableYourMediaStreamsOk returns a tuple with the StremioDisableYourMediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioDisableYourMediaStreams

`func (o *BaseSettingsModel) SetStremioDisableYourMediaStreams(v bool)`

SetStremioDisableYourMediaStreams sets StremioDisableYourMediaStreams field to given value.

### HasStremioDisableYourMediaStreams

`func (o *BaseSettingsModel) HasStremioDisableYourMediaStreams() bool`

HasStremioDisableYourMediaStreams returns a boolean if a field has been set.

### GetWebdavUseFolderView

`func (o *BaseSettingsModel) GetWebdavUseFolderView() bool`

GetWebdavUseFolderView returns the WebdavUseFolderView field if non-nil, zero value otherwise.

### GetWebdavUseFolderViewOk

`func (o *BaseSettingsModel) GetWebdavUseFolderViewOk() (*bool, bool)`

GetWebdavUseFolderViewOk returns a tuple with the WebdavUseFolderView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavUseFolderView

`func (o *BaseSettingsModel) SetWebdavUseFolderView(v bool)`

SetWebdavUseFolderView sets WebdavUseFolderView field to given value.

### HasWebdavUseFolderView

`func (o *BaseSettingsModel) HasWebdavUseFolderView() bool`

HasWebdavUseFolderView returns a boolean if a field has been set.

### GetWebdavFlatten

`func (o *BaseSettingsModel) GetWebdavFlatten() bool`

GetWebdavFlatten returns the WebdavFlatten field if non-nil, zero value otherwise.

### GetWebdavFlattenOk

`func (o *BaseSettingsModel) GetWebdavFlattenOk() (*bool, bool)`

GetWebdavFlattenOk returns a tuple with the WebdavFlatten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavFlatten

`func (o *BaseSettingsModel) SetWebdavFlatten(v bool)`

SetWebdavFlatten sets WebdavFlatten field to given value.

### HasWebdavFlatten

`func (o *BaseSettingsModel) HasWebdavFlatten() bool`

HasWebdavFlatten returns a boolean if a field has been set.

### GetStremioLimitPerResolutionTorrent

`func (o *BaseSettingsModel) GetStremioLimitPerResolutionTorrent() int64`

GetStremioLimitPerResolutionTorrent returns the StremioLimitPerResolutionTorrent field if non-nil, zero value otherwise.

### GetStremioLimitPerResolutionTorrentOk

`func (o *BaseSettingsModel) GetStremioLimitPerResolutionTorrentOk() (*int64, bool)`

GetStremioLimitPerResolutionTorrentOk returns a tuple with the StremioLimitPerResolutionTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLimitPerResolutionTorrent

`func (o *BaseSettingsModel) SetStremioLimitPerResolutionTorrent(v int64)`

SetStremioLimitPerResolutionTorrent sets StremioLimitPerResolutionTorrent field to given value.

### HasStremioLimitPerResolutionTorrent

`func (o *BaseSettingsModel) HasStremioLimitPerResolutionTorrent() bool`

HasStremioLimitPerResolutionTorrent returns a boolean if a field has been set.

### GetStremioLimitPerResolutionUsenet

`func (o *BaseSettingsModel) GetStremioLimitPerResolutionUsenet() int64`

GetStremioLimitPerResolutionUsenet returns the StremioLimitPerResolutionUsenet field if non-nil, zero value otherwise.

### GetStremioLimitPerResolutionUsenetOk

`func (o *BaseSettingsModel) GetStremioLimitPerResolutionUsenetOk() (*int64, bool)`

GetStremioLimitPerResolutionUsenetOk returns a tuple with the StremioLimitPerResolutionUsenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioLimitPerResolutionUsenet

`func (o *BaseSettingsModel) SetStremioLimitPerResolutionUsenet(v int64)`

SetStremioLimitPerResolutionUsenet sets StremioLimitPerResolutionUsenet field to given value.

### HasStremioLimitPerResolutionUsenet

`func (o *BaseSettingsModel) HasStremioLimitPerResolutionUsenet() bool`

HasStremioLimitPerResolutionUsenet returns a boolean if a field has been set.

### GetStremioSizeUpper

`func (o *BaseSettingsModel) GetStremioSizeUpper() int64`

GetStremioSizeUpper returns the StremioSizeUpper field if non-nil, zero value otherwise.

### GetStremioSizeUpperOk

`func (o *BaseSettingsModel) GetStremioSizeUpperOk() (*int64, bool)`

GetStremioSizeUpperOk returns a tuple with the StremioSizeUpper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioSizeUpper

`func (o *BaseSettingsModel) SetStremioSizeUpper(v int64)`

SetStremioSizeUpper sets StremioSizeUpper field to given value.

### HasStremioSizeUpper

`func (o *BaseSettingsModel) HasStremioSizeUpper() bool`

HasStremioSizeUpper returns a boolean if a field has been set.

### GetStremioTorrentSeedersCutoff

`func (o *BaseSettingsModel) GetStremioTorrentSeedersCutoff() int64`

GetStremioTorrentSeedersCutoff returns the StremioTorrentSeedersCutoff field if non-nil, zero value otherwise.

### GetStremioTorrentSeedersCutoffOk

`func (o *BaseSettingsModel) GetStremioTorrentSeedersCutoffOk() (*int64, bool)`

GetStremioTorrentSeedersCutoffOk returns a tuple with the StremioTorrentSeedersCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioTorrentSeedersCutoff

`func (o *BaseSettingsModel) SetStremioTorrentSeedersCutoff(v int64)`

SetStremioTorrentSeedersCutoff sets StremioTorrentSeedersCutoff field to given value.

### HasStremioTorrentSeedersCutoff

`func (o *BaseSettingsModel) HasStremioTorrentSeedersCutoff() bool`

HasStremioTorrentSeedersCutoff returns a boolean if a field has been set.

### GetPixeldrainApiKey

`func (o *BaseSettingsModel) GetPixeldrainApiKey() string`

GetPixeldrainApiKey returns the PixeldrainApiKey field if non-nil, zero value otherwise.

### GetPixeldrainApiKeyOk

`func (o *BaseSettingsModel) GetPixeldrainApiKeyOk() (*string, bool)`

GetPixeldrainApiKeyOk returns a tuple with the PixeldrainApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixeldrainApiKey

`func (o *BaseSettingsModel) SetPixeldrainApiKey(v string)`

SetPixeldrainApiKey sets PixeldrainApiKey field to given value.

### HasPixeldrainApiKey

`func (o *BaseSettingsModel) HasPixeldrainApiKey() bool`

HasPixeldrainApiKey returns a boolean if a field has been set.

### GetOnefichierApiKey

`func (o *BaseSettingsModel) GetOnefichierApiKey() string`

GetOnefichierApiKey returns the OnefichierApiKey field if non-nil, zero value otherwise.

### GetOnefichierApiKeyOk

`func (o *BaseSettingsModel) GetOnefichierApiKeyOk() (*string, bool)`

GetOnefichierApiKeyOk returns a tuple with the OnefichierApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnefichierApiKey

`func (o *BaseSettingsModel) SetOnefichierApiKey(v string)`

SetOnefichierApiKey sets OnefichierApiKey field to given value.

### HasOnefichierApiKey

`func (o *BaseSettingsModel) HasOnefichierApiKey() bool`

HasOnefichierApiKey returns a boolean if a field has been set.

### GetGofileApiKey

`func (o *BaseSettingsModel) GetGofileApiKey() string`

GetGofileApiKey returns the GofileApiKey field if non-nil, zero value otherwise.

### GetGofileApiKeyOk

`func (o *BaseSettingsModel) GetGofileApiKeyOk() (*string, bool)`

GetGofileApiKeyOk returns a tuple with the GofileApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGofileApiKey

`func (o *BaseSettingsModel) SetGofileApiKey(v string)`

SetGofileApiKey sets GofileApiKey field to given value.

### HasGofileApiKey

`func (o *BaseSettingsModel) HasGofileApiKey() bool`

HasGofileApiKey returns a boolean if a field has been set.

### GetPatreonId

`func (o *BaseSettingsModel) GetPatreonId() string`

GetPatreonId returns the PatreonId field if non-nil, zero value otherwise.

### GetPatreonIdOk

`func (o *BaseSettingsModel) GetPatreonIdOk() (*string, bool)`

GetPatreonIdOk returns a tuple with the PatreonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatreonId

`func (o *BaseSettingsModel) SetPatreonId(v string)`

SetPatreonId sets PatreonId field to given value.

### HasPatreonId

`func (o *BaseSettingsModel) HasPatreonId() bool`

HasPatreonId returns a boolean if a field has been set.

### GetStremioWaitForDownloadUsenet

`func (o *BaseSettingsModel) GetStremioWaitForDownloadUsenet() bool`

GetStremioWaitForDownloadUsenet returns the StremioWaitForDownloadUsenet field if non-nil, zero value otherwise.

### GetStremioWaitForDownloadUsenetOk

`func (o *BaseSettingsModel) GetStremioWaitForDownloadUsenetOk() (*bool, bool)`

GetStremioWaitForDownloadUsenetOk returns a tuple with the StremioWaitForDownloadUsenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioWaitForDownloadUsenet

`func (o *BaseSettingsModel) SetStremioWaitForDownloadUsenet(v bool)`

SetStremioWaitForDownloadUsenet sets StremioWaitForDownloadUsenet field to given value.

### HasStremioWaitForDownloadUsenet

`func (o *BaseSettingsModel) HasStremioWaitForDownloadUsenet() bool`

HasStremioWaitForDownloadUsenet returns a boolean if a field has been set.

### GetStremioWaitForDownloadTorrent

`func (o *BaseSettingsModel) GetStremioWaitForDownloadTorrent() bool`

GetStremioWaitForDownloadTorrent returns the StremioWaitForDownloadTorrent field if non-nil, zero value otherwise.

### GetStremioWaitForDownloadTorrentOk

`func (o *BaseSettingsModel) GetStremioWaitForDownloadTorrentOk() (*bool, bool)`

GetStremioWaitForDownloadTorrentOk returns a tuple with the StremioWaitForDownloadTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioWaitForDownloadTorrent

`func (o *BaseSettingsModel) SetStremioWaitForDownloadTorrent(v bool)`

SetStremioWaitForDownloadTorrent sets StremioWaitForDownloadTorrent field to given value.

### HasStremioWaitForDownloadTorrent

`func (o *BaseSettingsModel) HasStremioWaitForDownloadTorrent() bool`

HasStremioWaitForDownloadTorrent returns a boolean if a field has been set.

### GetStremioDisableFilteredNote

`func (o *BaseSettingsModel) GetStremioDisableFilteredNote() bool`

GetStremioDisableFilteredNote returns the StremioDisableFilteredNote field if non-nil, zero value otherwise.

### GetStremioDisableFilteredNoteOk

`func (o *BaseSettingsModel) GetStremioDisableFilteredNoteOk() (*bool, bool)`

GetStremioDisableFilteredNoteOk returns a tuple with the StremioDisableFilteredNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioDisableFilteredNote

`func (o *BaseSettingsModel) SetStremioDisableFilteredNote(v bool)`

SetStremioDisableFilteredNote sets StremioDisableFilteredNote field to given value.

### HasStremioDisableFilteredNote

`func (o *BaseSettingsModel) HasStremioDisableFilteredNote() bool`

HasStremioDisableFilteredNote returns a boolean if a field has been set.

### GetStremioEmojiInDescription

`func (o *BaseSettingsModel) GetStremioEmojiInDescription() bool`

GetStremioEmojiInDescription returns the StremioEmojiInDescription field if non-nil, zero value otherwise.

### GetStremioEmojiInDescriptionOk

`func (o *BaseSettingsModel) GetStremioEmojiInDescriptionOk() (*bool, bool)`

GetStremioEmojiInDescriptionOk returns a tuple with the StremioEmojiInDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStremioEmojiInDescription

`func (o *BaseSettingsModel) SetStremioEmojiInDescription(v bool)`

SetStremioEmojiInDescription sets StremioEmojiInDescription field to given value.

### HasStremioEmojiInDescription

`func (o *BaseSettingsModel) HasStremioEmojiInDescription() bool`

HasStremioEmojiInDescription returns a boolean if a field has been set.

### GetDashboardFilter

`func (o *BaseSettingsModel) GetDashboardFilter() map[string]interface{}`

GetDashboardFilter returns the DashboardFilter field if non-nil, zero value otherwise.

### GetDashboardFilterOk

`func (o *BaseSettingsModel) GetDashboardFilterOk() (*map[string]interface{}, bool)`

GetDashboardFilterOk returns a tuple with the DashboardFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilter

`func (o *BaseSettingsModel) SetDashboardFilter(v map[string]interface{})`

SetDashboardFilter sets DashboardFilter field to given value.

### HasDashboardFilter

`func (o *BaseSettingsModel) HasDashboardFilter() bool`

HasDashboardFilter returns a boolean if a field has been set.

### GetWebPlayerAlwaysTranscode

`func (o *BaseSettingsModel) GetWebPlayerAlwaysTranscode() bool`

GetWebPlayerAlwaysTranscode returns the WebPlayerAlwaysTranscode field if non-nil, zero value otherwise.

### GetWebPlayerAlwaysTranscodeOk

`func (o *BaseSettingsModel) GetWebPlayerAlwaysTranscodeOk() (*bool, bool)`

GetWebPlayerAlwaysTranscodeOk returns a tuple with the WebPlayerAlwaysTranscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerAlwaysTranscode

`func (o *BaseSettingsModel) SetWebPlayerAlwaysTranscode(v bool)`

SetWebPlayerAlwaysTranscode sets WebPlayerAlwaysTranscode field to given value.

### HasWebPlayerAlwaysTranscode

`func (o *BaseSettingsModel) HasWebPlayerAlwaysTranscode() bool`

HasWebPlayerAlwaysTranscode returns a boolean if a field has been set.

### GetWebPlayerAlwaysSkipIntro

`func (o *BaseSettingsModel) GetWebPlayerAlwaysSkipIntro() bool`

GetWebPlayerAlwaysSkipIntro returns the WebPlayerAlwaysSkipIntro field if non-nil, zero value otherwise.

### GetWebPlayerAlwaysSkipIntroOk

`func (o *BaseSettingsModel) GetWebPlayerAlwaysSkipIntroOk() (*bool, bool)`

GetWebPlayerAlwaysSkipIntroOk returns a tuple with the WebPlayerAlwaysSkipIntro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerAlwaysSkipIntro

`func (o *BaseSettingsModel) SetWebPlayerAlwaysSkipIntro(v bool)`

SetWebPlayerAlwaysSkipIntro sets WebPlayerAlwaysSkipIntro field to given value.

### HasWebPlayerAlwaysSkipIntro

`func (o *BaseSettingsModel) HasWebPlayerAlwaysSkipIntro() bool`

HasWebPlayerAlwaysSkipIntro returns a boolean if a field has been set.

### GetWebPlayerAudioPreferredLanguage

`func (o *BaseSettingsModel) GetWebPlayerAudioPreferredLanguage() string`

GetWebPlayerAudioPreferredLanguage returns the WebPlayerAudioPreferredLanguage field if non-nil, zero value otherwise.

### GetWebPlayerAudioPreferredLanguageOk

`func (o *BaseSettingsModel) GetWebPlayerAudioPreferredLanguageOk() (*string, bool)`

GetWebPlayerAudioPreferredLanguageOk returns a tuple with the WebPlayerAudioPreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerAudioPreferredLanguage

`func (o *BaseSettingsModel) SetWebPlayerAudioPreferredLanguage(v string)`

SetWebPlayerAudioPreferredLanguage sets WebPlayerAudioPreferredLanguage field to given value.

### HasWebPlayerAudioPreferredLanguage

`func (o *BaseSettingsModel) HasWebPlayerAudioPreferredLanguage() bool`

HasWebPlayerAudioPreferredLanguage returns a boolean if a field has been set.

### GetWebPlayerSubtitlePreferredLanguage

`func (o *BaseSettingsModel) GetWebPlayerSubtitlePreferredLanguage() string`

GetWebPlayerSubtitlePreferredLanguage returns the WebPlayerSubtitlePreferredLanguage field if non-nil, zero value otherwise.

### GetWebPlayerSubtitlePreferredLanguageOk

`func (o *BaseSettingsModel) GetWebPlayerSubtitlePreferredLanguageOk() (*string, bool)`

GetWebPlayerSubtitlePreferredLanguageOk returns a tuple with the WebPlayerSubtitlePreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerSubtitlePreferredLanguage

`func (o *BaseSettingsModel) SetWebPlayerSubtitlePreferredLanguage(v string)`

SetWebPlayerSubtitlePreferredLanguage sets WebPlayerSubtitlePreferredLanguage field to given value.

### HasWebPlayerSubtitlePreferredLanguage

`func (o *BaseSettingsModel) HasWebPlayerSubtitlePreferredLanguage() bool`

HasWebPlayerSubtitlePreferredLanguage returns a boolean if a field has been set.

### GetWebPlayerDisablePrestreamSelector

`func (o *BaseSettingsModel) GetWebPlayerDisablePrestreamSelector() bool`

GetWebPlayerDisablePrestreamSelector returns the WebPlayerDisablePrestreamSelector field if non-nil, zero value otherwise.

### GetWebPlayerDisablePrestreamSelectorOk

`func (o *BaseSettingsModel) GetWebPlayerDisablePrestreamSelectorOk() (*bool, bool)`

GetWebPlayerDisablePrestreamSelectorOk returns a tuple with the WebPlayerDisablePrestreamSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerDisablePrestreamSelector

`func (o *BaseSettingsModel) SetWebPlayerDisablePrestreamSelector(v bool)`

SetWebPlayerDisablePrestreamSelector sets WebPlayerDisablePrestreamSelector field to given value.

### HasWebPlayerDisablePrestreamSelector

`func (o *BaseSettingsModel) HasWebPlayerDisablePrestreamSelector() bool`

HasWebPlayerDisablePrestreamSelector returns a boolean if a field has been set.

### GetDashboardSort

`func (o *BaseSettingsModel) GetDashboardSort() string`

GetDashboardSort returns the DashboardSort field if non-nil, zero value otherwise.

### GetDashboardSortOk

`func (o *BaseSettingsModel) GetDashboardSortOk() (*string, bool)`

GetDashboardSortOk returns a tuple with the DashboardSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSort

`func (o *BaseSettingsModel) SetDashboardSort(v string)`

SetDashboardSort sets DashboardSort field to given value.

### HasDashboardSort

`func (o *BaseSettingsModel) HasDashboardSort() bool`

HasDashboardSort returns a boolean if a field has been set.

### GetWebPlayerDisableNextUpDialogue

`func (o *BaseSettingsModel) GetWebPlayerDisableNextUpDialogue() bool`

GetWebPlayerDisableNextUpDialogue returns the WebPlayerDisableNextUpDialogue field if non-nil, zero value otherwise.

### GetWebPlayerDisableNextUpDialogueOk

`func (o *BaseSettingsModel) GetWebPlayerDisableNextUpDialogueOk() (*bool, bool)`

GetWebPlayerDisableNextUpDialogueOk returns a tuple with the WebPlayerDisableNextUpDialogue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerDisableNextUpDialogue

`func (o *BaseSettingsModel) SetWebPlayerDisableNextUpDialogue(v bool)`

SetWebPlayerDisableNextUpDialogue sets WebPlayerDisableNextUpDialogue field to given value.

### HasWebPlayerDisableNextUpDialogue

`func (o *BaseSettingsModel) HasWebPlayerDisableNextUpDialogue() bool`

HasWebPlayerDisableNextUpDialogue returns a boolean if a field has been set.

### GetWebPlayerEnableScrobbling

`func (o *BaseSettingsModel) GetWebPlayerEnableScrobbling() bool`

GetWebPlayerEnableScrobbling returns the WebPlayerEnableScrobbling field if non-nil, zero value otherwise.

### GetWebPlayerEnableScrobblingOk

`func (o *BaseSettingsModel) GetWebPlayerEnableScrobblingOk() (*bool, bool)`

GetWebPlayerEnableScrobblingOk returns a tuple with the WebPlayerEnableScrobbling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPlayerEnableScrobbling

`func (o *BaseSettingsModel) SetWebPlayerEnableScrobbling(v bool)`

SetWebPlayerEnableScrobbling sets WebPlayerEnableScrobbling field to given value.

### HasWebPlayerEnableScrobbling

`func (o *BaseSettingsModel) HasWebPlayerEnableScrobbling() bool`

HasWebPlayerEnableScrobbling returns a boolean if a field has been set.

### GetAppendFilenameToLinks

`func (o *BaseSettingsModel) GetAppendFilenameToLinks() bool`

GetAppendFilenameToLinks returns the AppendFilenameToLinks field if non-nil, zero value otherwise.

### GetAppendFilenameToLinksOk

`func (o *BaseSettingsModel) GetAppendFilenameToLinksOk() (*bool, bool)`

GetAppendFilenameToLinksOk returns a tuple with the AppendFilenameToLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendFilenameToLinks

`func (o *BaseSettingsModel) SetAppendFilenameToLinks(v bool)`

SetAppendFilenameToLinks sets AppendFilenameToLinks field to given value.

### HasAppendFilenameToLinks

`func (o *BaseSettingsModel) HasAppendFilenameToLinks() bool`

HasAppendFilenameToLinks returns a boolean if a field has been set.

### GetSubscribedEmailSegments

`func (o *BaseSettingsModel) GetSubscribedEmailSegments() []string`

GetSubscribedEmailSegments returns the SubscribedEmailSegments field if non-nil, zero value otherwise.

### GetSubscribedEmailSegmentsOk

`func (o *BaseSettingsModel) GetSubscribedEmailSegmentsOk() (*[]string, bool)`

GetSubscribedEmailSegmentsOk returns a tuple with the SubscribedEmailSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEmailSegments

`func (o *BaseSettingsModel) SetSubscribedEmailSegments(v []string)`

SetSubscribedEmailSegments sets SubscribedEmailSegments field to given value.

### HasSubscribedEmailSegments

`func (o *BaseSettingsModel) HasSubscribedEmailSegments() bool`

HasSubscribedEmailSegments returns a boolean if a field has been set.

### GetDisableAi

`func (o *BaseSettingsModel) GetDisableAi() bool`

GetDisableAi returns the DisableAi field if non-nil, zero value otherwise.

### GetDisableAiOk

`func (o *BaseSettingsModel) GetDisableAiOk() (*bool, bool)`

GetDisableAiOk returns a tuple with the DisableAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAi

`func (o *BaseSettingsModel) SetDisableAi(v bool)`

SetDisableAi sets DisableAi field to given value.

### HasDisableAi

`func (o *BaseSettingsModel) HasDisableAi() bool`

HasDisableAi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


