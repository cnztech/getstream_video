# SortParamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **int32** | Direction of sorting, -1 for descending, 1 for ascending | [optional] 
**Field** | Pointer to **string** | Name of field to sort by | [optional] 

## Methods

### NewSortParamRequest

`func NewSortParamRequest() *SortParamRequest`

NewSortParamRequest instantiates a new SortParamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortParamRequestWithDefaults

`func NewSortParamRequestWithDefaults() *SortParamRequest`

NewSortParamRequestWithDefaults instantiates a new SortParamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SortParamRequest) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SortParamRequest) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SortParamRequest) SetDirection(v int32)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SortParamRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetField

`func (o *SortParamRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SortParamRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SortParamRequest) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SortParamRequest) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


