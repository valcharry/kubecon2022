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

// MetricReportDefinitionV140CalculationAlgorithmEnum The function to apply to the list of metric properties.
type MetricReportDefinitionV140CalculationAlgorithmEnum string

// List of MetricReportDefinition_v1_4_0_CalculationAlgorithmEnum
const (
	METRICREPORTDEFINITIONV140CALCULATIONALGORITHMENUM_AVERAGE MetricReportDefinitionV140CalculationAlgorithmEnum = "Average"
	METRICREPORTDEFINITIONV140CALCULATIONALGORITHMENUM_MAXIMUM MetricReportDefinitionV140CalculationAlgorithmEnum = "Maximum"
	METRICREPORTDEFINITIONV140CALCULATIONALGORITHMENUM_MINIMUM MetricReportDefinitionV140CalculationAlgorithmEnum = "Minimum"
	METRICREPORTDEFINITIONV140CALCULATIONALGORITHMENUM_SUMMATION MetricReportDefinitionV140CalculationAlgorithmEnum = "Summation"
)

func (v *MetricReportDefinitionV140CalculationAlgorithmEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricReportDefinitionV140CalculationAlgorithmEnum(value)
	for _, existing := range []MetricReportDefinitionV140CalculationAlgorithmEnum{ "Average", "Maximum", "Minimum", "Summation",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricReportDefinitionV140CalculationAlgorithmEnum", value)
}

// Ptr returns reference to MetricReportDefinition_v1_4_0_CalculationAlgorithmEnum value
func (v MetricReportDefinitionV140CalculationAlgorithmEnum) Ptr() *MetricReportDefinitionV140CalculationAlgorithmEnum {
	return &v
}

type NullableMetricReportDefinitionV140CalculationAlgorithmEnum struct {
	value *MetricReportDefinitionV140CalculationAlgorithmEnum
	isSet bool
}

func (v NullableMetricReportDefinitionV140CalculationAlgorithmEnum) Get() *MetricReportDefinitionV140CalculationAlgorithmEnum {
	return v.value
}

func (v *NullableMetricReportDefinitionV140CalculationAlgorithmEnum) Set(val *MetricReportDefinitionV140CalculationAlgorithmEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricReportDefinitionV140CalculationAlgorithmEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricReportDefinitionV140CalculationAlgorithmEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricReportDefinitionV140CalculationAlgorithmEnum(val *MetricReportDefinitionV140CalculationAlgorithmEnum) *NullableMetricReportDefinitionV140CalculationAlgorithmEnum {
	return &NullableMetricReportDefinitionV140CalculationAlgorithmEnum{value: val, isSet: true}
}

func (v NullableMetricReportDefinitionV140CalculationAlgorithmEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricReportDefinitionV140CalculationAlgorithmEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

