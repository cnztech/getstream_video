# QueryMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterConditions** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Prev** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**[]SortParamRequest**](SortParamRequest.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewQueryMembersRequest

`func NewQueryMembersRequest(id string, type_ string, ) *QueryMembersRequest`

NewQueryMembersRequest instantiates a new QueryMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMembersRequestWithDefaults

`func NewQueryMembersRequestWithDefaults() *QueryMembersRequest`

NewQueryMembersRequestWithDefaults instantiates a new QueryMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterConditions

`func (o *QueryMembersRequest) GetFilterConditions() map[string]interface{}`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *QueryMembersRequest) GetFilterConditionsOk() (*map[string]interface{}, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *QueryMembersRequest) SetFilterConditions(v map[string]interface{})`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *QueryMembersRequest) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### GetId

`func (o *QueryMembersRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryMembersRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryMembersRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLimit

`func (o *QueryMembersRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryMembersRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryMembersRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryMembersRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *QueryMembersRequest) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *QueryMembersRequest) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *QueryMembersRequest) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *QueryMembersRequest) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *QueryMembersRequest) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *QueryMembersRequest) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *QueryMembersRequest) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *QueryMembersRequest) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSort

`func (o *QueryMembersRequest) GetSort() []SortParamRequest`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryMembersRequest) GetSortOk() (*[]SortParamRequest, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryMembersRequest) SetSort(v []SortParamRequest)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryMembersRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetType

`func (o *QueryMembersRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryMembersRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryMembersRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


