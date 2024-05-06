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
	"time"
)

// checks if the TunnelTermination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelTermination{}

// TunnelTermination Adds support for custom fields and tags.
type TunnelTermination struct {
	Id                   int32                   `json:"id"`
	Url                  string                  `json:"url"`
	Display              string                  `json:"display"`
	Tunnel               NestedTunnel            `json:"tunnel"`
	Role                 TunnelTerminationRole   `json:"role"`
	TerminationType      string                  `json:"termination_type"`
	TerminationId        NullableInt64           `json:"termination_id,omitempty"`
	Termination          interface{}             `json:"termination"`
	OutsideIp            NullableNestedIPAddress `json:"outside_ip,omitempty"`
	Tags                 []NestedTag             `json:"tags,omitempty"`
	CustomFields         map[string]interface{}  `json:"custom_fields,omitempty"`
	Created              NullableTime            `json:"created"`
	LastUpdated          NullableTime            `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _TunnelTermination TunnelTermination

// NewTunnelTermination instantiates a new TunnelTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelTermination(id int32, url string, display string, tunnel NestedTunnel, role TunnelTerminationRole, terminationType string, termination interface{}, created NullableTime, lastUpdated NullableTime) *TunnelTermination {
	this := TunnelTermination{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Tunnel = tunnel
	this.Role = role
	this.TerminationType = terminationType
	this.Termination = termination
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewTunnelTerminationWithDefaults instantiates a new TunnelTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelTerminationWithDefaults() *TunnelTermination {
	this := TunnelTermination{}
	return &this
}

// GetId returns the Id field value
func (o *TunnelTermination) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TunnelTermination) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *TunnelTermination) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TunnelTermination) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *TunnelTermination) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *TunnelTermination) SetDisplay(v string) {
	o.Display = v
}

// GetTunnel returns the Tunnel field value
func (o *TunnelTermination) GetTunnel() NestedTunnel {
	if o == nil {
		var ret NestedTunnel
		return ret
	}

	return o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetTunnelOk() (*NestedTunnel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tunnel, true
}

// SetTunnel sets field value
func (o *TunnelTermination) SetTunnel(v NestedTunnel) {
	o.Tunnel = v
}

// GetRole returns the Role field value
func (o *TunnelTermination) GetRole() TunnelTerminationRole {
	if o == nil {
		var ret TunnelTerminationRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetRoleOk() (*TunnelTerminationRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *TunnelTermination) SetRole(v TunnelTerminationRole) {
	o.Role = v
}

// GetTerminationType returns the TerminationType field value
func (o *TunnelTermination) GetTerminationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationType
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetTerminationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationType, true
}

// SetTerminationType sets field value
func (o *TunnelTermination) SetTerminationType(v string) {
	o.TerminationType = v
}

// GetTerminationId returns the TerminationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelTermination) GetTerminationId() int64 {
	if o == nil || IsNil(o.TerminationId.Get()) {
		var ret int64
		return ret
	}
	return *o.TerminationId.Get()
}

// GetTerminationIdOk returns a tuple with the TerminationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTermination) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationId.Get(), o.TerminationId.IsSet()
}

// HasTerminationId returns a boolean if a field has been set.
func (o *TunnelTermination) HasTerminationId() bool {
	if o != nil && o.TerminationId.IsSet() {
		return true
	}

	return false
}

// SetTerminationId gets a reference to the given NullableInt64 and assigns it to the TerminationId field.
func (o *TunnelTermination) SetTerminationId(v int64) {
	o.TerminationId.Set(&v)
}

// SetTerminationIdNil sets the value for TerminationId to be an explicit nil
func (o *TunnelTermination) SetTerminationIdNil() {
	o.TerminationId.Set(nil)
}

// UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
func (o *TunnelTermination) UnsetTerminationId() {
	o.TerminationId.Unset()
}

// GetTermination returns the Termination field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TunnelTermination) GetTermination() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Termination
}

// GetTerminationOk returns a tuple with the Termination field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTermination) GetTerminationOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Termination) {
		return nil, false
	}
	return &o.Termination, true
}

// SetTermination sets field value
func (o *TunnelTermination) SetTermination(v interface{}) {
	o.Termination = v
}

// GetOutsideIp returns the OutsideIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelTermination) GetOutsideIp() NestedIPAddress {
	if o == nil || IsNil(o.OutsideIp.Get()) {
		var ret NestedIPAddress
		return ret
	}
	return *o.OutsideIp.Get()
}

// GetOutsideIpOk returns a tuple with the OutsideIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTermination) GetOutsideIpOk() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutsideIp.Get(), o.OutsideIp.IsSet()
}

// HasOutsideIp returns a boolean if a field has been set.
func (o *TunnelTermination) HasOutsideIp() bool {
	if o != nil && o.OutsideIp.IsSet() {
		return true
	}

	return false
}

// SetOutsideIp gets a reference to the given NullableNestedIPAddress and assigns it to the OutsideIp field.
func (o *TunnelTermination) SetOutsideIp(v NestedIPAddress) {
	o.OutsideIp.Set(&v)
}

// SetOutsideIpNil sets the value for OutsideIp to be an explicit nil
func (o *TunnelTermination) SetOutsideIpNil() {
	o.OutsideIp.Set(nil)
}

// UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
func (o *TunnelTermination) UnsetOutsideIp() {
	o.OutsideIp.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TunnelTermination) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TunnelTermination) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *TunnelTermination) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TunnelTermination) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelTermination) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TunnelTermination) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *TunnelTermination) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TunnelTermination) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTermination) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *TunnelTermination) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TunnelTermination) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTermination) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *TunnelTermination) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o TunnelTermination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelTermination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["tunnel"] = o.Tunnel
	toSerialize["role"] = o.Role
	toSerialize["termination_type"] = o.TerminationType
	if o.TerminationId.IsSet() {
		toSerialize["termination_id"] = o.TerminationId.Get()
	}
	if o.Termination != nil {
		toSerialize["termination"] = o.Termination
	}
	if o.OutsideIp.IsSet() {
		toSerialize["outside_ip"] = o.OutsideIp.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelTermination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"tunnel",
		"role",
		"termination_type",
		"termination",
		"created",
		"last_updated",
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

	varTunnelTermination := _TunnelTermination{}

	err = json.Unmarshal(data, &varTunnelTermination)

	if err != nil {
		return err
	}

	*o = TunnelTermination(varTunnelTermination)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "tunnel")
		delete(additionalProperties, "role")
		delete(additionalProperties, "termination_type")
		delete(additionalProperties, "termination_id")
		delete(additionalProperties, "termination")
		delete(additionalProperties, "outside_ip")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelTermination struct {
	value *TunnelTermination
	isSet bool
}

func (v NullableTunnelTermination) Get() *TunnelTermination {
	return v.value
}

func (v *NullableTunnelTermination) Set(val *TunnelTermination) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelTermination(val *TunnelTermination) *NullableTunnelTermination {
	return &NullableTunnelTermination{value: val, isSet: true}
}

func (v NullableTunnelTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
