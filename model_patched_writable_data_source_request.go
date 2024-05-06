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

// checks if the PatchedWritableDataSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableDataSourceRequest{}

// PatchedWritableDataSourceRequest Adds support for custom fields and tags.
type PatchedWritableDataSourceRequest struct {
	Name        *string     `json:"name,omitempty"`
	Type        *string     `json:"type,omitempty"`
	SourceUrl   *string     `json:"source_url,omitempty"`
	Enabled     *bool       `json:"enabled,omitempty"`
	Description *string     `json:"description,omitempty"`
	Comments    *string     `json:"comments,omitempty"`
	Parameters  interface{} `json:"parameters,omitempty"`
	// Patterns (one per line) matching files to ignore when syncing
	IgnoreRules          *string                `json:"ignore_rules,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableDataSourceRequest PatchedWritableDataSourceRequest

// NewPatchedWritableDataSourceRequest instantiates a new PatchedWritableDataSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableDataSourceRequest() *PatchedWritableDataSourceRequest {
	this := PatchedWritableDataSourceRequest{}
	return &this
}

// NewPatchedWritableDataSourceRequestWithDefaults instantiates a new PatchedWritableDataSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableDataSourceRequestWithDefaults() *PatchedWritableDataSourceRequest {
	this := PatchedWritableDataSourceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableDataSourceRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritableDataSourceRequest) SetType(v string) {
	o.Type = &v
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetSourceUrl() string {
	if o == nil || IsNil(o.SourceUrl) {
		var ret string
		return ret
	}
	return *o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetSourceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUrl) {
		return nil, false
	}
	return o.SourceUrl, true
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasSourceUrl() bool {
	if o != nil && !IsNil(o.SourceUrl) {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given string and assigns it to the SourceUrl field.
func (o *PatchedWritableDataSourceRequest) SetSourceUrl(v string) {
	o.SourceUrl = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedWritableDataSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableDataSourceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableDataSourceRequest) SetComments(v string) {
	o.Comments = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableDataSourceRequest) GetParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableDataSourceRequest) GetParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given interface{} and assigns it to the Parameters field.
func (o *PatchedWritableDataSourceRequest) SetParameters(v interface{}) {
	o.Parameters = v
}

// GetIgnoreRules returns the IgnoreRules field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetIgnoreRules() string {
	if o == nil || IsNil(o.IgnoreRules) {
		var ret string
		return ret
	}
	return *o.IgnoreRules
}

// GetIgnoreRulesOk returns a tuple with the IgnoreRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetIgnoreRulesOk() (*string, bool) {
	if o == nil || IsNil(o.IgnoreRules) {
		return nil, false
	}
	return o.IgnoreRules, true
}

// HasIgnoreRules returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasIgnoreRules() bool {
	if o != nil && !IsNil(o.IgnoreRules) {
		return true
	}

	return false
}

// SetIgnoreRules gets a reference to the given string and assigns it to the IgnoreRules field.
func (o *PatchedWritableDataSourceRequest) SetIgnoreRules(v string) {
	o.IgnoreRules = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableDataSourceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDataSourceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableDataSourceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableDataSourceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableDataSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.SourceUrl) {
		toSerialize["source_url"] = o.SourceUrl
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.IgnoreRules) {
		toSerialize["ignore_rules"] = o.IgnoreRules
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableDataSourceRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableDataSourceRequest := _PatchedWritableDataSourceRequest{}

	err = json.Unmarshal(data, &varPatchedWritableDataSourceRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableDataSourceRequest(varPatchedWritableDataSourceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "source_url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "ignore_rules")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableDataSourceRequest struct {
	value *PatchedWritableDataSourceRequest
	isSet bool
}

func (v NullablePatchedWritableDataSourceRequest) Get() *PatchedWritableDataSourceRequest {
	return v.value
}

func (v *NullablePatchedWritableDataSourceRequest) Set(val *PatchedWritableDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableDataSourceRequest(val *PatchedWritableDataSourceRequest) *NullablePatchedWritableDataSourceRequest {
	return &NullablePatchedWritableDataSourceRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
