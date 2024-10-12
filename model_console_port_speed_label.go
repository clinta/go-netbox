/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// ConsolePortSpeedLabel the model 'ConsolePortSpeedLabel'
type ConsolePortSpeedLabel string

// List of ConsolePort_speed_label
const (
	CONSOLEPORTSPEEDLABEL__1200_BPS   ConsolePortSpeedLabel = "1200 bps"
	CONSOLEPORTSPEEDLABEL__2400_BPS   ConsolePortSpeedLabel = "2400 bps"
	CONSOLEPORTSPEEDLABEL__4800_BPS   ConsolePortSpeedLabel = "4800 bps"
	CONSOLEPORTSPEEDLABEL__9600_BPS   ConsolePortSpeedLabel = "9600 bps"
	CONSOLEPORTSPEEDLABEL__19_2_KBPS  ConsolePortSpeedLabel = "19.2 kbps"
	CONSOLEPORTSPEEDLABEL__38_4_KBPS  ConsolePortSpeedLabel = "38.4 kbps"
	CONSOLEPORTSPEEDLABEL__57_6_KBPS  ConsolePortSpeedLabel = "57.6 kbps"
	CONSOLEPORTSPEEDLABEL__115_2_KBPS ConsolePortSpeedLabel = "115.2 kbps"
)

// All allowed values of ConsolePortSpeedLabel enum
var AllowedConsolePortSpeedLabelEnumValues = []ConsolePortSpeedLabel{
	"1200 bps",
	"2400 bps",
	"4800 bps",
	"9600 bps",
	"19.2 kbps",
	"38.4 kbps",
	"57.6 kbps",
	"115.2 kbps",
}

func (v *ConsolePortSpeedLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsolePortSpeedLabel(value)
	for _, existing := range AllowedConsolePortSpeedLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsolePortSpeedLabel", value)
}

// NewConsolePortSpeedLabelFromValue returns a pointer to a valid ConsolePortSpeedLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsolePortSpeedLabelFromValue(v string) (*ConsolePortSpeedLabel, error) {
	ev := ConsolePortSpeedLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsolePortSpeedLabel: valid values are %v", v, AllowedConsolePortSpeedLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsolePortSpeedLabel) IsValid() bool {
	for _, existing := range AllowedConsolePortSpeedLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsolePort_speed_label value
func (v ConsolePortSpeedLabel) Ptr() *ConsolePortSpeedLabel {
	return &v
}

type NullableConsolePortSpeedLabel struct {
	value *ConsolePortSpeedLabel
	isSet bool
}

func (v NullableConsolePortSpeedLabel) Get() *ConsolePortSpeedLabel {
	return v.value
}

func (v *NullableConsolePortSpeedLabel) Set(val *ConsolePortSpeedLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortSpeedLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortSpeedLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortSpeedLabel(val *ConsolePortSpeedLabel) *NullableConsolePortSpeedLabel {
	return &NullableConsolePortSpeedLabel{value: val, isSet: true}
}

func (v NullableConsolePortSpeedLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortSpeedLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
