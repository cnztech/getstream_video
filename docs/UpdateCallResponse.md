# UpdateCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | [**CallResponse**](CallResponse.md) |  | 
**Duration** | **string** |  | 
**Members** | [**[]MemberResponse**](MemberResponse.md) |  | 
**Membership** | Pointer to [**MemberResponse**](MemberResponse.md) |  | [optional] 
**OwnCapabilities** | [**[]OwnCapability**](OwnCapability.md) |  | 

## Methods

### NewUpdateCallResponse

`func NewUpdateCallResponse(call CallResponse, duration string, members []MemberResponse, ownCapabilities []OwnCapability, ) *UpdateCallResponse`

NewUpdateCallResponse instantiates a new UpdateCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallResponseWithDefaults

`func NewUpdateCallResponseWithDefaults() *UpdateCallResponse`

NewUpdateCallResponseWithDefaults instantiates a new UpdateCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *UpdateCallResponse) GetCall() CallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *UpdateCallResponse) GetCallOk() (*CallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *UpdateCallResponse) SetCall(v CallResponse)`

SetCall sets Call field to given value.


### GetDuration

`func (o *UpdateCallResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateCallResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateCallResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMembers

`func (o *UpdateCallResponse) GetMembers() []MemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateCallResponse) GetMembersOk() (*[]MemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateCallResponse) SetMembers(v []MemberResponse)`

SetMembers sets Members field to given value.


### GetMembership

`func (o *UpdateCallResponse) GetMembership() MemberResponse`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *UpdateCallResponse) GetMembershipOk() (*MemberResponse, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *UpdateCallResponse) SetMembership(v MemberResponse)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *UpdateCallResponse) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetOwnCapabilities

`func (o *UpdateCallResponse) GetOwnCapabilities() []OwnCapability`

GetOwnCapabilities returns the OwnCapabilities field if non-nil, zero value otherwise.

### GetOwnCapabilitiesOk

`func (o *UpdateCallResponse) GetOwnCapabilitiesOk() (*[]OwnCapability, bool)`

GetOwnCapabilitiesOk returns a tuple with the OwnCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnCapabilities

`func (o *UpdateCallResponse) SetOwnCapabilities(v []OwnCapability)`

SetOwnCapabilities sets OwnCapabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


