# HLSSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoOn** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Layout** | Pointer to [**LayoutSettingsRequest**](LayoutSettingsRequest.md) |  | [optional] 
**QualityTracks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHLSSettingsRequest

`func NewHLSSettingsRequest() *HLSSettingsRequest`

NewHLSSettingsRequest instantiates a new HLSSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHLSSettingsRequestWithDefaults

`func NewHLSSettingsRequestWithDefaults() *HLSSettingsRequest`

NewHLSSettingsRequestWithDefaults instantiates a new HLSSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoOn

`func (o *HLSSettingsRequest) GetAutoOn() bool`

GetAutoOn returns the AutoOn field if non-nil, zero value otherwise.

### GetAutoOnOk

`func (o *HLSSettingsRequest) GetAutoOnOk() (*bool, bool)`

GetAutoOnOk returns a tuple with the AutoOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOn

`func (o *HLSSettingsRequest) SetAutoOn(v bool)`

SetAutoOn sets AutoOn field to given value.

### HasAutoOn

`func (o *HLSSettingsRequest) HasAutoOn() bool`

HasAutoOn returns a boolean if a field has been set.

### GetEnabled

`func (o *HLSSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HLSSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HLSSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HLSSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLayout

`func (o *HLSSettingsRequest) GetLayout() LayoutSettingsRequest`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *HLSSettingsRequest) GetLayoutOk() (*LayoutSettingsRequest, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *HLSSettingsRequest) SetLayout(v LayoutSettingsRequest)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *HLSSettingsRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetQualityTracks

`func (o *HLSSettingsRequest) GetQualityTracks() []string`

GetQualityTracks returns the QualityTracks field if non-nil, zero value otherwise.

### GetQualityTracksOk

`func (o *HLSSettingsRequest) GetQualityTracksOk() (*[]string, bool)`

GetQualityTracksOk returns a tuple with the QualityTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityTracks

`func (o *HLSSettingsRequest) SetQualityTracks(v []string)`

SetQualityTracks sets QualityTracks field to given value.

### HasQualityTracks

`func (o *HLSSettingsRequest) HasQualityTracks() bool`

HasQualityTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


