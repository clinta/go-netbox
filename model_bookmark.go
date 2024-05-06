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

// checks if the Bookmark type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bookmark{}

// Bookmark Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Bookmark struct {
	Id                   int32       `json:"id"`
	Url                  string      `json:"url"`
	Display              string      `json:"display"`
	ObjectType           string      `json:"object_type"`
	ObjectId             int64       `json:"object_id"`
	Object               interface{} `json:"object"`
	User                 NestedUser  `json:"user"`
	Created              time.Time   `json:"created"`
	AdditionalProperties map[string]interface{}
}

type _Bookmark Bookmark

// NewBookmark instantiates a new Bookmark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmark(id int32, url string, display string, objectType string, objectId int64, object interface{}, user NestedUser, created time.Time) *Bookmark {
	this := Bookmark{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Object = object
	this.User = user
	this.Created = created
	return &this
}

// NewBookmarkWithDefaults instantiates a new Bookmark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkWithDefaults() *Bookmark {
	this := Bookmark{}
	return &this
}

// GetId returns the Id field value
func (o *Bookmark) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bookmark) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Bookmark) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Bookmark) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *Bookmark) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Bookmark) SetDisplay(v string) {
	o.Display = v
}

// GetObjectType returns the ObjectType field value
func (o *Bookmark) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *Bookmark) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *Bookmark) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *Bookmark) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Bookmark) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bookmark) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Bookmark) SetObject(v interface{}) {
	o.Object = v
}

// GetUser returns the User field value
func (o *Bookmark) GetUser() NestedUser {
	if o == nil {
		var ret NestedUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetUserOk() (*NestedUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Bookmark) SetUser(v NestedUser) {
	o.User = v
}

// GetCreated returns the Created field value
func (o *Bookmark) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Bookmark) SetCreated(v time.Time) {
	o.Created = v
}

func (o Bookmark) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bookmark) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	toSerialize["user"] = o.User
	toSerialize["created"] = o.Created

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Bookmark) UnmarshalJSON(data []byte) (err error) {
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
		"user",
		"created",
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

	varBookmark := _Bookmark{}

	err = json.Unmarshal(data, &varBookmark)

	if err != nil {
		return err
	}

	*o = Bookmark(varBookmark)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "user")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBookmark struct {
	value *Bookmark
	isSet bool
}

func (v NullableBookmark) Get() *Bookmark {
	return v.value
}

func (v *NullableBookmark) Set(val *Bookmark) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmark) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmark(val *Bookmark) *NullableBookmark {
	return &NullableBookmark{value: val, isSet: true}
}

func (v NullableBookmark) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
