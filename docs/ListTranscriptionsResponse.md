# ListTranscriptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** |  | 
**Transcriptions** | [**[]CallTranscription**](CallTranscription.md) |  | 

## Methods

### NewListTranscriptionsResponse

`func NewListTranscriptionsResponse(duration string, transcriptions []CallTranscription, ) *ListTranscriptionsResponse`

NewListTranscriptionsResponse instantiates a new ListTranscriptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTranscriptionsResponseWithDefaults

`func NewListTranscriptionsResponseWithDefaults() *ListTranscriptionsResponse`

NewListTranscriptionsResponseWithDefaults instantiates a new ListTranscriptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ListTranscriptionsResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListTranscriptionsResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListTranscriptionsResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetTranscriptions

`func (o *ListTranscriptionsResponse) GetTranscriptions() []CallTranscription`

GetTranscriptions returns the Transcriptions field if non-nil, zero value otherwise.

### GetTranscriptionsOk

`func (o *ListTranscriptionsResponse) GetTranscriptionsOk() (*[]CallTranscription, bool)`

GetTranscriptionsOk returns a tuple with the Transcriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptions

`func (o *ListTranscriptionsResponse) SetTranscriptions(v []CallTranscription)`

SetTranscriptions sets Transcriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


