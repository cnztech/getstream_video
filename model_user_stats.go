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

// checks if the UserStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStats{}

// UserStats struct for UserStats
type UserStats struct {
	Info UserInfoResponse `json:"info"`
	SessionStats map[string]UserSessionStats `json:"session_stats"`
	AdditionalProperties map[string]interface{}
}

type _UserStats UserStats

// NewUserStats instantiates a new UserStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStats(info UserInfoResponse, sessionStats map[string]UserSessionStats) *UserStats {
	this := UserStats{}
	this.Info = info
	this.SessionStats = sessionStats
	return &this
}

// NewUserStatsWithDefaults instantiates a new UserStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatsWithDefaults() *UserStats {
	this := UserStats{}
	return &this
}

// GetInfo returns the Info field value
func (o *UserStats) GetInfo() UserInfoResponse {
	if o == nil {
		var ret UserInfoResponse
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *UserStats) GetInfoOk() (*UserInfoResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *UserStats) SetInfo(v UserInfoResponse) {
	o.Info = v
}

// GetSessionStats returns the SessionStats field value
func (o *UserStats) GetSessionStats() map[string]UserSessionStats {
	if o == nil {
		var ret map[string]UserSessionStats
		return ret
	}

	return o.SessionStats
}

// GetSessionStatsOk returns a tuple with the SessionStats field value
// and a boolean to check if the value has been set.
func (o *UserStats) GetSessionStatsOk() (*map[string]UserSessionStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionStats, true
}

// SetSessionStats sets field value
func (o *UserStats) SetSessionStats(v map[string]UserSessionStats) {
	o.SessionStats = v
}

func (o UserStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["info"] = o.Info
	toSerialize["session_stats"] = o.SessionStats

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"info",
		"session_stats",
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

	varUserStats := _UserStats{}

	err = json.Unmarshal(data, &varUserStats)

	if err != nil {
		return err
	}

	*o = UserStats(varUserStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "info")
		delete(additionalProperties, "session_stats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserStats struct {
	value *UserStats
	isSet bool
}

func (v NullableUserStats) Get() *UserStats {
	return v.value
}

func (v *NullableUserStats) Set(val *UserStats) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStats) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStats(val *UserStats) *NullableUserStats {
	return &NullableUserStats{value: val, isSet: true}
}

func (v NullableUserStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


