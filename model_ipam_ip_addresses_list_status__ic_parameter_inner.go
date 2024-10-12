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

// IpamIpAddressesListStatusIcParameterInner the model 'IpamIpAddressesListStatusIcParameterInner'
type IpamIpAddressesListStatusIcParameterInner string

// List of ipam_ip_addresses_list_status__ic_parameter_inner
const (
	IPAMIPADDRESSESLISTSTATUSICPARAMETERINNER_ACTIVE     IpamIpAddressesListStatusIcParameterInner = "active"
	IPAMIPADDRESSESLISTSTATUSICPARAMETERINNER_DEPRECATED IpamIpAddressesListStatusIcParameterInner = "deprecated"
	IPAMIPADDRESSESLISTSTATUSICPARAMETERINNER_DHCP       IpamIpAddressesListStatusIcParameterInner = "dhcp"
	IPAMIPADDRESSESLISTSTATUSICPARAMETERINNER_RESERVED   IpamIpAddressesListStatusIcParameterInner = "reserved"
	IPAMIPADDRESSESLISTSTATUSICPARAMETERINNER_SLAAC      IpamIpAddressesListStatusIcParameterInner = "slaac"
)

// All allowed values of IpamIpAddressesListStatusIcParameterInner enum
var AllowedIpamIpAddressesListStatusIcParameterInnerEnumValues = []IpamIpAddressesListStatusIcParameterInner{
	"active",
	"deprecated",
	"dhcp",
	"reserved",
	"slaac",
}

func (v *IpamIpAddressesListStatusIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpamIpAddressesListStatusIcParameterInner(value)
	for _, existing := range AllowedIpamIpAddressesListStatusIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpamIpAddressesListStatusIcParameterInner", value)
}

// NewIpamIpAddressesListStatusIcParameterInnerFromValue returns a pointer to a valid IpamIpAddressesListStatusIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpamIpAddressesListStatusIcParameterInnerFromValue(v string) (*IpamIpAddressesListStatusIcParameterInner, error) {
	ev := IpamIpAddressesListStatusIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpamIpAddressesListStatusIcParameterInner: valid values are %v", v, AllowedIpamIpAddressesListStatusIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpamIpAddressesListStatusIcParameterInner) IsValid() bool {
	for _, existing := range AllowedIpamIpAddressesListStatusIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ipam_ip_addresses_list_status__ic_parameter_inner value
func (v IpamIpAddressesListStatusIcParameterInner) Ptr() *IpamIpAddressesListStatusIcParameterInner {
	return &v
}

type NullableIpamIpAddressesListStatusIcParameterInner struct {
	value *IpamIpAddressesListStatusIcParameterInner
	isSet bool
}

func (v NullableIpamIpAddressesListStatusIcParameterInner) Get() *IpamIpAddressesListStatusIcParameterInner {
	return v.value
}

func (v *NullableIpamIpAddressesListStatusIcParameterInner) Set(val *IpamIpAddressesListStatusIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamIpAddressesListStatusIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamIpAddressesListStatusIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamIpAddressesListStatusIcParameterInner(val *IpamIpAddressesListStatusIcParameterInner) *NullableIpamIpAddressesListStatusIcParameterInner {
	return &NullableIpamIpAddressesListStatusIcParameterInner{value: val, isSet: true}
}

func (v NullableIpamIpAddressesListStatusIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamIpAddressesListStatusIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
