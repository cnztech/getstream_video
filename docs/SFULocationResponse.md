# SFULocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | [**Coordinates**](Coordinates.md) |  | 
**Datacenter** | **string** |  | 
**Id** | **string** |  | 
**Location** | [**Location**](Location.md) |  | 

## Methods

### NewSFULocationResponse

`func NewSFULocationResponse(coordinates Coordinates, datacenter string, id string, location Location, ) *SFULocationResponse`

NewSFULocationResponse instantiates a new SFULocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFULocationResponseWithDefaults

`func NewSFULocationResponseWithDefaults() *SFULocationResponse`

NewSFULocationResponseWithDefaults instantiates a new SFULocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *SFULocationResponse) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *SFULocationResponse) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *SFULocationResponse) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.


### GetDatacenter

`func (o *SFULocationResponse) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *SFULocationResponse) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *SFULocationResponse) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetId

`func (o *SFULocationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SFULocationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SFULocationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *SFULocationResponse) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SFULocationResponse) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SFULocationResponse) SetLocation(v Location)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


