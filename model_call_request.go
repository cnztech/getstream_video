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
)

// checks if the CallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRequest{}

// CallRequest struct for CallRequest
type CallRequest struct {
	CreatedBy *UserRequest `json:"created_by,omitempty"`
	CreatedById *string `json:"created_by_id,omitempty"`
	Custom map[string]interface{} `json:"custom,omitempty"`
	Members []MemberRequest `json:"members,omitempty"`
	SettingsOverride *CallSettingsRequest `json:"settings_override,omitempty"`
	StartsAt *time.Time `json:"starts_at,omitempty"`
	Team *string `json:"team,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallRequest CallRequest

// NewCallRequest instantiates a new CallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallRequest() *CallRequest {
	this := CallRequest{}
	return &this
}

// NewCallRequestWithDefaults instantiates a new CallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRequestWithDefaults() *CallRequest {
	this := CallRequest{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CallRequest) GetCreatedBy() UserRequest {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserRequest
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetCreatedByOk() (*UserRequest, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CallRequest) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserRequest and assigns it to the CreatedBy field.
func (o *CallRequest) SetCreatedBy(v UserRequest) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *CallRequest) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *CallRequest) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *CallRequest) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *CallRequest) GetCustom() map[string]interface{} {
	if o == nil || IsNil(o.Custom) {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetCustomOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Custom) {
		return map[string]interface{}{}, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *CallRequest) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *CallRequest) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *CallRequest) GetMembers() []MemberRequest {
	if o == nil || IsNil(o.Members) {
		var ret []MemberRequest
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetMembersOk() ([]MemberRequest, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CallRequest) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MemberRequest and assigns it to the Members field.
func (o *CallRequest) SetMembers(v []MemberRequest) {
	o.Members = v
}

// GetSettingsOverride returns the SettingsOverride field value if set, zero value otherwise.
func (o *CallRequest) GetSettingsOverride() CallSettingsRequest {
	if o == nil || IsNil(o.SettingsOverride) {
		var ret CallSettingsRequest
		return ret
	}
	return *o.SettingsOverride
}

// GetSettingsOverrideOk returns a tuple with the SettingsOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetSettingsOverrideOk() (*CallSettingsRequest, bool) {
	if o == nil || IsNil(o.SettingsOverride) {
		return nil, false
	}
	return o.SettingsOverride, true
}

// HasSettingsOverride returns a boolean if a field has been set.
func (o *CallRequest) HasSettingsOverride() bool {
	if o != nil && !IsNil(o.SettingsOverride) {
		return true
	}

	return false
}

// SetSettingsOverride gets a reference to the given CallSettingsRequest and assigns it to the SettingsOverride field.
func (o *CallRequest) SetSettingsOverride(v CallSettingsRequest) {
	o.SettingsOverride = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *CallRequest) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *CallRequest) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *CallRequest) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *CallRequest) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *CallRequest) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *CallRequest) SetTeam(v string) {
	o.Team = &v
}

func (o CallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CreatedById) {
		toSerialize["created_by_id"] = o.CreatedById
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.SettingsOverride) {
		toSerialize["settings_override"] = o.SettingsOverride
	}
	if !IsNil(o.StartsAt) {
		toSerialize["starts_at"] = o.StartsAt
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallRequest) UnmarshalJSON(data []byte) (err error) {
	varCallRequest := _CallRequest{}

	err = json.Unmarshal(data, &varCallRequest)

	if err != nil {
		return err
	}

	*o = CallRequest(varCallRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "created_by_id")
		delete(additionalProperties, "custom")
		delete(additionalProperties, "members")
		delete(additionalProperties, "settings_override")
		delete(additionalProperties, "starts_at")
		delete(additionalProperties, "team")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallRequest struct {
	value *CallRequest
	isSet bool
}

func (v NullableCallRequest) Get() *CallRequest {
	return v.value
}

func (v *NullableCallRequest) Set(val *CallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRequest(val *CallRequest) *NullableCallRequest {
	return &NullableCallRequest{value: val, isSet: true}
}

func (v NullableCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

