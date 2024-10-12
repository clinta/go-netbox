/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the DeviceAirflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAirflow{}

// DeviceAirflow struct for DeviceAirflow
type DeviceAirflow struct {
	Value                *DeviceAirflowValue `json:"value,omitempty"`
	Label                *DeviceAirflowLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAirflow DeviceAirflow

// NewDeviceAirflow instantiates a new DeviceAirflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAirflow() *DeviceAirflow {
	this := DeviceAirflow{}
	return &this
}

// NewDeviceAirflowWithDefaults instantiates a new DeviceAirflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAirflowWithDefaults() *DeviceAirflow {
	this := DeviceAirflow{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceAirflow) GetValue() DeviceAirflowValue {
	if o == nil || IsNil(o.Value) {
		var ret DeviceAirflowValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAirflow) GetValueOk() (*DeviceAirflowValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceAirflow) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given DeviceAirflowValue and assigns it to the Value field.
func (o *DeviceAirflow) SetValue(v DeviceAirflowValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceAirflow) GetLabel() DeviceAirflowLabel {
	if o == nil || IsNil(o.Label) {
		var ret DeviceAirflowLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAirflow) GetLabelOk() (*DeviceAirflowLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceAirflow) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given DeviceAirflowLabel and assigns it to the Label field.
func (o *DeviceAirflow) SetLabel(v DeviceAirflowLabel) {
	o.Label = &v
}

func (o DeviceAirflow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAirflow) ToMap() (map[string]interface{}, error) {
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

func (o *DeviceAirflow) UnmarshalJSON(data []byte) (err error) {
	varDeviceAirflow := _DeviceAirflow{}

	err = json.Unmarshal(data, &varDeviceAirflow)

	if err != nil {
		return err
	}

	*o = DeviceAirflow(varDeviceAirflow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAirflow struct {
	value *DeviceAirflow
	isSet bool
}

func (v NullableDeviceAirflow) Get() *DeviceAirflow {
	return v.value
}

func (v *NullableDeviceAirflow) Set(val *DeviceAirflow) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAirflow) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAirflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAirflow(val *DeviceAirflow) *NullableDeviceAirflow {
	return &NullableDeviceAirflow{value: val, isSet: true}
}

func (v NullableDeviceAirflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAirflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
