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

// checks if the CallIngressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallIngressResponse{}

// CallIngressResponse struct for CallIngressResponse
type CallIngressResponse struct {
	Rtmp RTMPIngress `json:"rtmp"`
	AdditionalProperties map[string]interface{}
}

type _CallIngressResponse CallIngressResponse

// NewCallIngressResponse instantiates a new CallIngressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallIngressResponse(rtmp RTMPIngress) *CallIngressResponse {
	this := CallIngressResponse{}
	this.Rtmp = rtmp
	return &this
}

// NewCallIngressResponseWithDefaults instantiates a new CallIngressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallIngressResponseWithDefaults() *CallIngressResponse {
	this := CallIngressResponse{}
	return &this
}

// GetRtmp returns the Rtmp field value
func (o *CallIngressResponse) GetRtmp() RTMPIngress {
	if o == nil {
		var ret RTMPIngress
		return ret
	}

	return o.Rtmp
}

// GetRtmpOk returns a tuple with the Rtmp field value
// and a boolean to check if the value has been set.
func (o *CallIngressResponse) GetRtmpOk() (*RTMPIngress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rtmp, true
}

// SetRtmp sets field value
func (o *CallIngressResponse) SetRtmp(v RTMPIngress) {
	o.Rtmp = v
}

func (o CallIngressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallIngressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rtmp"] = o.Rtmp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallIngressResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rtmp",
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

	varCallIngressResponse := _CallIngressResponse{}

	err = json.Unmarshal(data, &varCallIngressResponse)

	if err != nil {
		return err
	}

	*o = CallIngressResponse(varCallIngressResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rtmp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallIngressResponse struct {
	value *CallIngressResponse
	isSet bool
}

func (v NullableCallIngressResponse) Get() *CallIngressResponse {
	return v.value
}

func (v *NullableCallIngressResponse) Set(val *CallIngressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallIngressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallIngressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallIngressResponse(val *CallIngressResponse) *NullableCallIngressResponse {
	return &NullableCallIngressResponse{value: val, isSet: true}
}

func (v NullableCallIngressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallIngressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


