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

// ScheduleV101DayOfWeek Days of the week.
type ScheduleV101DayOfWeek string

// List of Schedule_v1_0_1_DayOfWeek
const (
	SCHEDULEV101DAYOFWEEK_MONDAY ScheduleV101DayOfWeek = "Monday"
	SCHEDULEV101DAYOFWEEK_TUESDAY ScheduleV101DayOfWeek = "Tuesday"
	SCHEDULEV101DAYOFWEEK_WEDNESDAY ScheduleV101DayOfWeek = "Wednesday"
	SCHEDULEV101DAYOFWEEK_THURSDAY ScheduleV101DayOfWeek = "Thursday"
	SCHEDULEV101DAYOFWEEK_FRIDAY ScheduleV101DayOfWeek = "Friday"
	SCHEDULEV101DAYOFWEEK_SATURDAY ScheduleV101DayOfWeek = "Saturday"
	SCHEDULEV101DAYOFWEEK_SUNDAY ScheduleV101DayOfWeek = "Sunday"
	SCHEDULEV101DAYOFWEEK_EVERY ScheduleV101DayOfWeek = "Every"
)

func (v *ScheduleV101DayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleV101DayOfWeek(value)
	for _, existing := range []ScheduleV101DayOfWeek{ "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Every",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleV101DayOfWeek", value)
}

// Ptr returns reference to Schedule_v1_0_1_DayOfWeek value
func (v ScheduleV101DayOfWeek) Ptr() *ScheduleV101DayOfWeek {
	return &v
}

type NullableScheduleV101DayOfWeek struct {
	value *ScheduleV101DayOfWeek
	isSet bool
}

func (v NullableScheduleV101DayOfWeek) Get() *ScheduleV101DayOfWeek {
	return v.value
}

func (v *NullableScheduleV101DayOfWeek) Set(val *ScheduleV101DayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleV101DayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleV101DayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleV101DayOfWeek(val *ScheduleV101DayOfWeek) *NullableScheduleV101DayOfWeek {
	return &NullableScheduleV101DayOfWeek{value: val, isSet: true}
}

func (v NullableScheduleV101DayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleV101DayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

