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

// checks if the PatchedDashboardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDashboardRequest{}

// PatchedDashboardRequest struct for PatchedDashboardRequest
type PatchedDashboardRequest struct {
	Layout               interface{} `json:"layout,omitempty"`
	Config               interface{} `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedDashboardRequest PatchedDashboardRequest

// NewPatchedDashboardRequest instantiates a new PatchedDashboardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDashboardRequest() *PatchedDashboardRequest {
	this := PatchedDashboardRequest{}
	return &this
}

// NewPatchedDashboardRequestWithDefaults instantiates a new PatchedDashboardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDashboardRequestWithDefaults() *PatchedDashboardRequest {
	this := PatchedDashboardRequest{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDashboardRequest) GetLayout() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDashboardRequest) GetLayoutOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return &o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *PatchedDashboardRequest) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given interface{} and assigns it to the Layout field.
func (o *PatchedDashboardRequest) SetLayout(v interface{}) {
	o.Layout = v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDashboardRequest) GetConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDashboardRequest) GetConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PatchedDashboardRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *PatchedDashboardRequest) SetConfig(v interface{}) {
	o.Config = v
}

func (o PatchedDashboardRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDashboardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedDashboardRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedDashboardRequest := _PatchedDashboardRequest{}

	err = json.Unmarshal(data, &varPatchedDashboardRequest)

	if err != nil {
		return err
	}

	*o = PatchedDashboardRequest(varPatchedDashboardRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "layout")
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedDashboardRequest struct {
	value *PatchedDashboardRequest
	isSet bool
}

func (v NullablePatchedDashboardRequest) Get() *PatchedDashboardRequest {
	return v.value
}

func (v *NullablePatchedDashboardRequest) Set(val *PatchedDashboardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDashboardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDashboardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDashboardRequest(val *PatchedDashboardRequest) *NullablePatchedDashboardRequest {
	return &NullablePatchedDashboardRequest{value: val, isSet: true}
}

func (v NullablePatchedDashboardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDashboardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
