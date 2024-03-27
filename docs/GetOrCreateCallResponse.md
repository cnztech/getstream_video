# GetOrCreateCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | [**CallResponse**](CallResponse.md) |  | 
**Created** | **bool** |  | 
**Duration** | **string** |  | 
**Members** | [**[]MemberResponse**](MemberResponse.md) |  | 
**Membership** | Pointer to [**MemberResponse**](MemberResponse.md) |  | [optional] 
**OwnCapabilities** | [**[]OwnCapability**](OwnCapability.md) |  | 

## Methods

### NewGetOrCreateCallResponse

`func NewGetOrCreateCallResponse(call CallResponse, created bool, duration string, members []MemberResponse, ownCapabilities []OwnCapability, ) *GetOrCreateCallResponse`

NewGetOrCreateCallResponse instantiates a new GetOrCreateCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrCreateCallResponseWithDefaults

`func NewGetOrCreateCallResponseWithDefaults() *GetOrCreateCallResponse`

NewGetOrCreateCallResponseWithDefaults instantiates a new GetOrCreateCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *GetOrCreateCallResponse) GetCall() CallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *GetOrCreateCallResponse) GetCallOk() (*CallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *GetOrCreateCallResponse) SetCall(v CallResponse)`

SetCall sets Call field to given value.


### GetCreated

`func (o *GetOrCreateCallResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetOrCreateCallResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetOrCreateCallResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.


### GetDuration

`func (o *GetOrCreateCallResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetOrCreateCallResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetOrCreateCallResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMembers

`func (o *GetOrCreateCallResponse) GetMembers() []MemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetOrCreateCallResponse) GetMembersOk() (*[]MemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetOrCreateCallResponse) SetMembers(v []MemberResponse)`

SetMembers sets Members field to given value.


### GetMembership

`func (o *GetOrCreateCallResponse) GetMembership() MemberResponse`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *GetOrCreateCallResponse) GetMembershipOk() (*MemberResponse, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *GetOrCreateCallResponse) SetMembership(v MemberResponse)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *GetOrCreateCallResponse) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetOwnCapabilities

`func (o *GetOrCreateCallResponse) GetOwnCapabilities() []OwnCapability`

GetOwnCapabilities returns the OwnCapabilities field if non-nil, zero value otherwise.

### GetOwnCapabilitiesOk

`func (o *GetOrCreateCallResponse) GetOwnCapabilitiesOk() (*[]OwnCapability, bool)`

GetOwnCapabilitiesOk returns a tuple with the OwnCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnCapabilities

`func (o *GetOrCreateCallResponse) SetOwnCapabilities(v []OwnCapability)`

SetOwnCapabilities sets OwnCapabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


