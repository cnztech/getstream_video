# CreateCallTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalStorage** | Pointer to **string** |  | [optional] 
**Grants** | Pointer to **map[string][]string** |  | [optional] 
**Name** | **string** |  | 
**NotificationSettings** | Pointer to [**NotificationSettingsRequest**](NotificationSettingsRequest.md) |  | [optional] 
**Settings** | Pointer to [**CallSettingsRequest**](CallSettingsRequest.md) |  | [optional] 

## Methods

### NewCreateCallTypeRequest

`func NewCreateCallTypeRequest(name string, ) *CreateCallTypeRequest`

NewCreateCallTypeRequest instantiates a new CreateCallTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCallTypeRequestWithDefaults

`func NewCreateCallTypeRequestWithDefaults() *CreateCallTypeRequest`

NewCreateCallTypeRequestWithDefaults instantiates a new CreateCallTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalStorage

`func (o *CreateCallTypeRequest) GetExternalStorage() string`

GetExternalStorage returns the ExternalStorage field if non-nil, zero value otherwise.

### GetExternalStorageOk

`func (o *CreateCallTypeRequest) GetExternalStorageOk() (*string, bool)`

GetExternalStorageOk returns a tuple with the ExternalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorage

`func (o *CreateCallTypeRequest) SetExternalStorage(v string)`

SetExternalStorage sets ExternalStorage field to given value.

### HasExternalStorage

`func (o *CreateCallTypeRequest) HasExternalStorage() bool`

HasExternalStorage returns a boolean if a field has been set.

### GetGrants

`func (o *CreateCallTypeRequest) GetGrants() map[string][]string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *CreateCallTypeRequest) GetGrantsOk() (*map[string][]string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *CreateCallTypeRequest) SetGrants(v map[string][]string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *CreateCallTypeRequest) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetName

`func (o *CreateCallTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCallTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCallTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationSettings

`func (o *CreateCallTypeRequest) GetNotificationSettings() NotificationSettingsRequest`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *CreateCallTypeRequest) GetNotificationSettingsOk() (*NotificationSettingsRequest, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *CreateCallTypeRequest) SetNotificationSettings(v NotificationSettingsRequest)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *CreateCallTypeRequest) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetSettings

`func (o *CreateCallTypeRequest) GetSettings() CallSettingsRequest`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateCallTypeRequest) GetSettingsOk() (*CallSettingsRequest, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateCallTypeRequest) SetSettings(v CallSettingsRequest)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateCallTypeRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


