# QueryCallsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | [**[]CallStateResponseFields**](CallStateResponseFields.md) |  | 
**Duration** | **string** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Prev** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryCallsResponse

`func NewQueryCallsResponse(calls []CallStateResponseFields, duration string, ) *QueryCallsResponse`

NewQueryCallsResponse instantiates a new QueryCallsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCallsResponseWithDefaults

`func NewQueryCallsResponseWithDefaults() *QueryCallsResponse`

NewQueryCallsResponseWithDefaults instantiates a new QueryCallsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *QueryCallsResponse) GetCalls() []CallStateResponseFields`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *QueryCallsResponse) GetCallsOk() (*[]CallStateResponseFields, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *QueryCallsResponse) SetCalls(v []CallStateResponseFields)`

SetCalls sets Calls field to given value.


### GetDuration

`func (o *QueryCallsResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *QueryCallsResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *QueryCallsResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetNext

`func (o *QueryCallsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *QueryCallsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *QueryCallsResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *QueryCallsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *QueryCallsResponse) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *QueryCallsResponse) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *QueryCallsResponse) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *QueryCallsResponse) HasPrev() bool`

HasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


