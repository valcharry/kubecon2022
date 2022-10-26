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
)

// ResourceStatus The status and health of a resource and its children.
type ResourceStatus struct {
	// Conditions in this resource that require attention.
	Conditions *[]ResourceCondition `json:"Conditions,omitempty"`
	Health *ResourceHealth `json:"Health,omitempty"`
	HealthRollup *ResourceHealth `json:"HealthRollup,omitempty"`
	// The OEM extension.
	Oem *map[string]map[string]interface{} `json:"Oem,omitempty"`
	State *ResourceState `json:"State,omitempty"`
}

// NewResourceStatus instantiates a new ResourceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceStatus() *ResourceStatus {
	this := ResourceStatus{}
	return &this
}

// NewResourceStatusWithDefaults instantiates a new ResourceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceStatusWithDefaults() *ResourceStatus {
	this := ResourceStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ResourceStatus) GetConditions() []ResourceCondition {
	if o == nil || o.Conditions == nil {
		var ret []ResourceCondition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStatus) GetConditionsOk() (*[]ResourceCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ResourceStatus) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ResourceCondition and assigns it to the Conditions field.
func (o *ResourceStatus) SetConditions(v []ResourceCondition) {
	o.Conditions = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ResourceStatus) GetHealth() ResourceHealth {
	if o == nil || o.Health == nil {
		var ret ResourceHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStatus) GetHealthOk() (*ResourceHealth, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ResourceStatus) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given ResourceHealth and assigns it to the Health field.
func (o *ResourceStatus) SetHealth(v ResourceHealth) {
	o.Health = &v
}

// GetHealthRollup returns the HealthRollup field value if set, zero value otherwise.
func (o *ResourceStatus) GetHealthRollup() ResourceHealth {
	if o == nil || o.HealthRollup == nil {
		var ret ResourceHealth
		return ret
	}
	return *o.HealthRollup
}

// GetHealthRollupOk returns a tuple with the HealthRollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStatus) GetHealthRollupOk() (*ResourceHealth, bool) {
	if o == nil || o.HealthRollup == nil {
		return nil, false
	}
	return o.HealthRollup, true
}

// HasHealthRollup returns a boolean if a field has been set.
func (o *ResourceStatus) HasHealthRollup() bool {
	if o != nil && o.HealthRollup != nil {
		return true
	}

	return false
}

// SetHealthRollup gets a reference to the given ResourceHealth and assigns it to the HealthRollup field.
func (o *ResourceStatus) SetHealthRollup(v ResourceHealth) {
	o.HealthRollup = &v
}

// GetOem returns the Oem field value if set, zero value otherwise.
func (o *ResourceStatus) GetOem() map[string]map[string]interface{} {
	if o == nil || o.Oem == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Oem
}

// GetOemOk returns a tuple with the Oem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStatus) GetOemOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Oem == nil {
		return nil, false
	}
	return o.Oem, true
}

// HasOem returns a boolean if a field has been set.
func (o *ResourceStatus) HasOem() bool {
	if o != nil && o.Oem != nil {
		return true
	}

	return false
}

// SetOem gets a reference to the given map[string]map[string]interface{} and assigns it to the Oem field.
func (o *ResourceStatus) SetOem(v map[string]map[string]interface{}) {
	o.Oem = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ResourceStatus) GetState() ResourceState {
	if o == nil || o.State == nil {
		var ret ResourceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStatus) GetStateOk() (*ResourceState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ResourceStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given ResourceState and assigns it to the State field.
func (o *ResourceStatus) SetState(v ResourceState) {
	o.State = &v
}

func (o ResourceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["Conditions"] = o.Conditions
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.HealthRollup != nil {
		toSerialize["HealthRollup"] = o.HealthRollup
	}
	if o.Oem != nil {
		toSerialize["Oem"] = o.Oem
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableResourceStatus struct {
	value *ResourceStatus
	isSet bool
}

func (v NullableResourceStatus) Get() *ResourceStatus {
	return v.value
}

func (v *NullableResourceStatus) Set(val *ResourceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceStatus(val *ResourceStatus) *NullableResourceStatus {
	return &NullableResourceStatus{value: val, isSet: true}
}

func (v NullableResourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

