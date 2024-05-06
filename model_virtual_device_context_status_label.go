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

// VirtualDeviceContextStatusLabel the model 'VirtualDeviceContextStatusLabel'
type VirtualDeviceContextStatusLabel string

// List of VirtualDeviceContext_status_label
const (
	VIRTUALDEVICECONTEXTSTATUSLABEL_ACTIVE  VirtualDeviceContextStatusLabel = "Active"
	VIRTUALDEVICECONTEXTSTATUSLABEL_PLANNED VirtualDeviceContextStatusLabel = "Planned"
	VIRTUALDEVICECONTEXTSTATUSLABEL_OFFLINE VirtualDeviceContextStatusLabel = "Offline"
)

// All allowed values of VirtualDeviceContextStatusLabel enum
var AllowedVirtualDeviceContextStatusLabelEnumValues = []VirtualDeviceContextStatusLabel{
	"Active",
	"Planned",
	"Offline",
}

func (v *VirtualDeviceContextStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualDeviceContextStatusLabel(value)
	for _, existing := range AllowedVirtualDeviceContextStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualDeviceContextStatusLabel", value)
}

// NewVirtualDeviceContextStatusLabelFromValue returns a pointer to a valid VirtualDeviceContextStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualDeviceContextStatusLabelFromValue(v string) (*VirtualDeviceContextStatusLabel, error) {
	ev := VirtualDeviceContextStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualDeviceContextStatusLabel: valid values are %v", v, AllowedVirtualDeviceContextStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualDeviceContextStatusLabel) IsValid() bool {
	for _, existing := range AllowedVirtualDeviceContextStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualDeviceContext_status_label value
func (v VirtualDeviceContextStatusLabel) Ptr() *VirtualDeviceContextStatusLabel {
	return &v
}

type NullableVirtualDeviceContextStatusLabel struct {
	value *VirtualDeviceContextStatusLabel
	isSet bool
}

func (v NullableVirtualDeviceContextStatusLabel) Get() *VirtualDeviceContextStatusLabel {
	return v.value
}

func (v *NullableVirtualDeviceContextStatusLabel) Set(val *VirtualDeviceContextStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualDeviceContextStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualDeviceContextStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualDeviceContextStatusLabel(val *VirtualDeviceContextStatusLabel) *NullableVirtualDeviceContextStatusLabel {
	return &NullableVirtualDeviceContextStatusLabel{value: val, isSet: true}
}

func (v NullableVirtualDeviceContextStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualDeviceContextStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
