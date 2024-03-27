# UserStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | [**UserInfoResponse**](UserInfoResponse.md) |  | 
**SessionStats** | [**map[string]UserSessionStats**](UserSessionStats.md) |  | 

## Methods

### NewUserStats

`func NewUserStats(info UserInfoResponse, sessionStats map[string]UserSessionStats, ) *UserStats`

NewUserStats instantiates a new UserStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatsWithDefaults

`func NewUserStatsWithDefaults() *UserStats`

NewUserStatsWithDefaults instantiates a new UserStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *UserStats) GetInfo() UserInfoResponse`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UserStats) GetInfoOk() (*UserInfoResponse, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UserStats) SetInfo(v UserInfoResponse)`

SetInfo sets Info field to given value.


### GetSessionStats

`func (o *UserStats) GetSessionStats() map[string]UserSessionStats`

GetSessionStats returns the SessionStats field if non-nil, zero value otherwise.

### GetSessionStatsOk

`func (o *UserStats) GetSessionStatsOk() (*map[string]UserSessionStats, bool)`

GetSessionStatsOk returns a tuple with the SessionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStats

`func (o *UserStats) SetSessionStats(v map[string]UserSessionStats)`

SetSessionStats sets SessionStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


