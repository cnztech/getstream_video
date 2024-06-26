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

// checks if the ThumbnailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThumbnailResponse{}

// ThumbnailResponse struct for ThumbnailResponse
type ThumbnailResponse struct {
	ImageUrl string `json:"image_url"`
	AdditionalProperties map[string]interface{}
}

type _ThumbnailResponse ThumbnailResponse

// NewThumbnailResponse instantiates a new ThumbnailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbnailResponse(imageUrl string) *ThumbnailResponse {
	this := ThumbnailResponse{}
	this.ImageUrl = imageUrl
	return &this
}

// NewThumbnailResponseWithDefaults instantiates a new ThumbnailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbnailResponseWithDefaults() *ThumbnailResponse {
	this := ThumbnailResponse{}
	return &this
}

// GetImageUrl returns the ImageUrl field value
func (o *ThumbnailResponse) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *ThumbnailResponse) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *ThumbnailResponse) SetImageUrl(v string) {
	o.ImageUrl = v
}

func (o ThumbnailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThumbnailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image_url"] = o.ImageUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThumbnailResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image_url",
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

	varThumbnailResponse := _ThumbnailResponse{}

	err = json.Unmarshal(data, &varThumbnailResponse)

	if err != nil {
		return err
	}

	*o = ThumbnailResponse(varThumbnailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThumbnailResponse struct {
	value *ThumbnailResponse
	isSet bool
}

func (v NullableThumbnailResponse) Get() *ThumbnailResponse {
	return v.value
}

func (v *NullableThumbnailResponse) Set(val *ThumbnailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbnailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbnailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbnailResponse(val *ThumbnailResponse) *NullableThumbnailResponse {
	return &NullableThumbnailResponse{value: val, isSet: true}
}

func (v NullableThumbnailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbnailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


