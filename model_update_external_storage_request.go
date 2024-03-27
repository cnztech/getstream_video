/*
Stream API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v100.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getstream_video

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateExternalStorageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExternalStorageRequest{}

// UpdateExternalStorageRequest struct for UpdateExternalStorageRequest
type UpdateExternalStorageRequest struct {
	AwsS3 *S3Request `json:"aws_s3,omitempty"`
	AzureBlob *AzureRequest `json:"azure_blob,omitempty"`
	Bucket string `json:"bucket"`
	GcsCredentials *string `json:"gcs_credentials,omitempty"`
	Path *string `json:"path,omitempty"`
	StorageType string `json:"storage_type"`
	AdditionalProperties map[string]interface{}
}

type _UpdateExternalStorageRequest UpdateExternalStorageRequest

// NewUpdateExternalStorageRequest instantiates a new UpdateExternalStorageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExternalStorageRequest(bucket string, storageType string) *UpdateExternalStorageRequest {
	this := UpdateExternalStorageRequest{}
	this.Bucket = bucket
	this.StorageType = storageType
	return &this
}

// NewUpdateExternalStorageRequestWithDefaults instantiates a new UpdateExternalStorageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExternalStorageRequestWithDefaults() *UpdateExternalStorageRequest {
	this := UpdateExternalStorageRequest{}
	return &this
}

// GetAwsS3 returns the AwsS3 field value if set, zero value otherwise.
func (o *UpdateExternalStorageRequest) GetAwsS3() S3Request {
	if o == nil || IsNil(o.AwsS3) {
		var ret S3Request
		return ret
	}
	return *o.AwsS3
}

// GetAwsS3Ok returns a tuple with the AwsS3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageRequest) GetAwsS3Ok() (*S3Request, bool) {
	if o == nil || IsNil(o.AwsS3) {
		return nil, false
	}
	return o.AwsS3, true
}

// HasAwsS3 returns a boolean if a field has been set.
func (o *UpdateExternalStorageRequest) HasAwsS3() bool {
	if o != nil && !IsNil(o.AwsS3) {
		return true
	}

	return false
}

// SetAwsS3 gets a reference to the given S3Request and assigns it to the AwsS3 field.
func (o *UpdateExternalStorageRequest) SetAwsS3(v S3Request) {
	o.AwsS3 = &v
}

// GetAzureBlob returns the AzureBlob field value if set, zero value otherwise.
func (o *UpdateExternalStorageRequest) GetAzureBlob() AzureRequest {
	if o == nil || IsNil(o.AzureBlob) {
		var ret AzureRequest
		return ret
	}
	return *o.AzureBlob
}

// GetAzureBlobOk returns a tuple with the AzureBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageRequest) GetAzureBlobOk() (*AzureRequest, bool) {
	if o == nil || IsNil(o.AzureBlob) {
		return nil, false
	}
	return o.AzureBlob, true
}

// HasAzureBlob returns a boolean if a field has been set.
func (o *UpdateExternalStorageRequest) HasAzureBlob() bool {
	if o != nil && !IsNil(o.AzureBlob) {
		return true
	}

	return false
}

// SetAzureBlob gets a reference to the given AzureRequest and assigns it to the AzureBlob field.
func (o *UpdateExternalStorageRequest) SetAzureBlob(v AzureRequest) {
	o.AzureBlob = &v
}

// GetBucket returns the Bucket field value
func (o *UpdateExternalStorageRequest) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageRequest) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *UpdateExternalStorageRequest) SetBucket(v string) {
	o.Bucket = v
}

// GetGcsCredentials returns the GcsCredentials field value if set, zero value otherwise.
func (o *UpdateExternalStorageRequest) GetGcsCredentials() string {
	if o == nil || IsNil(o.GcsCredentials) {
		var ret string
		return ret
	}
	return *o.GcsCredentials
}

// GetGcsCredentialsOk returns a tuple with the GcsCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageRequest) GetGcsCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.GcsCredentials) {
		return nil, false
	}
	return o.GcsCredentials, true
}

// HasGcsCredentials returns a boolean if a field has been set.
func (o *UpdateExternalStorageRequest) HasGcsCredentials() bool {
	if o != nil && !IsNil(o.GcsCredentials) {
		return true
	}

	return false
}

// SetGcsCredentials gets a reference to the given string and assigns it to the GcsCredentials field.
func (o *UpdateExternalStorageRequest) SetGcsCredentials(v string) {
	o.GcsCredentials = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UpdateExternalStorageRequest) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageRequest) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UpdateExternalStorageRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UpdateExternalStorageRequest) SetPath(v string) {
	o.Path = &v
}

// GetStorageType returns the StorageType field value
func (o *UpdateExternalStorageRequest) GetStorageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageRequest) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *UpdateExternalStorageRequest) SetStorageType(v string) {
	o.StorageType = v
}

func (o UpdateExternalStorageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExternalStorageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsS3) {
		toSerialize["aws_s3"] = o.AwsS3
	}
	if !IsNil(o.AzureBlob) {
		toSerialize["azure_blob"] = o.AzureBlob
	}
	toSerialize["bucket"] = o.Bucket
	if !IsNil(o.GcsCredentials) {
		toSerialize["gcs_credentials"] = o.GcsCredentials
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	toSerialize["storage_type"] = o.StorageType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateExternalStorageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket",
		"storage_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateExternalStorageRequest := _UpdateExternalStorageRequest{}

	err = json.Unmarshal(data, &varUpdateExternalStorageRequest)

	if err != nil {
		return err
	}

	*o = UpdateExternalStorageRequest(varUpdateExternalStorageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aws_s3")
		delete(additionalProperties, "azure_blob")
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "gcs_credentials")
		delete(additionalProperties, "path")
		delete(additionalProperties, "storage_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateExternalStorageRequest struct {
	value *UpdateExternalStorageRequest
	isSet bool
}

func (v NullableUpdateExternalStorageRequest) Get() *UpdateExternalStorageRequest {
	return v.value
}

func (v *NullableUpdateExternalStorageRequest) Set(val *UpdateExternalStorageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExternalStorageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExternalStorageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExternalStorageRequest(val *UpdateExternalStorageRequest) *NullableUpdateExternalStorageRequest {
	return &NullableUpdateExternalStorageRequest{value: val, isSet: true}
}

func (v NullableUpdateExternalStorageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExternalStorageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

