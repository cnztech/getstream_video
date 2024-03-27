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

// checks if the TranscriptionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranscriptionSettings{}

// TranscriptionSettings struct for TranscriptionSettings
type TranscriptionSettings struct {
	ClosedCaptionMode string `json:"closed_caption_mode"`
	Mode string `json:"mode"`
	AdditionalProperties map[string]interface{}
}

type _TranscriptionSettings TranscriptionSettings

// NewTranscriptionSettings instantiates a new TranscriptionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscriptionSettings(closedCaptionMode string, mode string) *TranscriptionSettings {
	this := TranscriptionSettings{}
	this.ClosedCaptionMode = closedCaptionMode
	this.Mode = mode
	return &this
}

// NewTranscriptionSettingsWithDefaults instantiates a new TranscriptionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscriptionSettingsWithDefaults() *TranscriptionSettings {
	this := TranscriptionSettings{}
	return &this
}

// GetClosedCaptionMode returns the ClosedCaptionMode field value
func (o *TranscriptionSettings) GetClosedCaptionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClosedCaptionMode
}

// GetClosedCaptionModeOk returns a tuple with the ClosedCaptionMode field value
// and a boolean to check if the value has been set.
func (o *TranscriptionSettings) GetClosedCaptionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClosedCaptionMode, true
}

// SetClosedCaptionMode sets field value
func (o *TranscriptionSettings) SetClosedCaptionMode(v string) {
	o.ClosedCaptionMode = v
}

// GetMode returns the Mode field value
func (o *TranscriptionSettings) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TranscriptionSettings) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TranscriptionSettings) SetMode(v string) {
	o.Mode = v
}

func (o TranscriptionSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranscriptionSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["closed_caption_mode"] = o.ClosedCaptionMode
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TranscriptionSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"closed_caption_mode",
		"mode",
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

	varTranscriptionSettings := _TranscriptionSettings{}

	err = json.Unmarshal(data, &varTranscriptionSettings)

	if err != nil {
		return err
	}

	*o = TranscriptionSettings(varTranscriptionSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "closed_caption_mode")
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTranscriptionSettings struct {
	value *TranscriptionSettings
	isSet bool
}

func (v NullableTranscriptionSettings) Get() *TranscriptionSettings {
	return v.value
}

func (v *NullableTranscriptionSettings) Set(val *TranscriptionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscriptionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscriptionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscriptionSettings(val *TranscriptionSettings) *NullableTranscriptionSettings {
	return &NullableTranscriptionSettings{value: val, isSet: true}
}

func (v NullableTranscriptionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscriptionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


