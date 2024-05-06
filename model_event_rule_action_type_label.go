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

// EventRuleActionTypeLabel the model 'EventRuleActionTypeLabel'
type EventRuleActionTypeLabel string

// List of EventRule_action_type_label
const (
	EVENTRULEACTIONTYPELABEL_WEBHOOK EventRuleActionTypeLabel = "Webhook"
	EVENTRULEACTIONTYPELABEL_SCRIPT  EventRuleActionTypeLabel = "Script"
)

// All allowed values of EventRuleActionTypeLabel enum
var AllowedEventRuleActionTypeLabelEnumValues = []EventRuleActionTypeLabel{
	"Webhook",
	"Script",
}

func (v *EventRuleActionTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventRuleActionTypeLabel(value)
	for _, existing := range AllowedEventRuleActionTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventRuleActionTypeLabel", value)
}

// NewEventRuleActionTypeLabelFromValue returns a pointer to a valid EventRuleActionTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventRuleActionTypeLabelFromValue(v string) (*EventRuleActionTypeLabel, error) {
	ev := EventRuleActionTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventRuleActionTypeLabel: valid values are %v", v, AllowedEventRuleActionTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventRuleActionTypeLabel) IsValid() bool {
	for _, existing := range AllowedEventRuleActionTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventRule_action_type_label value
func (v EventRuleActionTypeLabel) Ptr() *EventRuleActionTypeLabel {
	return &v
}

type NullableEventRuleActionTypeLabel struct {
	value *EventRuleActionTypeLabel
	isSet bool
}

func (v NullableEventRuleActionTypeLabel) Get() *EventRuleActionTypeLabel {
	return v.value
}

func (v *NullableEventRuleActionTypeLabel) Set(val *EventRuleActionTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRuleActionTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRuleActionTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRuleActionTypeLabel(val *EventRuleActionTypeLabel) *NullableEventRuleActionTypeLabel {
	return &NullableEventRuleActionTypeLabel{value: val, isSet: true}
}

func (v NullableEventRuleActionTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRuleActionTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
