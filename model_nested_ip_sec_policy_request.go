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

// checks if the NestedIPSecPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedIPSecPolicyRequest{}

// NestedIPSecPolicyRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedIPSecPolicyRequest struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedIPSecPolicyRequest NestedIPSecPolicyRequest

// NewNestedIPSecPolicyRequest instantiates a new NestedIPSecPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedIPSecPolicyRequest(name string) *NestedIPSecPolicyRequest {
	this := NestedIPSecPolicyRequest{}
	this.Name = name
	return &this
}

// NewNestedIPSecPolicyRequestWithDefaults instantiates a new NestedIPSecPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedIPSecPolicyRequestWithDefaults() *NestedIPSecPolicyRequest {
	this := NestedIPSecPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedIPSecPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedIPSecPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedIPSecPolicyRequest) SetName(v string) {
	o.Name = v
}

func (o NestedIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedIPSecPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedIPSecPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNestedIPSecPolicyRequest := _NestedIPSecPolicyRequest{}

	err = json.Unmarshal(data, &varNestedIPSecPolicyRequest)

	if err != nil {
		return err
	}

	*o = NestedIPSecPolicyRequest(varNestedIPSecPolicyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedIPSecPolicyRequest struct {
	value *NestedIPSecPolicyRequest
	isSet bool
}

func (v NullableNestedIPSecPolicyRequest) Get() *NestedIPSecPolicyRequest {
	return v.value
}

func (v *NullableNestedIPSecPolicyRequest) Set(val *NestedIPSecPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedIPSecPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedIPSecPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedIPSecPolicyRequest(val *NestedIPSecPolicyRequest) *NullableNestedIPSecPolicyRequest {
	return &NullableNestedIPSecPolicyRequest{value: val, isSet: true}
}

func (v NullableNestedIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedIPSecPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
