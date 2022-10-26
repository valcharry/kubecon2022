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

// ScheduleV122DayOfWeek Days of the week.
type ScheduleV122DayOfWeek string

// List of Schedule_v1_2_2_DayOfWeek
const (
	SCHEDULEV122DAYOFWEEK_MONDAY ScheduleV122DayOfWeek = "Monday"
	SCHEDULEV122DAYOFWEEK_TUESDAY ScheduleV122DayOfWeek = "Tuesday"
	SCHEDULEV122DAYOFWEEK_WEDNESDAY ScheduleV122DayOfWeek = "Wednesday"
	SCHEDULEV122DAYOFWEEK_THURSDAY ScheduleV122DayOfWeek = "Thursday"
	SCHEDULEV122DAYOFWEEK_FRIDAY ScheduleV122DayOfWeek = "Friday"
	SCHEDULEV122DAYOFWEEK_SATURDAY ScheduleV122DayOfWeek = "Saturday"
	SCHEDULEV122DAYOFWEEK_SUNDAY ScheduleV122DayOfWeek = "Sunday"
	SCHEDULEV122DAYOFWEEK_EVERY ScheduleV122DayOfWeek = "Every"
)

func (v *ScheduleV122DayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleV122DayOfWeek(value)
	for _, existing := range []ScheduleV122DayOfWeek{ "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Every",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleV122DayOfWeek", value)
}

// Ptr returns reference to Schedule_v1_2_2_DayOfWeek value
func (v ScheduleV122DayOfWeek) Ptr() *ScheduleV122DayOfWeek {
	return &v
}

type NullableScheduleV122DayOfWeek struct {
	value *ScheduleV122DayOfWeek
	isSet bool
}

func (v NullableScheduleV122DayOfWeek) Get() *ScheduleV122DayOfWeek {
	return v.value
}

func (v *NullableScheduleV122DayOfWeek) Set(val *ScheduleV122DayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleV122DayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleV122DayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleV122DayOfWeek(val *ScheduleV122DayOfWeek) *NullableScheduleV122DayOfWeek {
	return &NullableScheduleV122DayOfWeek{value: val, isSet: true}
}

func (v NullableScheduleV122DayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleV122DayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
