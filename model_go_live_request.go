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

// checks if the GoLiveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoLiveRequest{}

// GoLiveRequest struct for GoLiveRequest
type GoLiveRequest struct {
	RecordingStorageName *string `json:"recording_storage_name,omitempty"`
	StartHls *bool `json:"start_hls,omitempty"`
	StartRecording *bool `json:"start_recording,omitempty"`
	StartTranscription *bool `json:"start_transcription,omitempty"`
	TranscriptionStorageName *string `json:"transcription_storage_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GoLiveRequest GoLiveRequest

// NewGoLiveRequest instantiates a new GoLiveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoLiveRequest() *GoLiveRequest {
	this := GoLiveRequest{}
	return &this
}

// NewGoLiveRequestWithDefaults instantiates a new GoLiveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoLiveRequestWithDefaults() *GoLiveRequest {
	this := GoLiveRequest{}
	return &this
}

// GetRecordingStorageName returns the RecordingStorageName field value if set, zero value otherwise.
func (o *GoLiveRequest) GetRecordingStorageName() string {
	if o == nil || IsNil(o.RecordingStorageName) {
		var ret string
		return ret
	}
	return *o.RecordingStorageName
}

// GetRecordingStorageNameOk returns a tuple with the RecordingStorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoLiveRequest) GetRecordingStorageNameOk() (*string, bool) {
	if o == nil || IsNil(o.RecordingStorageName) {
		return nil, false
	}
	return o.RecordingStorageName, true
}

// HasRecordingStorageName returns a boolean if a field has been set.
func (o *GoLiveRequest) HasRecordingStorageName() bool {
	if o != nil && !IsNil(o.RecordingStorageName) {
		return true
	}

	return false
}

// SetRecordingStorageName gets a reference to the given string and assigns it to the RecordingStorageName field.
func (o *GoLiveRequest) SetRecordingStorageName(v string) {
	o.RecordingStorageName = &v
}

// GetStartHls returns the StartHls field value if set, zero value otherwise.
func (o *GoLiveRequest) GetStartHls() bool {
	if o == nil || IsNil(o.StartHls) {
		var ret bool
		return ret
	}
	return *o.StartHls
}

// GetStartHlsOk returns a tuple with the StartHls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoLiveRequest) GetStartHlsOk() (*bool, bool) {
	if o == nil || IsNil(o.StartHls) {
		return nil, false
	}
	return o.StartHls, true
}

// HasStartHls returns a boolean if a field has been set.
func (o *GoLiveRequest) HasStartHls() bool {
	if o != nil && !IsNil(o.StartHls) {
		return true
	}

	return false
}

// SetStartHls gets a reference to the given bool and assigns it to the StartHls field.
func (o *GoLiveRequest) SetStartHls(v bool) {
	o.StartHls = &v
}

// GetStartRecording returns the StartRecording field value if set, zero value otherwise.
func (o *GoLiveRequest) GetStartRecording() bool {
	if o == nil || IsNil(o.StartRecording) {
		var ret bool
		return ret
	}
	return *o.StartRecording
}

// GetStartRecordingOk returns a tuple with the StartRecording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoLiveRequest) GetStartRecordingOk() (*bool, bool) {
	if o == nil || IsNil(o.StartRecording) {
		return nil, false
	}
	return o.StartRecording, true
}

// HasStartRecording returns a boolean if a field has been set.
func (o *GoLiveRequest) HasStartRecording() bool {
	if o != nil && !IsNil(o.StartRecording) {
		return true
	}

	return false
}

// SetStartRecording gets a reference to the given bool and assigns it to the StartRecording field.
func (o *GoLiveRequest) SetStartRecording(v bool) {
	o.StartRecording = &v
}

// GetStartTranscription returns the StartTranscription field value if set, zero value otherwise.
func (o *GoLiveRequest) GetStartTranscription() bool {
	if o == nil || IsNil(o.StartTranscription) {
		var ret bool
		return ret
	}
	return *o.StartTranscription
}

// GetStartTranscriptionOk returns a tuple with the StartTranscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoLiveRequest) GetStartTranscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.StartTranscription) {
		return nil, false
	}
	return o.StartTranscription, true
}

// HasStartTranscription returns a boolean if a field has been set.
func (o *GoLiveRequest) HasStartTranscription() bool {
	if o != nil && !IsNil(o.StartTranscription) {
		return true
	}

	return false
}

// SetStartTranscription gets a reference to the given bool and assigns it to the StartTranscription field.
func (o *GoLiveRequest) SetStartTranscription(v bool) {
	o.StartTranscription = &v
}

// GetTranscriptionStorageName returns the TranscriptionStorageName field value if set, zero value otherwise.
func (o *GoLiveRequest) GetTranscriptionStorageName() string {
	if o == nil || IsNil(o.TranscriptionStorageName) {
		var ret string
		return ret
	}
	return *o.TranscriptionStorageName
}

// GetTranscriptionStorageNameOk returns a tuple with the TranscriptionStorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoLiveRequest) GetTranscriptionStorageNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranscriptionStorageName) {
		return nil, false
	}
	return o.TranscriptionStorageName, true
}

// HasTranscriptionStorageName returns a boolean if a field has been set.
func (o *GoLiveRequest) HasTranscriptionStorageName() bool {
	if o != nil && !IsNil(o.TranscriptionStorageName) {
		return true
	}

	return false
}

// SetTranscriptionStorageName gets a reference to the given string and assigns it to the TranscriptionStorageName field.
func (o *GoLiveRequest) SetTranscriptionStorageName(v string) {
	o.TranscriptionStorageName = &v
}

func (o GoLiveRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoLiveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordingStorageName) {
		toSerialize["recording_storage_name"] = o.RecordingStorageName
	}
	if !IsNil(o.StartHls) {
		toSerialize["start_hls"] = o.StartHls
	}
	if !IsNil(o.StartRecording) {
		toSerialize["start_recording"] = o.StartRecording
	}
	if !IsNil(o.StartTranscription) {
		toSerialize["start_transcription"] = o.StartTranscription
	}
	if !IsNil(o.TranscriptionStorageName) {
		toSerialize["transcription_storage_name"] = o.TranscriptionStorageName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GoLiveRequest) UnmarshalJSON(data []byte) (err error) {
	varGoLiveRequest := _GoLiveRequest{}

	err = json.Unmarshal(data, &varGoLiveRequest)

	if err != nil {
		return err
	}

	*o = GoLiveRequest(varGoLiveRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recording_storage_name")
		delete(additionalProperties, "start_hls")
		delete(additionalProperties, "start_recording")
		delete(additionalProperties, "start_transcription")
		delete(additionalProperties, "transcription_storage_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGoLiveRequest struct {
	value *GoLiveRequest
	isSet bool
}

func (v NullableGoLiveRequest) Get() *GoLiveRequest {
	return v.value
}

func (v *NullableGoLiveRequest) Set(val *GoLiveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoLiveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoLiveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoLiveRequest(val *GoLiveRequest) *NullableGoLiveRequest {
	return &NullableGoLiveRequest{value: val, isSet: true}
}

func (v NullableGoLiveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoLiveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


