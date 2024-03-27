# CreateGuestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | the access token to authenticate the user | 
**Duration** | **string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 

## Methods

### NewCreateGuestResponse

`func NewCreateGuestResponse(accessToken string, duration string, user UserResponse, ) *CreateGuestResponse`

NewCreateGuestResponse instantiates a new CreateGuestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuestResponseWithDefaults

`func NewCreateGuestResponseWithDefaults() *CreateGuestResponse`

NewCreateGuestResponseWithDefaults instantiates a new CreateGuestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *CreateGuestResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CreateGuestResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CreateGuestResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetDuration

`func (o *CreateGuestResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateGuestResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateGuestResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetUser

`func (o *CreateGuestResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateGuestResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateGuestResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


