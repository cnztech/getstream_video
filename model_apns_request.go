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

// checks if the APNSRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APNSRequest{}

// APNSRequest struct for APNSRequest
type APNSRequest struct {
	Body *string `json:"body,omitempty"`
	Title *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APNSRequest APNSRequest

// NewAPNSRequest instantiates a new APNSRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPNSRequest() *APNSRequest {
	this := APNSRequest{}
	return &this
}

// NewAPNSRequestWithDefaults instantiates a new APNSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPNSRequestWithDefaults() *APNSRequest {
	this := APNSRequest{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *APNSRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *APNSRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *APNSRequest) SetBody(v string) {
	o.Body = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *APNSRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *APNSRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *APNSRequest) SetTitle(v string) {
	o.Title = &v
}

func (o APNSRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APNSRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *APNSRequest) UnmarshalJSON(data []byte) (err error) {
	varAPNSRequest := _APNSRequest{}

	err = json.Unmarshal(data, &varAPNSRequest)

	if err != nil {
		return err
	}

	*o = APNSRequest(varAPNSRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAPNSRequest struct {
	value *APNSRequest
	isSet bool
}

func (v NullableAPNSRequest) Get() *APNSRequest {
	return v.value
}

func (v *NullableAPNSRequest) Set(val *APNSRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAPNSRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAPNSRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPNSRequest(val *APNSRequest) *NullableAPNSRequest {
	return &NullableAPNSRequest{value: val, isSet: true}
}

func (v NullableAPNSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPNSRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


