# ListDevicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]Device**](Device.md) | List of devices | 
**Duration** | **string** |  | 

## Methods

### NewListDevicesResponse

`func NewListDevicesResponse(devices []Device, duration string, ) *ListDevicesResponse`

NewListDevicesResponse instantiates a new ListDevicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDevicesResponseWithDefaults

`func NewListDevicesResponseWithDefaults() *ListDevicesResponse`

NewListDevicesResponseWithDefaults instantiates a new ListDevicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *ListDevicesResponse) GetDevices() []Device`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ListDevicesResponse) GetDevicesOk() (*[]Device, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ListDevicesResponse) SetDevices(v []Device)`

SetDevices sets Devices field to given value.


### GetDuration

`func (o *ListDevicesResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListDevicesResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListDevicesResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


