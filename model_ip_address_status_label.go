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

// IPAddressStatusLabel the model 'IPAddressStatusLabel'
type IPAddressStatusLabel string

// List of IPAddress_status_label
const (
	IPADDRESSSTATUSLABEL_ACTIVE     IPAddressStatusLabel = "Active"
	IPADDRESSSTATUSLABEL_RESERVED   IPAddressStatusLabel = "Reserved"
	IPADDRESSSTATUSLABEL_DEPRECATED IPAddressStatusLabel = "Deprecated"
	IPADDRESSSTATUSLABEL_DHCP       IPAddressStatusLabel = "DHCP"
	IPADDRESSSTATUSLABEL_SLAAC      IPAddressStatusLabel = "SLAAC"
)

// All allowed values of IPAddressStatusLabel enum
var AllowedIPAddressStatusLabelEnumValues = []IPAddressStatusLabel{
	"Active",
	"Reserved",
	"Deprecated",
	"DHCP",
	"SLAAC",
}

func (v *IPAddressStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPAddressStatusLabel(value)
	for _, existing := range AllowedIPAddressStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPAddressStatusLabel", value)
}

// NewIPAddressStatusLabelFromValue returns a pointer to a valid IPAddressStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPAddressStatusLabelFromValue(v string) (*IPAddressStatusLabel, error) {
	ev := IPAddressStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPAddressStatusLabel: valid values are %v", v, AllowedIPAddressStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPAddressStatusLabel) IsValid() bool {
	for _, existing := range AllowedIPAddressStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPAddress_status_label value
func (v IPAddressStatusLabel) Ptr() *IPAddressStatusLabel {
	return &v
}

type NullableIPAddressStatusLabel struct {
	value *IPAddressStatusLabel
	isSet bool
}

func (v NullableIPAddressStatusLabel) Get() *IPAddressStatusLabel {
	return v.value
}

func (v *NullableIPAddressStatusLabel) Set(val *IPAddressStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressStatusLabel(val *IPAddressStatusLabel) *NullableIPAddressStatusLabel {
	return &NullableIPAddressStatusLabel{value: val, isSet: true}
}

func (v NullableIPAddressStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
