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

// PatchedWritableServiceRequestProtocol * `tcp` - TCP * `udp` - UDP * `sctp` - SCTP
type PatchedWritableServiceRequestProtocol string

// List of PatchedWritableServiceRequest_protocol
const (
	PATCHEDWRITABLESERVICEREQUESTPROTOCOL_TCP  PatchedWritableServiceRequestProtocol = "tcp"
	PATCHEDWRITABLESERVICEREQUESTPROTOCOL_UDP  PatchedWritableServiceRequestProtocol = "udp"
	PATCHEDWRITABLESERVICEREQUESTPROTOCOL_SCTP PatchedWritableServiceRequestProtocol = "sctp"
)

// All allowed values of PatchedWritableServiceRequestProtocol enum
var AllowedPatchedWritableServiceRequestProtocolEnumValues = []PatchedWritableServiceRequestProtocol{
	"tcp",
	"udp",
	"sctp",
}

func (v *PatchedWritableServiceRequestProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableServiceRequestProtocol(value)
	for _, existing := range AllowedPatchedWritableServiceRequestProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableServiceRequestProtocol", value)
}

// NewPatchedWritableServiceRequestProtocolFromValue returns a pointer to a valid PatchedWritableServiceRequestProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableServiceRequestProtocolFromValue(v string) (*PatchedWritableServiceRequestProtocol, error) {
	ev := PatchedWritableServiceRequestProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableServiceRequestProtocol: valid values are %v", v, AllowedPatchedWritableServiceRequestProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableServiceRequestProtocol) IsValid() bool {
	for _, existing := range AllowedPatchedWritableServiceRequestProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableServiceRequest_protocol value
func (v PatchedWritableServiceRequestProtocol) Ptr() *PatchedWritableServiceRequestProtocol {
	return &v
}

type NullablePatchedWritableServiceRequestProtocol struct {
	value *PatchedWritableServiceRequestProtocol
	isSet bool
}

func (v NullablePatchedWritableServiceRequestProtocol) Get() *PatchedWritableServiceRequestProtocol {
	return v.value
}

func (v *NullablePatchedWritableServiceRequestProtocol) Set(val *PatchedWritableServiceRequestProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableServiceRequestProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableServiceRequestProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableServiceRequestProtocol(val *PatchedWritableServiceRequestProtocol) *NullablePatchedWritableServiceRequestProtocol {
	return &NullablePatchedWritableServiceRequestProtocol{value: val, isSet: true}
}

func (v NullablePatchedWritableServiceRequestProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableServiceRequestProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
