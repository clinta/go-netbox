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

// checks if the CustomLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomLink{}

// CustomLink Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomLink struct {
	Id           int32    `json:"id"`
	Url          string   `json:"url"`
	Display      string   `json:"display"`
	ContentTypes []string `json:"content_types"`
	Name         string   `json:"name"`
	Enabled      *bool    `json:"enabled,omitempty"`
	// Jinja2 template code for link text
	LinkText string `json:"link_text"`
	// Jinja2 template code for link URL
	LinkUrl string `json:"link_url"`
	Weight  *int32 `json:"weight,omitempty"`
	// Links with the same group will appear as a dropdown menu
	GroupName   *string                `json:"group_name,omitempty"`
	ButtonClass *CustomLinkButtonClass `json:"button_class,omitempty"`
	// Force link to open in a new window
	NewWindow            *bool        `json:"new_window,omitempty"`
	Created              NullableTime `json:"created"`
	LastUpdated          NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _CustomLink CustomLink

// NewCustomLink instantiates a new CustomLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomLink(id int32, url string, display string, contentTypes []string, name string, linkText string, linkUrl string, created NullableTime, lastUpdated NullableTime) *CustomLink {
	this := CustomLink{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ContentTypes = contentTypes
	this.Name = name
	this.LinkText = linkText
	this.LinkUrl = linkUrl
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewCustomLinkWithDefaults instantiates a new CustomLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomLinkWithDefaults() *CustomLink {
	this := CustomLink{}
	return &this
}

// GetId returns the Id field value
func (o *CustomLink) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomLink) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CustomLink) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CustomLink) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *CustomLink) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CustomLink) SetDisplay(v string) {
	o.Display = v
}

// GetContentTypes returns the ContentTypes field value
func (o *CustomLink) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *CustomLink) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *CustomLink) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomLink) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomLink) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLink) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomLink) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomLink) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLinkText returns the LinkText field value
func (o *CustomLink) GetLinkText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkText
}

// GetLinkTextOk returns a tuple with the LinkText field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetLinkTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkText, true
}

// SetLinkText sets field value
func (o *CustomLink) SetLinkText(v string) {
	o.LinkText = v
}

// GetLinkUrl returns the LinkUrl field value
func (o *CustomLink) GetLinkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value
// and a boolean to check if the value has been set.
func (o *CustomLink) GetLinkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkUrl, true
}

// SetLinkUrl sets field value
func (o *CustomLink) SetLinkUrl(v string) {
	o.LinkUrl = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CustomLink) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLink) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CustomLink) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *CustomLink) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *CustomLink) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLink) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *CustomLink) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *CustomLink) SetGroupName(v string) {
	o.GroupName = &v
}

// GetButtonClass returns the ButtonClass field value if set, zero value otherwise.
func (o *CustomLink) GetButtonClass() CustomLinkButtonClass {
	if o == nil || IsNil(o.ButtonClass) {
		var ret CustomLinkButtonClass
		return ret
	}
	return *o.ButtonClass
}

// GetButtonClassOk returns a tuple with the ButtonClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLink) GetButtonClassOk() (*CustomLinkButtonClass, bool) {
	if o == nil || IsNil(o.ButtonClass) {
		return nil, false
	}
	return o.ButtonClass, true
}

// HasButtonClass returns a boolean if a field has been set.
func (o *CustomLink) HasButtonClass() bool {
	if o != nil && !IsNil(o.ButtonClass) {
		return true
	}

	return false
}

// SetButtonClass gets a reference to the given CustomLinkButtonClass and assigns it to the ButtonClass field.
func (o *CustomLink) SetButtonClass(v CustomLinkButtonClass) {
	o.ButtonClass = &v
}

// GetNewWindow returns the NewWindow field value if set, zero value otherwise.
func (o *CustomLink) GetNewWindow() bool {
	if o == nil || IsNil(o.NewWindow) {
		var ret bool
		return ret
	}
	return *o.NewWindow
}

// GetNewWindowOk returns a tuple with the NewWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLink) GetNewWindowOk() (*bool, bool) {
	if o == nil || IsNil(o.NewWindow) {
		return nil, false
	}
	return o.NewWindow, true
}

// HasNewWindow returns a boolean if a field has been set.
func (o *CustomLink) HasNewWindow() bool {
	if o != nil && !IsNil(o.NewWindow) {
		return true
	}

	return false
}

// SetNewWindow gets a reference to the given bool and assigns it to the NewWindow field.
func (o *CustomLink) SetNewWindow(v bool) {
	o.NewWindow = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomLink) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomLink) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *CustomLink) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomLink) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomLink) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CustomLink) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o CustomLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["link_text"] = o.LinkText
	toSerialize["link_url"] = o.LinkUrl
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.ButtonClass) {
		toSerialize["button_class"] = o.ButtonClass
	}
	if !IsNil(o.NewWindow) {
		toSerialize["new_window"] = o.NewWindow
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"content_types",
		"name",
		"link_text",
		"link_url",
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

	varCustomLink := _CustomLink{}

	err = json.Unmarshal(data, &varCustomLink)

	if err != nil {
		return err
	}

	*o = CustomLink(varCustomLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "link_text")
		delete(additionalProperties, "link_url")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "button_class")
		delete(additionalProperties, "new_window")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomLink struct {
	value *CustomLink
	isSet bool
}

func (v NullableCustomLink) Get() *CustomLink {
	return v.value
}

func (v *NullableCustomLink) Set(val *CustomLink) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomLink) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomLink(val *CustomLink) *NullableCustomLink {
	return &NullableCustomLink{value: val, isSet: true}
}

func (v NullableCustomLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
