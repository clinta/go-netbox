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

// checks if the WritableServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableServiceRequest{}

// WritableServiceRequest Adds support for custom fields and tags.
type WritableServiceRequest struct {
	Device         NullableInt32                         `json:"device,omitempty"`
	VirtualMachine NullableInt32                         `json:"virtual_machine,omitempty"`
	Name           string                                `json:"name"`
	Ports          []int32                               `json:"ports"`
	Protocol       PatchedWritableServiceRequestProtocol `json:"protocol"`
	// The specific IP addresses (if any) to which this service is bound
	Ipaddresses          []int32                `json:"ipaddresses,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableServiceRequest WritableServiceRequest

// NewWritableServiceRequest instantiates a new WritableServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableServiceRequest(name string, ports []int32, protocol PatchedWritableServiceRequestProtocol) *WritableServiceRequest {
	this := WritableServiceRequest{}
	this.Name = name
	this.Ports = ports
	this.Protocol = protocol
	return &this
}

// NewWritableServiceRequestWithDefaults instantiates a new WritableServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableServiceRequestWithDefaults() *WritableServiceRequest {
	this := WritableServiceRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableServiceRequest) GetDevice() int32 {
	if o == nil || IsNil(o.Device.Get()) {
		var ret int32
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableServiceRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableInt32 and assigns it to the Device field.
func (o *WritableServiceRequest) SetDevice(v int32) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *WritableServiceRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *WritableServiceRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableServiceRequest) GetVirtualMachine() int32 {
	if o == nil || IsNil(o.VirtualMachine.Get()) {
		var ret int32
		return ret
	}
	return *o.VirtualMachine.Get()
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableServiceRequest) GetVirtualMachineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualMachine.Get(), o.VirtualMachine.IsSet()
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine.IsSet() {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given NullableInt32 and assigns it to the VirtualMachine field.
func (o *WritableServiceRequest) SetVirtualMachine(v int32) {
	o.VirtualMachine.Set(&v)
}

// SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil
func (o *WritableServiceRequest) SetVirtualMachineNil() {
	o.VirtualMachine.Set(nil)
}

// UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil
func (o *WritableServiceRequest) UnsetVirtualMachine() {
	o.VirtualMachine.Unset()
}

// GetName returns the Name field value
func (o *WritableServiceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableServiceRequest) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value
func (o *WritableServiceRequest) GetPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *WritableServiceRequest) SetPorts(v []int32) {
	o.Ports = v
}

// GetProtocol returns the Protocol field value
func (o *WritableServiceRequest) GetProtocol() PatchedWritableServiceRequestProtocol {
	if o == nil {
		var ret PatchedWritableServiceRequestProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetProtocolOk() (*PatchedWritableServiceRequestProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *WritableServiceRequest) SetProtocol(v PatchedWritableServiceRequestProtocol) {
	o.Protocol = v
}

// GetIpaddresses returns the Ipaddresses field value if set, zero value otherwise.
func (o *WritableServiceRequest) GetIpaddresses() []int32 {
	if o == nil || IsNil(o.Ipaddresses) {
		var ret []int32
		return ret
	}
	return o.Ipaddresses
}

// GetIpaddressesOk returns a tuple with the Ipaddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetIpaddressesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ipaddresses) {
		return nil, false
	}
	return o.Ipaddresses, true
}

// HasIpaddresses returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasIpaddresses() bool {
	if o != nil && !IsNil(o.Ipaddresses) {
		return true
	}

	return false
}

// SetIpaddresses gets a reference to the given []int32 and assigns it to the Ipaddresses field.
func (o *WritableServiceRequest) SetIpaddresses(v []int32) {
	o.Ipaddresses = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableServiceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableServiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableServiceRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableServiceRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableServiceRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableServiceRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableServiceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableServiceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableServiceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.VirtualMachine.IsSet() {
		toSerialize["virtual_machine"] = o.VirtualMachine.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["ports"] = o.Ports
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.Ipaddresses) {
		toSerialize["ipaddresses"] = o.Ipaddresses
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableServiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"ports",
		"protocol",
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

	varWritableServiceRequest := _WritableServiceRequest{}

	err = json.Unmarshal(data, &varWritableServiceRequest)

	if err != nil {
		return err
	}

	*o = WritableServiceRequest(varWritableServiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "ipaddresses")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableServiceRequest struct {
	value *WritableServiceRequest
	isSet bool
}

func (v NullableWritableServiceRequest) Get() *WritableServiceRequest {
	return v.value
}

func (v *NullableWritableServiceRequest) Set(val *WritableServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableServiceRequest(val *WritableServiceRequest) *NullableWritableServiceRequest {
	return &NullableWritableServiceRequest{value: val, isSet: true}
}

func (v NullableWritableServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
