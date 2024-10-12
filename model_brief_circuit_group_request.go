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

// checks if the BriefCircuitGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCircuitGroupRequest{}

// BriefCircuitGroupRequest Adds support for custom fields and tags.
type BriefCircuitGroupRequest struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _BriefCircuitGroupRequest BriefCircuitGroupRequest

// NewBriefCircuitGroupRequest instantiates a new BriefCircuitGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCircuitGroupRequest(name string) *BriefCircuitGroupRequest {
	this := BriefCircuitGroupRequest{}
	this.Name = name
	return &this
}

// NewBriefCircuitGroupRequestWithDefaults instantiates a new BriefCircuitGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCircuitGroupRequestWithDefaults() *BriefCircuitGroupRequest {
	this := BriefCircuitGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefCircuitGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefCircuitGroupRequest) SetName(v string) {
	o.Name = v
}

func (o BriefCircuitGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCircuitGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCircuitGroupRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBriefCircuitGroupRequest := _BriefCircuitGroupRequest{}

	err = json.Unmarshal(data, &varBriefCircuitGroupRequest)

	if err != nil {
		return err
	}

	*o = BriefCircuitGroupRequest(varBriefCircuitGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCircuitGroupRequest struct {
	value *BriefCircuitGroupRequest
	isSet bool
}

func (v NullableBriefCircuitGroupRequest) Get() *BriefCircuitGroupRequest {
	return v.value
}

func (v *NullableBriefCircuitGroupRequest) Set(val *BriefCircuitGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCircuitGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCircuitGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCircuitGroupRequest(val *BriefCircuitGroupRequest) *NullableBriefCircuitGroupRequest {
	return &NullableBriefCircuitGroupRequest{value: val, isSet: true}
}

func (v NullableBriefCircuitGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCircuitGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
