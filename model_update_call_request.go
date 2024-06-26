/*
Stream API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v100.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getstream_video

import (
	"encoding/json"
	"time"
)

// checks if the UpdateCallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCallRequest{}

// UpdateCallRequest struct for UpdateCallRequest
type UpdateCallRequest struct {
	// Custom data for this object
	Custom map[string]interface{} `json:"custom,omitempty"`
	SettingsOverride *CallSettingsRequest `json:"settings_override,omitempty"`
	// the time the call is scheduled to start
	StartsAt *time.Time `json:"starts_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateCallRequest UpdateCallRequest

// NewUpdateCallRequest instantiates a new UpdateCallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCallRequest() *UpdateCallRequest {
	this := UpdateCallRequest{}
	return &this
}

// NewUpdateCallRequestWithDefaults instantiates a new UpdateCallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCallRequestWithDefaults() *UpdateCallRequest {
	this := UpdateCallRequest{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *UpdateCallRequest) GetCustom() map[string]interface{} {
	if o == nil || IsNil(o.Custom) {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCallRequest) GetCustomOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Custom) {
		return map[string]interface{}{}, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *UpdateCallRequest) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *UpdateCallRequest) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

// GetSettingsOverride returns the SettingsOverride field value if set, zero value otherwise.
func (o *UpdateCallRequest) GetSettingsOverride() CallSettingsRequest {
	if o == nil || IsNil(o.SettingsOverride) {
		var ret CallSettingsRequest
		return ret
	}
	return *o.SettingsOverride
}

// GetSettingsOverrideOk returns a tuple with the SettingsOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCallRequest) GetSettingsOverrideOk() (*CallSettingsRequest, bool) {
	if o == nil || IsNil(o.SettingsOverride) {
		return nil, false
	}
	return o.SettingsOverride, true
}

// HasSettingsOverride returns a boolean if a field has been set.
func (o *UpdateCallRequest) HasSettingsOverride() bool {
	if o != nil && !IsNil(o.SettingsOverride) {
		return true
	}

	return false
}

// SetSettingsOverride gets a reference to the given CallSettingsRequest and assigns it to the SettingsOverride field.
func (o *UpdateCallRequest) SetSettingsOverride(v CallSettingsRequest) {
	o.SettingsOverride = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *UpdateCallRequest) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCallRequest) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *UpdateCallRequest) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *UpdateCallRequest) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

func (o UpdateCallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.SettingsOverride) {
		toSerialize["settings_override"] = o.SettingsOverride
	}
	if !IsNil(o.StartsAt) {
		toSerialize["starts_at"] = o.StartsAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateCallRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateCallRequest := _UpdateCallRequest{}

	err = json.Unmarshal(data, &varUpdateCallRequest)

	if err != nil {
		return err
	}

	*o = UpdateCallRequest(varUpdateCallRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "custom")
		delete(additionalProperties, "settings_override")
		delete(additionalProperties, "starts_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateCallRequest struct {
	value *UpdateCallRequest
	isSet bool
}

func (v NullableUpdateCallRequest) Get() *UpdateCallRequest {
	return v.value
}

func (v *NullableUpdateCallRequest) Set(val *UpdateCallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCallRequest(val *UpdateCallRequest) *NullableUpdateCallRequest {
	return &NullableUpdateCallRequest{value: val, isSet: true}
}

func (v NullableUpdateCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


