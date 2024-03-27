# EventNotificationSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apns** | Pointer to [**APNSRequest**](APNSRequest.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewEventNotificationSettingsRequest

`func NewEventNotificationSettingsRequest() *EventNotificationSettingsRequest`

NewEventNotificationSettingsRequest instantiates a new EventNotificationSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationSettingsRequestWithDefaults

`func NewEventNotificationSettingsRequestWithDefaults() *EventNotificationSettingsRequest`

NewEventNotificationSettingsRequestWithDefaults instantiates a new EventNotificationSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApns

`func (o *EventNotificationSettingsRequest) GetApns() APNSRequest`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *EventNotificationSettingsRequest) GetApnsOk() (*APNSRequest, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *EventNotificationSettingsRequest) SetApns(v APNSRequest)`

SetApns sets Apns field to given value.

### HasApns

`func (o *EventNotificationSettingsRequest) HasApns() bool`

HasApns returns a boolean if a field has been set.

### GetEnabled

`func (o *EventNotificationSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EventNotificationSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EventNotificationSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EventNotificationSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


