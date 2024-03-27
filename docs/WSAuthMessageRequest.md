# WSAuthMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**UserDetails** | [**ConnectUserDetailsRequest**](ConnectUserDetailsRequest.md) |  | 

## Methods

### NewWSAuthMessageRequest

`func NewWSAuthMessageRequest(token string, userDetails ConnectUserDetailsRequest, ) *WSAuthMessageRequest`

NewWSAuthMessageRequest instantiates a new WSAuthMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWSAuthMessageRequestWithDefaults

`func NewWSAuthMessageRequestWithDefaults() *WSAuthMessageRequest`

NewWSAuthMessageRequestWithDefaults instantiates a new WSAuthMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *WSAuthMessageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WSAuthMessageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WSAuthMessageRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserDetails

`func (o *WSAuthMessageRequest) GetUserDetails() ConnectUserDetailsRequest`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *WSAuthMessageRequest) GetUserDetailsOk() (*ConnectUserDetailsRequest, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *WSAuthMessageRequest) SetUserDetails(v ConnectUserDetailsRequest)`

SetUserDetails sets UserDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


