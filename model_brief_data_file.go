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

// checks if the BriefDataFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefDataFile{}

// BriefDataFile Adds support for custom fields and tags.
type BriefDataFile struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// File path relative to the data source's root
	Path                 string `json:"path"`
	AdditionalProperties map[string]interface{}
}

type _BriefDataFile BriefDataFile

// NewBriefDataFile instantiates a new BriefDataFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefDataFile(id int32, url string, display string, path string) *BriefDataFile {
	this := BriefDataFile{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Path = path
	return &this
}

// NewBriefDataFileWithDefaults instantiates a new BriefDataFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefDataFileWithDefaults() *BriefDataFile {
	this := BriefDataFile{}
	return &this
}

// GetId returns the Id field value
func (o *BriefDataFile) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefDataFile) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefDataFile) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefDataFile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefDataFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefDataFile) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefDataFile) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefDataFile) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefDataFile) SetDisplay(v string) {
	o.Display = v
}

// GetPath returns the Path field value
func (o *BriefDataFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *BriefDataFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *BriefDataFile) SetPath(v string) {
	o.Path = v
}

func (o BriefDataFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefDataFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefDataFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"path",
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

	varBriefDataFile := _BriefDataFile{}

	err = json.Unmarshal(data, &varBriefDataFile)

	if err != nil {
		return err
	}

	*o = BriefDataFile(varBriefDataFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefDataFile struct {
	value *BriefDataFile
	isSet bool
}

func (v NullableBriefDataFile) Get() *BriefDataFile {
	return v.value
}

func (v *NullableBriefDataFile) Set(val *BriefDataFile) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefDataFile) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefDataFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefDataFile(val *BriefDataFile) *NullableBriefDataFile {
	return &NullableBriefDataFile{value: val, isSet: true}
}

func (v NullableBriefDataFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefDataFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
