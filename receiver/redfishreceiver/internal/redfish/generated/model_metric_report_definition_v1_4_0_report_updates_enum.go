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

// MetricReportDefinitionV140ReportUpdatesEnum Handling of subsequent metric reports when a metric report exists.
type MetricReportDefinitionV140ReportUpdatesEnum string

// List of MetricReportDefinition_v1_4_0_ReportUpdatesEnum
const (
	METRICREPORTDEFINITIONV140REPORTUPDATESENUM_OVERWRITE MetricReportDefinitionV140ReportUpdatesEnum = "Overwrite"
	METRICREPORTDEFINITIONV140REPORTUPDATESENUM_APPEND_WRAPS_WHEN_FULL MetricReportDefinitionV140ReportUpdatesEnum = "AppendWrapsWhenFull"
	METRICREPORTDEFINITIONV140REPORTUPDATESENUM_APPEND_STOPS_WHEN_FULL MetricReportDefinitionV140ReportUpdatesEnum = "AppendStopsWhenFull"
	METRICREPORTDEFINITIONV140REPORTUPDATESENUM_NEW_REPORT MetricReportDefinitionV140ReportUpdatesEnum = "NewReport"
)

func (v *MetricReportDefinitionV140ReportUpdatesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricReportDefinitionV140ReportUpdatesEnum(value)
	for _, existing := range []MetricReportDefinitionV140ReportUpdatesEnum{ "Overwrite", "AppendWrapsWhenFull", "AppendStopsWhenFull", "NewReport",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricReportDefinitionV140ReportUpdatesEnum", value)
}

// Ptr returns reference to MetricReportDefinition_v1_4_0_ReportUpdatesEnum value
func (v MetricReportDefinitionV140ReportUpdatesEnum) Ptr() *MetricReportDefinitionV140ReportUpdatesEnum {
	return &v
}

type NullableMetricReportDefinitionV140ReportUpdatesEnum struct {
	value *MetricReportDefinitionV140ReportUpdatesEnum
	isSet bool
}

func (v NullableMetricReportDefinitionV140ReportUpdatesEnum) Get() *MetricReportDefinitionV140ReportUpdatesEnum {
	return v.value
}

func (v *NullableMetricReportDefinitionV140ReportUpdatesEnum) Set(val *MetricReportDefinitionV140ReportUpdatesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricReportDefinitionV140ReportUpdatesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricReportDefinitionV140ReportUpdatesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricReportDefinitionV140ReportUpdatesEnum(val *MetricReportDefinitionV140ReportUpdatesEnum) *NullableMetricReportDefinitionV140ReportUpdatesEnum {
	return &NullableMetricReportDefinitionV140ReportUpdatesEnum{value: val, isSet: true}
}

func (v NullableMetricReportDefinitionV140ReportUpdatesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricReportDefinitionV140ReportUpdatesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

