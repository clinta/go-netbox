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

// DataSourceTypeValue * `None` - --------- * `local` - Local * `git` - Git * `amazon-s3` - Amazon S3
type DataSourceTypeValue string

// List of DataSource_type_value
const (
	DATASOURCETYPEVALUE_LOCAL     DataSourceTypeValue = "local"
	DATASOURCETYPEVALUE_GIT       DataSourceTypeValue = "git"
	DATASOURCETYPEVALUE_AMAZON_S3 DataSourceTypeValue = "amazon-s3"
)

// All allowed values of DataSourceTypeValue enum
var AllowedDataSourceTypeValueEnumValues = []DataSourceTypeValue{
	"local",
	"git",
	"amazon-s3",
}

func (v *DataSourceTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceTypeValue(value)
	for _, existing := range AllowedDataSourceTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceTypeValue", value)
}

// NewDataSourceTypeValueFromValue returns a pointer to a valid DataSourceTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceTypeValueFromValue(v string) (*DataSourceTypeValue, error) {
	ev := DataSourceTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceTypeValue: valid values are %v", v, AllowedDataSourceTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceTypeValue) IsValid() bool {
	for _, existing := range AllowedDataSourceTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSource_type_value value
func (v DataSourceTypeValue) Ptr() *DataSourceTypeValue {
	return &v
}

type NullableDataSourceTypeValue struct {
	value *DataSourceTypeValue
	isSet bool
}

func (v NullableDataSourceTypeValue) Get() *DataSourceTypeValue {
	return v.value
}

func (v *NullableDataSourceTypeValue) Set(val *DataSourceTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceTypeValue(val *DataSourceTypeValue) *NullableDataSourceTypeValue {
	return &NullableDataSourceTypeValue{value: val, isSet: true}
}

func (v NullableDataSourceTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
