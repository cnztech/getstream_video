# NotificationSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallLiveStarted** | Pointer to [**EventNotificationSettingsRequest**](EventNotificationSettingsRequest.md) |  | [optional] 
**CallNotification** | Pointer to [**EventNotificationSettingsRequest**](EventNotificationSettingsRequest.md) |  | [optional] 
**CallRing** | Pointer to [**EventNotificationSettingsRequest**](EventNotificationSettingsRequest.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SessionStarted** | Pointer to [**EventNotificationSettingsRequest**](EventNotificationSettingsRequest.md) |  | [optional] 

## Methods

### NewNotificationSettingsRequest

`func NewNotificationSettingsRequest() *NotificationSettingsRequest`

NewNotificationSettingsRequest instantiates a new NotificationSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsRequestWithDefaults

`func NewNotificationSettingsRequestWithDefaults() *NotificationSettingsRequest`

NewNotificationSettingsRequestWithDefaults instantiates a new NotificationSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallLiveStarted

`func (o *NotificationSettingsRequest) GetCallLiveStarted() EventNotificationSettingsRequest`

GetCallLiveStarted returns the CallLiveStarted field if non-nil, zero value otherwise.

### GetCallLiveStartedOk

`func (o *NotificationSettingsRequest) GetCallLiveStartedOk() (*EventNotificationSettingsRequest, bool)`

GetCallLiveStartedOk returns a tuple with the CallLiveStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLiveStarted

`func (o *NotificationSettingsRequest) SetCallLiveStarted(v EventNotificationSettingsRequest)`

SetCallLiveStarted sets CallLiveStarted field to given value.

### HasCallLiveStarted

`func (o *NotificationSettingsRequest) HasCallLiveStarted() bool`

HasCallLiveStarted returns a boolean if a field has been set.

### GetCallNotification

`func (o *NotificationSettingsRequest) GetCallNotification() EventNotificationSettingsRequest`

GetCallNotification returns the CallNotification field if non-nil, zero value otherwise.

### GetCallNotificationOk

`func (o *NotificationSettingsRequest) GetCallNotificationOk() (*EventNotificationSettingsRequest, bool)`

GetCallNotificationOk returns a tuple with the CallNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallNotification

`func (o *NotificationSettingsRequest) SetCallNotification(v EventNotificationSettingsRequest)`

SetCallNotification sets CallNotification field to given value.

### HasCallNotification

`func (o *NotificationSettingsRequest) HasCallNotification() bool`

HasCallNotification returns a boolean if a field has been set.

### GetCallRing

`func (o *NotificationSettingsRequest) GetCallRing() EventNotificationSettingsRequest`

GetCallRing returns the CallRing field if non-nil, zero value otherwise.

### GetCallRingOk

`func (o *NotificationSettingsRequest) GetCallRingOk() (*EventNotificationSettingsRequest, bool)`

GetCallRingOk returns a tuple with the CallRing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRing

`func (o *NotificationSettingsRequest) SetCallRing(v EventNotificationSettingsRequest)`

SetCallRing sets CallRing field to given value.

### HasCallRing

`func (o *NotificationSettingsRequest) HasCallRing() bool`

HasCallRing returns a boolean if a field has been set.

### GetEnabled

`func (o *NotificationSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSessionStarted

`func (o *NotificationSettingsRequest) GetSessionStarted() EventNotificationSettingsRequest`

GetSessionStarted returns the SessionStarted field if non-nil, zero value otherwise.

### GetSessionStartedOk

`func (o *NotificationSettingsRequest) GetSessionStartedOk() (*EventNotificationSettingsRequest, bool)`

GetSessionStartedOk returns a tuple with the SessionStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStarted

`func (o *NotificationSettingsRequest) SetSessionStarted(v EventNotificationSettingsRequest)`

SetSessionStarted sets SessionStarted field to given value.

### HasSessionStarted

`func (o *NotificationSettingsRequest) HasSessionStarted() bool`

HasSessionStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


