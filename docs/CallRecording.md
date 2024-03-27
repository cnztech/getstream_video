# CallRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **time.Time** |  | 
**Filename** | **string** |  | 
**StartTime** | **time.Time** |  | 
**Url** | **string** |  | 

## Methods

### NewCallRecording

`func NewCallRecording(endTime time.Time, filename string, startTime time.Time, url string, ) *CallRecording`

NewCallRecording instantiates a new CallRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordingWithDefaults

`func NewCallRecordingWithDefaults() *CallRecording`

NewCallRecordingWithDefaults instantiates a new CallRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *CallRecording) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CallRecording) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CallRecording) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetFilename

`func (o *CallRecording) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CallRecording) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CallRecording) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetStartTime

`func (o *CallRecording) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallRecording) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallRecording) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetUrl

`func (o *CallRecording) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CallRecording) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CallRecording) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


