# CallTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**ExternalStorage** | Pointer to **string** |  | [optional] 
**Grants** | **map[string][]string** |  | 
**Name** | **string** |  | 
**NotificationSettings** | [**NotificationSettings**](NotificationSettings.md) |  | 
**Settings** | [**CallSettingsResponse**](CallSettingsResponse.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCallTypeResponse

`func NewCallTypeResponse(createdAt time.Time, grants map[string][]string, name string, notificationSettings NotificationSettings, settings CallSettingsResponse, updatedAt time.Time, ) *CallTypeResponse`

NewCallTypeResponse instantiates a new CallTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTypeResponseWithDefaults

`func NewCallTypeResponseWithDefaults() *CallTypeResponse`

NewCallTypeResponseWithDefaults instantiates a new CallTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CallTypeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CallTypeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CallTypeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExternalStorage

`func (o *CallTypeResponse) GetExternalStorage() string`

GetExternalStorage returns the ExternalStorage field if non-nil, zero value otherwise.

### GetExternalStorageOk

`func (o *CallTypeResponse) GetExternalStorageOk() (*string, bool)`

GetExternalStorageOk returns a tuple with the ExternalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorage

`func (o *CallTypeResponse) SetExternalStorage(v string)`

SetExternalStorage sets ExternalStorage field to given value.

### HasExternalStorage

`func (o *CallTypeResponse) HasExternalStorage() bool`

HasExternalStorage returns a boolean if a field has been set.

### GetGrants

`func (o *CallTypeResponse) GetGrants() map[string][]string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *CallTypeResponse) GetGrantsOk() (*map[string][]string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *CallTypeResponse) SetGrants(v map[string][]string)`

SetGrants sets Grants field to given value.


### GetName

`func (o *CallTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallTypeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationSettings

`func (o *CallTypeResponse) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *CallTypeResponse) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *CallTypeResponse) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.


### GetSettings

`func (o *CallTypeResponse) GetSettings() CallSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CallTypeResponse) GetSettingsOk() (*CallSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CallTypeResponse) SetSettings(v CallSettingsResponse)`

SetSettings sets Settings field to given value.


### GetUpdatedAt

`func (o *CallTypeResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CallTypeResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CallTypeResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


