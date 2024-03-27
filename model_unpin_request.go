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

// checks if the UnpinRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnpinRequest{}

// UnpinRequest struct for UnpinRequest
type UnpinRequest struct {
	SessionId string `json:"session_id"`
	UserId string `json:"user_id"`
	AdditionalProperties map[string]interface{}
}

type _UnpinRequest UnpinRequest

// NewUnpinRequest instantiates a new UnpinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnpinRequest(sessionId string, userId string) *UnpinRequest {
	this := UnpinRequest{}
	this.SessionId = sessionId
	this.UserId = userId
	return &this
}

// NewUnpinRequestWithDefaults instantiates a new UnpinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnpinRequestWithDefaults() *UnpinRequest {
	this := UnpinRequest{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *UnpinRequest) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *UnpinRequest) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *UnpinRequest) SetSessionId(v string) {
	o.SessionId = v
}

// GetUserId returns the UserId field value
func (o *UnpinRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnpinRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnpinRequest) SetUserId(v string) {
	o.UserId = v
}

func (o UnpinRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnpinRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session_id"] = o.SessionId
	toSerialize["user_id"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnpinRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session_id",
		"user_id",
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

	varUnpinRequest := _UnpinRequest{}

	err = json.Unmarshal(data, &varUnpinRequest)

	if err != nil {
		return err
	}

	*o = UnpinRequest(varUnpinRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "session_id")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnpinRequest struct {
	value *UnpinRequest
	isSet bool
}

func (v NullableUnpinRequest) Get() *UnpinRequest {
	return v.value
}

func (v *NullableUnpinRequest) Set(val *UnpinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnpinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnpinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnpinRequest(val *UnpinRequest) *NullableUnpinRequest {
	return &NullableUnpinRequest{value: val, isSet: true}
}

func (v NullableUnpinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnpinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


