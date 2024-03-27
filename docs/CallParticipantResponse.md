# CallParticipantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinedAt** | **time.Time** |  | 
**Role** | **string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 
**UserSessionId** | **string** |  | 

## Methods

### NewCallParticipantResponse

`func NewCallParticipantResponse(joinedAt time.Time, role string, user UserResponse, userSessionId string, ) *CallParticipantResponse`

NewCallParticipantResponse instantiates a new CallParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallParticipantResponseWithDefaults

`func NewCallParticipantResponseWithDefaults() *CallParticipantResponse`

NewCallParticipantResponseWithDefaults instantiates a new CallParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinedAt

`func (o *CallParticipantResponse) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *CallParticipantResponse) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *CallParticipantResponse) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetRole

`func (o *CallParticipantResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CallParticipantResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CallParticipantResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetUser

`func (o *CallParticipantResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CallParticipantResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CallParticipantResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.


### GetUserSessionId

`func (o *CallParticipantResponse) GetUserSessionId() string`

GetUserSessionId returns the UserSessionId field if non-nil, zero value otherwise.

### GetUserSessionIdOk

`func (o *CallParticipantResponse) GetUserSessionIdOk() (*string, bool)`

GetUserSessionIdOk returns a tuple with the UserSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionId

`func (o *CallParticipantResponse) SetUserSessionId(v string)`

SetUserSessionId sets UserSessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


