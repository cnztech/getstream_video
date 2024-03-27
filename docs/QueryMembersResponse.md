# QueryMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** | Duration of the request in human-readable format | 
**Members** | [**[]MemberResponse**](MemberResponse.md) |  | 
**Next** | Pointer to **string** |  | [optional] 
**Prev** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryMembersResponse

`func NewQueryMembersResponse(duration string, members []MemberResponse, ) *QueryMembersResponse`

NewQueryMembersResponse instantiates a new QueryMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMembersResponseWithDefaults

`func NewQueryMembersResponseWithDefaults() *QueryMembersResponse`

NewQueryMembersResponseWithDefaults instantiates a new QueryMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *QueryMembersResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *QueryMembersResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *QueryMembersResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMembers

`func (o *QueryMembersResponse) GetMembers() []MemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *QueryMembersResponse) GetMembersOk() (*[]MemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *QueryMembersResponse) SetMembers(v []MemberResponse)`

SetMembers sets Members field to given value.


### GetNext

`func (o *QueryMembersResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *QueryMembersResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *QueryMembersResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *QueryMembersResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *QueryMembersResponse) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *QueryMembersResponse) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *QueryMembersResponse) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *QueryMembersResponse) HasPrev() bool`

HasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


