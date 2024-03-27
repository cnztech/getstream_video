# VideoSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRequestEnabled** | **bool** |  | 
**CameraDefaultOn** | **bool** |  | 
**CameraFacing** | **string** |  | 
**Enabled** | **bool** |  | 
**TargetResolution** | [**TargetResolution**](TargetResolution.md) |  | 

## Methods

### NewVideoSettings

`func NewVideoSettings(accessRequestEnabled bool, cameraDefaultOn bool, cameraFacing string, enabled bool, targetResolution TargetResolution, ) *VideoSettings`

NewVideoSettings instantiates a new VideoSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSettingsWithDefaults

`func NewVideoSettingsWithDefaults() *VideoSettings`

NewVideoSettingsWithDefaults instantiates a new VideoSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRequestEnabled

`func (o *VideoSettings) GetAccessRequestEnabled() bool`

GetAccessRequestEnabled returns the AccessRequestEnabled field if non-nil, zero value otherwise.

### GetAccessRequestEnabledOk

`func (o *VideoSettings) GetAccessRequestEnabledOk() (*bool, bool)`

GetAccessRequestEnabledOk returns a tuple with the AccessRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestEnabled

`func (o *VideoSettings) SetAccessRequestEnabled(v bool)`

SetAccessRequestEnabled sets AccessRequestEnabled field to given value.


### GetCameraDefaultOn

`func (o *VideoSettings) GetCameraDefaultOn() bool`

GetCameraDefaultOn returns the CameraDefaultOn field if non-nil, zero value otherwise.

### GetCameraDefaultOnOk

`func (o *VideoSettings) GetCameraDefaultOnOk() (*bool, bool)`

GetCameraDefaultOnOk returns a tuple with the CameraDefaultOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraDefaultOn

`func (o *VideoSettings) SetCameraDefaultOn(v bool)`

SetCameraDefaultOn sets CameraDefaultOn field to given value.


### GetCameraFacing

`func (o *VideoSettings) GetCameraFacing() string`

GetCameraFacing returns the CameraFacing field if non-nil, zero value otherwise.

### GetCameraFacingOk

`func (o *VideoSettings) GetCameraFacingOk() (*string, bool)`

GetCameraFacingOk returns a tuple with the CameraFacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraFacing

`func (o *VideoSettings) SetCameraFacing(v string)`

SetCameraFacing sets CameraFacing field to given value.


### GetEnabled

`func (o *VideoSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VideoSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VideoSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTargetResolution

`func (o *VideoSettings) GetTargetResolution() TargetResolution`

GetTargetResolution returns the TargetResolution field if non-nil, zero value otherwise.

### GetTargetResolutionOk

`func (o *VideoSettings) GetTargetResolutionOk() (*TargetResolution, bool)`

GetTargetResolutionOk returns a tuple with the TargetResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResolution

`func (o *VideoSettings) SetTargetResolution(v TargetResolution)`

SetTargetResolution sets TargetResolution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


