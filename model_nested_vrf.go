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

// checks if the NestedVRF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVRF{}

// NestedVRF Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVRF struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Name    string `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd                   NullableString `json:"rd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedVRF NestedVRF

// NewNestedVRF instantiates a new NestedVRF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVRF(id int32, url string, display string, name string) *NestedVRF {
	this := NestedVRF{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedVRFWithDefaults instantiates a new NestedVRF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVRFWithDefaults() *NestedVRF {
	this := NestedVRF{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVRF) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVRF) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVRF) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedVRF) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVRF) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVRF) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedVRF) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedVRF) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedVRF) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedVRF) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVRF) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVRF) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedVRF) GetRd() string {
	if o == nil || IsNil(o.Rd.Get()) {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedVRF) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *NestedVRF) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *NestedVRF) SetRd(v string) {
	o.Rd.Set(&v)
}

// SetRdNil sets the value for Rd to be an explicit nil
func (o *NestedVRF) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *NestedVRF) UnsetRd() {
	o.Rd.Unset()
}

func (o NestedVRF) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVRF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if o.Rd.IsSet() {
		toSerialize["rd"] = o.Rd.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVRF) UnmarshalJSON(data []byte) (err error) {
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

	varNestedVRF := _NestedVRF{}

	err = json.Unmarshal(data, &varNestedVRF)

	if err != nil {
		return err
	}

	*o = NestedVRF(varNestedVRF)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVRF struct {
	value *NestedVRF
	isSet bool
}

func (v NullableNestedVRF) Get() *NestedVRF {
	return v.value
}

func (v *NullableNestedVRF) Set(val *NestedVRF) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVRF) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVRF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVRF(val *NestedVRF) *NullableNestedVRF {
	return &NullableNestedVRF{value: val, isSet: true}
}

func (v NullableNestedVRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVRF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
