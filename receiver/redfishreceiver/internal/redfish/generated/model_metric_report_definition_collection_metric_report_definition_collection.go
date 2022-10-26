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

// MetricReportDefinitionCollectionMetricReportDefinitionCollection The collection of MetricReportDefinition resource instances.
type MetricReportDefinitionCollectionMetricReportDefinitionCollection struct {
	// The OData description of a payload.
	OdataContext *string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag *string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The description of this resource.  Used for commonality in the schema definitions.
	Description *string `json:"Description,omitempty"`
	// The members of this collection.
	Members []OdataV4IdRef `json:"Members"`
	// The number of items in a collection.
	MembersodataCount int64 `json:"Members@odata.count"`
	// The URI to the resource containing the next set of partial members.
	MembersodataNextLink *string `json:"Members@odata.nextLink,omitempty"`
	// The name of the resource or array member.
	Name string `json:"Name"`
	// The OEM extension.
	Oem *map[string]map[string]interface{} `json:"Oem,omitempty"`
}

// NewMetricReportDefinitionCollectionMetricReportDefinitionCollection instantiates a new MetricReportDefinitionCollectionMetricReportDefinitionCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricReportDefinitionCollectionMetricReportDefinitionCollection(odataId string, odataType string, members []OdataV4IdRef, membersodataCount int64, name string, ) *MetricReportDefinitionCollectionMetricReportDefinitionCollection {
	this := MetricReportDefinitionCollectionMetricReportDefinitionCollection{}
	this.OdataId = odataId
	this.OdataType = odataType
	this.Members = members
	this.MembersodataCount = membersodataCount
	this.Name = name
	return &this
}

// NewMetricReportDefinitionCollectionMetricReportDefinitionCollectionWithDefaults instantiates a new MetricReportDefinitionCollectionMetricReportDefinitionCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricReportDefinitionCollectionMetricReportDefinitionCollectionWithDefaults() *MetricReportDefinitionCollectionMetricReportDefinitionCollection {
	this := MetricReportDefinitionCollectionMetricReportDefinitionCollection{}
	return &this
}

// GetOdataContext returns the OdataContext field value if set, zero value otherwise.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataContext() string {
	if o == nil || o.OdataContext == nil {
		var ret string
		return ret
	}
	return *o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataContextOk() (*string, bool) {
	if o == nil || o.OdataContext == nil {
		return nil, false
	}
	return o.OdataContext, true
}

// HasOdataContext returns a boolean if a field has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) HasOdataContext() bool {
	if o != nil && o.OdataContext != nil {
		return true
	}

	return false
}

// SetOdataContext gets a reference to the given string and assigns it to the OdataContext field.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetOdataContext(v string) {
	o.OdataContext = &v
}

// GetOdataEtag returns the OdataEtag field value if set, zero value otherwise.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataEtag() string {
	if o == nil || o.OdataEtag == nil {
		var ret string
		return ret
	}
	return *o.OdataEtag
}

// GetOdataEtagOk returns a tuple with the OdataEtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataEtagOk() (*string, bool) {
	if o == nil || o.OdataEtag == nil {
		return nil, false
	}
	return o.OdataEtag, true
}

// HasOdataEtag returns a boolean if a field has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) HasOdataEtag() bool {
	if o != nil && o.OdataEtag != nil {
		return true
	}

	return false
}

// SetOdataEtag gets a reference to the given string and assigns it to the OdataEtag field.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetOdataEtag(v string) {
	o.OdataEtag = &v
}

// GetOdataId returns the OdataId field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOdataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetOdataType(v string) {
	o.OdataType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetDescription(v string) {
	o.Description = &v
}

// GetMembers returns the Members field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetMembers() []OdataV4IdRef {
	if o == nil  {
		var ret []OdataV4IdRef
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetMembersOk() (*[]OdataV4IdRef, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetMembers(v []OdataV4IdRef) {
	o.Members = v
}

// GetMembersodataCount returns the MembersodataCount field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetMembersodataCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.MembersodataCount
}

// GetMembersodataCountOk returns a tuple with the MembersodataCount field value
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetMembersodataCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MembersodataCount, true
}

// SetMembersodataCount sets field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetMembersodataCount(v int64) {
	o.MembersodataCount = v
}

// GetMembersodataNextLink returns the MembersodataNextLink field value if set, zero value otherwise.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetMembersodataNextLink() string {
	if o == nil || o.MembersodataNextLink == nil {
		var ret string
		return ret
	}
	return *o.MembersodataNextLink
}

// GetMembersodataNextLinkOk returns a tuple with the MembersodataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetMembersodataNextLinkOk() (*string, bool) {
	if o == nil || o.MembersodataNextLink == nil {
		return nil, false
	}
	return o.MembersodataNextLink, true
}

// HasMembersodataNextLink returns a boolean if a field has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) HasMembersodataNextLink() bool {
	if o != nil && o.MembersodataNextLink != nil {
		return true
	}

	return false
}

// SetMembersodataNextLink gets a reference to the given string and assigns it to the MembersodataNextLink field.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetMembersodataNextLink(v string) {
	o.MembersodataNextLink = &v
}

// GetName returns the Name field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetName(v string) {
	o.Name = v
}

// GetOem returns the Oem field value if set, zero value otherwise.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOem() map[string]map[string]interface{} {
	if o == nil || o.Oem == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Oem
}

// GetOemOk returns a tuple with the Oem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) GetOemOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Oem == nil {
		return nil, false
	}
	return o.Oem, true
}

// HasOem returns a boolean if a field has been set.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) HasOem() bool {
	if o != nil && o.Oem != nil {
		return true
	}

	return false
}

// SetOem gets a reference to the given map[string]map[string]interface{} and assigns it to the Oem field.
func (o *MetricReportDefinitionCollectionMetricReportDefinitionCollection) SetOem(v map[string]map[string]interface{}) {
	o.Oem = &v
}

func (o MetricReportDefinitionCollectionMetricReportDefinitionCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OdataContext != nil {
		toSerialize["@odata.context"] = o.OdataContext
	}
	if o.OdataEtag != nil {
		toSerialize["@odata.etag"] = o.OdataEtag
	}
	if true {
		toSerialize["@odata.id"] = o.OdataId
	}
	if true {
		toSerialize["@odata.type"] = o.OdataType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["Members"] = o.Members
	}
	if true {
		toSerialize["Members@odata.count"] = o.MembersodataCount
	}
	if o.MembersodataNextLink != nil {
		toSerialize["Members@odata.nextLink"] = o.MembersodataNextLink
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.Oem != nil {
		toSerialize["Oem"] = o.Oem
	}
	return json.Marshal(toSerialize)
}

type NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection struct {
	value *MetricReportDefinitionCollectionMetricReportDefinitionCollection
	isSet bool
}

func (v NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection) Get() *MetricReportDefinitionCollectionMetricReportDefinitionCollection {
	return v.value
}

func (v *NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection) Set(val *MetricReportDefinitionCollectionMetricReportDefinitionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricReportDefinitionCollectionMetricReportDefinitionCollection(val *MetricReportDefinitionCollectionMetricReportDefinitionCollection) *NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection {
	return &NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection{value: val, isSet: true}
}

func (v NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricReportDefinitionCollectionMetricReportDefinitionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

