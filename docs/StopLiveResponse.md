# StopLiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | [**CallResponse**](CallResponse.md) |  | 
**Duration** | **string** | Duration of the request in human-readable format | 

## Methods

### NewStopLiveResponse

`func NewStopLiveResponse(call CallResponse, duration string, ) *StopLiveResponse`

NewStopLiveResponse instantiates a new StopLiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopLiveResponseWithDefaults

`func NewStopLiveResponseWithDefaults() *StopLiveResponse`

NewStopLiveResponseWithDefaults instantiates a new StopLiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *StopLiveResponse) GetCall() CallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *StopLiveResponse) GetCallOk() (*CallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *StopLiveResponse) SetCall(v CallResponse)`

SetCall sets Call field to given value.


### GetDuration

`func (o *StopLiveResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *StopLiveResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *StopLiveResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


