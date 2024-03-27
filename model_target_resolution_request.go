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

// checks if the TargetResolutionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetResolutionRequest{}

// TargetResolutionRequest struct for TargetResolutionRequest
type TargetResolutionRequest struct {
	Bitrate *int32 `json:"bitrate,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Width *int32 `json:"width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TargetResolutionRequest TargetResolutionRequest

// NewTargetResolutionRequest instantiates a new TargetResolutionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetResolutionRequest() *TargetResolutionRequest {
	this := TargetResolutionRequest{}
	return &this
}

// NewTargetResolutionRequestWithDefaults instantiates a new TargetResolutionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetResolutionRequestWithDefaults() *TargetResolutionRequest {
	this := TargetResolutionRequest{}
	return &this
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *TargetResolutionRequest) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate) {
		var ret int32
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResolutionRequest) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate) {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *TargetResolutionRequest) HasBitrate() bool {
	if o != nil && !IsNil(o.Bitrate) {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.
func (o *TargetResolutionRequest) SetBitrate(v int32) {
	o.Bitrate = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TargetResolutionRequest) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResolutionRequest) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TargetResolutionRequest) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *TargetResolutionRequest) SetHeight(v int32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *TargetResolutionRequest) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResolutionRequest) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *TargetResolutionRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *TargetResolutionRequest) SetWidth(v int32) {
	o.Width = &v
}

func (o TargetResolutionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetResolutionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bitrate) {
		toSerialize["bitrate"] = o.Bitrate
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TargetResolutionRequest) UnmarshalJSON(data []byte) (err error) {
	varTargetResolutionRequest := _TargetResolutionRequest{}

	err = json.Unmarshal(data, &varTargetResolutionRequest)

	if err != nil {
		return err
	}

	*o = TargetResolutionRequest(varTargetResolutionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bitrate")
		delete(additionalProperties, "height")
		delete(additionalProperties, "width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTargetResolutionRequest struct {
	value *TargetResolutionRequest
	isSet bool
}

func (v NullableTargetResolutionRequest) Get() *TargetResolutionRequest {
	return v.value
}

func (v *NullableTargetResolutionRequest) Set(val *TargetResolutionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetResolutionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetResolutionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetResolutionRequest(val *TargetResolutionRequest) *NullableTargetResolutionRequest {
	return &NullableTargetResolutionRequest{value: val, isSet: true}
}

func (v NullableTargetResolutionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetResolutionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

