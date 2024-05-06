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

// checks if the PaginatedTunnelGroupList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedTunnelGroupList{}

// PaginatedTunnelGroupList struct for PaginatedTunnelGroupList
type PaginatedTunnelGroupList struct {
	Count                *int32         `json:"count,omitempty"`
	Next                 NullableString `json:"next,omitempty"`
	Previous             NullableString `json:"previous,omitempty"`
	Results              []TunnelGroup  `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedTunnelGroupList PaginatedTunnelGroupList

// NewPaginatedTunnelGroupList instantiates a new PaginatedTunnelGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTunnelGroupList() *PaginatedTunnelGroupList {
	this := PaginatedTunnelGroupList{}
	return &this
}

// NewPaginatedTunnelGroupListWithDefaults instantiates a new PaginatedTunnelGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTunnelGroupListWithDefaults() *PaginatedTunnelGroupList {
	this := PaginatedTunnelGroupList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedTunnelGroupList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTunnelGroupList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedTunnelGroupList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedTunnelGroupList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedTunnelGroupList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedTunnelGroupList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedTunnelGroupList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedTunnelGroupList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedTunnelGroupList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedTunnelGroupList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedTunnelGroupList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedTunnelGroupList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedTunnelGroupList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedTunnelGroupList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedTunnelGroupList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedTunnelGroupList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedTunnelGroupList) GetResults() []TunnelGroup {
	if o == nil || IsNil(o.Results) {
		var ret []TunnelGroup
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTunnelGroupList) GetResultsOk() ([]TunnelGroup, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedTunnelGroupList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TunnelGroup and assigns it to the Results field.
func (o *PaginatedTunnelGroupList) SetResults(v []TunnelGroup) {
	o.Results = v
}

func (o PaginatedTunnelGroupList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedTunnelGroupList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedTunnelGroupList) UnmarshalJSON(data []byte) (err error) {
	varPaginatedTunnelGroupList := _PaginatedTunnelGroupList{}

	err = json.Unmarshal(data, &varPaginatedTunnelGroupList)

	if err != nil {
		return err
	}

	*o = PaginatedTunnelGroupList(varPaginatedTunnelGroupList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedTunnelGroupList struct {
	value *PaginatedTunnelGroupList
	isSet bool
}

func (v NullablePaginatedTunnelGroupList) Get() *PaginatedTunnelGroupList {
	return v.value
}

func (v *NullablePaginatedTunnelGroupList) Set(val *PaginatedTunnelGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTunnelGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTunnelGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTunnelGroupList(val *PaginatedTunnelGroupList) *NullablePaginatedTunnelGroupList {
	return &NullablePaginatedTunnelGroupList{value: val, isSet: true}
}

func (v NullablePaginatedTunnelGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTunnelGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
