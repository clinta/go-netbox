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

// checks if the ConsolePortSpeed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsolePortSpeed{}

// ConsolePortSpeed struct for ConsolePortSpeed
type ConsolePortSpeed struct {
	Value                *ConsolePortSpeedValue `json:"value,omitempty"`
	Label                *ConsolePortSpeedLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsolePortSpeed ConsolePortSpeed

// NewConsolePortSpeed instantiates a new ConsolePortSpeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsolePortSpeed() *ConsolePortSpeed {
	this := ConsolePortSpeed{}
	return &this
}

// NewConsolePortSpeedWithDefaults instantiates a new ConsolePortSpeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsolePortSpeedWithDefaults() *ConsolePortSpeed {
	this := ConsolePortSpeed{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConsolePortSpeed) GetValue() ConsolePortSpeedValue {
	if o == nil || IsNil(o.Value) {
		var ret ConsolePortSpeedValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortSpeed) GetValueOk() (*ConsolePortSpeedValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConsolePortSpeed) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ConsolePortSpeedValue and assigns it to the Value field.
func (o *ConsolePortSpeed) SetValue(v ConsolePortSpeedValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsolePortSpeed) GetLabel() ConsolePortSpeedLabel {
	if o == nil || IsNil(o.Label) {
		var ret ConsolePortSpeedLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortSpeed) GetLabelOk() (*ConsolePortSpeedLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsolePortSpeed) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ConsolePortSpeedLabel and assigns it to the Label field.
func (o *ConsolePortSpeed) SetLabel(v ConsolePortSpeedLabel) {
	o.Label = &v
}

func (o ConsolePortSpeed) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsolePortSpeed) ToMap() (map[string]interface{}, error) {
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

func (o *ConsolePortSpeed) UnmarshalJSON(data []byte) (err error) {
	varConsolePortSpeed := _ConsolePortSpeed{}

	err = json.Unmarshal(data, &varConsolePortSpeed)

	if err != nil {
		return err
	}

	*o = ConsolePortSpeed(varConsolePortSpeed)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsolePortSpeed struct {
	value *ConsolePortSpeed
	isSet bool
}

func (v NullableConsolePortSpeed) Get() *ConsolePortSpeed {
	return v.value
}

func (v *NullableConsolePortSpeed) Set(val *ConsolePortSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortSpeed(val *ConsolePortSpeed) *NullableConsolePortSpeed {
	return &NullableConsolePortSpeed{value: val, isSet: true}
}

func (v NullableConsolePortSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
