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

// checks if the UpdateUserPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserPermissionsRequest{}

// UpdateUserPermissionsRequest struct for UpdateUserPermissionsRequest
type UpdateUserPermissionsRequest struct {
	GrantPermissions []string `json:"grant_permissions,omitempty"`
	RevokePermissions []string `json:"revoke_permissions,omitempty"`
	UserId string `json:"user_id"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserPermissionsRequest UpdateUserPermissionsRequest

// NewUpdateUserPermissionsRequest instantiates a new UpdateUserPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserPermissionsRequest(userId string) *UpdateUserPermissionsRequest {
	this := UpdateUserPermissionsRequest{}
	this.UserId = userId
	return &this
}

// NewUpdateUserPermissionsRequestWithDefaults instantiates a new UpdateUserPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserPermissionsRequestWithDefaults() *UpdateUserPermissionsRequest {
	this := UpdateUserPermissionsRequest{}
	return &this
}

// GetGrantPermissions returns the GrantPermissions field value if set, zero value otherwise.
func (o *UpdateUserPermissionsRequest) GetGrantPermissions() []string {
	if o == nil || IsNil(o.GrantPermissions) {
		var ret []string
		return ret
	}
	return o.GrantPermissions
}

// GetGrantPermissionsOk returns a tuple with the GrantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPermissionsRequest) GetGrantPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GrantPermissions) {
		return nil, false
	}
	return o.GrantPermissions, true
}

// HasGrantPermissions returns a boolean if a field has been set.
func (o *UpdateUserPermissionsRequest) HasGrantPermissions() bool {
	if o != nil && !IsNil(o.GrantPermissions) {
		return true
	}

	return false
}

// SetGrantPermissions gets a reference to the given []string and assigns it to the GrantPermissions field.
func (o *UpdateUserPermissionsRequest) SetGrantPermissions(v []string) {
	o.GrantPermissions = v
}

// GetRevokePermissions returns the RevokePermissions field value if set, zero value otherwise.
func (o *UpdateUserPermissionsRequest) GetRevokePermissions() []string {
	if o == nil || IsNil(o.RevokePermissions) {
		var ret []string
		return ret
	}
	return o.RevokePermissions
}

// GetRevokePermissionsOk returns a tuple with the RevokePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPermissionsRequest) GetRevokePermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.RevokePermissions) {
		return nil, false
	}
	return o.RevokePermissions, true
}

// HasRevokePermissions returns a boolean if a field has been set.
func (o *UpdateUserPermissionsRequest) HasRevokePermissions() bool {
	if o != nil && !IsNil(o.RevokePermissions) {
		return true
	}

	return false
}

// SetRevokePermissions gets a reference to the given []string and assigns it to the RevokePermissions field.
func (o *UpdateUserPermissionsRequest) SetRevokePermissions(v []string) {
	o.RevokePermissions = v
}

// GetUserId returns the UserId field value
func (o *UpdateUserPermissionsRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UpdateUserPermissionsRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UpdateUserPermissionsRequest) SetUserId(v string) {
	o.UserId = v
}

func (o UpdateUserPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GrantPermissions) {
		toSerialize["grant_permissions"] = o.GrantPermissions
	}
	if !IsNil(o.RevokePermissions) {
		toSerialize["revoke_permissions"] = o.RevokePermissions
	}
	toSerialize["user_id"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateUserPermissionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varUpdateUserPermissionsRequest := _UpdateUserPermissionsRequest{}

	err = json.Unmarshal(data, &varUpdateUserPermissionsRequest)

	if err != nil {
		return err
	}

	*o = UpdateUserPermissionsRequest(varUpdateUserPermissionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grant_permissions")
		delete(additionalProperties, "revoke_permissions")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserPermissionsRequest struct {
	value *UpdateUserPermissionsRequest
	isSet bool
}

func (v NullableUpdateUserPermissionsRequest) Get() *UpdateUserPermissionsRequest {
	return v.value
}

func (v *NullableUpdateUserPermissionsRequest) Set(val *UpdateUserPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserPermissionsRequest(val *UpdateUserPermissionsRequest) *NullableUpdateUserPermissionsRequest {
	return &NullableUpdateUserPermissionsRequest{value: val, isSet: true}
}

func (v NullableUpdateUserPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

