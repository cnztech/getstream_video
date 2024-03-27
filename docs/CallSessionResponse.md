# CallSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedBy** | [**map[string]time.Time**](time.Time.md) |  | 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**LiveEndedAt** | Pointer to **time.Time** |  | [optional] 
**LiveStartedAt** | Pointer to **time.Time** |  | [optional] 
**Participants** | [**[]CallParticipantResponse**](CallParticipantResponse.md) |  | 
**ParticipantsCountByRole** | **map[string]int32** |  | 
**RejectedBy** | [**map[string]time.Time**](time.Time.md) |  | 
**StartedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCallSessionResponse

`func NewCallSessionResponse(acceptedBy map[string]time.Time, id string, participants []CallParticipantResponse, participantsCountByRole map[string]int32, rejectedBy map[string]time.Time, ) *CallSessionResponse`

NewCallSessionResponse instantiates a new CallSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallSessionResponseWithDefaults

`func NewCallSessionResponseWithDefaults() *CallSessionResponse`

NewCallSessionResponseWithDefaults instantiates a new CallSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedBy

`func (o *CallSessionResponse) GetAcceptedBy() map[string]time.Time`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *CallSessionResponse) GetAcceptedByOk() (*map[string]time.Time, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *CallSessionResponse) SetAcceptedBy(v map[string]time.Time)`

SetAcceptedBy sets AcceptedBy field to given value.


### GetEndedAt

`func (o *CallSessionResponse) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CallSessionResponse) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CallSessionResponse) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CallSessionResponse) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetId

`func (o *CallSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallSessionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLiveEndedAt

`func (o *CallSessionResponse) GetLiveEndedAt() time.Time`

GetLiveEndedAt returns the LiveEndedAt field if non-nil, zero value otherwise.

### GetLiveEndedAtOk

`func (o *CallSessionResponse) GetLiveEndedAtOk() (*time.Time, bool)`

GetLiveEndedAtOk returns a tuple with the LiveEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEndedAt

`func (o *CallSessionResponse) SetLiveEndedAt(v time.Time)`

SetLiveEndedAt sets LiveEndedAt field to given value.

### HasLiveEndedAt

`func (o *CallSessionResponse) HasLiveEndedAt() bool`

HasLiveEndedAt returns a boolean if a field has been set.

### GetLiveStartedAt

`func (o *CallSessionResponse) GetLiveStartedAt() time.Time`

GetLiveStartedAt returns the LiveStartedAt field if non-nil, zero value otherwise.

### GetLiveStartedAtOk

`func (o *CallSessionResponse) GetLiveStartedAtOk() (*time.Time, bool)`

GetLiveStartedAtOk returns a tuple with the LiveStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStartedAt

`func (o *CallSessionResponse) SetLiveStartedAt(v time.Time)`

SetLiveStartedAt sets LiveStartedAt field to given value.

### HasLiveStartedAt

`func (o *CallSessionResponse) HasLiveStartedAt() bool`

HasLiveStartedAt returns a boolean if a field has been set.

### GetParticipants

`func (o *CallSessionResponse) GetParticipants() []CallParticipantResponse`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *CallSessionResponse) GetParticipantsOk() (*[]CallParticipantResponse, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *CallSessionResponse) SetParticipants(v []CallParticipantResponse)`

SetParticipants sets Participants field to given value.


### GetParticipantsCountByRole

`func (o *CallSessionResponse) GetParticipantsCountByRole() map[string]int32`

GetParticipantsCountByRole returns the ParticipantsCountByRole field if non-nil, zero value otherwise.

### GetParticipantsCountByRoleOk

`func (o *CallSessionResponse) GetParticipantsCountByRoleOk() (*map[string]int32, bool)`

GetParticipantsCountByRoleOk returns a tuple with the ParticipantsCountByRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantsCountByRole

`func (o *CallSessionResponse) SetParticipantsCountByRole(v map[string]int32)`

SetParticipantsCountByRole sets ParticipantsCountByRole field to given value.


### GetRejectedBy

`func (o *CallSessionResponse) GetRejectedBy() map[string]time.Time`

GetRejectedBy returns the RejectedBy field if non-nil, zero value otherwise.

### GetRejectedByOk

`func (o *CallSessionResponse) GetRejectedByOk() (*map[string]time.Time, bool)`

GetRejectedByOk returns a tuple with the RejectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedBy

`func (o *CallSessionResponse) SetRejectedBy(v map[string]time.Time)`

SetRejectedBy sets RejectedBy field to given value.


### GetStartedAt

`func (o *CallSessionResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CallSessionResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CallSessionResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CallSessionResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


