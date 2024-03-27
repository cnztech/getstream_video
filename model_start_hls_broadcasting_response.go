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

// checks if the StartHLSBroadcastingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartHLSBroadcastingResponse{}

// StartHLSBroadcastingResponse struct for StartHLSBroadcastingResponse
type StartHLSBroadcastingResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
	PlaylistUrl string `json:"playlist_url"`
	AdditionalProperties map[string]interface{}
}

type _StartHLSBroadcastingResponse StartHLSBroadcastingResponse

// NewStartHLSBroadcastingResponse instantiates a new StartHLSBroadcastingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartHLSBroadcastingResponse(duration string, playlistUrl string) *StartHLSBroadcastingResponse {
	this := StartHLSBroadcastingResponse{}
	this.Duration = duration
	this.PlaylistUrl = playlistUrl
	return &this
}

// NewStartHLSBroadcastingResponseWithDefaults instantiates a new StartHLSBroadcastingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartHLSBroadcastingResponseWithDefaults() *StartHLSBroadcastingResponse {
	this := StartHLSBroadcastingResponse{}
	return &this
}

// GetDuration returns the Duration field value
func (o *StartHLSBroadcastingResponse) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *StartHLSBroadcastingResponse) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *StartHLSBroadcastingResponse) SetDuration(v string) {
	o.Duration = v
}

// GetPlaylistUrl returns the PlaylistUrl field value
func (o *StartHLSBroadcastingResponse) GetPlaylistUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaylistUrl
}

// GetPlaylistUrlOk returns a tuple with the PlaylistUrl field value
// and a boolean to check if the value has been set.
func (o *StartHLSBroadcastingResponse) GetPlaylistUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaylistUrl, true
}

// SetPlaylistUrl sets field value
func (o *StartHLSBroadcastingResponse) SetPlaylistUrl(v string) {
	o.PlaylistUrl = v
}

func (o StartHLSBroadcastingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartHLSBroadcastingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["playlist_url"] = o.PlaylistUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartHLSBroadcastingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration",
		"playlist_url",
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

	varStartHLSBroadcastingResponse := _StartHLSBroadcastingResponse{}

	err = json.Unmarshal(data, &varStartHLSBroadcastingResponse)

	if err != nil {
		return err
	}

	*o = StartHLSBroadcastingResponse(varStartHLSBroadcastingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "playlist_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartHLSBroadcastingResponse struct {
	value *StartHLSBroadcastingResponse
	isSet bool
}

func (v NullableStartHLSBroadcastingResponse) Get() *StartHLSBroadcastingResponse {
	return v.value
}

func (v *NullableStartHLSBroadcastingResponse) Set(val *StartHLSBroadcastingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartHLSBroadcastingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartHLSBroadcastingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartHLSBroadcastingResponse(val *StartHLSBroadcastingResponse) *NullableStartHLSBroadcastingResponse {
	return &NullableStartHLSBroadcastingResponse{value: val, isSet: true}
}

func (v NullableStartHLSBroadcastingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartHLSBroadcastingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

