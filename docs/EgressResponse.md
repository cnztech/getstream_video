# EgressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Broadcasting** | **bool** |  | 
**Hls** | Pointer to [**EgressHLSResponse**](EgressHLSResponse.md) |  | [optional] 
**Rtmps** | [**[]EgressRTMPResponse**](EgressRTMPResponse.md) |  | 

## Methods

### NewEgressResponse

`func NewEgressResponse(broadcasting bool, rtmps []EgressRTMPResponse, ) *EgressResponse`

NewEgressResponse instantiates a new EgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEgressResponseWithDefaults

`func NewEgressResponseWithDefaults() *EgressResponse`

NewEgressResponseWithDefaults instantiates a new EgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcasting

`func (o *EgressResponse) GetBroadcasting() bool`

GetBroadcasting returns the Broadcasting field if non-nil, zero value otherwise.

### GetBroadcastingOk

`func (o *EgressResponse) GetBroadcastingOk() (*bool, bool)`

GetBroadcastingOk returns a tuple with the Broadcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasting

`func (o *EgressResponse) SetBroadcasting(v bool)`

SetBroadcasting sets Broadcasting field to given value.


### GetHls

`func (o *EgressResponse) GetHls() EgressHLSResponse`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *EgressResponse) GetHlsOk() (*EgressHLSResponse, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *EgressResponse) SetHls(v EgressHLSResponse)`

SetHls sets Hls field to given value.

### HasHls

`func (o *EgressResponse) HasHls() bool`

HasHls returns a boolean if a field has been set.

### GetRtmps

`func (o *EgressResponse) GetRtmps() []EgressRTMPResponse`

GetRtmps returns the Rtmps field if non-nil, zero value otherwise.

### GetRtmpsOk

`func (o *EgressResponse) GetRtmpsOk() (*[]EgressRTMPResponse, bool)`

GetRtmpsOk returns a tuple with the Rtmps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtmps

`func (o *EgressResponse) SetRtmps(v []EgressRTMPResponse)`

SetRtmps sets Rtmps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


