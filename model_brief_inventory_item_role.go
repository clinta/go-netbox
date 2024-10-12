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

// checks if the BriefInventoryItemRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefInventoryItemRole{}

// BriefInventoryItemRole Adds support for custom fields and tags.
type BriefInventoryItemRole struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Description          *string `json:"description,omitempty"`
	InventoryitemCount   int64   `json:"inventoryitem_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefInventoryItemRole BriefInventoryItemRole

// NewBriefInventoryItemRole instantiates a new BriefInventoryItemRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefInventoryItemRole(id int32, url string, display string, name string, slug string, inventoryitemCount int64) *BriefInventoryItemRole {
	this := BriefInventoryItemRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.InventoryitemCount = inventoryitemCount
	return &this
}

// NewBriefInventoryItemRoleWithDefaults instantiates a new BriefInventoryItemRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefInventoryItemRoleWithDefaults() *BriefInventoryItemRole {
	this := BriefInventoryItemRole{}
	return &this
}

// GetId returns the Id field value
func (o *BriefInventoryItemRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefInventoryItemRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefInventoryItemRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefInventoryItemRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefInventoryItemRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefInventoryItemRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefInventoryItemRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefInventoryItemRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefInventoryItemRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefInventoryItemRole) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefInventoryItemRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefInventoryItemRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefInventoryItemRole) SetDescription(v string) {
	o.Description = &v
}

// GetInventoryitemCount returns the InventoryitemCount field value
func (o *BriefInventoryItemRole) GetInventoryitemCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InventoryitemCount
}

// GetInventoryitemCountOk returns a tuple with the InventoryitemCount field value
// and a boolean to check if the value has been set.
func (o *BriefInventoryItemRole) GetInventoryitemCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryitemCount, true
}

// SetInventoryitemCount sets field value
func (o *BriefInventoryItemRole) SetInventoryitemCount(v int64) {
	o.InventoryitemCount = v
}

func (o BriefInventoryItemRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefInventoryItemRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["inventoryitem_count"] = o.InventoryitemCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefInventoryItemRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"inventoryitem_count",
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

	varBriefInventoryItemRole := _BriefInventoryItemRole{}

	err = json.Unmarshal(data, &varBriefInventoryItemRole)

	if err != nil {
		return err
	}

	*o = BriefInventoryItemRole(varBriefInventoryItemRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "inventoryitem_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefInventoryItemRole struct {
	value *BriefInventoryItemRole
	isSet bool
}

func (v NullableBriefInventoryItemRole) Get() *BriefInventoryItemRole {
	return v.value
}

func (v *NullableBriefInventoryItemRole) Set(val *BriefInventoryItemRole) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefInventoryItemRole) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefInventoryItemRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefInventoryItemRole(val *BriefInventoryItemRole) *NullableBriefInventoryItemRole {
	return &NullableBriefInventoryItemRole{value: val, isSet: true}
}

func (v NullableBriefInventoryItemRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefInventoryItemRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
