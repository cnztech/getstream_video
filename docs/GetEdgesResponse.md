# GetEdgesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** | Duration of the request in human-readable format | 
**Edges** | [**[]EdgeResponse**](EdgeResponse.md) |  | 

## Methods

### NewGetEdgesResponse

`func NewGetEdgesResponse(duration string, edges []EdgeResponse, ) *GetEdgesResponse`

NewGetEdgesResponse instantiates a new GetEdgesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEdgesResponseWithDefaults

`func NewGetEdgesResponseWithDefaults() *GetEdgesResponse`

NewGetEdgesResponseWithDefaults instantiates a new GetEdgesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *GetEdgesResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetEdgesResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetEdgesResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetEdges

`func (o *GetEdgesResponse) GetEdges() []EdgeResponse`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GetEdgesResponse) GetEdgesOk() (*[]EdgeResponse, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GetEdgesResponse) SetEdges(v []EdgeResponse)`

SetEdges sets Edges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


