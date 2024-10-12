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

// checks if the NestedVirtualMachine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVirtualMachine{}

// NestedVirtualMachine Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVirtualMachine struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	DisplayUrl           string `json:"display_url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedVirtualMachine NestedVirtualMachine

// NewNestedVirtualMachine instantiates a new NestedVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVirtualMachine(id int32, url string, displayUrl string, display string, name string) *NestedVirtualMachine {
	this := NestedVirtualMachine{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedVirtualMachineWithDefaults instantiates a new NestedVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVirtualMachineWithDefaults() *NestedVirtualMachine {
	this := NestedVirtualMachine{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVirtualMachine) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVirtualMachine) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedVirtualMachine) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVirtualMachine) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *NestedVirtualMachine) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *NestedVirtualMachine) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *NestedVirtualMachine) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedVirtualMachine) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedVirtualMachine) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVirtualMachine) SetName(v string) {
	o.Name = v
}

func (o NestedVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVirtualMachine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVirtualMachine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
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

	varNestedVirtualMachine := _NestedVirtualMachine{}

	err = json.Unmarshal(data, &varNestedVirtualMachine)

	if err != nil {
		return err
	}

	*o = NestedVirtualMachine(varNestedVirtualMachine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVirtualMachine struct {
	value *NestedVirtualMachine
	isSet bool
}

func (v NullableNestedVirtualMachine) Get() *NestedVirtualMachine {
	return v.value
}

func (v *NullableNestedVirtualMachine) Set(val *NestedVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVirtualMachine(val *NestedVirtualMachine) *NullableNestedVirtualMachine {
	return &NullableNestedVirtualMachine{value: val, isSet: true}
}

func (v NullableNestedVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
