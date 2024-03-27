# RecordSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioOnly** | **bool** |  | 
**Layout** | [**LayoutSettings**](LayoutSettings.md) |  | 
**Mode** | **string** |  | 
**Quality** | **string** |  | 

## Methods

### NewRecordSettingsResponse

`func NewRecordSettingsResponse(audioOnly bool, layout LayoutSettings, mode string, quality string, ) *RecordSettingsResponse`

NewRecordSettingsResponse instantiates a new RecordSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSettingsResponseWithDefaults

`func NewRecordSettingsResponseWithDefaults() *RecordSettingsResponse`

NewRecordSettingsResponseWithDefaults instantiates a new RecordSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioOnly

`func (o *RecordSettingsResponse) GetAudioOnly() bool`

GetAudioOnly returns the AudioOnly field if non-nil, zero value otherwise.

### GetAudioOnlyOk

`func (o *RecordSettingsResponse) GetAudioOnlyOk() (*bool, bool)`

GetAudioOnlyOk returns a tuple with the AudioOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioOnly

`func (o *RecordSettingsResponse) SetAudioOnly(v bool)`

SetAudioOnly sets AudioOnly field to given value.


### GetLayout

`func (o *RecordSettingsResponse) GetLayout() LayoutSettings`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *RecordSettingsResponse) GetLayoutOk() (*LayoutSettings, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *RecordSettingsResponse) SetLayout(v LayoutSettings)`

SetLayout sets Layout field to given value.


### GetMode

`func (o *RecordSettingsResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RecordSettingsResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RecordSettingsResponse) SetMode(v string)`

SetMode sets Mode field to given value.


### GetQuality

`func (o *RecordSettingsResponse) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *RecordSettingsResponse) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *RecordSettingsResponse) SetQuality(v string)`

SetQuality sets Quality field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


