# EventNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apns** | [**APNS**](APNS.md) |  | 
**Enabled** | **bool** |  | 

## Methods

### NewEventNotificationSettings

`func NewEventNotificationSettings(apns APNS, enabled bool, ) *EventNotificationSettings`

NewEventNotificationSettings instantiates a new EventNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationSettingsWithDefaults

`func NewEventNotificationSettingsWithDefaults() *EventNotificationSettings`

NewEventNotificationSettingsWithDefaults instantiates a new EventNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApns

`func (o *EventNotificationSettings) GetApns() APNS`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *EventNotificationSettings) GetApnsOk() (*APNS, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *EventNotificationSettings) SetApns(v APNS)`

SetApns sets Apns field to given value.


### GetEnabled

`func (o *EventNotificationSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EventNotificationSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EventNotificationSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


