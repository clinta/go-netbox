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

// DeviceTypeSubdeviceRoleValue * `parent` - Parent * `child` - Child
type DeviceTypeSubdeviceRoleValue string

// List of DeviceType_subdevice_role_value
const (
	DEVICETYPESUBDEVICEROLEVALUE_PARENT DeviceTypeSubdeviceRoleValue = "parent"
	DEVICETYPESUBDEVICEROLEVALUE_CHILD  DeviceTypeSubdeviceRoleValue = "child"
	DEVICETYPESUBDEVICEROLEVALUE_EMPTY  DeviceTypeSubdeviceRoleValue = ""
)

// All allowed values of DeviceTypeSubdeviceRoleValue enum
var AllowedDeviceTypeSubdeviceRoleValueEnumValues = []DeviceTypeSubdeviceRoleValue{
	"parent",
	"child",
	"",
}

func (v *DeviceTypeSubdeviceRoleValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeSubdeviceRoleValue(value)
	for _, existing := range AllowedDeviceTypeSubdeviceRoleValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeSubdeviceRoleValue", value)
}

// NewDeviceTypeSubdeviceRoleValueFromValue returns a pointer to a valid DeviceTypeSubdeviceRoleValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeSubdeviceRoleValueFromValue(v string) (*DeviceTypeSubdeviceRoleValue, error) {
	ev := DeviceTypeSubdeviceRoleValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeSubdeviceRoleValue: valid values are %v", v, AllowedDeviceTypeSubdeviceRoleValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeSubdeviceRoleValue) IsValid() bool {
	for _, existing := range AllowedDeviceTypeSubdeviceRoleValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceType_subdevice_role_value value
func (v DeviceTypeSubdeviceRoleValue) Ptr() *DeviceTypeSubdeviceRoleValue {
	return &v
}

type NullableDeviceTypeSubdeviceRoleValue struct {
	value *DeviceTypeSubdeviceRoleValue
	isSet bool
}

func (v NullableDeviceTypeSubdeviceRoleValue) Get() *DeviceTypeSubdeviceRoleValue {
	return v.value
}

func (v *NullableDeviceTypeSubdeviceRoleValue) Set(val *DeviceTypeSubdeviceRoleValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeSubdeviceRoleValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeSubdeviceRoleValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeSubdeviceRoleValue(val *DeviceTypeSubdeviceRoleValue) *NullableDeviceTypeSubdeviceRoleValue {
	return &NullableDeviceTypeSubdeviceRoleValue{value: val, isSet: true}
}

func (v NullableDeviceTypeSubdeviceRoleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeSubdeviceRoleValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
