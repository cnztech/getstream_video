/*
Stream API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v100.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getstream_video

import (
	"encoding/json"
)

// checks if the ScreensharingSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScreensharingSettingsRequest{}

// ScreensharingSettingsRequest struct for ScreensharingSettingsRequest
type ScreensharingSettingsRequest struct {
	AccessRequestEnabled *bool `json:"access_request_enabled,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScreensharingSettingsRequest ScreensharingSettingsRequest

// NewScreensharingSettingsRequest instantiates a new ScreensharingSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreensharingSettingsRequest() *ScreensharingSettingsRequest {
	this := ScreensharingSettingsRequest{}
	return &this
}

// NewScreensharingSettingsRequestWithDefaults instantiates a new ScreensharingSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreensharingSettingsRequestWithDefaults() *ScreensharingSettingsRequest {
	this := ScreensharingSettingsRequest{}
	return &this
}

// GetAccessRequestEnabled returns the AccessRequestEnabled field value if set, zero value otherwise.
func (o *ScreensharingSettingsRequest) GetAccessRequestEnabled() bool {
	if o == nil || IsNil(o.AccessRequestEnabled) {
		var ret bool
		return ret
	}
	return *o.AccessRequestEnabled
}

// GetAccessRequestEnabledOk returns a tuple with the AccessRequestEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreensharingSettingsRequest) GetAccessRequestEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessRequestEnabled) {
		return nil, false
	}
	return o.AccessRequestEnabled, true
}

// HasAccessRequestEnabled returns a boolean if a field has been set.
func (o *ScreensharingSettingsRequest) HasAccessRequestEnabled() bool {
	if o != nil && !IsNil(o.AccessRequestEnabled) {
		return true
	}

	return false
}

// SetAccessRequestEnabled gets a reference to the given bool and assigns it to the AccessRequestEnabled field.
func (o *ScreensharingSettingsRequest) SetAccessRequestEnabled(v bool) {
	o.AccessRequestEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScreensharingSettingsRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreensharingSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScreensharingSettingsRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScreensharingSettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ScreensharingSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScreensharingSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessRequestEnabled) {
		toSerialize["access_request_enabled"] = o.AccessRequestEnabled
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScreensharingSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	varScreensharingSettingsRequest := _ScreensharingSettingsRequest{}

	err = json.Unmarshal(data, &varScreensharingSettingsRequest)

	if err != nil {
		return err
	}

	*o = ScreensharingSettingsRequest(varScreensharingSettingsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_request_enabled")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScreensharingSettingsRequest struct {
	value *ScreensharingSettingsRequest
	isSet bool
}

func (v NullableScreensharingSettingsRequest) Get() *ScreensharingSettingsRequest {
	return v.value
}

func (v *NullableScreensharingSettingsRequest) Set(val *ScreensharingSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScreensharingSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScreensharingSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreensharingSettingsRequest(val *ScreensharingSettingsRequest) *NullableScreensharingSettingsRequest {
	return &NullableScreensharingSettingsRequest{value: val, isSet: true}
}

func (v NullableScreensharingSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreensharingSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


