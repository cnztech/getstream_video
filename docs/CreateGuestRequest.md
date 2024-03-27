# CreateGuestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserRequest**](UserRequest.md) |  | 

## Methods

### NewCreateGuestRequest

`func NewCreateGuestRequest(user UserRequest, ) *CreateGuestRequest`

NewCreateGuestRequest instantiates a new CreateGuestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuestRequestWithDefaults

`func NewCreateGuestRequestWithDefaults() *CreateGuestRequest`

NewCreateGuestRequestWithDefaults instantiates a new CreateGuestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *CreateGuestRequest) GetUser() UserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateGuestRequest) GetUserOk() (*UserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateGuestRequest) SetUser(v UserRequest)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


