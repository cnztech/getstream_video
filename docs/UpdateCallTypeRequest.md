# UpdateCallTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalStorage** | Pointer to **string** |  | [optional] 
**Grants** | Pointer to **map[string][]string** |  | [optional] 
**NotificationSettings** | Pointer to [**NotificationSettingsRequest**](NotificationSettingsRequest.md) |  | [optional] 
**Settings** | Pointer to [**CallSettingsRequest**](CallSettingsRequest.md) |  | [optional] 

## Methods

### NewUpdateCallTypeRequest

`func NewUpdateCallTypeRequest() *UpdateCallTypeRequest`

NewUpdateCallTypeRequest instantiates a new UpdateCallTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallTypeRequestWithDefaults

`func NewUpdateCallTypeRequestWithDefaults() *UpdateCallTypeRequest`

NewUpdateCallTypeRequestWithDefaults instantiates a new UpdateCallTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalStorage

`func (o *UpdateCallTypeRequest) GetExternalStorage() string`

GetExternalStorage returns the ExternalStorage field if non-nil, zero value otherwise.

### GetExternalStorageOk

`func (o *UpdateCallTypeRequest) GetExternalStorageOk() (*string, bool)`

GetExternalStorageOk returns a tuple with the ExternalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorage

`func (o *UpdateCallTypeRequest) SetExternalStorage(v string)`

SetExternalStorage sets ExternalStorage field to given value.

### HasExternalStorage

`func (o *UpdateCallTypeRequest) HasExternalStorage() bool`

HasExternalStorage returns a boolean if a field has been set.

### GetGrants

`func (o *UpdateCallTypeRequest) GetGrants() map[string][]string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *UpdateCallTypeRequest) GetGrantsOk() (*map[string][]string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *UpdateCallTypeRequest) SetGrants(v map[string][]string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *UpdateCallTypeRequest) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetNotificationSettings

`func (o *UpdateCallTypeRequest) GetNotificationSettings() NotificationSettingsRequest`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *UpdateCallTypeRequest) GetNotificationSettingsOk() (*NotificationSettingsRequest, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *UpdateCallTypeRequest) SetNotificationSettings(v NotificationSettingsRequest)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *UpdateCallTypeRequest) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateCallTypeRequest) GetSettings() CallSettingsRequest`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateCallTypeRequest) GetSettingsOk() (*CallSettingsRequest, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateCallTypeRequest) SetSettings(v CallSettingsRequest)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateCallTypeRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


