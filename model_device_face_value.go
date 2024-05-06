/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.7 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// DeviceFaceValue * `front` - Front * `rear` - Rear
type DeviceFaceValue string

// List of Device_face_value
const (
	DEVICEFACEVALUE_FRONT DeviceFaceValue = "front"
	DEVICEFACEVALUE_REAR  DeviceFaceValue = "rear"
	DEVICEFACEVALUE_EMPTY DeviceFaceValue = ""
)

// All allowed values of DeviceFaceValue enum
var AllowedDeviceFaceValueEnumValues = []DeviceFaceValue{
	"front",
	"rear",
	"",
}

func (v *DeviceFaceValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceFaceValue(value)
	for _, existing := range AllowedDeviceFaceValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceFaceValue", value)
}

// NewDeviceFaceValueFromValue returns a pointer to a valid DeviceFaceValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceFaceValueFromValue(v string) (*DeviceFaceValue, error) {
	ev := DeviceFaceValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceFaceValue: valid values are %v", v, AllowedDeviceFaceValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceFaceValue) IsValid() bool {
	for _, existing := range AllowedDeviceFaceValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Device_face_value value
func (v DeviceFaceValue) Ptr() *DeviceFaceValue {
	return &v
}

type NullableDeviceFaceValue struct {
	value *DeviceFaceValue
	isSet bool
}

func (v NullableDeviceFaceValue) Get() *DeviceFaceValue {
	return v.value
}

func (v *NullableDeviceFaceValue) Set(val *DeviceFaceValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceFaceValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceFaceValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceFaceValue(val *DeviceFaceValue) *NullableDeviceFaceValue {
	return &NullableDeviceFaceValue{value: val, isSet: true}
}

func (v NullableDeviceFaceValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceFaceValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
