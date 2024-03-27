# QueryCallsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterConditions** | Pointer to **map[string]interface{}** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Prev** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**[]SortParamRequest**](SortParamRequest.md) |  | [optional] 
**Watch** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueryCallsRequest

`func NewQueryCallsRequest() *QueryCallsRequest`

NewQueryCallsRequest instantiates a new QueryCallsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCallsRequestWithDefaults

`func NewQueryCallsRequestWithDefaults() *QueryCallsRequest`

NewQueryCallsRequestWithDefaults instantiates a new QueryCallsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterConditions

`func (o *QueryCallsRequest) GetFilterConditions() map[string]interface{}`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *QueryCallsRequest) GetFilterConditionsOk() (*map[string]interface{}, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *QueryCallsRequest) SetFilterConditions(v map[string]interface{})`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *QueryCallsRequest) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### GetLimit

`func (o *QueryCallsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryCallsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryCallsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryCallsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *QueryCallsRequest) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *QueryCallsRequest) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *QueryCallsRequest) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *QueryCallsRequest) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *QueryCallsRequest) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *QueryCallsRequest) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *QueryCallsRequest) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *QueryCallsRequest) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSort

`func (o *QueryCallsRequest) GetSort() []SortParamRequest`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryCallsRequest) GetSortOk() (*[]SortParamRequest, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryCallsRequest) SetSort(v []SortParamRequest)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryCallsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetWatch

`func (o *QueryCallsRequest) GetWatch() bool`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *QueryCallsRequest) GetWatchOk() (*bool, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *QueryCallsRequest) SetWatch(v bool)`

SetWatch sets Watch field to given value.

### HasWatch

`func (o *QueryCallsRequest) HasWatch() bool`

HasWatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


