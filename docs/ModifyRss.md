# ModifyRss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RssFeedId** | **int64** | The ID of the RSS feed you want to modify. | 
**Name** | Pointer to **string** | Display name of the rss feed you want to add. | [optional] 
**DoRegex** | Pointer to **string** | Regex of items you want to make sure are added. Optional. | [optional] 
**DontRegex** | Pointer to **string** | Regex of items you dont want added. Optional. | [optional] 
**ScanInterval** | Pointer to **int64** | Time in minutes with the minimum being 10 which you want the api to scan in interval. Optional. Default is 60. | [optional] [default to 60]
**DontOlderThan** | Pointer to **int64** | Time in days that tells the parser to skip if the torrent is older than that many days. Optional. | [optional] 
**RssType** | Pointer to **string** | The type of downloads inside of this RSS feed. Optional. Default is torrent. Options are \&quot;torrent\&quot;, \&quot;usenet\&quot;, \&quot;webdl\&quot; | [optional] [default to "torrent"]
**TorrentSeeding** | Pointer to **int64** | Determines if torrents inside of this RSS feed will seed or not. Optional. Default is 1, for auto. Options are 1 for auto, 2 for always, and 3 for never. | [optional] [default to 1]

## Methods

### NewModifyRss

`func NewModifyRss(rssFeedId int64, ) *ModifyRss`

NewModifyRss instantiates a new ModifyRss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyRssWithDefaults

`func NewModifyRssWithDefaults() *ModifyRss`

NewModifyRssWithDefaults instantiates a new ModifyRss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRssFeedId

`func (o *ModifyRss) GetRssFeedId() int64`

GetRssFeedId returns the RssFeedId field if non-nil, zero value otherwise.

### GetRssFeedIdOk

`func (o *ModifyRss) GetRssFeedIdOk() (*int64, bool)`

GetRssFeedIdOk returns a tuple with the RssFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssFeedId

`func (o *ModifyRss) SetRssFeedId(v int64)`

SetRssFeedId sets RssFeedId field to given value.


### GetName

`func (o *ModifyRss) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyRss) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyRss) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyRss) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDoRegex

`func (o *ModifyRss) GetDoRegex() string`

GetDoRegex returns the DoRegex field if non-nil, zero value otherwise.

### GetDoRegexOk

`func (o *ModifyRss) GetDoRegexOk() (*string, bool)`

GetDoRegexOk returns a tuple with the DoRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoRegex

`func (o *ModifyRss) SetDoRegex(v string)`

SetDoRegex sets DoRegex field to given value.

### HasDoRegex

`func (o *ModifyRss) HasDoRegex() bool`

HasDoRegex returns a boolean if a field has been set.

### GetDontRegex

`func (o *ModifyRss) GetDontRegex() string`

GetDontRegex returns the DontRegex field if non-nil, zero value otherwise.

### GetDontRegexOk

`func (o *ModifyRss) GetDontRegexOk() (*string, bool)`

GetDontRegexOk returns a tuple with the DontRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontRegex

`func (o *ModifyRss) SetDontRegex(v string)`

SetDontRegex sets DontRegex field to given value.

### HasDontRegex

`func (o *ModifyRss) HasDontRegex() bool`

HasDontRegex returns a boolean if a field has been set.

### GetScanInterval

`func (o *ModifyRss) GetScanInterval() int64`

GetScanInterval returns the ScanInterval field if non-nil, zero value otherwise.

### GetScanIntervalOk

`func (o *ModifyRss) GetScanIntervalOk() (*int64, bool)`

GetScanIntervalOk returns a tuple with the ScanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInterval

`func (o *ModifyRss) SetScanInterval(v int64)`

SetScanInterval sets ScanInterval field to given value.

### HasScanInterval

`func (o *ModifyRss) HasScanInterval() bool`

HasScanInterval returns a boolean if a field has been set.

### GetDontOlderThan

`func (o *ModifyRss) GetDontOlderThan() int64`

GetDontOlderThan returns the DontOlderThan field if non-nil, zero value otherwise.

### GetDontOlderThanOk

`func (o *ModifyRss) GetDontOlderThanOk() (*int64, bool)`

GetDontOlderThanOk returns a tuple with the DontOlderThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontOlderThan

`func (o *ModifyRss) SetDontOlderThan(v int64)`

SetDontOlderThan sets DontOlderThan field to given value.

### HasDontOlderThan

`func (o *ModifyRss) HasDontOlderThan() bool`

HasDontOlderThan returns a boolean if a field has been set.

### GetRssType

`func (o *ModifyRss) GetRssType() string`

GetRssType returns the RssType field if non-nil, zero value otherwise.

### GetRssTypeOk

`func (o *ModifyRss) GetRssTypeOk() (*string, bool)`

GetRssTypeOk returns a tuple with the RssType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssType

`func (o *ModifyRss) SetRssType(v string)`

SetRssType sets RssType field to given value.

### HasRssType

`func (o *ModifyRss) HasRssType() bool`

HasRssType returns a boolean if a field has been set.

### GetTorrentSeeding

`func (o *ModifyRss) GetTorrentSeeding() int64`

GetTorrentSeeding returns the TorrentSeeding field if non-nil, zero value otherwise.

### GetTorrentSeedingOk

`func (o *ModifyRss) GetTorrentSeedingOk() (*int64, bool)`

GetTorrentSeedingOk returns a tuple with the TorrentSeeding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentSeeding

`func (o *ModifyRss) SetTorrentSeeding(v int64)`

SetTorrentSeeding sets TorrentSeeding field to given value.

### HasTorrentSeeding

`func (o *ModifyRss) HasTorrentSeeding() bool`

HasTorrentSeeding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


