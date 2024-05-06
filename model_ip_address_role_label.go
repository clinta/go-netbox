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

// IPAddressRoleLabel the model 'IPAddressRoleLabel'
type IPAddressRoleLabel string

// List of IPAddress_role_label
const (
	IPADDRESSROLELABEL_LOOPBACK  IPAddressRoleLabel = "Loopback"
	IPADDRESSROLELABEL_SECONDARY IPAddressRoleLabel = "Secondary"
	IPADDRESSROLELABEL_ANYCAST   IPAddressRoleLabel = "Anycast"
	IPADDRESSROLELABEL_VIP       IPAddressRoleLabel = "VIP"
	IPADDRESSROLELABEL_VRRP      IPAddressRoleLabel = "VRRP"
	IPADDRESSROLELABEL_HSRP      IPAddressRoleLabel = "HSRP"
	IPADDRESSROLELABEL_GLBP      IPAddressRoleLabel = "GLBP"
	IPADDRESSROLELABEL_CARP      IPAddressRoleLabel = "CARP"
)

// All allowed values of IPAddressRoleLabel enum
var AllowedIPAddressRoleLabelEnumValues = []IPAddressRoleLabel{
	"Loopback",
	"Secondary",
	"Anycast",
	"VIP",
	"VRRP",
	"HSRP",
	"GLBP",
	"CARP",
}

func (v *IPAddressRoleLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPAddressRoleLabel(value)
	for _, existing := range AllowedIPAddressRoleLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPAddressRoleLabel", value)
}

// NewIPAddressRoleLabelFromValue returns a pointer to a valid IPAddressRoleLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPAddressRoleLabelFromValue(v string) (*IPAddressRoleLabel, error) {
	ev := IPAddressRoleLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPAddressRoleLabel: valid values are %v", v, AllowedIPAddressRoleLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPAddressRoleLabel) IsValid() bool {
	for _, existing := range AllowedIPAddressRoleLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPAddress_role_label value
func (v IPAddressRoleLabel) Ptr() *IPAddressRoleLabel {
	return &v
}

type NullableIPAddressRoleLabel struct {
	value *IPAddressRoleLabel
	isSet bool
}

func (v NullableIPAddressRoleLabel) Get() *IPAddressRoleLabel {
	return v.value
}

func (v *NullableIPAddressRoleLabel) Set(val *IPAddressRoleLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressRoleLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressRoleLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressRoleLabel(val *IPAddressRoleLabel) *NullableIPAddressRoleLabel {
	return &NullableIPAddressRoleLabel{value: val, isSet: true}
}

func (v NullableIPAddressRoleLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressRoleLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
