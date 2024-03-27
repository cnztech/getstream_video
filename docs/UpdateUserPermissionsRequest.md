# UpdateUserPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantPermissions** | Pointer to **[]string** |  | [optional] 
**RevokePermissions** | Pointer to **[]string** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewUpdateUserPermissionsRequest

`func NewUpdateUserPermissionsRequest(userId string, ) *UpdateUserPermissionsRequest`

NewUpdateUserPermissionsRequest instantiates a new UpdateUserPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPermissionsRequestWithDefaults

`func NewUpdateUserPermissionsRequestWithDefaults() *UpdateUserPermissionsRequest`

NewUpdateUserPermissionsRequestWithDefaults instantiates a new UpdateUserPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantPermissions

`func (o *UpdateUserPermissionsRequest) GetGrantPermissions() []string`

GetGrantPermissions returns the GrantPermissions field if non-nil, zero value otherwise.

### GetGrantPermissionsOk

`func (o *UpdateUserPermissionsRequest) GetGrantPermissionsOk() (*[]string, bool)`

GetGrantPermissionsOk returns a tuple with the GrantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantPermissions

`func (o *UpdateUserPermissionsRequest) SetGrantPermissions(v []string)`

SetGrantPermissions sets GrantPermissions field to given value.

### HasGrantPermissions

`func (o *UpdateUserPermissionsRequest) HasGrantPermissions() bool`

HasGrantPermissions returns a boolean if a field has been set.

### GetRevokePermissions

`func (o *UpdateUserPermissionsRequest) GetRevokePermissions() []string`

GetRevokePermissions returns the RevokePermissions field if non-nil, zero value otherwise.

### GetRevokePermissionsOk

`func (o *UpdateUserPermissionsRequest) GetRevokePermissionsOk() (*[]string, bool)`

GetRevokePermissionsOk returns a tuple with the RevokePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokePermissions

`func (o *UpdateUserPermissionsRequest) SetRevokePermissions(v []string)`

SetRevokePermissions sets RevokePermissions field to given value.

### HasRevokePermissions

`func (o *UpdateUserPermissionsRequest) HasRevokePermissions() bool`

HasRevokePermissions returns a boolean if a field has been set.

### GetUserId

`func (o *UpdateUserPermissionsRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateUserPermissionsRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateUserPermissionsRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


