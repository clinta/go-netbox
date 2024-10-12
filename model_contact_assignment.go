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
	"time"
)

// checks if the ContactAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactAssignment{}

// ContactAssignment Adds support for custom fields and tags.
type ContactAssignment struct {
	Id                   int32                                          `json:"id"`
	Url                  string                                         `json:"url"`
	Display              string                                         `json:"display"`
	ObjectType           string                                         `json:"object_type"`
	ObjectId             int64                                          `json:"object_id"`
	Object               map[string]interface{}                         `json:"object"`
	Contact              BriefContact                                   `json:"contact"`
	Role                 NullableBriefContactRole                       `json:"role,omitempty"`
	Priority             *BriefCircuitGroupAssignmentSerializerPriority `json:"priority,omitempty"`
	Tags                 []NestedTag                                    `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                         `json:"custom_fields,omitempty"`
	Created              NullableTime                                   `json:"created"`
	LastUpdated          NullableTime                                   `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ContactAssignment ContactAssignment

// NewContactAssignment instantiates a new ContactAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactAssignment(id int32, url string, display string, objectType string, objectId int64, object map[string]interface{}, contact BriefContact, created NullableTime, lastUpdated NullableTime) *ContactAssignment {
	this := ContactAssignment{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Object = object
	this.Contact = contact
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewContactAssignmentWithDefaults instantiates a new ContactAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactAssignmentWithDefaults() *ContactAssignment {
	this := ContactAssignment{}
	return &this
}

// GetId returns the Id field value
func (o *ContactAssignment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContactAssignment) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ContactAssignment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ContactAssignment) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ContactAssignment) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ContactAssignment) SetDisplay(v string) {
	o.Display = v
}

// GetObjectType returns the ObjectType field value
func (o *ContactAssignment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ContactAssignment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *ContactAssignment) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ContactAssignment) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetObject returns the Object field value
func (o *ContactAssignment) GetObject() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetObjectOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Object, true
}

// SetObject sets field value
func (o *ContactAssignment) SetObject(v map[string]interface{}) {
	o.Object = v
}

// GetContact returns the Contact field value
func (o *ContactAssignment) GetContact() BriefContact {
	if o == nil {
		var ret BriefContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetContactOk() (*BriefContact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ContactAssignment) SetContact(v BriefContact) {
	o.Contact = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactAssignment) GetRole() BriefContactRole {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefContactRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactAssignment) GetRoleOk() (*BriefContactRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *ContactAssignment) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefContactRole and assigns it to the Role field.
func (o *ContactAssignment) SetRole(v BriefContactRole) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *ContactAssignment) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *ContactAssignment) UnsetRole() {
	o.Role.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ContactAssignment) GetPriority() BriefCircuitGroupAssignmentSerializerPriority {
	if o == nil || IsNil(o.Priority) {
		var ret BriefCircuitGroupAssignmentSerializerPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ContactAssignment) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given BriefCircuitGroupAssignmentSerializerPriority and assigns it to the Priority field.
func (o *ContactAssignment) SetPriority(v BriefCircuitGroupAssignmentSerializerPriority) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContactAssignment) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContactAssignment) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ContactAssignment) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ContactAssignment) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignment) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ContactAssignment) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ContactAssignment) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ContactAssignment) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactAssignment) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ContactAssignment) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ContactAssignment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactAssignment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ContactAssignment) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ContactAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["object"] = o.Object
	toSerialize["contact"] = o.Contact
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
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

func (o *ContactAssignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"object_type",
		"object_id",
		"object",
		"contact",
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

	varContactAssignment := _ContactAssignment{}

	err = json.Unmarshal(data, &varContactAssignment)

	if err != nil {
		return err
	}

	*o = ContactAssignment(varContactAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "role")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactAssignment struct {
	value *ContactAssignment
	isSet bool
}

func (v NullableContactAssignment) Get() *ContactAssignment {
	return v.value
}

func (v *NullableContactAssignment) Set(val *ContactAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableContactAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableContactAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactAssignment(val *ContactAssignment) *NullableContactAssignment {
	return &NullableContactAssignment{value: val, isSet: true}
}

func (v NullableContactAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
