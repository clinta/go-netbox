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

// DeviceAirflowLabel the model 'DeviceAirflowLabel'
type DeviceAirflowLabel string

// List of Device_airflow_label
const (
	DEVICEAIRFLOWLABEL_FRONT_TO_REAR DeviceAirflowLabel = "Front to rear"
	DEVICEAIRFLOWLABEL_REAR_TO_FRONT DeviceAirflowLabel = "Rear to front"
	DEVICEAIRFLOWLABEL_LEFT_TO_RIGHT DeviceAirflowLabel = "Left to right"
	DEVICEAIRFLOWLABEL_RIGHT_TO_LEFT DeviceAirflowLabel = "Right to left"
	DEVICEAIRFLOWLABEL_SIDE_TO_REAR  DeviceAirflowLabel = "Side to rear"
	DEVICEAIRFLOWLABEL_PASSIVE       DeviceAirflowLabel = "Passive"
	DEVICEAIRFLOWLABEL_MIXED         DeviceAirflowLabel = "Mixed"
)

// All allowed values of DeviceAirflowLabel enum
var AllowedDeviceAirflowLabelEnumValues = []DeviceAirflowLabel{
	"Front to rear",
	"Rear to front",
	"Left to right",
	"Right to left",
	"Side to rear",
	"Passive",
	"Mixed",
}

func (v *DeviceAirflowLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceAirflowLabel(value)
	for _, existing := range AllowedDeviceAirflowLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceAirflowLabel", value)
}

// NewDeviceAirflowLabelFromValue returns a pointer to a valid DeviceAirflowLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceAirflowLabelFromValue(v string) (*DeviceAirflowLabel, error) {
	ev := DeviceAirflowLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceAirflowLabel: valid values are %v", v, AllowedDeviceAirflowLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceAirflowLabel) IsValid() bool {
	for _, existing := range AllowedDeviceAirflowLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Device_airflow_label value
func (v DeviceAirflowLabel) Ptr() *DeviceAirflowLabel {
	return &v
}

type NullableDeviceAirflowLabel struct {
	value *DeviceAirflowLabel
	isSet bool
}

func (v NullableDeviceAirflowLabel) Get() *DeviceAirflowLabel {
	return v.value
}

func (v *NullableDeviceAirflowLabel) Set(val *DeviceAirflowLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAirflowLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAirflowLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAirflowLabel(val *DeviceAirflowLabel) *NullableDeviceAirflowLabel {
	return &NullableDeviceAirflowLabel{value: val, isSet: true}
}

func (v NullableDeviceAirflowLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAirflowLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
