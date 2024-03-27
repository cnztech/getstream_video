# CallTranscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **time.Time** |  | 
**Filename** | **string** |  | 
**StartTime** | **time.Time** |  | 
**Url** | **string** |  | 

## Methods

### NewCallTranscription

`func NewCallTranscription(endTime time.Time, filename string, startTime time.Time, url string, ) *CallTranscription`

NewCallTranscription instantiates a new CallTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTranscriptionWithDefaults

`func NewCallTranscriptionWithDefaults() *CallTranscription`

NewCallTranscriptionWithDefaults instantiates a new CallTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *CallTranscription) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CallTranscription) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CallTranscription) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetFilename

`func (o *CallTranscription) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CallTranscription) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CallTranscription) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetStartTime

`func (o *CallTranscription) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallTranscription) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallTranscription) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetUrl

`func (o *CallTranscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CallTranscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CallTranscription) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


