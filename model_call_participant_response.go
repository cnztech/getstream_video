/*
Stream API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v100.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getstream_video

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CallParticipantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallParticipantResponse{}

// CallParticipantResponse struct for CallParticipantResponse
type CallParticipantResponse struct {
	JoinedAt time.Time `json:"joined_at"`
	Role string `json:"role"`
	User UserResponse `json:"user"`
	UserSessionId string `json:"user_session_id"`
	AdditionalProperties map[string]interface{}
}

type _CallParticipantResponse CallParticipantResponse

// NewCallParticipantResponse instantiates a new CallParticipantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallParticipantResponse(joinedAt time.Time, role string, user UserResponse, userSessionId string) *CallParticipantResponse {
	this := CallParticipantResponse{}
	this.JoinedAt = joinedAt
	this.Role = role
	this.User = user
	this.UserSessionId = userSessionId
	return &this
}

// NewCallParticipantResponseWithDefaults instantiates a new CallParticipantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallParticipantResponseWithDefaults() *CallParticipantResponse {
	this := CallParticipantResponse{}
	return &this
}

// GetJoinedAt returns the JoinedAt field value
func (o *CallParticipantResponse) GetJoinedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value
// and a boolean to check if the value has been set.
func (o *CallParticipantResponse) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinedAt, true
}

// SetJoinedAt sets field value
func (o *CallParticipantResponse) SetJoinedAt(v time.Time) {
	o.JoinedAt = v
}

// GetRole returns the Role field value
func (o *CallParticipantResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CallParticipantResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CallParticipantResponse) SetRole(v string) {
	o.Role = v
}

// GetUser returns the User field value
func (o *CallParticipantResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *CallParticipantResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *CallParticipantResponse) SetUser(v UserResponse) {
	o.User = v
}

// GetUserSessionId returns the UserSessionId field value
func (o *CallParticipantResponse) GetUserSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserSessionId
}

// GetUserSessionIdOk returns a tuple with the UserSessionId field value
// and a boolean to check if the value has been set.
func (o *CallParticipantResponse) GetUserSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserSessionId, true
}

// SetUserSessionId sets field value
func (o *CallParticipantResponse) SetUserSessionId(v string) {
	o.UserSessionId = v
}

func (o CallParticipantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallParticipantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["joined_at"] = o.JoinedAt
	toSerialize["role"] = o.Role
	toSerialize["user"] = o.User
	toSerialize["user_session_id"] = o.UserSessionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallParticipantResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"joined_at",
		"role",
		"user",
		"user_session_id",
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

	varCallParticipantResponse := _CallParticipantResponse{}

	err = json.Unmarshal(data, &varCallParticipantResponse)

	if err != nil {
		return err
	}

	*o = CallParticipantResponse(varCallParticipantResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "joined_at")
		delete(additionalProperties, "role")
		delete(additionalProperties, "user")
		delete(additionalProperties, "user_session_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallParticipantResponse struct {
	value *CallParticipantResponse
	isSet bool
}

func (v NullableCallParticipantResponse) Get() *CallParticipantResponse {
	return v.value
}

func (v *NullableCallParticipantResponse) Set(val *CallParticipantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallParticipantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallParticipantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallParticipantResponse(val *CallParticipantResponse) *NullableCallParticipantResponse {
	return &NullableCallParticipantResponse{value: val, isSet: true}
}

func (v NullableCallParticipantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallParticipantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

