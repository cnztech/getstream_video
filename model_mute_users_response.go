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

// checks if the MuteUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteUsersResponse{}

// MuteUsersResponse struct for MuteUsersResponse
type MuteUsersResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
	AdditionalProperties map[string]interface{}
}

type _MuteUsersResponse MuteUsersResponse

// NewMuteUsersResponse instantiates a new MuteUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteUsersResponse(duration string) *MuteUsersResponse {
	this := MuteUsersResponse{}
	this.Duration = duration
	return &this
}

// NewMuteUsersResponseWithDefaults instantiates a new MuteUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteUsersResponseWithDefaults() *MuteUsersResponse {
	this := MuteUsersResponse{}
	return &this
}

// GetDuration returns the Duration field value
func (o *MuteUsersResponse) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *MuteUsersResponse) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *MuteUsersResponse) SetDuration(v string) {
	o.Duration = v
}

func (o MuteUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MuteUsersResponse) UnmarshalJSON(data []byte) (err error) {
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

	varMuteUsersResponse := _MuteUsersResponse{}

	err = json.Unmarshal(data, &varMuteUsersResponse)

	if err != nil {
		return err
	}

	*o = MuteUsersResponse(varMuteUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMuteUsersResponse struct {
	value *MuteUsersResponse
	isSet bool
}

func (v NullableMuteUsersResponse) Get() *MuteUsersResponse {
	return v.value
}

func (v *NullableMuteUsersResponse) Set(val *MuteUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteUsersResponse(val *MuteUsersResponse) *NullableMuteUsersResponse {
	return &NullableMuteUsersResponse{value: val, isSet: true}
}

func (v NullableMuteUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


