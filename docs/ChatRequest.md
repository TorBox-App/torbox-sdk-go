# ChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]ChatFile**](ChatFile.md) |  | [optional] 
**Blocks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChatRequest

`func NewChatRequest(message string, ) *ChatRequest`

NewChatRequest instantiates a new ChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatRequestWithDefaults

`func NewChatRequestWithDefaults() *ChatRequest`

NewChatRequestWithDefaults instantiates a new ChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ChatRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSessionId

`func (o *ChatRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetFiles

`func (o *ChatRequest) GetFiles() []ChatFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ChatRequest) GetFilesOk() (*[]ChatFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ChatRequest) SetFiles(v []ChatFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ChatRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetBlocks

`func (o *ChatRequest) GetBlocks() []string`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ChatRequest) GetBlocksOk() (*[]string, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ChatRequest) SetBlocks(v []string)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ChatRequest) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


