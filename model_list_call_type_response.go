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

// checks if the ListCallTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCallTypeResponse{}

// ListCallTypeResponse struct for ListCallTypeResponse
type ListCallTypeResponse struct {
	CallTypes map[string]CallTypeResponse `json:"call_types"`
	Duration string `json:"duration"`
	AdditionalProperties map[string]interface{}
}

type _ListCallTypeResponse ListCallTypeResponse

// NewListCallTypeResponse instantiates a new ListCallTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCallTypeResponse(callTypes map[string]CallTypeResponse, duration string) *ListCallTypeResponse {
	this := ListCallTypeResponse{}
	this.CallTypes = callTypes
	this.Duration = duration
	return &this
}

// NewListCallTypeResponseWithDefaults instantiates a new ListCallTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCallTypeResponseWithDefaults() *ListCallTypeResponse {
	this := ListCallTypeResponse{}
	return &this
}

// GetCallTypes returns the CallTypes field value
func (o *ListCallTypeResponse) GetCallTypes() map[string]CallTypeResponse {
	if o == nil {
		var ret map[string]CallTypeResponse
		return ret
	}

	return o.CallTypes
}

// GetCallTypesOk returns a tuple with the CallTypes field value
// and a boolean to check if the value has been set.
func (o *ListCallTypeResponse) GetCallTypesOk() (*map[string]CallTypeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallTypes, true
}

// SetCallTypes sets field value
func (o *ListCallTypeResponse) SetCallTypes(v map[string]CallTypeResponse) {
	o.CallTypes = v
}

// GetDuration returns the Duration field value
func (o *ListCallTypeResponse) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ListCallTypeResponse) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ListCallTypeResponse) SetDuration(v string) {
	o.Duration = v
}

func (o ListCallTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCallTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["call_types"] = o.CallTypes
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCallTypeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"call_types",
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

	varListCallTypeResponse := _ListCallTypeResponse{}

	err = json.Unmarshal(data, &varListCallTypeResponse)

	if err != nil {
		return err
	}

	*o = ListCallTypeResponse(varListCallTypeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "call_types")
		delete(additionalProperties, "duration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCallTypeResponse struct {
	value *ListCallTypeResponse
	isSet bool
}

func (v NullableListCallTypeResponse) Get() *ListCallTypeResponse {
	return v.value
}

func (v *NullableListCallTypeResponse) Set(val *ListCallTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCallTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCallTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCallTypeResponse(val *ListCallTypeResponse) *NullableListCallTypeResponse {
	return &NullableListCallTypeResponse{value: val, isSet: true}
}

func (v NullableListCallTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCallTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


