# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | Response HTTP status code | 
**Code** | **int32** | API error code | 
**Details** | **[]int32** | Additional error-specific information | 
**Duration** | **string** | Request duration | 
**ExceptionFields** | Pointer to **map[string]string** | Additional error info | [optional] 
**Message** | **string** | Message describing an error | 
**MoreInfo** | **string** | URL with additional information | 

## Methods

### NewAPIError

`func NewAPIError(statusCode int32, code int32, details []int32, duration string, message string, moreInfo string, ) *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *APIError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *APIError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *APIError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetCode

`func (o *APIError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDetails

`func (o *APIError) GetDetails() []int32`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *APIError) GetDetailsOk() (*[]int32, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *APIError) SetDetails(v []int32)`

SetDetails sets Details field to given value.


### GetDuration

`func (o *APIError) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *APIError) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *APIError) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetExceptionFields

`func (o *APIError) GetExceptionFields() map[string]string`

GetExceptionFields returns the ExceptionFields field if non-nil, zero value otherwise.

### GetExceptionFieldsOk

`func (o *APIError) GetExceptionFieldsOk() (*map[string]string, bool)`

GetExceptionFieldsOk returns a tuple with the ExceptionFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionFields

`func (o *APIError) SetExceptionFields(v map[string]string)`

SetExceptionFields sets ExceptionFields field to given value.

### HasExceptionFields

`func (o *APIError) HasExceptionFields() bool`

HasExceptionFields returns a boolean if a field has been set.

### GetMessage

`func (o *APIError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *APIError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *APIError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMoreInfo

`func (o *APIError) GetMoreInfo() string`

GetMoreInfo returns the MoreInfo field if non-nil, zero value otherwise.

### GetMoreInfoOk

`func (o *APIError) GetMoreInfoOk() (*string, bool)`

GetMoreInfoOk returns a tuple with the MoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreInfo

`func (o *APIError) SetMoreInfo(v string)`

SetMoreInfo sets MoreInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


