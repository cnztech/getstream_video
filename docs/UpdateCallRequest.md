# UpdateCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **map[string]interface{}** | Custom data for this object | [optional] 
**SettingsOverride** | Pointer to [**CallSettingsRequest**](CallSettingsRequest.md) |  | [optional] 
**StartsAt** | Pointer to **time.Time** | the time the call is scheduled to start | [optional] 

## Methods

### NewUpdateCallRequest

`func NewUpdateCallRequest() *UpdateCallRequest`

NewUpdateCallRequest instantiates a new UpdateCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallRequestWithDefaults

`func NewUpdateCallRequestWithDefaults() *UpdateCallRequest`

NewUpdateCallRequestWithDefaults instantiates a new UpdateCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *UpdateCallRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UpdateCallRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UpdateCallRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UpdateCallRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetSettingsOverride

`func (o *UpdateCallRequest) GetSettingsOverride() CallSettingsRequest`

GetSettingsOverride returns the SettingsOverride field if non-nil, zero value otherwise.

### GetSettingsOverrideOk

`func (o *UpdateCallRequest) GetSettingsOverrideOk() (*CallSettingsRequest, bool)`

GetSettingsOverrideOk returns a tuple with the SettingsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsOverride

`func (o *UpdateCallRequest) SetSettingsOverride(v CallSettingsRequest)`

SetSettingsOverride sets SettingsOverride field to given value.

### HasSettingsOverride

`func (o *UpdateCallRequest) HasSettingsOverride() bool`

HasSettingsOverride returns a boolean if a field has been set.

### GetStartsAt

`func (o *UpdateCallRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateCallRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateCallRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateCallRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


