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

// PatchedWritableCustomFieldRequestUiVisible Specifies whether the custom field is displayed in the UI  * `always` - Always * `if-set` - If set * `hidden` - Hidden
type PatchedWritableCustomFieldRequestUiVisible string

// List of PatchedWritableCustomFieldRequest_ui_visible
const (
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBLE_ALWAYS PatchedWritableCustomFieldRequestUiVisible = "always"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBLE_IF_SET PatchedWritableCustomFieldRequestUiVisible = "if-set"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBLE_HIDDEN PatchedWritableCustomFieldRequestUiVisible = "hidden"
)

// All allowed values of PatchedWritableCustomFieldRequestUiVisible enum
var AllowedPatchedWritableCustomFieldRequestUiVisibleEnumValues = []PatchedWritableCustomFieldRequestUiVisible{
	"always",
	"if-set",
	"hidden",
}

func (v *PatchedWritableCustomFieldRequestUiVisible) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableCustomFieldRequestUiVisible(value)
	for _, existing := range AllowedPatchedWritableCustomFieldRequestUiVisibleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableCustomFieldRequestUiVisible", value)
}

// NewPatchedWritableCustomFieldRequestUiVisibleFromValue returns a pointer to a valid PatchedWritableCustomFieldRequestUiVisible
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableCustomFieldRequestUiVisibleFromValue(v string) (*PatchedWritableCustomFieldRequestUiVisible, error) {
	ev := PatchedWritableCustomFieldRequestUiVisible(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableCustomFieldRequestUiVisible: valid values are %v", v, AllowedPatchedWritableCustomFieldRequestUiVisibleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableCustomFieldRequestUiVisible) IsValid() bool {
	for _, existing := range AllowedPatchedWritableCustomFieldRequestUiVisibleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableCustomFieldRequest_ui_visible value
func (v PatchedWritableCustomFieldRequestUiVisible) Ptr() *PatchedWritableCustomFieldRequestUiVisible {
	return &v
}

type NullablePatchedWritableCustomFieldRequestUiVisible struct {
	value *PatchedWritableCustomFieldRequestUiVisible
	isSet bool
}

func (v NullablePatchedWritableCustomFieldRequestUiVisible) Get() *PatchedWritableCustomFieldRequestUiVisible {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldRequestUiVisible) Set(val *PatchedWritableCustomFieldRequestUiVisible) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldRequestUiVisible) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldRequestUiVisible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldRequestUiVisible(val *PatchedWritableCustomFieldRequestUiVisible) *NullablePatchedWritableCustomFieldRequestUiVisible {
	return &NullablePatchedWritableCustomFieldRequestUiVisible{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldRequestUiVisible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldRequestUiVisible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
