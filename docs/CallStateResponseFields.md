# CallStateResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | [**CallResponse**](CallResponse.md) |  | 
**Members** | [**[]MemberResponse**](MemberResponse.md) | List of call members | 
**Membership** | Pointer to [**MemberResponse**](MemberResponse.md) |  | [optional] 
**OwnCapabilities** | [**[]OwnCapability**](OwnCapability.md) |  | 

## Methods

### NewCallStateResponseFields

`func NewCallStateResponseFields(call CallResponse, members []MemberResponse, ownCapabilities []OwnCapability, ) *CallStateResponseFields`

NewCallStateResponseFields instantiates a new CallStateResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStateResponseFieldsWithDefaults

`func NewCallStateResponseFieldsWithDefaults() *CallStateResponseFields`

NewCallStateResponseFieldsWithDefaults instantiates a new CallStateResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *CallStateResponseFields) GetCall() CallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CallStateResponseFields) GetCallOk() (*CallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CallStateResponseFields) SetCall(v CallResponse)`

SetCall sets Call field to given value.


### GetMembers

`func (o *CallStateResponseFields) GetMembers() []MemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CallStateResponseFields) GetMembersOk() (*[]MemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CallStateResponseFields) SetMembers(v []MemberResponse)`

SetMembers sets Members field to given value.


### GetMembership

`func (o *CallStateResponseFields) GetMembership() MemberResponse`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *CallStateResponseFields) GetMembershipOk() (*MemberResponse, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *CallStateResponseFields) SetMembership(v MemberResponse)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *CallStateResponseFields) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetOwnCapabilities

`func (o *CallStateResponseFields) GetOwnCapabilities() []OwnCapability`

GetOwnCapabilities returns the OwnCapabilities field if non-nil, zero value otherwise.

### GetOwnCapabilitiesOk

`func (o *CallStateResponseFields) GetOwnCapabilitiesOk() (*[]OwnCapability, bool)`

GetOwnCapabilitiesOk returns a tuple with the OwnCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnCapabilities

`func (o *CallStateResponseFields) SetOwnCapabilities(v []OwnCapability)`

SetOwnCapabilities sets OwnCapabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


