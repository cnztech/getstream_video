# AudioSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRequestEnabled** | Pointer to **bool** |  | [optional] 
**DefaultDevice** | **string** |  | 
**MicDefaultOn** | Pointer to **bool** |  | [optional] 
**OpusDtxEnabled** | Pointer to **bool** |  | [optional] 
**RedundantCodingEnabled** | Pointer to **bool** |  | [optional] 
**SpeakerDefaultOn** | Pointer to **bool** |  | [optional] 

## Methods

### NewAudioSettingsRequest

`func NewAudioSettingsRequest(defaultDevice string, ) *AudioSettingsRequest`

NewAudioSettingsRequest instantiates a new AudioSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioSettingsRequestWithDefaults

`func NewAudioSettingsRequestWithDefaults() *AudioSettingsRequest`

NewAudioSettingsRequestWithDefaults instantiates a new AudioSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRequestEnabled

`func (o *AudioSettingsRequest) GetAccessRequestEnabled() bool`

GetAccessRequestEnabled returns the AccessRequestEnabled field if non-nil, zero value otherwise.

### GetAccessRequestEnabledOk

`func (o *AudioSettingsRequest) GetAccessRequestEnabledOk() (*bool, bool)`

GetAccessRequestEnabledOk returns a tuple with the AccessRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestEnabled

`func (o *AudioSettingsRequest) SetAccessRequestEnabled(v bool)`

SetAccessRequestEnabled sets AccessRequestEnabled field to given value.

### HasAccessRequestEnabled

`func (o *AudioSettingsRequest) HasAccessRequestEnabled() bool`

HasAccessRequestEnabled returns a boolean if a field has been set.

### GetDefaultDevice

`func (o *AudioSettingsRequest) GetDefaultDevice() string`

GetDefaultDevice returns the DefaultDevice field if non-nil, zero value otherwise.

### GetDefaultDeviceOk

`func (o *AudioSettingsRequest) GetDefaultDeviceOk() (*string, bool)`

GetDefaultDeviceOk returns a tuple with the DefaultDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDevice

`func (o *AudioSettingsRequest) SetDefaultDevice(v string)`

SetDefaultDevice sets DefaultDevice field to given value.


### GetMicDefaultOn

`func (o *AudioSettingsRequest) GetMicDefaultOn() bool`

GetMicDefaultOn returns the MicDefaultOn field if non-nil, zero value otherwise.

### GetMicDefaultOnOk

`func (o *AudioSettingsRequest) GetMicDefaultOnOk() (*bool, bool)`

GetMicDefaultOnOk returns a tuple with the MicDefaultOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicDefaultOn

`func (o *AudioSettingsRequest) SetMicDefaultOn(v bool)`

SetMicDefaultOn sets MicDefaultOn field to given value.

### HasMicDefaultOn

`func (o *AudioSettingsRequest) HasMicDefaultOn() bool`

HasMicDefaultOn returns a boolean if a field has been set.

### GetOpusDtxEnabled

`func (o *AudioSettingsRequest) GetOpusDtxEnabled() bool`

GetOpusDtxEnabled returns the OpusDtxEnabled field if non-nil, zero value otherwise.

### GetOpusDtxEnabledOk

`func (o *AudioSettingsRequest) GetOpusDtxEnabledOk() (*bool, bool)`

GetOpusDtxEnabledOk returns a tuple with the OpusDtxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpusDtxEnabled

`func (o *AudioSettingsRequest) SetOpusDtxEnabled(v bool)`

SetOpusDtxEnabled sets OpusDtxEnabled field to given value.

### HasOpusDtxEnabled

`func (o *AudioSettingsRequest) HasOpusDtxEnabled() bool`

HasOpusDtxEnabled returns a boolean if a field has been set.

### GetRedundantCodingEnabled

`func (o *AudioSettingsRequest) GetRedundantCodingEnabled() bool`

GetRedundantCodingEnabled returns the RedundantCodingEnabled field if non-nil, zero value otherwise.

### GetRedundantCodingEnabledOk

`func (o *AudioSettingsRequest) GetRedundantCodingEnabledOk() (*bool, bool)`

GetRedundantCodingEnabledOk returns a tuple with the RedundantCodingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantCodingEnabled

`func (o *AudioSettingsRequest) SetRedundantCodingEnabled(v bool)`

SetRedundantCodingEnabled sets RedundantCodingEnabled field to given value.

### HasRedundantCodingEnabled

`func (o *AudioSettingsRequest) HasRedundantCodingEnabled() bool`

HasRedundantCodingEnabled returns a boolean if a field has been set.

### GetSpeakerDefaultOn

`func (o *AudioSettingsRequest) GetSpeakerDefaultOn() bool`

GetSpeakerDefaultOn returns the SpeakerDefaultOn field if non-nil, zero value otherwise.

### GetSpeakerDefaultOnOk

`func (o *AudioSettingsRequest) GetSpeakerDefaultOnOk() (*bool, bool)`

GetSpeakerDefaultOnOk returns a tuple with the SpeakerDefaultOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerDefaultOn

`func (o *AudioSettingsRequest) SetSpeakerDefaultOn(v bool)`

SetSpeakerDefaultOn sets SpeakerDefaultOn field to given value.

### HasSpeakerDefaultOn

`func (o *AudioSettingsRequest) HasSpeakerDefaultOn() bool`

HasSpeakerDefaultOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


