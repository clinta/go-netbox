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

// checks if the NestedVirtualChassisRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVirtualChassisRequest{}

// NestedVirtualChassisRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVirtualChassisRequest struct {
	Name                 string              `json:"name"`
	Master               NestedDeviceRequest `json:"master"`
	AdditionalProperties map[string]interface{}
}

type _NestedVirtualChassisRequest NestedVirtualChassisRequest

// NewNestedVirtualChassisRequest instantiates a new NestedVirtualChassisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVirtualChassisRequest(name string, master NestedDeviceRequest) *NestedVirtualChassisRequest {
	this := NestedVirtualChassisRequest{}
	this.Name = name
	this.Master = master
	return &this
}

// NewNestedVirtualChassisRequestWithDefaults instantiates a new NestedVirtualChassisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVirtualChassisRequestWithDefaults() *NestedVirtualChassisRequest {
	this := NestedVirtualChassisRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedVirtualChassisRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassisRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVirtualChassisRequest) SetName(v string) {
	o.Name = v
}

// GetMaster returns the Master field value
func (o *NestedVirtualChassisRequest) GetMaster() NestedDeviceRequest {
	if o == nil {
		var ret NestedDeviceRequest
		return ret
	}

	return o.Master
}

// GetMasterOk returns a tuple with the Master field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassisRequest) GetMasterOk() (*NestedDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Master, true
}

// SetMaster sets field value
func (o *NestedVirtualChassisRequest) SetMaster(v NestedDeviceRequest) {
	o.Master = v
}

func (o NestedVirtualChassisRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVirtualChassisRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["master"] = o.Master

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVirtualChassisRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"master",
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

	varNestedVirtualChassisRequest := _NestedVirtualChassisRequest{}

	err = json.Unmarshal(data, &varNestedVirtualChassisRequest)

	if err != nil {
		return err
	}

	*o = NestedVirtualChassisRequest(varNestedVirtualChassisRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "master")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVirtualChassisRequest struct {
	value *NestedVirtualChassisRequest
	isSet bool
}

func (v NullableNestedVirtualChassisRequest) Get() *NestedVirtualChassisRequest {
	return v.value
}

func (v *NullableNestedVirtualChassisRequest) Set(val *NestedVirtualChassisRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVirtualChassisRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVirtualChassisRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVirtualChassisRequest(val *NestedVirtualChassisRequest) *NullableNestedVirtualChassisRequest {
	return &NullableNestedVirtualChassisRequest{value: val, isSet: true}
}

func (v NullableNestedVirtualChassisRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVirtualChassisRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
