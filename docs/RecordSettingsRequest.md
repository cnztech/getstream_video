# RecordSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioOnly** | Pointer to **bool** |  | [optional] 
**Layout** | Pointer to [**LayoutSettingsRequest**](LayoutSettingsRequest.md) |  | [optional] 
**Mode** | **string** |  | 
**Quality** | Pointer to **string** |  | [optional] 

## Methods

### NewRecordSettingsRequest

`func NewRecordSettingsRequest(mode string, ) *RecordSettingsRequest`

NewRecordSettingsRequest instantiates a new RecordSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSettingsRequestWithDefaults

`func NewRecordSettingsRequestWithDefaults() *RecordSettingsRequest`

NewRecordSettingsRequestWithDefaults instantiates a new RecordSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioOnly

`func (o *RecordSettingsRequest) GetAudioOnly() bool`

GetAudioOnly returns the AudioOnly field if non-nil, zero value otherwise.

### GetAudioOnlyOk

`func (o *RecordSettingsRequest) GetAudioOnlyOk() (*bool, bool)`

GetAudioOnlyOk returns a tuple with the AudioOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioOnly

`func (o *RecordSettingsRequest) SetAudioOnly(v bool)`

SetAudioOnly sets AudioOnly field to given value.

### HasAudioOnly

`func (o *RecordSettingsRequest) HasAudioOnly() bool`

HasAudioOnly returns a boolean if a field has been set.

### GetLayout

`func (o *RecordSettingsRequest) GetLayout() LayoutSettingsRequest`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *RecordSettingsRequest) GetLayoutOk() (*LayoutSettingsRequest, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *RecordSettingsRequest) SetLayout(v LayoutSettingsRequest)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *RecordSettingsRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetMode

`func (o *RecordSettingsRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RecordSettingsRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RecordSettingsRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetQuality

`func (o *RecordSettingsRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *RecordSettingsRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *RecordSettingsRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *RecordSettingsRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


