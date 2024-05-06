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

// checks if the WritableContactAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableContactAssignmentRequest{}

// WritableContactAssignmentRequest Adds support for custom fields and tags.
type WritableContactAssignmentRequest struct {
	ContentType          string                          `json:"content_type"`
	ObjectId             int64                           `json:"object_id"`
	Contact              int32                           `json:"contact"`
	Role                 int32                           `json:"role"`
	Priority             *ContactAssignmentPriorityValue `json:"priority,omitempty"`
	Tags                 []NestedTagRequest              `json:"tags,omitempty"`
	CustomFields         map[string]interface{}          `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableContactAssignmentRequest WritableContactAssignmentRequest

// NewWritableContactAssignmentRequest instantiates a new WritableContactAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableContactAssignmentRequest(contentType string, objectId int64, contact int32, role int32) *WritableContactAssignmentRequest {
	this := WritableContactAssignmentRequest{}
	this.ContentType = contentType
	this.ObjectId = objectId
	this.Contact = contact
	this.Role = role
	return &this
}

// NewWritableContactAssignmentRequestWithDefaults instantiates a new WritableContactAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableContactAssignmentRequestWithDefaults() *WritableContactAssignmentRequest {
	this := WritableContactAssignmentRequest{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *WritableContactAssignmentRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *WritableContactAssignmentRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetObjectId returns the ObjectId field value
func (o *WritableContactAssignmentRequest) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *WritableContactAssignmentRequest) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetContact returns the Contact field value
func (o *WritableContactAssignmentRequest) GetContact() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetContactOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *WritableContactAssignmentRequest) SetContact(v int32) {
	o.Contact = v
}

// GetRole returns the Role field value
func (o *WritableContactAssignmentRequest) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *WritableContactAssignmentRequest) SetRole(v int32) {
	o.Role = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *WritableContactAssignmentRequest) GetPriority() ContactAssignmentPriorityValue {
	if o == nil || IsNil(o.Priority) {
		var ret ContactAssignmentPriorityValue
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetPriorityOk() (*ContactAssignmentPriorityValue, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *WritableContactAssignmentRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given ContactAssignmentPriorityValue and assigns it to the Priority field.
func (o *WritableContactAssignmentRequest) SetPriority(v ContactAssignmentPriorityValue) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableContactAssignmentRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableContactAssignmentRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableContactAssignmentRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableContactAssignmentRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableContactAssignmentRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableContactAssignmentRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableContactAssignmentRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableContactAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableContactAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content_type"] = o.ContentType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["contact"] = o.Contact
	toSerialize["role"] = o.Role
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
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

func (o *WritableContactAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content_type",
		"object_id",
		"contact",
		"role",
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

	varWritableContactAssignmentRequest := _WritableContactAssignmentRequest{}

	err = json.Unmarshal(data, &varWritableContactAssignmentRequest)

	if err != nil {
		return err
	}

	*o = WritableContactAssignmentRequest(varWritableContactAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "role")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableContactAssignmentRequest struct {
	value *WritableContactAssignmentRequest
	isSet bool
}

func (v NullableWritableContactAssignmentRequest) Get() *WritableContactAssignmentRequest {
	return v.value
}

func (v *NullableWritableContactAssignmentRequest) Set(val *WritableContactAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableContactAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableContactAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableContactAssignmentRequest(val *WritableContactAssignmentRequest) *NullableWritableContactAssignmentRequest {
	return &NullableWritableContactAssignmentRequest{value: val, isSet: true}
}

func (v NullableWritableContactAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableContactAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
