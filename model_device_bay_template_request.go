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

// checks if the DeviceBayTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceBayTemplateRequest{}

// DeviceBayTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type DeviceBayTemplateRequest struct {
	DeviceType NestedDeviceTypeRequest `json:"device_type"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string `json:"label,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceBayTemplateRequest DeviceBayTemplateRequest

// NewDeviceBayTemplateRequest instantiates a new DeviceBayTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceBayTemplateRequest(deviceType NestedDeviceTypeRequest, name string) *DeviceBayTemplateRequest {
	this := DeviceBayTemplateRequest{}
	this.DeviceType = deviceType
	this.Name = name
	return &this
}

// NewDeviceBayTemplateRequestWithDefaults instantiates a new DeviceBayTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceBayTemplateRequestWithDefaults() *DeviceBayTemplateRequest {
	this := DeviceBayTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value
func (o *DeviceBayTemplateRequest) GetDeviceType() NestedDeviceTypeRequest {
	if o == nil {
		var ret NestedDeviceTypeRequest
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *DeviceBayTemplateRequest) GetDeviceTypeOk() (*NestedDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *DeviceBayTemplateRequest) SetDeviceType(v NestedDeviceTypeRequest) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *DeviceBayTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceBayTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceBayTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceBayTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBayTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceBayTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DeviceBayTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceBayTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBayTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceBayTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceBayTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o DeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceBayTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_type"] = o.DeviceType
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceBayTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device_type",
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

	varDeviceBayTemplateRequest := _DeviceBayTemplateRequest{}

	err = json.Unmarshal(data, &varDeviceBayTemplateRequest)

	if err != nil {
		return err
	}

	*o = DeviceBayTemplateRequest(varDeviceBayTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceBayTemplateRequest struct {
	value *DeviceBayTemplateRequest
	isSet bool
}

func (v NullableDeviceBayTemplateRequest) Get() *DeviceBayTemplateRequest {
	return v.value
}

func (v *NullableDeviceBayTemplateRequest) Set(val *DeviceBayTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceBayTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceBayTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceBayTemplateRequest(val *DeviceBayTemplateRequest) *NullableDeviceBayTemplateRequest {
	return &NullableDeviceBayTemplateRequest{value: val, isSet: true}
}

func (v NullableDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceBayTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
