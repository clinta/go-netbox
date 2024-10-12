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

// BriefJobStatusLabel the model 'BriefJobStatusLabel'
type BriefJobStatusLabel string

// List of BriefJob_status_label
const (
	BRIEFJOBSTATUSLABEL_PENDING   BriefJobStatusLabel = "Pending"
	BRIEFJOBSTATUSLABEL_SCHEDULED BriefJobStatusLabel = "Scheduled"
	BRIEFJOBSTATUSLABEL_RUNNING   BriefJobStatusLabel = "Running"
	BRIEFJOBSTATUSLABEL_COMPLETED BriefJobStatusLabel = "Completed"
	BRIEFJOBSTATUSLABEL_ERRORED   BriefJobStatusLabel = "Errored"
	BRIEFJOBSTATUSLABEL_FAILED    BriefJobStatusLabel = "Failed"
)

// All allowed values of BriefJobStatusLabel enum
var AllowedBriefJobStatusLabelEnumValues = []BriefJobStatusLabel{
	"Pending",
	"Scheduled",
	"Running",
	"Completed",
	"Errored",
	"Failed",
}

func (v *BriefJobStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BriefJobStatusLabel(value)
	for _, existing := range AllowedBriefJobStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BriefJobStatusLabel", value)
}

// NewBriefJobStatusLabelFromValue returns a pointer to a valid BriefJobStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBriefJobStatusLabelFromValue(v string) (*BriefJobStatusLabel, error) {
	ev := BriefJobStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BriefJobStatusLabel: valid values are %v", v, AllowedBriefJobStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BriefJobStatusLabel) IsValid() bool {
	for _, existing := range AllowedBriefJobStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BriefJob_status_label value
func (v BriefJobStatusLabel) Ptr() *BriefJobStatusLabel {
	return &v
}

type NullableBriefJobStatusLabel struct {
	value *BriefJobStatusLabel
	isSet bool
}

func (v NullableBriefJobStatusLabel) Get() *BriefJobStatusLabel {
	return v.value
}

func (v *NullableBriefJobStatusLabel) Set(val *BriefJobStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefJobStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefJobStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefJobStatusLabel(val *BriefJobStatusLabel) *NullableBriefJobStatusLabel {
	return &NullableBriefJobStatusLabel{value: val, isSet: true}
}

func (v NullableBriefJobStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefJobStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
