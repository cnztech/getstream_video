# CreateCallTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Duration** | **string** |  | 
**ExternalStorage** | Pointer to **string** |  | [optional] 
**Grants** | **map[string][]string** |  | 
**Name** | **string** |  | 
**NotificationSettings** | [**NotificationSettings**](NotificationSettings.md) |  | 
**Settings** | [**CallSettingsResponse**](CallSettingsResponse.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCreateCallTypeResponse

`func NewCreateCallTypeResponse(createdAt time.Time, duration string, grants map[string][]string, name string, notificationSettings NotificationSettings, settings CallSettingsResponse, updatedAt time.Time, ) *CreateCallTypeResponse`

NewCreateCallTypeResponse instantiates a new CreateCallTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCallTypeResponseWithDefaults

`func NewCreateCallTypeResponseWithDefaults() *CreateCallTypeResponse`

NewCreateCallTypeResponseWithDefaults instantiates a new CreateCallTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CreateCallTypeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCallTypeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCallTypeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDuration

`func (o *CreateCallTypeResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateCallTypeResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateCallTypeResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetExternalStorage

`func (o *CreateCallTypeResponse) GetExternalStorage() string`

GetExternalStorage returns the ExternalStorage field if non-nil, zero value otherwise.

### GetExternalStorageOk

`func (o *CreateCallTypeResponse) GetExternalStorageOk() (*string, bool)`

GetExternalStorageOk returns a tuple with the ExternalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorage

`func (o *CreateCallTypeResponse) SetExternalStorage(v string)`

SetExternalStorage sets ExternalStorage field to given value.

### HasExternalStorage

`func (o *CreateCallTypeResponse) HasExternalStorage() bool`

HasExternalStorage returns a boolean if a field has been set.

### GetGrants

`func (o *CreateCallTypeResponse) GetGrants() map[string][]string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *CreateCallTypeResponse) GetGrantsOk() (*map[string][]string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *CreateCallTypeResponse) SetGrants(v map[string][]string)`

SetGrants sets Grants field to given value.


### GetName

`func (o *CreateCallTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCallTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCallTypeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationSettings

`func (o *CreateCallTypeResponse) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *CreateCallTypeResponse) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *CreateCallTypeResponse) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.


### GetSettings

`func (o *CreateCallTypeResponse) GetSettings() CallSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateCallTypeResponse) GetSettingsOk() (*CallSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateCallTypeResponse) SetSettings(v CallSettingsResponse)`

SetSettings sets Settings field to given value.


### GetUpdatedAt

`func (o *CreateCallTypeResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateCallTypeResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateCallTypeResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


