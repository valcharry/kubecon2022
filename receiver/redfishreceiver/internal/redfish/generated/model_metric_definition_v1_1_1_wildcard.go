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

// MetricDefinitionV111Wildcard The wildcard and its substitution values.
type MetricDefinitionV111Wildcard struct {
	// The string used as a wildcard.
	Name NullableString `json:"Name,omitempty"`
	// An array of values to substitute for the wildcard.
	Values *[]string `json:"Values,omitempty"`
}

// NewMetricDefinitionV111Wildcard instantiates a new MetricDefinitionV111Wildcard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDefinitionV111Wildcard() *MetricDefinitionV111Wildcard {
	this := MetricDefinitionV111Wildcard{}
	return &this
}

// NewMetricDefinitionV111WildcardWithDefaults instantiates a new MetricDefinitionV111Wildcard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDefinitionV111WildcardWithDefaults() *MetricDefinitionV111Wildcard {
	this := MetricDefinitionV111Wildcard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricDefinitionV111Wildcard) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricDefinitionV111Wildcard) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetricDefinitionV111Wildcard) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetricDefinitionV111Wildcard) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetricDefinitionV111Wildcard) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetricDefinitionV111Wildcard) UnsetName() {
	o.Name.Unset()
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MetricDefinitionV111Wildcard) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDefinitionV111Wildcard) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MetricDefinitionV111Wildcard) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *MetricDefinitionV111Wildcard) SetValues(v []string) {
	o.Values = &v
}

func (o MetricDefinitionV111Wildcard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableMetricDefinitionV111Wildcard struct {
	value *MetricDefinitionV111Wildcard
	isSet bool
}

func (v NullableMetricDefinitionV111Wildcard) Get() *MetricDefinitionV111Wildcard {
	return v.value
}

func (v *NullableMetricDefinitionV111Wildcard) Set(val *MetricDefinitionV111Wildcard) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDefinitionV111Wildcard) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDefinitionV111Wildcard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDefinitionV111Wildcard(val *MetricDefinitionV111Wildcard) *NullableMetricDefinitionV111Wildcard {
	return &NullableMetricDefinitionV111Wildcard{value: val, isSet: true}
}

func (v NullableMetricDefinitionV111Wildcard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDefinitionV111Wildcard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


