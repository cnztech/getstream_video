# HLSSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoOn** | **bool** |  | 
**Enabled** | **bool** |  | 
**Layout** | [**LayoutSettings**](LayoutSettings.md) |  | 
**QualityTracks** | **[]string** |  | 

## Methods

### NewHLSSettingsResponse

`func NewHLSSettingsResponse(autoOn bool, enabled bool, layout LayoutSettings, qualityTracks []string, ) *HLSSettingsResponse`

NewHLSSettingsResponse instantiates a new HLSSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHLSSettingsResponseWithDefaults

`func NewHLSSettingsResponseWithDefaults() *HLSSettingsResponse`

NewHLSSettingsResponseWithDefaults instantiates a new HLSSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoOn

`func (o *HLSSettingsResponse) GetAutoOn() bool`

GetAutoOn returns the AutoOn field if non-nil, zero value otherwise.

### GetAutoOnOk

`func (o *HLSSettingsResponse) GetAutoOnOk() (*bool, bool)`

GetAutoOnOk returns a tuple with the AutoOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOn

`func (o *HLSSettingsResponse) SetAutoOn(v bool)`

SetAutoOn sets AutoOn field to given value.


### GetEnabled

`func (o *HLSSettingsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HLSSettingsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HLSSettingsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLayout

`func (o *HLSSettingsResponse) GetLayout() LayoutSettings`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *HLSSettingsResponse) GetLayoutOk() (*LayoutSettings, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *HLSSettingsResponse) SetLayout(v LayoutSettings)`

SetLayout sets Layout field to given value.


### GetQualityTracks

`func (o *HLSSettingsResponse) GetQualityTracks() []string`

GetQualityTracks returns the QualityTracks field if non-nil, zero value otherwise.

### GetQualityTracksOk

`func (o *HLSSettingsResponse) GetQualityTracksOk() (*[]string, bool)`

GetQualityTracksOk returns a tuple with the QualityTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityTracks

`func (o *HLSSettingsResponse) SetQualityTracks(v []string)`

SetQualityTracks sets QualityTracks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


