# S3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3ApiKey** | Pointer to **string** |  | [optional] 
**S3Region** | **string** |  | 
**S3Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewS3Request

`func NewS3Request(s3Region string, ) *S3Request`

NewS3Request instantiates a new S3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3RequestWithDefaults

`func NewS3RequestWithDefaults() *S3Request`

NewS3RequestWithDefaults instantiates a new S3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3ApiKey

`func (o *S3Request) GetS3ApiKey() string`

GetS3ApiKey returns the S3ApiKey field if non-nil, zero value otherwise.

### GetS3ApiKeyOk

`func (o *S3Request) GetS3ApiKeyOk() (*string, bool)`

GetS3ApiKeyOk returns a tuple with the S3ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ApiKey

`func (o *S3Request) SetS3ApiKey(v string)`

SetS3ApiKey sets S3ApiKey field to given value.

### HasS3ApiKey

`func (o *S3Request) HasS3ApiKey() bool`

HasS3ApiKey returns a boolean if a field has been set.

### GetS3Region

`func (o *S3Request) GetS3Region() string`

GetS3Region returns the S3Region field if non-nil, zero value otherwise.

### GetS3RegionOk

`func (o *S3Request) GetS3RegionOk() (*string, bool)`

GetS3RegionOk returns a tuple with the S3Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Region

`func (o *S3Request) SetS3Region(v string)`

SetS3Region sets S3Region field to given value.


### GetS3Secret

`func (o *S3Request) GetS3Secret() string`

GetS3Secret returns the S3Secret field if non-nil, zero value otherwise.

### GetS3SecretOk

`func (o *S3Request) GetS3SecretOk() (*string, bool)`

GetS3SecretOk returns a tuple with the S3Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Secret

`func (o *S3Request) SetS3Secret(v string)`

SetS3Secret sets S3Secret field to given value.

### HasS3Secret

`func (o *S3Request) HasS3Secret() bool`

HasS3Secret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


