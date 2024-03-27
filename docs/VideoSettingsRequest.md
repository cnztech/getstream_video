# VideoSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRequestEnabled** | Pointer to **bool** |  | [optional] 
**CameraDefaultOn** | Pointer to **bool** |  | [optional] 
**CameraFacing** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TargetResolution** | Pointer to [**TargetResolutionRequest**](TargetResolutionRequest.md) |  | [optional] 

## Methods

### NewVideoSettingsRequest

`func NewVideoSettingsRequest() *VideoSettingsRequest`

NewVideoSettingsRequest instantiates a new VideoSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSettingsRequestWithDefaults

`func NewVideoSettingsRequestWithDefaults() *VideoSettingsRequest`

NewVideoSettingsRequestWithDefaults instantiates a new VideoSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRequestEnabled

`func (o *VideoSettingsRequest) GetAccessRequestEnabled() bool`

GetAccessRequestEnabled returns the AccessRequestEnabled field if non-nil, zero value otherwise.

### GetAccessRequestEnabledOk

`func (o *VideoSettingsRequest) GetAccessRequestEnabledOk() (*bool, bool)`

GetAccessRequestEnabledOk returns a tuple with the AccessRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestEnabled

`func (o *VideoSettingsRequest) SetAccessRequestEnabled(v bool)`

SetAccessRequestEnabled sets AccessRequestEnabled field to given value.

### HasAccessRequestEnabled

`func (o *VideoSettingsRequest) HasAccessRequestEnabled() bool`

HasAccessRequestEnabled returns a boolean if a field has been set.

### GetCameraDefaultOn

`func (o *VideoSettingsRequest) GetCameraDefaultOn() bool`

GetCameraDefaultOn returns the CameraDefaultOn field if non-nil, zero value otherwise.

### GetCameraDefaultOnOk

`func (o *VideoSettingsRequest) GetCameraDefaultOnOk() (*bool, bool)`

GetCameraDefaultOnOk returns a tuple with the CameraDefaultOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraDefaultOn

`func (o *VideoSettingsRequest) SetCameraDefaultOn(v bool)`

SetCameraDefaultOn sets CameraDefaultOn field to given value.

### HasCameraDefaultOn

`func (o *VideoSettingsRequest) HasCameraDefaultOn() bool`

HasCameraDefaultOn returns a boolean if a field has been set.

### GetCameraFacing

`func (o *VideoSettingsRequest) GetCameraFacing() string`

GetCameraFacing returns the CameraFacing field if non-nil, zero value otherwise.

### GetCameraFacingOk

`func (o *VideoSettingsRequest) GetCameraFacingOk() (*string, bool)`

GetCameraFacingOk returns a tuple with the CameraFacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraFacing

`func (o *VideoSettingsRequest) SetCameraFacing(v string)`

SetCameraFacing sets CameraFacing field to given value.

### HasCameraFacing

`func (o *VideoSettingsRequest) HasCameraFacing() bool`

HasCameraFacing returns a boolean if a field has been set.

### GetEnabled

`func (o *VideoSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VideoSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VideoSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VideoSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTargetResolution

`func (o *VideoSettingsRequest) GetTargetResolution() TargetResolutionRequest`

GetTargetResolution returns the TargetResolution field if non-nil, zero value otherwise.

### GetTargetResolutionOk

`func (o *VideoSettingsRequest) GetTargetResolutionOk() (*TargetResolutionRequest, bool)`

GetTargetResolutionOk returns a tuple with the TargetResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResolution

`func (o *VideoSettingsRequest) SetTargetResolution(v TargetResolutionRequest)`

SetTargetResolution sets TargetResolution field to given value.

### HasTargetResolution

`func (o *VideoSettingsRequest) HasTargetResolution() bool`

HasTargetResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


