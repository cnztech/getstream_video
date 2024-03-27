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

// checks if the ThumbnailsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThumbnailsSettings{}

// ThumbnailsSettings struct for ThumbnailsSettings
type ThumbnailsSettings struct {
	Enabled bool `json:"enabled"`
	AdditionalProperties map[string]interface{}
}

type _ThumbnailsSettings ThumbnailsSettings

// NewThumbnailsSettings instantiates a new ThumbnailsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbnailsSettings(enabled bool) *ThumbnailsSettings {
	this := ThumbnailsSettings{}
	this.Enabled = enabled
	return &this
}

// NewThumbnailsSettingsWithDefaults instantiates a new ThumbnailsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbnailsSettingsWithDefaults() *ThumbnailsSettings {
	this := ThumbnailsSettings{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ThumbnailsSettings) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThumbnailsSettings) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThumbnailsSettings) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThumbnailsSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThumbnailsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThumbnailsSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varThumbnailsSettings := _ThumbnailsSettings{}

	err = json.Unmarshal(data, &varThumbnailsSettings)

	if err != nil {
		return err
	}

	*o = ThumbnailsSettings(varThumbnailsSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThumbnailsSettings struct {
	value *ThumbnailsSettings
	isSet bool
}

func (v NullableThumbnailsSettings) Get() *ThumbnailsSettings {
	return v.value
}

func (v *NullableThumbnailsSettings) Set(val *ThumbnailsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbnailsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbnailsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbnailsSettings(val *ThumbnailsSettings) *NullableThumbnailsSettings {
	return &NullableThumbnailsSettings{value: val, isSet: true}
}

func (v NullableThumbnailsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbnailsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

