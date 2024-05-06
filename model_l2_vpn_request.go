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

// checks if the L2VPNRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2VPNRequest{}

// L2VPNRequest Adds support for custom fields and tags.
type L2VPNRequest struct {
	Identifier           NullableInt64               `json:"identifier,omitempty"`
	Name                 string                      `json:"name"`
	Slug                 string                      `json:"slug"`
	Type                 *L2VPNTypeValue             `json:"type,omitempty"`
	ImportTargets        []int32                     `json:"import_targets,omitempty"`
	ExportTargets        []int32                     `json:"export_targets,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	Comments             *string                     `json:"comments,omitempty"`
	Tenant               NullableNestedTenantRequest `json:"tenant,omitempty"`
	Tags                 []NestedTagRequest          `json:"tags,omitempty"`
	CustomFields         map[string]interface{}      `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _L2VPNRequest L2VPNRequest

// NewL2VPNRequest instantiates a new L2VPNRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2VPNRequest(name string, slug string) *L2VPNRequest {
	this := L2VPNRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewL2VPNRequestWithDefaults instantiates a new L2VPNRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2VPNRequestWithDefaults() *L2VPNRequest {
	this := L2VPNRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *L2VPNRequest) GetIdentifier() int64 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int64
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPNRequest) GetIdentifierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *L2VPNRequest) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt64 and assigns it to the Identifier field.
func (o *L2VPNRequest) SetIdentifier(v int64) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *L2VPNRequest) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *L2VPNRequest) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetName returns the Name field value
func (o *L2VPNRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *L2VPNRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *L2VPNRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *L2VPNRequest) SetSlug(v string) {
	o.Slug = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *L2VPNRequest) GetType() L2VPNTypeValue {
	if o == nil || IsNil(o.Type) {
		var ret L2VPNTypeValue
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetTypeOk() (*L2VPNTypeValue, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *L2VPNRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given L2VPNTypeValue and assigns it to the Type field.
func (o *L2VPNRequest) SetType(v L2VPNTypeValue) {
	o.Type = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *L2VPNRequest) GetImportTargets() []int32 {
	if o == nil || IsNil(o.ImportTargets) {
		var ret []int32
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetImportTargetsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ImportTargets) {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *L2VPNRequest) HasImportTargets() bool {
	if o != nil && !IsNil(o.ImportTargets) {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []int32 and assigns it to the ImportTargets field.
func (o *L2VPNRequest) SetImportTargets(v []int32) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *L2VPNRequest) GetExportTargets() []int32 {
	if o == nil || IsNil(o.ExportTargets) {
		var ret []int32
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetExportTargetsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ExportTargets) {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *L2VPNRequest) HasExportTargets() bool {
	if o != nil && !IsNil(o.ExportTargets) {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []int32 and assigns it to the ExportTargets field.
func (o *L2VPNRequest) SetExportTargets(v []int32) {
	o.ExportTargets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *L2VPNRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *L2VPNRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *L2VPNRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *L2VPNRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *L2VPNRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *L2VPNRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *L2VPNRequest) GetTenant() NestedTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPNRequest) GetTenantOk() (*NestedTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *L2VPNRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenantRequest and assigns it to the Tenant field.
func (o *L2VPNRequest) SetTenant(v NestedTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *L2VPNRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *L2VPNRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *L2VPNRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *L2VPNRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *L2VPNRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *L2VPNRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *L2VPNRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *L2VPNRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o L2VPNRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2VPNRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ImportTargets) {
		toSerialize["import_targets"] = o.ImportTargets
	}
	if !IsNil(o.ExportTargets) {
		toSerialize["export_targets"] = o.ExportTargets
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

func (o *L2VPNRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varL2VPNRequest := _L2VPNRequest{}

	err = json.Unmarshal(data, &varL2VPNRequest)

	if err != nil {
		return err
	}

	*o = L2VPNRequest(varL2VPNRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "type")
		delete(additionalProperties, "import_targets")
		delete(additionalProperties, "export_targets")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableL2VPNRequest struct {
	value *L2VPNRequest
	isSet bool
}

func (v NullableL2VPNRequest) Get() *L2VPNRequest {
	return v.value
}

func (v *NullableL2VPNRequest) Set(val *L2VPNRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNRequest(val *L2VPNRequest) *NullableL2VPNRequest {
	return &NullableL2VPNRequest{value: val, isSet: true}
}

func (v NullableL2VPNRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
