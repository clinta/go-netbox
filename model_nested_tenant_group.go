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

// checks if the NestedTenantGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedTenantGroup{}

// NestedTenantGroup Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedTenantGroup struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	DisplayUrl           string `json:"display_url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	Depth                int32  `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _NestedTenantGroup NestedTenantGroup

// NewNestedTenantGroup instantiates a new NestedTenantGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTenantGroup(id int32, url string, displayUrl string, display string, name string, slug string, depth int32) *NestedTenantGroup {
	this := NestedTenantGroup{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Depth = depth
	return &this
}

// NewNestedTenantGroupWithDefaults instantiates a new NestedTenantGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTenantGroupWithDefaults() *NestedTenantGroup {
	this := NestedTenantGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NestedTenantGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedTenantGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedTenantGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedTenantGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *NestedTenantGroup) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *NestedTenantGroup) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *NestedTenantGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedTenantGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedTenantGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTenantGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedTenantGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedTenantGroup) SetSlug(v string) {
	o.Slug = v
}

// GetDepth returns the Depth field value
func (o *NestedTenantGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *NestedTenantGroup) SetDepth(v int32) {
	o.Depth = v
}

func (o NestedTenantGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedTenantGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["_depth"] = o.Depth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedTenantGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"slug",
		"_depth",
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

	varNestedTenantGroup := _NestedTenantGroup{}

	err = json.Unmarshal(data, &varNestedTenantGroup)

	if err != nil {
		return err
	}

	*o = NestedTenantGroup(varNestedTenantGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedTenantGroup struct {
	value *NestedTenantGroup
	isSet bool
}

func (v NullableNestedTenantGroup) Get() *NestedTenantGroup {
	return v.value
}

func (v *NullableNestedTenantGroup) Set(val *NestedTenantGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTenantGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTenantGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTenantGroup(val *NestedTenantGroup) *NullableNestedTenantGroup {
	return &NullableNestedTenantGroup{value: val, isSet: true}
}

func (v NullableNestedTenantGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTenantGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
