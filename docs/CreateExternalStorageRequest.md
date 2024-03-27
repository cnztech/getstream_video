# CreateExternalStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsS3** | Pointer to [**S3Request**](S3Request.md) |  | [optional] 
**AzureBlob** | Pointer to [**AzureRequest**](AzureRequest.md) |  | [optional] 
**Bucket** | **string** |  | 
**GcsCredentials** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 
**StorageType** | **string** |  | 

## Methods

### NewCreateExternalStorageRequest

`func NewCreateExternalStorageRequest(bucket string, name string, storageType string, ) *CreateExternalStorageRequest`

NewCreateExternalStorageRequest instantiates a new CreateExternalStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalStorageRequestWithDefaults

`func NewCreateExternalStorageRequestWithDefaults() *CreateExternalStorageRequest`

NewCreateExternalStorageRequestWithDefaults instantiates a new CreateExternalStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsS3

`func (o *CreateExternalStorageRequest) GetAwsS3() S3Request`

GetAwsS3 returns the AwsS3 field if non-nil, zero value otherwise.

### GetAwsS3Ok

`func (o *CreateExternalStorageRequest) GetAwsS3Ok() (*S3Request, bool)`

GetAwsS3Ok returns a tuple with the AwsS3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3

`func (o *CreateExternalStorageRequest) SetAwsS3(v S3Request)`

SetAwsS3 sets AwsS3 field to given value.

### HasAwsS3

`func (o *CreateExternalStorageRequest) HasAwsS3() bool`

HasAwsS3 returns a boolean if a field has been set.

### GetAzureBlob

`func (o *CreateExternalStorageRequest) GetAzureBlob() AzureRequest`

GetAzureBlob returns the AzureBlob field if non-nil, zero value otherwise.

### GetAzureBlobOk

`func (o *CreateExternalStorageRequest) GetAzureBlobOk() (*AzureRequest, bool)`

GetAzureBlobOk returns a tuple with the AzureBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlob

`func (o *CreateExternalStorageRequest) SetAzureBlob(v AzureRequest)`

SetAzureBlob sets AzureBlob field to given value.

### HasAzureBlob

`func (o *CreateExternalStorageRequest) HasAzureBlob() bool`

HasAzureBlob returns a boolean if a field has been set.

### GetBucket

`func (o *CreateExternalStorageRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CreateExternalStorageRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CreateExternalStorageRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetGcsCredentials

`func (o *CreateExternalStorageRequest) GetGcsCredentials() string`

GetGcsCredentials returns the GcsCredentials field if non-nil, zero value otherwise.

### GetGcsCredentialsOk

`func (o *CreateExternalStorageRequest) GetGcsCredentialsOk() (*string, bool)`

GetGcsCredentialsOk returns a tuple with the GcsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsCredentials

`func (o *CreateExternalStorageRequest) SetGcsCredentials(v string)`

SetGcsCredentials sets GcsCredentials field to given value.

### HasGcsCredentials

`func (o *CreateExternalStorageRequest) HasGcsCredentials() bool`

HasGcsCredentials returns a boolean if a field has been set.

### GetName

`func (o *CreateExternalStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExternalStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExternalStorageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *CreateExternalStorageRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateExternalStorageRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateExternalStorageRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateExternalStorageRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStorageType

`func (o *CreateExternalStorageRequest) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CreateExternalStorageRequest) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CreateExternalStorageRequest) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


