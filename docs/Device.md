# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Date/time of creation | 
**Disabled** | Pointer to **bool** | Whether device is disabled or not | [optional] 
**DisabledReason** | Pointer to **string** | Reason explaining why device had been disabled | [optional] 
**Id** | **string** |  | 
**PushProvider** | **string** |  | 
**PushProviderName** | Pointer to **string** |  | [optional] 
**Voip** | Pointer to **bool** | When true the token is for Apple VoIP push notifications | [optional] 

## Methods

### NewDevice

`func NewDevice(createdAt time.Time, id string, pushProvider string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Device) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisabled

`func (o *Device) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Device) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Device) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Device) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledReason

`func (o *Device) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *Device) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *Device) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.

### HasDisabledReason

`func (o *Device) HasDisabledReason() bool`

HasDisabledReason returns a boolean if a field has been set.

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetPushProvider

`func (o *Device) GetPushProvider() string`

GetPushProvider returns the PushProvider field if non-nil, zero value otherwise.

### GetPushProviderOk

`func (o *Device) GetPushProviderOk() (*string, bool)`

GetPushProviderOk returns a tuple with the PushProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProvider

`func (o *Device) SetPushProvider(v string)`

SetPushProvider sets PushProvider field to given value.


### GetPushProviderName

`func (o *Device) GetPushProviderName() string`

GetPushProviderName returns the PushProviderName field if non-nil, zero value otherwise.

### GetPushProviderNameOk

`func (o *Device) GetPushProviderNameOk() (*string, bool)`

GetPushProviderNameOk returns a tuple with the PushProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProviderName

`func (o *Device) SetPushProviderName(v string)`

SetPushProviderName sets PushProviderName field to given value.

### HasPushProviderName

`func (o *Device) HasPushProviderName() bool`

HasPushProviderName returns a boolean if a field has been set.

### GetVoip

`func (o *Device) GetVoip() bool`

GetVoip returns the Voip field if non-nil, zero value otherwise.

### GetVoipOk

`func (o *Device) GetVoipOk() (*bool, bool)`

GetVoipOk returns a tuple with the Voip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoip

`func (o *Device) SetVoip(v bool)`

SetVoip sets Voip field to given value.

### HasVoip

`func (o *Device) HasVoip() bool`

HasVoip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


