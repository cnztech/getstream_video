# RingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCancelTimeoutMs** | **int32** |  | 
**IncomingCallTimeoutMs** | **int32** |  | 

## Methods

### NewRingSettings

`func NewRingSettings(autoCancelTimeoutMs int32, incomingCallTimeoutMs int32, ) *RingSettings`

NewRingSettings instantiates a new RingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRingSettingsWithDefaults

`func NewRingSettingsWithDefaults() *RingSettings`

NewRingSettingsWithDefaults instantiates a new RingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCancelTimeoutMs

`func (o *RingSettings) GetAutoCancelTimeoutMs() int32`

GetAutoCancelTimeoutMs returns the AutoCancelTimeoutMs field if non-nil, zero value otherwise.

### GetAutoCancelTimeoutMsOk

`func (o *RingSettings) GetAutoCancelTimeoutMsOk() (*int32, bool)`

GetAutoCancelTimeoutMsOk returns a tuple with the AutoCancelTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelTimeoutMs

`func (o *RingSettings) SetAutoCancelTimeoutMs(v int32)`

SetAutoCancelTimeoutMs sets AutoCancelTimeoutMs field to given value.


### GetIncomingCallTimeoutMs

`func (o *RingSettings) GetIncomingCallTimeoutMs() int32`

GetIncomingCallTimeoutMs returns the IncomingCallTimeoutMs field if non-nil, zero value otherwise.

### GetIncomingCallTimeoutMsOk

`func (o *RingSettings) GetIncomingCallTimeoutMsOk() (*int32, bool)`

GetIncomingCallTimeoutMsOk returns a tuple with the IncomingCallTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallTimeoutMs

`func (o *RingSettings) SetIncomingCallTimeoutMs(v int32)`

SetIncomingCallTimeoutMs sets IncomingCallTimeoutMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


