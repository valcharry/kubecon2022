/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2020.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redfish

import (
	"encoding/json"
	"fmt"
)

// ResourceState the model 'ResourceState'
type ResourceState string

// List of Resource_State
const (
	RESOURCESTATE_ENABLED ResourceState = "Enabled"
	RESOURCESTATE_DISABLED ResourceState = "Disabled"
	RESOURCESTATE_STANDBY_OFFLINE ResourceState = "StandbyOffline"
	RESOURCESTATE_STANDBY_SPARE ResourceState = "StandbySpare"
	RESOURCESTATE_IN_TEST ResourceState = "InTest"
	RESOURCESTATE_STARTING ResourceState = "Starting"
	RESOURCESTATE_ABSENT ResourceState = "Absent"
	RESOURCESTATE_UNAVAILABLE_OFFLINE ResourceState = "UnavailableOffline"
	RESOURCESTATE_DEFERRING ResourceState = "Deferring"
	RESOURCESTATE_QUIESCED ResourceState = "Quiesced"
	RESOURCESTATE_UPDATING ResourceState = "Updating"
	RESOURCESTATE_QUALIFIED ResourceState = "Qualified"
)

func (v *ResourceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceState(value)
	for _, existing := range []ResourceState{ "Enabled", "Disabled", "StandbyOffline", "StandbySpare", "InTest", "Starting", "Absent", "UnavailableOffline", "Deferring", "Quiesced", "Updating", "Qualified",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceState", value)
}

// Ptr returns reference to Resource_State value
func (v ResourceState) Ptr() *ResourceState {
	return &v
}

type NullableResourceState struct {
	value *ResourceState
	isSet bool
}

func (v NullableResourceState) Get() *ResourceState {
	return v.value
}

func (v *NullableResourceState) Set(val *ResourceState) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceState) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceState(val *ResourceState) *NullableResourceState {
	return &NullableResourceState{value: val, isSet: true}
}

func (v NullableResourceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

