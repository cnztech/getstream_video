# ListCallTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallTypes** | [**map[string]CallTypeResponse**](CallTypeResponse.md) |  | 
**Duration** | **string** |  | 

## Methods

### NewListCallTypeResponse

`func NewListCallTypeResponse(callTypes map[string]CallTypeResponse, duration string, ) *ListCallTypeResponse`

NewListCallTypeResponse instantiates a new ListCallTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCallTypeResponseWithDefaults

`func NewListCallTypeResponseWithDefaults() *ListCallTypeResponse`

NewListCallTypeResponseWithDefaults instantiates a new ListCallTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallTypes

`func (o *ListCallTypeResponse) GetCallTypes() map[string]CallTypeResponse`

GetCallTypes returns the CallTypes field if non-nil, zero value otherwise.

### GetCallTypesOk

`func (o *ListCallTypeResponse) GetCallTypesOk() (*map[string]CallTypeResponse, bool)`

GetCallTypesOk returns a tuple with the CallTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTypes

`func (o *ListCallTypeResponse) SetCallTypes(v map[string]CallTypeResponse)`

SetCallTypes sets CallTypes field to given value.


### GetDuration

`func (o *ListCallTypeResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListCallTypeResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListCallTypeResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


