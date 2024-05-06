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

// checks if the NestedPowerPortTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedPowerPortTemplate{}

// NestedPowerPortTemplate Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedPowerPortTemplate struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedPowerPortTemplate NestedPowerPortTemplate

// NewNestedPowerPortTemplate instantiates a new NestedPowerPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedPowerPortTemplate(id int32, url string, display string, name string) *NestedPowerPortTemplate {
	this := NestedPowerPortTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedPowerPortTemplateWithDefaults instantiates a new NestedPowerPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedPowerPortTemplateWithDefaults() *NestedPowerPortTemplate {
	this := NestedPowerPortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *NestedPowerPortTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPortTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedPowerPortTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedPowerPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedPowerPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedPowerPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedPowerPortTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedPowerPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedPowerPortTemplate) SetName(v string) {
	o.Name = v
}

func (o NestedPowerPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedPowerPortTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedPowerPortTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
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

	varNestedPowerPortTemplate := _NestedPowerPortTemplate{}

	err = json.Unmarshal(data, &varNestedPowerPortTemplate)

	if err != nil {
		return err
	}

	*o = NestedPowerPortTemplate(varNestedPowerPortTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedPowerPortTemplate struct {
	value *NestedPowerPortTemplate
	isSet bool
}

func (v NullableNestedPowerPortTemplate) Get() *NestedPowerPortTemplate {
	return v.value
}

func (v *NullableNestedPowerPortTemplate) Set(val *NestedPowerPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedPowerPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedPowerPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedPowerPortTemplate(val *NestedPowerPortTemplate) *NullableNestedPowerPortTemplate {
	return &NullableNestedPowerPortTemplate{value: val, isSet: true}
}

func (v NullableNestedPowerPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedPowerPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
