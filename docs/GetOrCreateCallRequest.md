# GetOrCreateCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CallRequest**](CallRequest.md) |  | [optional] 
**MembersLimit** | Pointer to **int32** |  | [optional] 
**Notify** | Pointer to **bool** | if provided it sends a notification event to the members for this call | [optional] 
**Ring** | Pointer to **bool** | if provided it sends a ring event to the members for this call | [optional] 

## Methods

### NewGetOrCreateCallRequest

`func NewGetOrCreateCallRequest() *GetOrCreateCallRequest`

NewGetOrCreateCallRequest instantiates a new GetOrCreateCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrCreateCallRequestWithDefaults

`func NewGetOrCreateCallRequestWithDefaults() *GetOrCreateCallRequest`

NewGetOrCreateCallRequestWithDefaults instantiates a new GetOrCreateCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOrCreateCallRequest) GetData() CallRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrCreateCallRequest) GetDataOk() (*CallRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrCreateCallRequest) SetData(v CallRequest)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrCreateCallRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMembersLimit

`func (o *GetOrCreateCallRequest) GetMembersLimit() int32`

GetMembersLimit returns the MembersLimit field if non-nil, zero value otherwise.

### GetMembersLimitOk

`func (o *GetOrCreateCallRequest) GetMembersLimitOk() (*int32, bool)`

GetMembersLimitOk returns a tuple with the MembersLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersLimit

`func (o *GetOrCreateCallRequest) SetMembersLimit(v int32)`

SetMembersLimit sets MembersLimit field to given value.

### HasMembersLimit

`func (o *GetOrCreateCallRequest) HasMembersLimit() bool`

HasMembersLimit returns a boolean if a field has been set.

### GetNotify

`func (o *GetOrCreateCallRequest) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *GetOrCreateCallRequest) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *GetOrCreateCallRequest) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *GetOrCreateCallRequest) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetRing

`func (o *GetOrCreateCallRequest) GetRing() bool`

GetRing returns the Ring field if non-nil, zero value otherwise.

### GetRingOk

`func (o *GetOrCreateCallRequest) GetRingOk() (*bool, bool)`

GetRingOk returns a tuple with the Ring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing

`func (o *GetOrCreateCallRequest) SetRing(v bool)`

SetRing sets Ring field to given value.

### HasRing

`func (o *GetOrCreateCallRequest) HasRing() bool`

HasRing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


