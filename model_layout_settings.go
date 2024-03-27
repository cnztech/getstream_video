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

// checks if the LayoutSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LayoutSettings{}

// LayoutSettings struct for LayoutSettings
type LayoutSettings struct {
	ExternalAppUrl string `json:"external_app_url"`
	ExternalCssUrl string `json:"external_css_url"`
	Name string `json:"name"`
	Options map[string]interface{} `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LayoutSettings LayoutSettings

// NewLayoutSettings instantiates a new LayoutSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutSettings(externalAppUrl string, externalCssUrl string, name string) *LayoutSettings {
	this := LayoutSettings{}
	this.ExternalAppUrl = externalAppUrl
	this.ExternalCssUrl = externalCssUrl
	this.Name = name
	return &this
}

// NewLayoutSettingsWithDefaults instantiates a new LayoutSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutSettingsWithDefaults() *LayoutSettings {
	this := LayoutSettings{}
	return &this
}

// GetExternalAppUrl returns the ExternalAppUrl field value
func (o *LayoutSettings) GetExternalAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalAppUrl
}

// GetExternalAppUrlOk returns a tuple with the ExternalAppUrl field value
// and a boolean to check if the value has been set.
func (o *LayoutSettings) GetExternalAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalAppUrl, true
}

// SetExternalAppUrl sets field value
func (o *LayoutSettings) SetExternalAppUrl(v string) {
	o.ExternalAppUrl = v
}

// GetExternalCssUrl returns the ExternalCssUrl field value
func (o *LayoutSettings) GetExternalCssUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCssUrl
}

// GetExternalCssUrlOk returns a tuple with the ExternalCssUrl field value
// and a boolean to check if the value has been set.
func (o *LayoutSettings) GetExternalCssUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCssUrl, true
}

// SetExternalCssUrl sets field value
func (o *LayoutSettings) SetExternalCssUrl(v string) {
	o.ExternalCssUrl = v
}

// GetName returns the Name field value
func (o *LayoutSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LayoutSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LayoutSettings) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LayoutSettings) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutSettings) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LayoutSettings) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *LayoutSettings) SetOptions(v map[string]interface{}) {
	o.Options = v
}

func (o LayoutSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LayoutSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["external_app_url"] = o.ExternalAppUrl
	toSerialize["external_css_url"] = o.ExternalCssUrl
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LayoutSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_app_url",
		"external_css_url",
		"name",
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

	varLayoutSettings := _LayoutSettings{}

	err = json.Unmarshal(data, &varLayoutSettings)

	if err != nil {
		return err
	}

	*o = LayoutSettings(varLayoutSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "external_app_url")
		delete(additionalProperties, "external_css_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLayoutSettings struct {
	value *LayoutSettings
	isSet bool
}

func (v NullableLayoutSettings) Get() *LayoutSettings {
	return v.value
}

func (v *NullableLayoutSettings) Set(val *LayoutSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutSettings(val *LayoutSettings) *NullableLayoutSettings {
	return &NullableLayoutSettings{value: val, isSet: true}
}

func (v NullableLayoutSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


