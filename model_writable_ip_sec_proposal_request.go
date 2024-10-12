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

// checks if the WritableIPSecProposalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableIPSecProposalRequest{}

// WritableIPSecProposalRequest Adds support for custom fields and tags.
type WritableIPSecProposalRequest struct {
	Name                    string           `json:"name"`
	Description             *string          `json:"description,omitempty"`
	EncryptionAlgorithm     *Encryption1     `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *Authentication1 `json:"authentication_algorithm,omitempty"`
	// Security association lifetime (seconds)
	SaLifetimeSeconds NullableInt32 `json:"sa_lifetime_seconds,omitempty"`
	// Security association lifetime (in kilobytes)
	SaLifetimeData       NullableInt32          `json:"sa_lifetime_data,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableIPSecProposalRequest WritableIPSecProposalRequest

// NewWritableIPSecProposalRequest instantiates a new WritableIPSecProposalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableIPSecProposalRequest(name string) *WritableIPSecProposalRequest {
	this := WritableIPSecProposalRequest{}
	this.Name = name
	return &this
}

// NewWritableIPSecProposalRequestWithDefaults instantiates a new WritableIPSecProposalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableIPSecProposalRequestWithDefaults() *WritableIPSecProposalRequest {
	this := WritableIPSecProposalRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableIPSecProposalRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableIPSecProposalRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableIPSecProposalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableIPSecProposalRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value if set, zero value otherwise.
func (o *WritableIPSecProposalRequest) GetEncryptionAlgorithm() Encryption1 {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		var ret Encryption1
		return ret
	}
	return *o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetEncryptionAlgorithmOk() (*Encryption1, bool) {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		return nil, false
	}
	return o.EncryptionAlgorithm, true
}

// HasEncryptionAlgorithm returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.EncryptionAlgorithm) {
		return true
	}

	return false
}

// SetEncryptionAlgorithm gets a reference to the given Encryption1 and assigns it to the EncryptionAlgorithm field.
func (o *WritableIPSecProposalRequest) SetEncryptionAlgorithm(v Encryption1) {
	o.EncryptionAlgorithm = &v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value if set, zero value otherwise.
func (o *WritableIPSecProposalRequest) GetAuthenticationAlgorithm() Authentication1 {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		var ret Authentication1
		return ret
	}
	return *o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetAuthenticationAlgorithmOk() (*Authentication1, bool) {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		return nil, false
	}
	return o.AuthenticationAlgorithm, true
}

// HasAuthenticationAlgorithm returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasAuthenticationAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationAlgorithm gets a reference to the given Authentication1 and assigns it to the AuthenticationAlgorithm field.
func (o *WritableIPSecProposalRequest) SetAuthenticationAlgorithm(v Authentication1) {
	o.AuthenticationAlgorithm = &v
}

// GetSaLifetimeSeconds returns the SaLifetimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPSecProposalRequest) GetSaLifetimeSeconds() int32 {
	if o == nil || IsNil(o.SaLifetimeSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetimeSeconds.Get()
}

// GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPSecProposalRequest) GetSaLifetimeSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeSeconds.Get(), o.SaLifetimeSeconds.IsSet()
}

// HasSaLifetimeSeconds returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasSaLifetimeSeconds() bool {
	if o != nil && o.SaLifetimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeSeconds gets a reference to the given NullableInt32 and assigns it to the SaLifetimeSeconds field.
func (o *WritableIPSecProposalRequest) SetSaLifetimeSeconds(v int32) {
	o.SaLifetimeSeconds.Set(&v)
}

// SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil
func (o *WritableIPSecProposalRequest) SetSaLifetimeSecondsNil() {
	o.SaLifetimeSeconds.Set(nil)
}

// UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
func (o *WritableIPSecProposalRequest) UnsetSaLifetimeSeconds() {
	o.SaLifetimeSeconds.Unset()
}

// GetSaLifetimeData returns the SaLifetimeData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPSecProposalRequest) GetSaLifetimeData() int32 {
	if o == nil || IsNil(o.SaLifetimeData.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetimeData.Get()
}

// GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPSecProposalRequest) GetSaLifetimeDataOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeData.Get(), o.SaLifetimeData.IsSet()
}

// HasSaLifetimeData returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasSaLifetimeData() bool {
	if o != nil && o.SaLifetimeData.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeData gets a reference to the given NullableInt32 and assigns it to the SaLifetimeData field.
func (o *WritableIPSecProposalRequest) SetSaLifetimeData(v int32) {
	o.SaLifetimeData.Set(&v)
}

// SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil
func (o *WritableIPSecProposalRequest) SetSaLifetimeDataNil() {
	o.SaLifetimeData.Set(nil)
}

// UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
func (o *WritableIPSecProposalRequest) UnsetSaLifetimeData() {
	o.SaLifetimeData.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableIPSecProposalRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableIPSecProposalRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableIPSecProposalRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableIPSecProposalRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableIPSecProposalRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecProposalRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableIPSecProposalRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableIPSecProposalRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableIPSecProposalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableIPSecProposalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EncryptionAlgorithm) {
		toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	}
	if !IsNil(o.AuthenticationAlgorithm) {
		toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	}
	if o.SaLifetimeSeconds.IsSet() {
		toSerialize["sa_lifetime_seconds"] = o.SaLifetimeSeconds.Get()
	}
	if o.SaLifetimeData.IsSet() {
		toSerialize["sa_lifetime_data"] = o.SaLifetimeData.Get()
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

func (o *WritableIPSecProposalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varWritableIPSecProposalRequest := _WritableIPSecProposalRequest{}

	err = json.Unmarshal(data, &varWritableIPSecProposalRequest)

	if err != nil {
		return err
	}

	*o = WritableIPSecProposalRequest(varWritableIPSecProposalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "sa_lifetime_seconds")
		delete(additionalProperties, "sa_lifetime_data")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableIPSecProposalRequest struct {
	value *WritableIPSecProposalRequest
	isSet bool
}

func (v NullableWritableIPSecProposalRequest) Get() *WritableIPSecProposalRequest {
	return v.value
}

func (v *NullableWritableIPSecProposalRequest) Set(val *WritableIPSecProposalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableIPSecProposalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableIPSecProposalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableIPSecProposalRequest(val *WritableIPSecProposalRequest) *NullableWritableIPSecProposalRequest {
	return &NullableWritableIPSecProposalRequest{value: val, isSet: true}
}

func (v NullableWritableIPSecProposalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableIPSecProposalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
