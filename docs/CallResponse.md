# CallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backstage** | **bool** |  | 
**BlockedUserIds** | **[]string** |  | 
**Cid** | **string** | The unique identifier for a call (&lt;type&gt;:&lt;id&gt;) | 
**CreatedAt** | **time.Time** | Date/time of creation | 
**CreatedBy** | [**UserResponse**](UserResponse.md) |  | 
**CurrentSessionId** | **string** |  | 
**Custom** | **map[string]interface{}** | Custom data for this object | 
**Egress** | [**EgressResponse**](EgressResponse.md) |  | 
**EndedAt** | Pointer to **time.Time** | Date/time when the call ended | [optional] 
**Id** | **string** | Call ID | 
**Ingress** | [**CallIngressResponse**](CallIngressResponse.md) |  | 
**Recording** | **bool** |  | 
**Session** | Pointer to [**CallSessionResponse**](CallSessionResponse.md) |  | [optional] 
**Settings** | [**CallSettingsResponse**](CallSettingsResponse.md) |  | 
**StartsAt** | Pointer to **time.Time** | Date/time when the call will start | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Thumbnails** | Pointer to [**ThumbnailResponse**](ThumbnailResponse.md) |  | [optional] 
**Transcribing** | **bool** |  | 
**Type** | **string** | The type of call | 
**UpdatedAt** | **time.Time** | Date/time of the last update | 

## Methods

### NewCallResponse

`func NewCallResponse(backstage bool, blockedUserIds []string, cid string, createdAt time.Time, createdBy UserResponse, currentSessionId string, custom map[string]interface{}, egress EgressResponse, id string, ingress CallIngressResponse, recording bool, settings CallSettingsResponse, transcribing bool, type_ string, updatedAt time.Time, ) *CallResponse`

NewCallResponse instantiates a new CallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallResponseWithDefaults

`func NewCallResponseWithDefaults() *CallResponse`

NewCallResponseWithDefaults instantiates a new CallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackstage

`func (o *CallResponse) GetBackstage() bool`

GetBackstage returns the Backstage field if non-nil, zero value otherwise.

### GetBackstageOk

`func (o *CallResponse) GetBackstageOk() (*bool, bool)`

GetBackstageOk returns a tuple with the Backstage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackstage

`func (o *CallResponse) SetBackstage(v bool)`

SetBackstage sets Backstage field to given value.


### GetBlockedUserIds

`func (o *CallResponse) GetBlockedUserIds() []string`

GetBlockedUserIds returns the BlockedUserIds field if non-nil, zero value otherwise.

### GetBlockedUserIdsOk

`func (o *CallResponse) GetBlockedUserIdsOk() (*[]string, bool)`

GetBlockedUserIdsOk returns a tuple with the BlockedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUserIds

`func (o *CallResponse) SetBlockedUserIds(v []string)`

SetBlockedUserIds sets BlockedUserIds field to given value.


### GetCid

`func (o *CallResponse) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *CallResponse) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *CallResponse) SetCid(v string)`

SetCid sets Cid field to given value.


### GetCreatedAt

`func (o *CallResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CallResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CallResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *CallResponse) GetCreatedBy() UserResponse`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CallResponse) GetCreatedByOk() (*UserResponse, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CallResponse) SetCreatedBy(v UserResponse)`

SetCreatedBy sets CreatedBy field to given value.


### GetCurrentSessionId

`func (o *CallResponse) GetCurrentSessionId() string`

GetCurrentSessionId returns the CurrentSessionId field if non-nil, zero value otherwise.

### GetCurrentSessionIdOk

`func (o *CallResponse) GetCurrentSessionIdOk() (*string, bool)`

GetCurrentSessionIdOk returns a tuple with the CurrentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSessionId

`func (o *CallResponse) SetCurrentSessionId(v string)`

SetCurrentSessionId sets CurrentSessionId field to given value.


### GetCustom

`func (o *CallResponse) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CallResponse) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CallResponse) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.


### GetEgress

`func (o *CallResponse) GetEgress() EgressResponse`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *CallResponse) GetEgressOk() (*EgressResponse, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *CallResponse) SetEgress(v EgressResponse)`

SetEgress sets Egress field to given value.


### GetEndedAt

`func (o *CallResponse) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CallResponse) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CallResponse) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CallResponse) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetId

`func (o *CallResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIngress

`func (o *CallResponse) GetIngress() CallIngressResponse`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *CallResponse) GetIngressOk() (*CallIngressResponse, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *CallResponse) SetIngress(v CallIngressResponse)`

SetIngress sets Ingress field to given value.


### GetRecording

`func (o *CallResponse) GetRecording() bool`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *CallResponse) GetRecordingOk() (*bool, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *CallResponse) SetRecording(v bool)`

SetRecording sets Recording field to given value.


### GetSession

`func (o *CallResponse) GetSession() CallSessionResponse`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CallResponse) GetSessionOk() (*CallSessionResponse, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CallResponse) SetSession(v CallSessionResponse)`

SetSession sets Session field to given value.

### HasSession

`func (o *CallResponse) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSettings

`func (o *CallResponse) GetSettings() CallSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CallResponse) GetSettingsOk() (*CallSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CallResponse) SetSettings(v CallSettingsResponse)`

SetSettings sets Settings field to given value.


### GetStartsAt

`func (o *CallResponse) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CallResponse) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CallResponse) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *CallResponse) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetTeam

`func (o *CallResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CallResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CallResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *CallResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetThumbnails

`func (o *CallResponse) GetThumbnails() ThumbnailResponse`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *CallResponse) GetThumbnailsOk() (*ThumbnailResponse, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *CallResponse) SetThumbnails(v ThumbnailResponse)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *CallResponse) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetTranscribing

`func (o *CallResponse) GetTranscribing() bool`

GetTranscribing returns the Transcribing field if non-nil, zero value otherwise.

### GetTranscribingOk

`func (o *CallResponse) GetTranscribingOk() (*bool, bool)`

GetTranscribingOk returns a tuple with the Transcribing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscribing

`func (o *CallResponse) SetTranscribing(v bool)`

SetTranscribing sets Transcribing field to given value.


### GetType

`func (o *CallResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CallResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CallResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *CallResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CallResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CallResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


