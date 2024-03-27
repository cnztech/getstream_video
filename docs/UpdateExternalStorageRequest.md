# UpdateExternalStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsS3** | Pointer to [**S3Request**](S3Request.md) |  | [optional] 
**AzureBlob** | Pointer to [**AzureRequest**](AzureRequest.md) |  | [optional] 
**Bucket** | **string** |  | 
**GcsCredentials** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**StorageType** | **string** |  | 

## Methods

### NewUpdateExternalStorageRequest

`func NewUpdateExternalStorageRequest(bucket string, storageType string, ) *UpdateExternalStorageRequest`

NewUpdateExternalStorageRequest instantiates a new UpdateExternalStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExternalStorageRequestWithDefaults

`func NewUpdateExternalStorageRequestWithDefaults() *UpdateExternalStorageRequest`

NewUpdateExternalStorageRequestWithDefaults instantiates a new UpdateExternalStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsS3

`func (o *UpdateExternalStorageRequest) GetAwsS3() S3Request`

GetAwsS3 returns the AwsS3 field if non-nil, zero value otherwise.

### GetAwsS3Ok

`func (o *UpdateExternalStorageRequest) GetAwsS3Ok() (*S3Request, bool)`

GetAwsS3Ok returns a tuple with the AwsS3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3

`func (o *UpdateExternalStorageRequest) SetAwsS3(v S3Request)`

SetAwsS3 sets AwsS3 field to given value.

### HasAwsS3

`func (o *UpdateExternalStorageRequest) HasAwsS3() bool`

HasAwsS3 returns a boolean if a field has been set.

### GetAzureBlob

`func (o *UpdateExternalStorageRequest) GetAzureBlob() AzureRequest`

GetAzureBlob returns the AzureBlob field if non-nil, zero value otherwise.

### GetAzureBlobOk

`func (o *UpdateExternalStorageRequest) GetAzureBlobOk() (*AzureRequest, bool)`

GetAzureBlobOk returns a tuple with the AzureBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlob

`func (o *UpdateExternalStorageRequest) SetAzureBlob(v AzureRequest)`

SetAzureBlob sets AzureBlob field to given value.

### HasAzureBlob

`func (o *UpdateExternalStorageRequest) HasAzureBlob() bool`

HasAzureBlob returns a boolean if a field has been set.

### GetBucket

`func (o *UpdateExternalStorageRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *UpdateExternalStorageRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *UpdateExternalStorageRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetGcsCredentials

`func (o *UpdateExternalStorageRequest) GetGcsCredentials() string`

GetGcsCredentials returns the GcsCredentials field if non-nil, zero value otherwise.

### GetGcsCredentialsOk

`func (o *UpdateExternalStorageRequest) GetGcsCredentialsOk() (*string, bool)`

GetGcsCredentialsOk returns a tuple with the GcsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsCredentials

`func (o *UpdateExternalStorageRequest) SetGcsCredentials(v string)`

SetGcsCredentials sets GcsCredentials field to given value.

### HasGcsCredentials

`func (o *UpdateExternalStorageRequest) HasGcsCredentials() bool`

HasGcsCredentials returns a boolean if a field has been set.

### GetPath

`func (o *UpdateExternalStorageRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateExternalStorageRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateExternalStorageRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateExternalStorageRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStorageType

`func (o *UpdateExternalStorageRequest) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *UpdateExternalStorageRequest) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *UpdateExternalStorageRequest) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


