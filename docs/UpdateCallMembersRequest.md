# UpdateCallMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveMembers** | Pointer to **[]string** | List of userID to remove | [optional] 
**UpdateMembers** | Pointer to [**[]MemberRequest**](MemberRequest.md) | List of members to update or insert | [optional] 

## Methods

### NewUpdateCallMembersRequest

`func NewUpdateCallMembersRequest() *UpdateCallMembersRequest`

NewUpdateCallMembersRequest instantiates a new UpdateCallMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallMembersRequestWithDefaults

`func NewUpdateCallMembersRequestWithDefaults() *UpdateCallMembersRequest`

NewUpdateCallMembersRequestWithDefaults instantiates a new UpdateCallMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveMembers

`func (o *UpdateCallMembersRequest) GetRemoveMembers() []string`

GetRemoveMembers returns the RemoveMembers field if non-nil, zero value otherwise.

### GetRemoveMembersOk

`func (o *UpdateCallMembersRequest) GetRemoveMembersOk() (*[]string, bool)`

GetRemoveMembersOk returns a tuple with the RemoveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveMembers

`func (o *UpdateCallMembersRequest) SetRemoveMembers(v []string)`

SetRemoveMembers sets RemoveMembers field to given value.

### HasRemoveMembers

`func (o *UpdateCallMembersRequest) HasRemoveMembers() bool`

HasRemoveMembers returns a boolean if a field has been set.

### GetUpdateMembers

`func (o *UpdateCallMembersRequest) GetUpdateMembers() []MemberRequest`

GetUpdateMembers returns the UpdateMembers field if non-nil, zero value otherwise.

### GetUpdateMembersOk

`func (o *UpdateCallMembersRequest) GetUpdateMembersOk() (*[]MemberRequest, bool)`

GetUpdateMembersOk returns a tuple with the UpdateMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMembers

`func (o *UpdateCallMembersRequest) SetUpdateMembers(v []MemberRequest)`

SetUpdateMembers sets UpdateMembers field to given value.

### HasUpdateMembers

`func (o *UpdateCallMembersRequest) HasUpdateMembers() bool`

HasUpdateMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


