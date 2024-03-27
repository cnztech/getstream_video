# NotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallLiveStarted** | [**EventNotificationSettings**](EventNotificationSettings.md) |  | 
**CallNotification** | [**EventNotificationSettings**](EventNotificationSettings.md) |  | 
**CallRing** | [**EventNotificationSettings**](EventNotificationSettings.md) |  | 
**Enabled** | **bool** |  | 
**SessionStarted** | [**EventNotificationSettings**](EventNotificationSettings.md) |  | 

## Methods

### NewNotificationSettings

`func NewNotificationSettings(callLiveStarted EventNotificationSettings, callNotification EventNotificationSettings, callRing EventNotificationSettings, enabled bool, sessionStarted EventNotificationSettings, ) *NotificationSettings`

NewNotificationSettings instantiates a new NotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsWithDefaults

`func NewNotificationSettingsWithDefaults() *NotificationSettings`

NewNotificationSettingsWithDefaults instantiates a new NotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallLiveStarted

`func (o *NotificationSettings) GetCallLiveStarted() EventNotificationSettings`

GetCallLiveStarted returns the CallLiveStarted field if non-nil, zero value otherwise.

### GetCallLiveStartedOk

`func (o *NotificationSettings) GetCallLiveStartedOk() (*EventNotificationSettings, bool)`

GetCallLiveStartedOk returns a tuple with the CallLiveStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLiveStarted

`func (o *NotificationSettings) SetCallLiveStarted(v EventNotificationSettings)`

SetCallLiveStarted sets CallLiveStarted field to given value.


### GetCallNotification

`func (o *NotificationSettings) GetCallNotification() EventNotificationSettings`

GetCallNotification returns the CallNotification field if non-nil, zero value otherwise.

### GetCallNotificationOk

`func (o *NotificationSettings) GetCallNotificationOk() (*EventNotificationSettings, bool)`

GetCallNotificationOk returns a tuple with the CallNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallNotification

`func (o *NotificationSettings) SetCallNotification(v EventNotificationSettings)`

SetCallNotification sets CallNotification field to given value.


### GetCallRing

`func (o *NotificationSettings) GetCallRing() EventNotificationSettings`

GetCallRing returns the CallRing field if non-nil, zero value otherwise.

### GetCallRingOk

`func (o *NotificationSettings) GetCallRingOk() (*EventNotificationSettings, bool)`

GetCallRingOk returns a tuple with the CallRing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRing

`func (o *NotificationSettings) SetCallRing(v EventNotificationSettings)`

SetCallRing sets CallRing field to given value.


### GetEnabled

`func (o *NotificationSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSessionStarted

`func (o *NotificationSettings) GetSessionStarted() EventNotificationSettings`

GetSessionStarted returns the SessionStarted field if non-nil, zero value otherwise.

### GetSessionStartedOk

`func (o *NotificationSettings) GetSessionStartedOk() (*EventNotificationSettings, bool)`

GetSessionStartedOk returns a tuple with the SessionStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStarted

`func (o *NotificationSettings) SetSessionStarted(v EventNotificationSettings)`

SetSessionStarted sets SessionStarted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


