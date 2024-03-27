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

// checks if the EgressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EgressResponse{}

// EgressResponse struct for EgressResponse
type EgressResponse struct {
	Broadcasting bool `json:"broadcasting"`
	Hls *EgressHLSResponse `json:"hls,omitempty"`
	Rtmps []EgressRTMPResponse `json:"rtmps"`
	AdditionalProperties map[string]interface{}
}

type _EgressResponse EgressResponse

// NewEgressResponse instantiates a new EgressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEgressResponse(broadcasting bool, rtmps []EgressRTMPResponse) *EgressResponse {
	this := EgressResponse{}
	this.Broadcasting = broadcasting
	this.Rtmps = rtmps
	return &this
}

// NewEgressResponseWithDefaults instantiates a new EgressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEgressResponseWithDefaults() *EgressResponse {
	this := EgressResponse{}
	return &this
}

// GetBroadcasting returns the Broadcasting field value
func (o *EgressResponse) GetBroadcasting() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Broadcasting
}

// GetBroadcastingOk returns a tuple with the Broadcasting field value
// and a boolean to check if the value has been set.
func (o *EgressResponse) GetBroadcastingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broadcasting, true
}

// SetBroadcasting sets field value
func (o *EgressResponse) SetBroadcasting(v bool) {
	o.Broadcasting = v
}

// GetHls returns the Hls field value if set, zero value otherwise.
func (o *EgressResponse) GetHls() EgressHLSResponse {
	if o == nil || IsNil(o.Hls) {
		var ret EgressHLSResponse
		return ret
	}
	return *o.Hls
}

// GetHlsOk returns a tuple with the Hls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EgressResponse) GetHlsOk() (*EgressHLSResponse, bool) {
	if o == nil || IsNil(o.Hls) {
		return nil, false
	}
	return o.Hls, true
}

// HasHls returns a boolean if a field has been set.
func (o *EgressResponse) HasHls() bool {
	if o != nil && !IsNil(o.Hls) {
		return true
	}

	return false
}

// SetHls gets a reference to the given EgressHLSResponse and assigns it to the Hls field.
func (o *EgressResponse) SetHls(v EgressHLSResponse) {
	o.Hls = &v
}

// GetRtmps returns the Rtmps field value
func (o *EgressResponse) GetRtmps() []EgressRTMPResponse {
	if o == nil {
		var ret []EgressRTMPResponse
		return ret
	}

	return o.Rtmps
}

// GetRtmpsOk returns a tuple with the Rtmps field value
// and a boolean to check if the value has been set.
func (o *EgressResponse) GetRtmpsOk() ([]EgressRTMPResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rtmps, true
}

// SetRtmps sets field value
func (o *EgressResponse) SetRtmps(v []EgressRTMPResponse) {
	o.Rtmps = v
}

func (o EgressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EgressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["broadcasting"] = o.Broadcasting
	if !IsNil(o.Hls) {
		toSerialize["hls"] = o.Hls
	}
	toSerialize["rtmps"] = o.Rtmps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EgressResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"broadcasting",
		"rtmps",
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

	varEgressResponse := _EgressResponse{}

	err = json.Unmarshal(data, &varEgressResponse)

	if err != nil {
		return err
	}

	*o = EgressResponse(varEgressResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "broadcasting")
		delete(additionalProperties, "hls")
		delete(additionalProperties, "rtmps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEgressResponse struct {
	value *EgressResponse
	isSet bool
}

func (v NullableEgressResponse) Get() *EgressResponse {
	return v.value
}

func (v *NullableEgressResponse) Set(val *EgressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEgressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEgressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEgressResponse(val *EgressResponse) *NullableEgressResponse {
	return &NullableEgressResponse{value: val, isSet: true}
}

func (v NullableEgressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEgressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


