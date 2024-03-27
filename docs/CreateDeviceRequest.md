# CreateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PushProvider** | Pointer to **string** |  | [optional] 
**PushProviderName** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserRequest**](UserRequest.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**VoipToken** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateDeviceRequest

`func NewCreateDeviceRequest() *CreateDeviceRequest`

NewCreateDeviceRequest instantiates a new CreateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceRequestWithDefaults

`func NewCreateDeviceRequestWithDefaults() *CreateDeviceRequest`

NewCreateDeviceRequestWithDefaults instantiates a new CreateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateDeviceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateDeviceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateDeviceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateDeviceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPushProvider

`func (o *CreateDeviceRequest) GetPushProvider() string`

GetPushProvider returns the PushProvider field if non-nil, zero value otherwise.

### GetPushProviderOk

`func (o *CreateDeviceRequest) GetPushProviderOk() (*string, bool)`

GetPushProviderOk returns a tuple with the PushProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProvider

`func (o *CreateDeviceRequest) SetPushProvider(v string)`

SetPushProvider sets PushProvider field to given value.

### HasPushProvider

`func (o *CreateDeviceRequest) HasPushProvider() bool`

HasPushProvider returns a boolean if a field has been set.

### GetPushProviderName

`func (o *CreateDeviceRequest) GetPushProviderName() string`

GetPushProviderName returns the PushProviderName field if non-nil, zero value otherwise.

### GetPushProviderNameOk

`func (o *CreateDeviceRequest) GetPushProviderNameOk() (*string, bool)`

GetPushProviderNameOk returns a tuple with the PushProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushProviderName

`func (o *CreateDeviceRequest) SetPushProviderName(v string)`

SetPushProviderName sets PushProviderName field to given value.

### HasPushProviderName

`func (o *CreateDeviceRequest) HasPushProviderName() bool`

HasPushProviderName returns a boolean if a field has been set.

### GetUser

`func (o *CreateDeviceRequest) GetUser() UserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateDeviceRequest) GetUserOk() (*UserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateDeviceRequest) SetUser(v UserRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateDeviceRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *CreateDeviceRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateDeviceRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateDeviceRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateDeviceRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVoipToken

`func (o *CreateDeviceRequest) GetVoipToken() bool`

GetVoipToken returns the VoipToken field if non-nil, zero value otherwise.

### GetVoipTokenOk

`func (o *CreateDeviceRequest) GetVoipTokenOk() (*bool, bool)`

GetVoipTokenOk returns a tuple with the VoipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipToken

`func (o *CreateDeviceRequest) SetVoipToken(v bool)`

SetVoipToken sets VoipToken field to given value.

### HasVoipToken

`func (o *CreateDeviceRequest) HasVoipToken() bool`

HasVoipToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


