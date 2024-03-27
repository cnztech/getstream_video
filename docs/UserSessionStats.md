# UserSessionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | **string** |  | 
**Sdk** | **string** |  | 
**SdkVersion** | **string** |  | 
**Timeline** | Pointer to [**CallTimeline**](CallTimeline.md) |  | [optional] 
**WebrtcVersion** | **string** |  | 

## Methods

### NewUserSessionStats

`func NewUserSessionStats(os string, sdk string, sdkVersion string, webrtcVersion string, ) *UserSessionStats`

NewUserSessionStats instantiates a new UserSessionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionStatsWithDefaults

`func NewUserSessionStatsWithDefaults() *UserSessionStats`

NewUserSessionStatsWithDefaults instantiates a new UserSessionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *UserSessionStats) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *UserSessionStats) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *UserSessionStats) SetOs(v string)`

SetOs sets Os field to given value.


### GetSdk

`func (o *UserSessionStats) GetSdk() string`

GetSdk returns the Sdk field if non-nil, zero value otherwise.

### GetSdkOk

`func (o *UserSessionStats) GetSdkOk() (*string, bool)`

GetSdkOk returns a tuple with the Sdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdk

`func (o *UserSessionStats) SetSdk(v string)`

SetSdk sets Sdk field to given value.


### GetSdkVersion

`func (o *UserSessionStats) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *UserSessionStats) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *UserSessionStats) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.


### GetTimeline

`func (o *UserSessionStats) GetTimeline() CallTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *UserSessionStats) GetTimelineOk() (*CallTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *UserSessionStats) SetTimeline(v CallTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *UserSessionStats) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetWebrtcVersion

`func (o *UserSessionStats) GetWebrtcVersion() string`

GetWebrtcVersion returns the WebrtcVersion field if non-nil, zero value otherwise.

### GetWebrtcVersionOk

`func (o *UserSessionStats) GetWebrtcVersionOk() (*string, bool)`

GetWebrtcVersionOk returns a tuple with the WebrtcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebrtcVersion

`func (o *UserSessionStats) SetWebrtcVersion(v string)`

SetWebrtcVersion sets WebrtcVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


