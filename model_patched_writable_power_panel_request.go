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

// checks if the PatchedWritablePowerPanelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritablePowerPanelRequest{}

// PatchedWritablePowerPanelRequest Adds support for custom fields and tags.
type PatchedWritablePowerPanelRequest struct {
	Site                 *int32                 `json:"site,omitempty"`
	Location             NullableInt32          `json:"location,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritablePowerPanelRequest PatchedWritablePowerPanelRequest

// NewPatchedWritablePowerPanelRequest instantiates a new PatchedWritablePowerPanelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerPanelRequest() *PatchedWritablePowerPanelRequest {
	this := PatchedWritablePowerPanelRequest{}
	return &this
}

// NewPatchedWritablePowerPanelRequestWithDefaults instantiates a new PatchedWritablePowerPanelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerPanelRequestWithDefaults() *PatchedWritablePowerPanelRequest {
	this := PatchedWritablePowerPanelRequest{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanelRequest) GetSite() int32 {
	if o == nil || IsNil(o.Site) {
		var ret int32
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanelRequest) GetSiteOk() (*int32, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given int32 and assigns it to the Site field.
func (o *PatchedWritablePowerPanelRequest) SetSite(v int32) {
	o.Site = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPanelRequest) GetLocation() int32 {
	if o == nil || IsNil(o.Location.Get()) {
		var ret int32
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPanelRequest) GetLocationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableInt32 and assigns it to the Location field.
func (o *PatchedWritablePowerPanelRequest) SetLocation(v int32) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *PatchedWritablePowerPanelRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *PatchedWritablePowerPanelRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerPanelRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanelRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritablePowerPanelRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanelRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanelRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritablePowerPanelRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanelRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanelRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritablePowerPanelRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanelRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanelRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanelRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePowerPanelRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritablePowerPanelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritablePowerPanelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

func (o *PatchedWritablePowerPanelRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritablePowerPanelRequest := _PatchedWritablePowerPanelRequest{}

	err = json.Unmarshal(data, &varPatchedWritablePowerPanelRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritablePowerPanelRequest(varPatchedWritablePowerPanelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritablePowerPanelRequest struct {
	value *PatchedWritablePowerPanelRequest
	isSet bool
}

func (v NullablePatchedWritablePowerPanelRequest) Get() *PatchedWritablePowerPanelRequest {
	return v.value
}

func (v *NullablePatchedWritablePowerPanelRequest) Set(val *PatchedWritablePowerPanelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPanelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPanelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPanelRequest(val *PatchedWritablePowerPanelRequest) *NullablePatchedWritablePowerPanelRequest {
	return &NullablePatchedWritablePowerPanelRequest{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPanelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPanelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
