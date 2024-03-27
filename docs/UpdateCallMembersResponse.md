# UpdateCallMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** | Duration of the request in human-readable format | 
**Members** | [**[]MemberResponse**](MemberResponse.md) |  | 

## Methods

### NewUpdateCallMembersResponse

`func NewUpdateCallMembersResponse(duration string, members []MemberResponse, ) *UpdateCallMembersResponse`

NewUpdateCallMembersResponse instantiates a new UpdateCallMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallMembersResponseWithDefaults

`func NewUpdateCallMembersResponseWithDefaults() *UpdateCallMembersResponse`

NewUpdateCallMembersResponseWithDefaults instantiates a new UpdateCallMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *UpdateCallMembersResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateCallMembersResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateCallMembersResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMembers

`func (o *UpdateCallMembersResponse) GetMembers() []MemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateCallMembersResponse) GetMembersOk() (*[]MemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateCallMembersResponse) SetMembers(v []MemberResponse)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


