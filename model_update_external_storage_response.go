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

// checks if the UpdateExternalStorageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExternalStorageResponse{}

// UpdateExternalStorageResponse struct for UpdateExternalStorageResponse
type UpdateExternalStorageResponse struct {
	Bucket string `json:"bucket"`
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _UpdateExternalStorageResponse UpdateExternalStorageResponse

// NewUpdateExternalStorageResponse instantiates a new UpdateExternalStorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExternalStorageResponse(bucket string, duration string, name string, path string, type_ string) *UpdateExternalStorageResponse {
	this := UpdateExternalStorageResponse{}
	this.Bucket = bucket
	this.Duration = duration
	this.Name = name
	this.Path = path
	this.Type = type_
	return &this
}

// NewUpdateExternalStorageResponseWithDefaults instantiates a new UpdateExternalStorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExternalStorageResponseWithDefaults() *UpdateExternalStorageResponse {
	this := UpdateExternalStorageResponse{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *UpdateExternalStorageResponse) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageResponse) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *UpdateExternalStorageResponse) SetBucket(v string) {
	o.Bucket = v
}

// GetDuration returns the Duration field value
func (o *UpdateExternalStorageResponse) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageResponse) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *UpdateExternalStorageResponse) SetDuration(v string) {
	o.Duration = v
}

// GetName returns the Name field value
func (o *UpdateExternalStorageResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateExternalStorageResponse) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *UpdateExternalStorageResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *UpdateExternalStorageResponse) SetPath(v string) {
	o.Path = v
}

// GetType returns the Type field value
func (o *UpdateExternalStorageResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateExternalStorageResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateExternalStorageResponse) SetType(v string) {
	o.Type = v
}

func (o UpdateExternalStorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExternalStorageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	toSerialize["duration"] = o.Duration
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateExternalStorageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket",
		"duration",
		"name",
		"path",
		"type",
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

	varUpdateExternalStorageResponse := _UpdateExternalStorageResponse{}

	err = json.Unmarshal(data, &varUpdateExternalStorageResponse)

	if err != nil {
		return err
	}

	*o = UpdateExternalStorageResponse(varUpdateExternalStorageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateExternalStorageResponse struct {
	value *UpdateExternalStorageResponse
	isSet bool
}

func (v NullableUpdateExternalStorageResponse) Get() *UpdateExternalStorageResponse {
	return v.value
}

func (v *NullableUpdateExternalStorageResponse) Set(val *UpdateExternalStorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExternalStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExternalStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExternalStorageResponse(val *UpdateExternalStorageResponse) *NullableUpdateExternalStorageResponse {
	return &NullableUpdateExternalStorageResponse{value: val, isSet: true}
}

func (v NullableUpdateExternalStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExternalStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

