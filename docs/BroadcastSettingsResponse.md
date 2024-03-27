# BroadcastSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Hls** | [**HLSSettingsResponse**](HLSSettingsResponse.md) |  | 

## Methods

### NewBroadcastSettingsResponse

`func NewBroadcastSettingsResponse(enabled bool, hls HLSSettingsResponse, ) *BroadcastSettingsResponse`

NewBroadcastSettingsResponse instantiates a new BroadcastSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastSettingsResponseWithDefaults

`func NewBroadcastSettingsResponseWithDefaults() *BroadcastSettingsResponse`

NewBroadcastSettingsResponseWithDefaults instantiates a new BroadcastSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BroadcastSettingsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BroadcastSettingsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BroadcastSettingsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHls

`func (o *BroadcastSettingsResponse) GetHls() HLSSettingsResponse`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *BroadcastSettingsResponse) GetHlsOk() (*HLSSettingsResponse, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *BroadcastSettingsResponse) SetHls(v HLSSettingsResponse)`

SetHls sets Hls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


