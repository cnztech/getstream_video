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

// checks if the GoLiveResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoLiveResponse{}

// GoLiveResponse struct for GoLiveResponse
type GoLiveResponse struct {
	Call CallResponse `json:"call"`
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
	AdditionalProperties map[string]interface{}
}

type _GoLiveResponse GoLiveResponse

// NewGoLiveResponse instantiates a new GoLiveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoLiveResponse(call CallResponse, duration string) *GoLiveResponse {
	this := GoLiveResponse{}
	this.Call = call
	this.Duration = duration
	return &this
}

// NewGoLiveResponseWithDefaults instantiates a new GoLiveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoLiveResponseWithDefaults() *GoLiveResponse {
	this := GoLiveResponse{}
	return &this
}

// GetCall returns the Call field value
func (o *GoLiveResponse) GetCall() CallResponse {
	if o == nil {
		var ret CallResponse
		return ret
	}

	return o.Call
}

// GetCallOk returns a tuple with the Call field value
// and a boolean to check if the value has been set.
func (o *GoLiveResponse) GetCallOk() (*CallResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Call, true
}

// SetCall sets field value
func (o *GoLiveResponse) SetCall(v CallResponse) {
	o.Call = v
}

// GetDuration returns the Duration field value
func (o *GoLiveResponse) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *GoLiveResponse) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *GoLiveResponse) SetDuration(v string) {
	o.Duration = v
}

func (o GoLiveResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoLiveResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["call"] = o.Call
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GoLiveResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"call",
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

	varGoLiveResponse := _GoLiveResponse{}

	err = json.Unmarshal(data, &varGoLiveResponse)

	if err != nil {
		return err
	}

	*o = GoLiveResponse(varGoLiveResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "call")
		delete(additionalProperties, "duration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGoLiveResponse struct {
	value *GoLiveResponse
	isSet bool
}

func (v NullableGoLiveResponse) Get() *GoLiveResponse {
	return v.value
}

func (v *NullableGoLiveResponse) Set(val *GoLiveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGoLiveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGoLiveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoLiveResponse(val *GoLiveResponse) *NullableGoLiveResponse {
	return &NullableGoLiveResponse{value: val, isSet: true}
}

func (v NullableGoLiveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoLiveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


