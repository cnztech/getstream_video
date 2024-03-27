# RingSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCancelTimeoutMs** | Pointer to **int32** |  | [optional] 
**IncomingCallTimeoutMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewRingSettingsRequest

`func NewRingSettingsRequest() *RingSettingsRequest`

NewRingSettingsRequest instantiates a new RingSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRingSettingsRequestWithDefaults

`func NewRingSettingsRequestWithDefaults() *RingSettingsRequest`

NewRingSettingsRequestWithDefaults instantiates a new RingSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCancelTimeoutMs

`func (o *RingSettingsRequest) GetAutoCancelTimeoutMs() int32`

GetAutoCancelTimeoutMs returns the AutoCancelTimeoutMs field if non-nil, zero value otherwise.

### GetAutoCancelTimeoutMsOk

`func (o *RingSettingsRequest) GetAutoCancelTimeoutMsOk() (*int32, bool)`

GetAutoCancelTimeoutMsOk returns a tuple with the AutoCancelTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelTimeoutMs

`func (o *RingSettingsRequest) SetAutoCancelTimeoutMs(v int32)`

SetAutoCancelTimeoutMs sets AutoCancelTimeoutMs field to given value.

### HasAutoCancelTimeoutMs

`func (o *RingSettingsRequest) HasAutoCancelTimeoutMs() bool`

HasAutoCancelTimeoutMs returns a boolean if a field has been set.

### GetIncomingCallTimeoutMs

`func (o *RingSettingsRequest) GetIncomingCallTimeoutMs() int32`

GetIncomingCallTimeoutMs returns the IncomingCallTimeoutMs field if non-nil, zero value otherwise.

### GetIncomingCallTimeoutMsOk

`func (o *RingSettingsRequest) GetIncomingCallTimeoutMsOk() (*int32, bool)`

GetIncomingCallTimeoutMsOk returns a tuple with the IncomingCallTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallTimeoutMs

`func (o *RingSettingsRequest) SetIncomingCallTimeoutMs(v int32)`

SetIncomingCallTimeoutMs sets IncomingCallTimeoutMs field to given value.

### HasIncomingCallTimeoutMs

`func (o *RingSettingsRequest) HasIncomingCallTimeoutMs() bool`

HasIncomingCallTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


