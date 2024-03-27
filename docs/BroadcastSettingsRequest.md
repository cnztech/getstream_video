# BroadcastSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Hls** | Pointer to [**HLSSettingsRequest**](HLSSettingsRequest.md) |  | [optional] 

## Methods

### NewBroadcastSettingsRequest

`func NewBroadcastSettingsRequest() *BroadcastSettingsRequest`

NewBroadcastSettingsRequest instantiates a new BroadcastSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastSettingsRequestWithDefaults

`func NewBroadcastSettingsRequestWithDefaults() *BroadcastSettingsRequest`

NewBroadcastSettingsRequestWithDefaults instantiates a new BroadcastSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BroadcastSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BroadcastSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BroadcastSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BroadcastSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHls

`func (o *BroadcastSettingsRequest) GetHls() HLSSettingsRequest`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *BroadcastSettingsRequest) GetHlsOk() (*HLSSettingsRequest, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *BroadcastSettingsRequest) SetHls(v HLSSettingsRequest)`

SetHls sets Hls field to given value.

### HasHls

`func (o *BroadcastSettingsRequest) HasHls() bool`

HasHls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


