# UserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **string** | User ID | 
**Image** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Optional name of user | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserRequest

`func NewUserRequest(id string, ) *UserRequest`

NewUserRequest instantiates a new UserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestWithDefaults

`func NewUserRequestWithDefaults() *UserRequest`

NewUserRequestWithDefaults instantiates a new UserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *UserRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UserRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UserRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UserRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetId

`func (o *UserRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRequest) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *UserRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UserRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UserRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UserRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLanguage

`func (o *UserRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *UserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *UserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeams

`func (o *UserRequest) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *UserRequest) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *UserRequest) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *UserRequest) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


