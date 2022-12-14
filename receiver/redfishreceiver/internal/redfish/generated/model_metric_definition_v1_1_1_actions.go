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

// MetricDefinitionV111Actions The available actions for this resource.
type MetricDefinitionV111Actions struct {
	// The available OEM-specific actions for this resource.
	Oem *map[string]map[string]interface{} `json:"Oem,omitempty"`
}

// NewMetricDefinitionV111Actions instantiates a new MetricDefinitionV111Actions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDefinitionV111Actions() *MetricDefinitionV111Actions {
	this := MetricDefinitionV111Actions{}
	return &this
}

// NewMetricDefinitionV111ActionsWithDefaults instantiates a new MetricDefinitionV111Actions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDefinitionV111ActionsWithDefaults() *MetricDefinitionV111Actions {
	this := MetricDefinitionV111Actions{}
	return &this
}

// GetOem returns the Oem field value if set, zero value otherwise.
func (o *MetricDefinitionV111Actions) GetOem() map[string]map[string]interface{} {
	if o == nil || o.Oem == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Oem
}

// GetOemOk returns a tuple with the Oem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDefinitionV111Actions) GetOemOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Oem == nil {
		return nil, false
	}
	return o.Oem, true
}

// HasOem returns a boolean if a field has been set.
func (o *MetricDefinitionV111Actions) HasOem() bool {
	if o != nil && o.Oem != nil {
		return true
	}

	return false
}

// SetOem gets a reference to the given map[string]map[string]interface{} and assigns it to the Oem field.
func (o *MetricDefinitionV111Actions) SetOem(v map[string]map[string]interface{}) {
	o.Oem = &v
}

func (o MetricDefinitionV111Actions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Oem != nil {
		toSerialize["Oem"] = o.Oem
	}
	return json.Marshal(toSerialize)
}

type NullableMetricDefinitionV111Actions struct {
	value *MetricDefinitionV111Actions
	isSet bool
}

func (v NullableMetricDefinitionV111Actions) Get() *MetricDefinitionV111Actions {
	return v.value
}

func (v *NullableMetricDefinitionV111Actions) Set(val *MetricDefinitionV111Actions) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDefinitionV111Actions) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDefinitionV111Actions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDefinitionV111Actions(val *MetricDefinitionV111Actions) *NullableMetricDefinitionV111Actions {
	return &NullableMetricDefinitionV111Actions{value: val, isSet: true}
}

func (v NullableMetricDefinitionV111Actions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDefinitionV111Actions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


