/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.7 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the InterfaceRfChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceRfChannel{}

// InterfaceRfChannel struct for InterfaceRfChannel
type InterfaceRfChannel struct {
	Value                *InterfaceRfChannelValue `json:"value,omitempty"`
	Label                *InterfaceRfChannelLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceRfChannel InterfaceRfChannel

// NewInterfaceRfChannel instantiates a new InterfaceRfChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceRfChannel() *InterfaceRfChannel {
	this := InterfaceRfChannel{}
	return &this
}

// NewInterfaceRfChannelWithDefaults instantiates a new InterfaceRfChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceRfChannelWithDefaults() *InterfaceRfChannel {
	this := InterfaceRfChannel{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceRfChannel) GetValue() InterfaceRfChannelValue {
	if o == nil || IsNil(o.Value) {
		var ret InterfaceRfChannelValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceRfChannel) GetValueOk() (*InterfaceRfChannelValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceRfChannel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given InterfaceRfChannelValue and assigns it to the Value field.
func (o *InterfaceRfChannel) SetValue(v InterfaceRfChannelValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceRfChannel) GetLabel() InterfaceRfChannelLabel {
	if o == nil || IsNil(o.Label) {
		var ret InterfaceRfChannelLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceRfChannel) GetLabelOk() (*InterfaceRfChannelLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceRfChannel) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given InterfaceRfChannelLabel and assigns it to the Label field.
func (o *InterfaceRfChannel) SetLabel(v InterfaceRfChannelLabel) {
	o.Label = &v
}

func (o InterfaceRfChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceRfChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceRfChannel) UnmarshalJSON(data []byte) (err error) {
	varInterfaceRfChannel := _InterfaceRfChannel{}

	err = json.Unmarshal(data, &varInterfaceRfChannel)

	if err != nil {
		return err
	}

	*o = InterfaceRfChannel(varInterfaceRfChannel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceRfChannel struct {
	value *InterfaceRfChannel
	isSet bool
}

func (v NullableInterfaceRfChannel) Get() *InterfaceRfChannel {
	return v.value
}

func (v *NullableInterfaceRfChannel) Set(val *InterfaceRfChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRfChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRfChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRfChannel(val *InterfaceRfChannel) *NullableInterfaceRfChannel {
	return &NullableInterfaceRfChannel{value: val, isSet: true}
}

func (v NullableInterfaceRfChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRfChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
