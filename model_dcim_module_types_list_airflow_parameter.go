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

// DcimModuleTypesListAirflowParameter the model 'DcimModuleTypesListAirflowParameter'
type DcimModuleTypesListAirflowParameter string

// List of dcim_module_types_list_airflow_parameter
const (
	DCIMMODULETYPESLISTAIRFLOWPARAMETER_FRONT_TO_REAR DcimModuleTypesListAirflowParameter = "front-to-rear"
	DCIMMODULETYPESLISTAIRFLOWPARAMETER_LEFT_TO_RIGHT DcimModuleTypesListAirflowParameter = "left-to-right"
	DCIMMODULETYPESLISTAIRFLOWPARAMETER_PASSIVE       DcimModuleTypesListAirflowParameter = "passive"
	DCIMMODULETYPESLISTAIRFLOWPARAMETER_REAR_TO_FRONT DcimModuleTypesListAirflowParameter = "rear-to-front"
	DCIMMODULETYPESLISTAIRFLOWPARAMETER_RIGHT_TO_LEFT DcimModuleTypesListAirflowParameter = "right-to-left"
	DCIMMODULETYPESLISTAIRFLOWPARAMETER_SIDE_TO_REAR  DcimModuleTypesListAirflowParameter = "side-to-rear"
)

// All allowed values of DcimModuleTypesListAirflowParameter enum
var AllowedDcimModuleTypesListAirflowParameterEnumValues = []DcimModuleTypesListAirflowParameter{
	"front-to-rear",
	"left-to-right",
	"passive",
	"rear-to-front",
	"right-to-left",
	"side-to-rear",
}

func (v *DcimModuleTypesListAirflowParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimModuleTypesListAirflowParameter(value)
	for _, existing := range AllowedDcimModuleTypesListAirflowParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimModuleTypesListAirflowParameter", value)
}

// NewDcimModuleTypesListAirflowParameterFromValue returns a pointer to a valid DcimModuleTypesListAirflowParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimModuleTypesListAirflowParameterFromValue(v string) (*DcimModuleTypesListAirflowParameter, error) {
	ev := DcimModuleTypesListAirflowParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimModuleTypesListAirflowParameter: valid values are %v", v, AllowedDcimModuleTypesListAirflowParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimModuleTypesListAirflowParameter) IsValid() bool {
	for _, existing := range AllowedDcimModuleTypesListAirflowParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_module_types_list_airflow_parameter value
func (v DcimModuleTypesListAirflowParameter) Ptr() *DcimModuleTypesListAirflowParameter {
	return &v
}

type NullableDcimModuleTypesListAirflowParameter struct {
	value *DcimModuleTypesListAirflowParameter
	isSet bool
}

func (v NullableDcimModuleTypesListAirflowParameter) Get() *DcimModuleTypesListAirflowParameter {
	return v.value
}

func (v *NullableDcimModuleTypesListAirflowParameter) Set(val *DcimModuleTypesListAirflowParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimModuleTypesListAirflowParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimModuleTypesListAirflowParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimModuleTypesListAirflowParameter(val *DcimModuleTypesListAirflowParameter) *NullableDcimModuleTypesListAirflowParameter {
	return &NullableDcimModuleTypesListAirflowParameter{value: val, isSet: true}
}

func (v NullableDcimModuleTypesListAirflowParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimModuleTypesListAirflowParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
