# LayoutSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAppUrl** | Pointer to **string** |  | [optional] 
**ExternalCssUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLayoutSettingsRequest

`func NewLayoutSettingsRequest(name string, ) *LayoutSettingsRequest`

NewLayoutSettingsRequest instantiates a new LayoutSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutSettingsRequestWithDefaults

`func NewLayoutSettingsRequestWithDefaults() *LayoutSettingsRequest`

NewLayoutSettingsRequestWithDefaults instantiates a new LayoutSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAppUrl

`func (o *LayoutSettingsRequest) GetExternalAppUrl() string`

GetExternalAppUrl returns the ExternalAppUrl field if non-nil, zero value otherwise.

### GetExternalAppUrlOk

`func (o *LayoutSettingsRequest) GetExternalAppUrlOk() (*string, bool)`

GetExternalAppUrlOk returns a tuple with the ExternalAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAppUrl

`func (o *LayoutSettingsRequest) SetExternalAppUrl(v string)`

SetExternalAppUrl sets ExternalAppUrl field to given value.

### HasExternalAppUrl

`func (o *LayoutSettingsRequest) HasExternalAppUrl() bool`

HasExternalAppUrl returns a boolean if a field has been set.

### GetExternalCssUrl

`func (o *LayoutSettingsRequest) GetExternalCssUrl() string`

GetExternalCssUrl returns the ExternalCssUrl field if non-nil, zero value otherwise.

### GetExternalCssUrlOk

`func (o *LayoutSettingsRequest) GetExternalCssUrlOk() (*string, bool)`

GetExternalCssUrlOk returns a tuple with the ExternalCssUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCssUrl

`func (o *LayoutSettingsRequest) SetExternalCssUrl(v string)`

SetExternalCssUrl sets ExternalCssUrl field to given value.

### HasExternalCssUrl

`func (o *LayoutSettingsRequest) HasExternalCssUrl() bool`

HasExternalCssUrl returns a boolean if a field has been set.

### GetName

`func (o *LayoutSettingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayoutSettingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayoutSettingsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *LayoutSettingsRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *LayoutSettingsRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *LayoutSettingsRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *LayoutSettingsRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


