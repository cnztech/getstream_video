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

// checks if the BlockUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockUserRequest{}

// BlockUserRequest struct for BlockUserRequest
type BlockUserRequest struct {
	// the user to block
	UserId string `json:"user_id"`
	AdditionalProperties map[string]interface{}
}

type _BlockUserRequest BlockUserRequest

// NewBlockUserRequest instantiates a new BlockUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockUserRequest(userId string) *BlockUserRequest {
	this := BlockUserRequest{}
	this.UserId = userId
	return &this
}

// NewBlockUserRequestWithDefaults instantiates a new BlockUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockUserRequestWithDefaults() *BlockUserRequest {
	this := BlockUserRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *BlockUserRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *BlockUserRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *BlockUserRequest) SetUserId(v string) {
	o.UserId = v
}

func (o BlockUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varBlockUserRequest := _BlockUserRequest{}

	err = json.Unmarshal(data, &varBlockUserRequest)

	if err != nil {
		return err
	}

	*o = BlockUserRequest(varBlockUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockUserRequest struct {
	value *BlockUserRequest
	isSet bool
}

func (v NullableBlockUserRequest) Get() *BlockUserRequest {
	return v.value
}

func (v *NullableBlockUserRequest) Set(val *BlockUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockUserRequest(val *BlockUserRequest) *NullableBlockUserRequest {
	return &NullableBlockUserRequest{value: val, isSet: true}
}

func (v NullableBlockUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


