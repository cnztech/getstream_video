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

// checks if the CreateExternalStorageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExternalStorageResponse{}

// CreateExternalStorageResponse struct for CreateExternalStorageResponse
type CreateExternalStorageResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
	AdditionalProperties map[string]interface{}
}

type _CreateExternalStorageResponse CreateExternalStorageResponse

// NewCreateExternalStorageResponse instantiates a new CreateExternalStorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExternalStorageResponse(duration string) *CreateExternalStorageResponse {
	this := CreateExternalStorageResponse{}
	this.Duration = duration
	return &this
}

// NewCreateExternalStorageResponseWithDefaults instantiates a new CreateExternalStorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExternalStorageResponseWithDefaults() *CreateExternalStorageResponse {
	this := CreateExternalStorageResponse{}
	return &this
}

// GetDuration returns the Duration field value
func (o *CreateExternalStorageResponse) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *CreateExternalStorageResponse) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *CreateExternalStorageResponse) SetDuration(v string) {
	o.Duration = v
}

func (o CreateExternalStorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateExternalStorageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateExternalStorageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration",
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

	varCreateExternalStorageResponse := _CreateExternalStorageResponse{}

	err = json.Unmarshal(data, &varCreateExternalStorageResponse)

	if err != nil {
		return err
	}

	*o = CreateExternalStorageResponse(varCreateExternalStorageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateExternalStorageResponse struct {
	value *CreateExternalStorageResponse
	isSet bool
}

func (v NullableCreateExternalStorageResponse) Get() *CreateExternalStorageResponse {
	return v.value
}

func (v *NullableCreateExternalStorageResponse) Set(val *CreateExternalStorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExternalStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExternalStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExternalStorageResponse(val *CreateExternalStorageResponse) *NullableCreateExternalStorageResponse {
	return &NullableCreateExternalStorageResponse{value: val, isSet: true}
}

func (v NullableCreateExternalStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExternalStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


