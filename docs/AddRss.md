# AddRss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the rss feed you want to add. | 
**Name** | **string** | Display name of the rss feed you want to add. | 
**DoRegex** | Pointer to **string** | Regex of items you want to make sure are added. Optional. | [optional] 
**DontRegex** | Pointer to **string** | Regex of items you dont want added. Optional. | [optional] 
**DontOlderThan** | Pointer to **int64** | Time in days that tells the parser to skip if the torrent is older than that many days. Optional. | [optional] 
**PassCheck** | Pointer to **bool** | Boolean which tells to allow multiple of the same url if it already exists on your account | [optional] [default to false]
**ScanInterval** | Pointer to **int64** | Time in minutes with the minimum being 10 which you want the api to scan in interval. Optional. Default is 60. | [optional] [default to 60]
**RssType** | Pointer to **string** | The type of downloads inside of this RSS feed. Optional. Default is torrent. Options are \&quot;torrent\&quot;, \&quot;usenet\&quot;, \&quot;webdl\&quot; | [optional] [default to "torrent"]
**TorrentSeeding** | Pointer to **int64** | Determines if torrents inside of this RSS feed will seed or not. Optional. Default is 1, for auto. Options are 1 for auto, 2 for always, and 3 for never. | [optional] [default to 1]

## Methods

### NewAddRss

`func NewAddRss(url string, name string, ) *AddRss`

NewAddRss instantiates a new AddRss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRssWithDefaults

`func NewAddRssWithDefaults() *AddRss`

NewAddRssWithDefaults instantiates a new AddRss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AddRss) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddRss) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddRss) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *AddRss) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddRss) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddRss) SetName(v string)`

SetName sets Name field to given value.


### GetDoRegex

`func (o *AddRss) GetDoRegex() string`

GetDoRegex returns the DoRegex field if non-nil, zero value otherwise.

### GetDoRegexOk

`func (o *AddRss) GetDoRegexOk() (*string, bool)`

GetDoRegexOk returns a tuple with the DoRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoRegex

`func (o *AddRss) SetDoRegex(v string)`

SetDoRegex sets DoRegex field to given value.

### HasDoRegex

`func (o *AddRss) HasDoRegex() bool`

HasDoRegex returns a boolean if a field has been set.

### GetDontRegex

`func (o *AddRss) GetDontRegex() string`

GetDontRegex returns the DontRegex field if non-nil, zero value otherwise.

### GetDontRegexOk

`func (o *AddRss) GetDontRegexOk() (*string, bool)`

GetDontRegexOk returns a tuple with the DontRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontRegex

`func (o *AddRss) SetDontRegex(v string)`

SetDontRegex sets DontRegex field to given value.

### HasDontRegex

`func (o *AddRss) HasDontRegex() bool`

HasDontRegex returns a boolean if a field has been set.

### GetDontOlderThan

`func (o *AddRss) GetDontOlderThan() int64`

GetDontOlderThan returns the DontOlderThan field if non-nil, zero value otherwise.

### GetDontOlderThanOk

`func (o *AddRss) GetDontOlderThanOk() (*int64, bool)`

GetDontOlderThanOk returns a tuple with the DontOlderThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontOlderThan

`func (o *AddRss) SetDontOlderThan(v int64)`

SetDontOlderThan sets DontOlderThan field to given value.

### HasDontOlderThan

`func (o *AddRss) HasDontOlderThan() bool`

HasDontOlderThan returns a boolean if a field has been set.

### GetPassCheck

`func (o *AddRss) GetPassCheck() bool`

GetPassCheck returns the PassCheck field if non-nil, zero value otherwise.

### GetPassCheckOk

`func (o *AddRss) GetPassCheckOk() (*bool, bool)`

GetPassCheckOk returns a tuple with the PassCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCheck

`func (o *AddRss) SetPassCheck(v bool)`

SetPassCheck sets PassCheck field to given value.

### HasPassCheck

`func (o *AddRss) HasPassCheck() bool`

HasPassCheck returns a boolean if a field has been set.

### GetScanInterval

`func (o *AddRss) GetScanInterval() int64`

GetScanInterval returns the ScanInterval field if non-nil, zero value otherwise.

### GetScanIntervalOk

`func (o *AddRss) GetScanIntervalOk() (*int64, bool)`

GetScanIntervalOk returns a tuple with the ScanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInterval

`func (o *AddRss) SetScanInterval(v int64)`

SetScanInterval sets ScanInterval field to given value.

### HasScanInterval

`func (o *AddRss) HasScanInterval() bool`

HasScanInterval returns a boolean if a field has been set.

### GetRssType

`func (o *AddRss) GetRssType() string`

GetRssType returns the RssType field if non-nil, zero value otherwise.

### GetRssTypeOk

`func (o *AddRss) GetRssTypeOk() (*string, bool)`

GetRssTypeOk returns a tuple with the RssType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssType

`func (o *AddRss) SetRssType(v string)`

SetRssType sets RssType field to given value.

### HasRssType

`func (o *AddRss) HasRssType() bool`

HasRssType returns a boolean if a field has been set.

### GetTorrentSeeding

`func (o *AddRss) GetTorrentSeeding() int64`

GetTorrentSeeding returns the TorrentSeeding field if non-nil, zero value otherwise.

### GetTorrentSeedingOk

`func (o *AddRss) GetTorrentSeedingOk() (*int64, bool)`

GetTorrentSeedingOk returns a tuple with the TorrentSeeding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentSeeding

`func (o *AddRss) SetTorrentSeeding(v int64)`

SetTorrentSeeding sets TorrentSeeding field to given value.

### HasTorrentSeeding

`func (o *AddRss) HasTorrentSeeding() bool`

HasTorrentSeeding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


