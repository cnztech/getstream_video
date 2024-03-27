# CallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**UserRequest**](UserRequest.md) |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**Custom** | Pointer to **map[string]interface{}** |  | [optional] 
**Members** | Pointer to [**[]MemberRequest**](MemberRequest.md) |  | [optional] 
**SettingsOverride** | Pointer to [**CallSettingsRequest**](CallSettingsRequest.md) |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 

## Methods

### NewCallRequest

`func NewCallRequest() *CallRequest`

NewCallRequest instantiates a new CallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestWithDefaults

`func NewCallRequestWithDefaults() *CallRequest`

NewCallRequestWithDefaults instantiates a new CallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *CallRequest) GetCreatedBy() UserRequest`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CallRequest) GetCreatedByOk() (*UserRequest, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CallRequest) SetCreatedBy(v UserRequest)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CallRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *CallRequest) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *CallRequest) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *CallRequest) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *CallRequest) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCustom

`func (o *CallRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CallRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CallRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CallRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMembers

`func (o *CallRequest) GetMembers() []MemberRequest`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CallRequest) GetMembersOk() (*[]MemberRequest, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CallRequest) SetMembers(v []MemberRequest)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CallRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSettingsOverride

`func (o *CallRequest) GetSettingsOverride() CallSettingsRequest`

GetSettingsOverride returns the SettingsOverride field if non-nil, zero value otherwise.

### GetSettingsOverrideOk

`func (o *CallRequest) GetSettingsOverrideOk() (*CallSettingsRequest, bool)`

GetSettingsOverrideOk returns a tuple with the SettingsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsOverride

`func (o *CallRequest) SetSettingsOverride(v CallSettingsRequest)`

SetSettingsOverride sets SettingsOverride field to given value.

### HasSettingsOverride

`func (o *CallRequest) HasSettingsOverride() bool`

HasSettingsOverride returns a boolean if a field has been set.

### GetStartsAt

`func (o *CallRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CallRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CallRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *CallRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetTeam

`func (o *CallRequest) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CallRequest) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CallRequest) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *CallRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


