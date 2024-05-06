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

// checks if the PatchedWritableCustomFieldChoiceSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableCustomFieldChoiceSetRequest{}

// PatchedWritableCustomFieldChoiceSetRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableCustomFieldChoiceSetRequest struct {
	Name         *string                                                `json:"name,omitempty"`
	Description  *string                                                `json:"description,omitempty"`
	BaseChoices  *PatchedWritableCustomFieldChoiceSetRequestBaseChoices `json:"base_choices,omitempty"`
	ExtraChoices [][]interface{}                                        `json:"extra_choices,omitempty"`
	// Choices are automatically ordered alphabetically
	OrderAlphabetically  *bool `json:"order_alphabetically,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableCustomFieldChoiceSetRequest PatchedWritableCustomFieldChoiceSetRequest

// NewPatchedWritableCustomFieldChoiceSetRequest instantiates a new PatchedWritableCustomFieldChoiceSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCustomFieldChoiceSetRequest() *PatchedWritableCustomFieldChoiceSetRequest {
	this := PatchedWritableCustomFieldChoiceSetRequest{}
	return &this
}

// NewPatchedWritableCustomFieldChoiceSetRequestWithDefaults instantiates a new PatchedWritableCustomFieldChoiceSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCustomFieldChoiceSetRequestWithDefaults() *PatchedWritableCustomFieldChoiceSetRequest {
	this := PatchedWritableCustomFieldChoiceSetRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableCustomFieldChoiceSetRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCustomFieldChoiceSetRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBaseChoices returns the BaseChoices field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetBaseChoices() PatchedWritableCustomFieldChoiceSetRequestBaseChoices {
	if o == nil || IsNil(o.BaseChoices) {
		var ret PatchedWritableCustomFieldChoiceSetRequestBaseChoices
		return ret
	}
	return *o.BaseChoices
}

// GetBaseChoicesOk returns a tuple with the BaseChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetBaseChoicesOk() (*PatchedWritableCustomFieldChoiceSetRequestBaseChoices, bool) {
	if o == nil || IsNil(o.BaseChoices) {
		return nil, false
	}
	return o.BaseChoices, true
}

// HasBaseChoices returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) HasBaseChoices() bool {
	if o != nil && !IsNil(o.BaseChoices) {
		return true
	}

	return false
}

// SetBaseChoices gets a reference to the given PatchedWritableCustomFieldChoiceSetRequestBaseChoices and assigns it to the BaseChoices field.
func (o *PatchedWritableCustomFieldChoiceSetRequest) SetBaseChoices(v PatchedWritableCustomFieldChoiceSetRequestBaseChoices) {
	o.BaseChoices = &v
}

// GetExtraChoices returns the ExtraChoices field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{} {
	if o == nil || IsNil(o.ExtraChoices) {
		var ret [][]interface{}
		return ret
	}
	return o.ExtraChoices
}

// GetExtraChoicesOk returns a tuple with the ExtraChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetExtraChoicesOk() ([][]interface{}, bool) {
	if o == nil || IsNil(o.ExtraChoices) {
		return nil, false
	}
	return o.ExtraChoices, true
}

// HasExtraChoices returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) HasExtraChoices() bool {
	if o != nil && !IsNil(o.ExtraChoices) {
		return true
	}

	return false
}

// SetExtraChoices gets a reference to the given [][]interface{} and assigns it to the ExtraChoices field.
func (o *PatchedWritableCustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{}) {
	o.ExtraChoices = v
}

// GetOrderAlphabetically returns the OrderAlphabetically field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetOrderAlphabetically() bool {
	if o == nil || IsNil(o.OrderAlphabetically) {
		var ret bool
		return ret
	}
	return *o.OrderAlphabetically
}

// GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderAlphabetically) {
		return nil, false
	}
	return o.OrderAlphabetically, true
}

// HasOrderAlphabetically returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldChoiceSetRequest) HasOrderAlphabetically() bool {
	if o != nil && !IsNil(o.OrderAlphabetically) {
		return true
	}

	return false
}

// SetOrderAlphabetically gets a reference to the given bool and assigns it to the OrderAlphabetically field.
func (o *PatchedWritableCustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool) {
	o.OrderAlphabetically = &v
}

func (o PatchedWritableCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableCustomFieldChoiceSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BaseChoices) {
		toSerialize["base_choices"] = o.BaseChoices
	}
	if !IsNil(o.ExtraChoices) {
		toSerialize["extra_choices"] = o.ExtraChoices
	}
	if !IsNil(o.OrderAlphabetically) {
		toSerialize["order_alphabetically"] = o.OrderAlphabetically
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableCustomFieldChoiceSetRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableCustomFieldChoiceSetRequest := _PatchedWritableCustomFieldChoiceSetRequest{}

	err = json.Unmarshal(data, &varPatchedWritableCustomFieldChoiceSetRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableCustomFieldChoiceSetRequest(varPatchedWritableCustomFieldChoiceSetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "base_choices")
		delete(additionalProperties, "extra_choices")
		delete(additionalProperties, "order_alphabetically")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableCustomFieldChoiceSetRequest struct {
	value *PatchedWritableCustomFieldChoiceSetRequest
	isSet bool
}

func (v NullablePatchedWritableCustomFieldChoiceSetRequest) Get() *PatchedWritableCustomFieldChoiceSetRequest {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldChoiceSetRequest) Set(val *PatchedWritableCustomFieldChoiceSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldChoiceSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldChoiceSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldChoiceSetRequest(val *PatchedWritableCustomFieldChoiceSetRequest) *NullablePatchedWritableCustomFieldChoiceSetRequest {
	return &NullablePatchedWritableCustomFieldChoiceSetRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldChoiceSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
