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

// WirelessWirelessLinksListDistanceUnitParameter the model 'WirelessWirelessLinksListDistanceUnitParameter'
type WirelessWirelessLinksListDistanceUnitParameter string

// List of wireless_wireless_links_list_distance_unit_parameter
const (
	WIRELESSWIRELESSLINKSLISTDISTANCEUNITPARAMETER_FT WirelessWirelessLinksListDistanceUnitParameter = "ft"
	WIRELESSWIRELESSLINKSLISTDISTANCEUNITPARAMETER_KM WirelessWirelessLinksListDistanceUnitParameter = "km"
	WIRELESSWIRELESSLINKSLISTDISTANCEUNITPARAMETER_M  WirelessWirelessLinksListDistanceUnitParameter = "m"
	WIRELESSWIRELESSLINKSLISTDISTANCEUNITPARAMETER_MI WirelessWirelessLinksListDistanceUnitParameter = "mi"
)

// All allowed values of WirelessWirelessLinksListDistanceUnitParameter enum
var AllowedWirelessWirelessLinksListDistanceUnitParameterEnumValues = []WirelessWirelessLinksListDistanceUnitParameter{
	"ft",
	"km",
	"m",
	"mi",
}

func (v *WirelessWirelessLinksListDistanceUnitParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessWirelessLinksListDistanceUnitParameter(value)
	for _, existing := range AllowedWirelessWirelessLinksListDistanceUnitParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessWirelessLinksListDistanceUnitParameter", value)
}

// NewWirelessWirelessLinksListDistanceUnitParameterFromValue returns a pointer to a valid WirelessWirelessLinksListDistanceUnitParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessWirelessLinksListDistanceUnitParameterFromValue(v string) (*WirelessWirelessLinksListDistanceUnitParameter, error) {
	ev := WirelessWirelessLinksListDistanceUnitParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessWirelessLinksListDistanceUnitParameter: valid values are %v", v, AllowedWirelessWirelessLinksListDistanceUnitParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessWirelessLinksListDistanceUnitParameter) IsValid() bool {
	for _, existing := range AllowedWirelessWirelessLinksListDistanceUnitParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wireless_wireless_links_list_distance_unit_parameter value
func (v WirelessWirelessLinksListDistanceUnitParameter) Ptr() *WirelessWirelessLinksListDistanceUnitParameter {
	return &v
}

type NullableWirelessWirelessLinksListDistanceUnitParameter struct {
	value *WirelessWirelessLinksListDistanceUnitParameter
	isSet bool
}

func (v NullableWirelessWirelessLinksListDistanceUnitParameter) Get() *WirelessWirelessLinksListDistanceUnitParameter {
	return v.value
}

func (v *NullableWirelessWirelessLinksListDistanceUnitParameter) Set(val *WirelessWirelessLinksListDistanceUnitParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessWirelessLinksListDistanceUnitParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessWirelessLinksListDistanceUnitParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessWirelessLinksListDistanceUnitParameter(val *WirelessWirelessLinksListDistanceUnitParameter) *NullableWirelessWirelessLinksListDistanceUnitParameter {
	return &NullableWirelessWirelessLinksListDistanceUnitParameter{value: val, isSet: true}
}

func (v NullableWirelessWirelessLinksListDistanceUnitParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessWirelessLinksListDistanceUnitParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
