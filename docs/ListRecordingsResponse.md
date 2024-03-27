# ListRecordingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** |  | 
**Recordings** | [**[]CallRecording**](CallRecording.md) |  | 

## Methods

### NewListRecordingsResponse

`func NewListRecordingsResponse(duration string, recordings []CallRecording, ) *ListRecordingsResponse`

NewListRecordingsResponse instantiates a new ListRecordingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRecordingsResponseWithDefaults

`func NewListRecordingsResponseWithDefaults() *ListRecordingsResponse`

NewListRecordingsResponseWithDefaults instantiates a new ListRecordingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ListRecordingsResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListRecordingsResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListRecordingsResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetRecordings

`func (o *ListRecordingsResponse) GetRecordings() []CallRecording`

GetRecordings returns the Recordings field if non-nil, zero value otherwise.

### GetRecordingsOk

`func (o *ListRecordingsResponse) GetRecordingsOk() (*[]CallRecording, bool)`

GetRecordingsOk returns a tuple with the Recordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordings

`func (o *ListRecordingsResponse) SetRecordings(v []CallRecording)`

SetRecordings sets Recordings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


