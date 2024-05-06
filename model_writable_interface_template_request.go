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

// checks if the WritableInterfaceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableInterfaceTemplateRequest{}

// WritableInterfaceTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableInterfaceTemplateRequest struct {
	DeviceType NullableInt32 `json:"device_type,omitempty"`
	ModuleType NullableInt32 `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string                `json:"label,omitempty"`
	Type                 InterfaceTypeValue     `json:"type"`
	Enabled              *bool                  `json:"enabled,omitempty"`
	MgmtOnly             *bool                  `json:"mgmt_only,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Bridge               NullableInt32          `json:"bridge,omitempty"`
	PoeMode              *InterfacePoeModeValue `json:"poe_mode,omitempty"`
	PoeType              *InterfacePoeTypeValue `json:"poe_type,omitempty"`
	RfRole               *WirelessRole          `json:"rf_role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableInterfaceTemplateRequest WritableInterfaceTemplateRequest

// NewWritableInterfaceTemplateRequest instantiates a new WritableInterfaceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableInterfaceTemplateRequest(name string, type_ InterfaceTypeValue) *WritableInterfaceTemplateRequest {
	this := WritableInterfaceTemplateRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewWritableInterfaceTemplateRequestWithDefaults instantiates a new WritableInterfaceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableInterfaceTemplateRequestWithDefaults() *WritableInterfaceTemplateRequest {
	this := WritableInterfaceTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableInterfaceTemplateRequest) GetDeviceType() int32 {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret int32
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableInterfaceTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableInt32 and assigns it to the DeviceType field.
func (o *WritableInterfaceTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *WritableInterfaceTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *WritableInterfaceTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableInterfaceTemplateRequest) GetModuleType() int32 {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret int32
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableInterfaceTemplateRequest) GetModuleTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableInt32 and assigns it to the ModuleType field.
func (o *WritableInterfaceTemplateRequest) SetModuleType(v int32) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *WritableInterfaceTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *WritableInterfaceTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *WritableInterfaceTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableInterfaceTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableInterfaceTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *WritableInterfaceTemplateRequest) GetType() InterfaceTypeValue {
	if o == nil {
		var ret InterfaceTypeValue
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetTypeOk() (*InterfaceTypeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WritableInterfaceTemplateRequest) SetType(v InterfaceTypeValue) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WritableInterfaceTemplateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetMgmtOnly() bool {
	if o == nil || IsNil(o.MgmtOnly) {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MgmtOnly) {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasMgmtOnly() bool {
	if o != nil && !IsNil(o.MgmtOnly) {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *WritableInterfaceTemplateRequest) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableInterfaceTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableInterfaceTemplateRequest) GetBridge() int32 {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret int32
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableInterfaceTemplateRequest) GetBridgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableInt32 and assigns it to the Bridge field.
func (o *WritableInterfaceTemplateRequest) SetBridge(v int32) {
	o.Bridge.Set(&v)
}

// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *WritableInterfaceTemplateRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *WritableInterfaceTemplateRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetPoeMode returns the PoeMode field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetPoeMode() InterfacePoeModeValue {
	if o == nil || IsNil(o.PoeMode) {
		var ret InterfacePoeModeValue
		return ret
	}
	return *o.PoeMode
}

// GetPoeModeOk returns a tuple with the PoeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetPoeModeOk() (*InterfacePoeModeValue, bool) {
	if o == nil || IsNil(o.PoeMode) {
		return nil, false
	}
	return o.PoeMode, true
}

// HasPoeMode returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasPoeMode() bool {
	if o != nil && !IsNil(o.PoeMode) {
		return true
	}

	return false
}

// SetPoeMode gets a reference to the given InterfacePoeModeValue and assigns it to the PoeMode field.
func (o *WritableInterfaceTemplateRequest) SetPoeMode(v InterfacePoeModeValue) {
	o.PoeMode = &v
}

// GetPoeType returns the PoeType field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetPoeType() InterfacePoeTypeValue {
	if o == nil || IsNil(o.PoeType) {
		var ret InterfacePoeTypeValue
		return ret
	}
	return *o.PoeType
}

// GetPoeTypeOk returns a tuple with the PoeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetPoeTypeOk() (*InterfacePoeTypeValue, bool) {
	if o == nil || IsNil(o.PoeType) {
		return nil, false
	}
	return o.PoeType, true
}

// HasPoeType returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasPoeType() bool {
	if o != nil && !IsNil(o.PoeType) {
		return true
	}

	return false
}

// SetPoeType gets a reference to the given InterfacePoeTypeValue and assigns it to the PoeType field.
func (o *WritableInterfaceTemplateRequest) SetPoeType(v InterfacePoeTypeValue) {
	o.PoeType = &v
}

// GetRfRole returns the RfRole field value if set, zero value otherwise.
func (o *WritableInterfaceTemplateRequest) GetRfRole() WirelessRole {
	if o == nil || IsNil(o.RfRole) {
		var ret WirelessRole
		return ret
	}
	return *o.RfRole
}

// GetRfRoleOk returns a tuple with the RfRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableInterfaceTemplateRequest) GetRfRoleOk() (*WirelessRole, bool) {
	if o == nil || IsNil(o.RfRole) {
		return nil, false
	}
	return o.RfRole, true
}

// HasRfRole returns a boolean if a field has been set.
func (o *WritableInterfaceTemplateRequest) HasRfRole() bool {
	if o != nil && !IsNil(o.RfRole) {
		return true
	}

	return false
}

// SetRfRole gets a reference to the given WirelessRole and assigns it to the RfRole field.
func (o *WritableInterfaceTemplateRequest) SetRfRole(v WirelessRole) {
	o.RfRole = &v
}

func (o WritableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableInterfaceTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MgmtOnly) {
		toSerialize["mgmt_only"] = o.MgmtOnly
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if !IsNil(o.PoeMode) {
		toSerialize["poe_mode"] = o.PoeMode
	}
	if !IsNil(o.PoeType) {
		toSerialize["poe_type"] = o.PoeType
	}
	if !IsNil(o.RfRole) {
		toSerialize["rf_role"] = o.RfRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableInterfaceTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varWritableInterfaceTemplateRequest := _WritableInterfaceTemplateRequest{}

	err = json.Unmarshal(data, &varWritableInterfaceTemplateRequest)

	if err != nil {
		return err
	}

	*o = WritableInterfaceTemplateRequest(varWritableInterfaceTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "mgmt_only")
		delete(additionalProperties, "description")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "poe_mode")
		delete(additionalProperties, "poe_type")
		delete(additionalProperties, "rf_role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableInterfaceTemplateRequest struct {
	value *WritableInterfaceTemplateRequest
	isSet bool
}

func (v NullableWritableInterfaceTemplateRequest) Get() *WritableInterfaceTemplateRequest {
	return v.value
}

func (v *NullableWritableInterfaceTemplateRequest) Set(val *WritableInterfaceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableInterfaceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableInterfaceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableInterfaceTemplateRequest(val *WritableInterfaceTemplateRequest) *NullableWritableInterfaceTemplateRequest {
	return &NullableWritableInterfaceTemplateRequest{value: val, isSet: true}
}

func (v NullableWritableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableInterfaceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
