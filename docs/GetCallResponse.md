# GetCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | [**CallResponse**](CallResponse.md) |  | 
**Duration** | **string** |  | 
**Members** | [**[]MemberResponse**](MemberResponse.md) |  | 
**Membership** | Pointer to [**MemberResponse**](MemberResponse.md) |  | [optional] 
**OwnCapabilities** | [**[]OwnCapability**](OwnCapability.md) |  | 

## Methods

### NewGetCallResponse

`func NewGetCallResponse(call CallResponse, duration string, members []MemberResponse, ownCapabilities []OwnCapability, ) *GetCallResponse`

NewGetCallResponse instantiates a new GetCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCallResponseWithDefaults

`func NewGetCallResponseWithDefaults() *GetCallResponse`

NewGetCallResponseWithDefaults instantiates a new GetCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *GetCallResponse) GetCall() CallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *GetCallResponse) GetCallOk() (*CallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *GetCallResponse) SetCall(v CallResponse)`

SetCall sets Call field to given value.


### GetDuration

`func (o *GetCallResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetCallResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetCallResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMembers

`func (o *GetCallResponse) GetMembers() []MemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetCallResponse) GetMembersOk() (*[]MemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetCallResponse) SetMembers(v []MemberResponse)`

SetMembers sets Members field to given value.


### GetMembership

`func (o *GetCallResponse) GetMembership() MemberResponse`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *GetCallResponse) GetMembershipOk() (*MemberResponse, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *GetCallResponse) SetMembership(v MemberResponse)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *GetCallResponse) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetOwnCapabilities

`func (o *GetCallResponse) GetOwnCapabilities() []OwnCapability`

GetOwnCapabilities returns the OwnCapabilities field if non-nil, zero value otherwise.

### GetOwnCapabilitiesOk

`func (o *GetCallResponse) GetOwnCapabilitiesOk() (*[]OwnCapability, bool)`

GetOwnCapabilitiesOk returns a tuple with the OwnCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnCapabilities

`func (o *GetCallResponse) SetOwnCapabilities(v []OwnCapability)`

SetOwnCapabilities sets OwnCapabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


