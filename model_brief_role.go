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

// checks if the BriefRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefRole{}

// BriefRole Adds support for custom fields and tags.
type BriefRole struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Description          *string `json:"description,omitempty"`
	PrefixCount          int64   `json:"prefix_count"`
	VlanCount            int64   `json:"vlan_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefRole BriefRole

// NewBriefRole instantiates a new BriefRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefRole(id int32, url string, display string, name string, slug string, prefixCount int64, vlanCount int64) *BriefRole {
	this := BriefRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.PrefixCount = prefixCount
	this.VlanCount = vlanCount
	return &this
}

// NewBriefRoleWithDefaults instantiates a new BriefRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefRoleWithDefaults() *BriefRole {
	this := BriefRole{}
	return &this
}

// GetId returns the Id field value
func (o *BriefRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefRole) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefRole) SetDescription(v string) {
	o.Description = &v
}

// GetPrefixCount returns the PrefixCount field value
func (o *BriefRole) GetPrefixCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetPrefixCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *BriefRole) SetPrefixCount(v int64) {
	o.PrefixCount = v
}

// GetVlanCount returns the VlanCount field value
func (o *BriefRole) GetVlanCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value
// and a boolean to check if the value has been set.
func (o *BriefRole) GetVlanCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanCount, true
}

// SetVlanCount sets field value
func (o *BriefRole) SetVlanCount(v int64) {
	o.VlanCount = v
}

func (o BriefRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["prefix_count"] = o.PrefixCount
	toSerialize["vlan_count"] = o.VlanCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"prefix_count",
		"vlan_count",
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

	varBriefRole := _BriefRole{}

	err = json.Unmarshal(data, &varBriefRole)

	if err != nil {
		return err
	}

	*o = BriefRole(varBriefRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "prefix_count")
		delete(additionalProperties, "vlan_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefRole struct {
	value *BriefRole
	isSet bool
}

func (v NullableBriefRole) Get() *BriefRole {
	return v.value
}

func (v *NullableBriefRole) Set(val *BriefRole) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefRole) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefRole(val *BriefRole) *NullableBriefRole {
	return &NullableBriefRole{value: val, isSet: true}
}

func (v NullableBriefRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
