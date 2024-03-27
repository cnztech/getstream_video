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

// checks if the ThumbnailsSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThumbnailsSettingsRequest{}

// ThumbnailsSettingsRequest struct for ThumbnailsSettingsRequest
type ThumbnailsSettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThumbnailsSettingsRequest ThumbnailsSettingsRequest

// NewThumbnailsSettingsRequest instantiates a new ThumbnailsSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbnailsSettingsRequest() *ThumbnailsSettingsRequest {
	this := ThumbnailsSettingsRequest{}
	return &this
}

// NewThumbnailsSettingsRequestWithDefaults instantiates a new ThumbnailsSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbnailsSettingsRequestWithDefaults() *ThumbnailsSettingsRequest {
	this := ThumbnailsSettingsRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ThumbnailsSettingsRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThumbnailsSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ThumbnailsSettingsRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ThumbnailsSettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ThumbnailsSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThumbnailsSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThumbnailsSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	varThumbnailsSettingsRequest := _ThumbnailsSettingsRequest{}

	err = json.Unmarshal(data, &varThumbnailsSettingsRequest)

	if err != nil {
		return err
	}

	*o = ThumbnailsSettingsRequest(varThumbnailsSettingsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThumbnailsSettingsRequest struct {
	value *ThumbnailsSettingsRequest
	isSet bool
}

func (v NullableThumbnailsSettingsRequest) Get() *ThumbnailsSettingsRequest {
	return v.value
}

func (v *NullableThumbnailsSettingsRequest) Set(val *ThumbnailsSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbnailsSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbnailsSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbnailsSettingsRequest(val *ThumbnailsSettingsRequest) *NullableThumbnailsSettingsRequest {
	return &NullableThumbnailsSettingsRequest{value: val, isSet: true}
}

func (v NullableThumbnailsSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbnailsSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

