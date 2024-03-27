# MemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **map[string]interface{}** | Custom data for this object | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewMemberRequest

`func NewMemberRequest(userId string, ) *MemberRequest`

NewMemberRequest instantiates a new MemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberRequestWithDefaults

`func NewMemberRequestWithDefaults() *MemberRequest`

NewMemberRequestWithDefaults instantiates a new MemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *MemberRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *MemberRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *MemberRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *MemberRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetRole

`func (o *MemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserId

`func (o *MemberRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


