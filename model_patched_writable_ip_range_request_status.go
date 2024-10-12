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

// PatchedWritableIPRangeRequestStatus Operational status of this range  * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
type PatchedWritableIPRangeRequestStatus string

// List of PatchedWritableIPRangeRequest_status
const (
	PATCHEDWRITABLEIPRANGEREQUESTSTATUS_ACTIVE     PatchedWritableIPRangeRequestStatus = "active"
	PATCHEDWRITABLEIPRANGEREQUESTSTATUS_RESERVED   PatchedWritableIPRangeRequestStatus = "reserved"
	PATCHEDWRITABLEIPRANGEREQUESTSTATUS_DEPRECATED PatchedWritableIPRangeRequestStatus = "deprecated"
)

// All allowed values of PatchedWritableIPRangeRequestStatus enum
var AllowedPatchedWritableIPRangeRequestStatusEnumValues = []PatchedWritableIPRangeRequestStatus{
	"active",
	"reserved",
	"deprecated",
}

func (v *PatchedWritableIPRangeRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableIPRangeRequestStatus(value)
	for _, existing := range AllowedPatchedWritableIPRangeRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableIPRangeRequestStatus", value)
}

// NewPatchedWritableIPRangeRequestStatusFromValue returns a pointer to a valid PatchedWritableIPRangeRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableIPRangeRequestStatusFromValue(v string) (*PatchedWritableIPRangeRequestStatus, error) {
	ev := PatchedWritableIPRangeRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableIPRangeRequestStatus: valid values are %v", v, AllowedPatchedWritableIPRangeRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableIPRangeRequestStatus) IsValid() bool {
	for _, existing := range AllowedPatchedWritableIPRangeRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableIPRangeRequest_status value
func (v PatchedWritableIPRangeRequestStatus) Ptr() *PatchedWritableIPRangeRequestStatus {
	return &v
}

type NullablePatchedWritableIPRangeRequestStatus struct {
	value *PatchedWritableIPRangeRequestStatus
	isSet bool
}

func (v NullablePatchedWritableIPRangeRequestStatus) Get() *PatchedWritableIPRangeRequestStatus {
	return v.value
}

func (v *NullablePatchedWritableIPRangeRequestStatus) Set(val *PatchedWritableIPRangeRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIPRangeRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIPRangeRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIPRangeRequestStatus(val *PatchedWritableIPRangeRequestStatus) *NullablePatchedWritableIPRangeRequestStatus {
	return &NullablePatchedWritableIPRangeRequestStatus{value: val, isSet: true}
}

func (v NullablePatchedWritableIPRangeRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIPRangeRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
