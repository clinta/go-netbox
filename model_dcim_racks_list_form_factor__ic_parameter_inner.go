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

// DcimRacksListFormFactorIcParameterInner the model 'DcimRacksListFormFactorIcParameterInner'
type DcimRacksListFormFactorIcParameterInner string

// List of dcim_racks_list_form_factor__ic_parameter_inner
const (
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER_EMPTY                 DcimRacksListFormFactorIcParameterInner = ""
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER__2_POST_FRAME         DcimRacksListFormFactorIcParameterInner = "2-post-frame"
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER__4_POST_CABINET       DcimRacksListFormFactorIcParameterInner = "4-post-cabinet"
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER__4_POST_FRAME         DcimRacksListFormFactorIcParameterInner = "4-post-frame"
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER_WALL_CABINET          DcimRacksListFormFactorIcParameterInner = "wall-cabinet"
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER_WALL_CABINET_VERTICAL DcimRacksListFormFactorIcParameterInner = "wall-cabinet-vertical"
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER_WALL_FRAME            DcimRacksListFormFactorIcParameterInner = "wall-frame"
	DCIMRACKSLISTFORMFACTORICPARAMETERINNER_WALL_FRAME_VERTICAL   DcimRacksListFormFactorIcParameterInner = "wall-frame-vertical"
)

// All allowed values of DcimRacksListFormFactorIcParameterInner enum
var AllowedDcimRacksListFormFactorIcParameterInnerEnumValues = []DcimRacksListFormFactorIcParameterInner{
	"",
	"2-post-frame",
	"4-post-cabinet",
	"4-post-frame",
	"wall-cabinet",
	"wall-cabinet-vertical",
	"wall-frame",
	"wall-frame-vertical",
}

func (v *DcimRacksListFormFactorIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRacksListFormFactorIcParameterInner(value)
	for _, existing := range AllowedDcimRacksListFormFactorIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRacksListFormFactorIcParameterInner", value)
}

// NewDcimRacksListFormFactorIcParameterInnerFromValue returns a pointer to a valid DcimRacksListFormFactorIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRacksListFormFactorIcParameterInnerFromValue(v string) (*DcimRacksListFormFactorIcParameterInner, error) {
	ev := DcimRacksListFormFactorIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRacksListFormFactorIcParameterInner: valid values are %v", v, AllowedDcimRacksListFormFactorIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRacksListFormFactorIcParameterInner) IsValid() bool {
	for _, existing := range AllowedDcimRacksListFormFactorIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_racks_list_form_factor__ic_parameter_inner value
func (v DcimRacksListFormFactorIcParameterInner) Ptr() *DcimRacksListFormFactorIcParameterInner {
	return &v
}

type NullableDcimRacksListFormFactorIcParameterInner struct {
	value *DcimRacksListFormFactorIcParameterInner
	isSet bool
}

func (v NullableDcimRacksListFormFactorIcParameterInner) Get() *DcimRacksListFormFactorIcParameterInner {
	return v.value
}

func (v *NullableDcimRacksListFormFactorIcParameterInner) Set(val *DcimRacksListFormFactorIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRacksListFormFactorIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRacksListFormFactorIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRacksListFormFactorIcParameterInner(val *DcimRacksListFormFactorIcParameterInner) *NullableDcimRacksListFormFactorIcParameterInner {
	return &NullableDcimRacksListFormFactorIcParameterInner{value: val, isSet: true}
}

func (v NullableDcimRacksListFormFactorIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRacksListFormFactorIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
