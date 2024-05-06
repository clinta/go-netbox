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
	"os"
)

// checks if the DeviceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceTypeRequest{}

// DeviceTypeRequest Adds support for custom fields and tags.
type DeviceTypeRequest struct {
	Manufacturer    NestedManufacturerRequest     `json:"manufacturer"`
	DefaultPlatform NullableNestedPlatformRequest `json:"default_platform,omitempty"`
	Model           string                        `json:"model"`
	Slug            string                        `json:"slug"`
	// Discrete part number (optional)
	PartNumber *string  `json:"part_number,omitempty"`
	UHeight    *float64 `json:"u_height,omitempty"`
	// Devices of this type are excluded when calculating rack utilization.
	ExcludeFromUtilization *bool `json:"exclude_from_utilization,omitempty"`
	// Device consumes both front and rear rack faces.
	IsFullDepth          *bool                                  `json:"is_full_depth,omitempty"`
	SubdeviceRole        NullableDeviceTypeRequestSubdeviceRole `json:"subdevice_role,omitempty"`
	Airflow              NullableDeviceTypeRequestAirflow       `json:"airflow,omitempty"`
	Weight               NullableFloat64                        `json:"weight,omitempty"`
	WeightUnit           NullableDeviceTypeRequestWeightUnit    `json:"weight_unit,omitempty"`
	FrontImage           **os.File                              `json:"front_image,omitempty"`
	RearImage            **os.File                              `json:"rear_image,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	Comments             *string                                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                     `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                 `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTypeRequest DeviceTypeRequest

// NewDeviceTypeRequest instantiates a new DeviceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTypeRequest(manufacturer NestedManufacturerRequest, model string, slug string) *DeviceTypeRequest {
	this := DeviceTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	var uHeight float64 = 1.0
	this.UHeight = &uHeight
	return &this
}

// NewDeviceTypeRequestWithDefaults instantiates a new DeviceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTypeRequestWithDefaults() *DeviceTypeRequest {
	this := DeviceTypeRequest{}
	var uHeight float64 = 1.0
	this.UHeight = &uHeight
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *DeviceTypeRequest) GetManufacturer() NestedManufacturerRequest {
	if o == nil {
		var ret NestedManufacturerRequest
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetManufacturerOk() (*NestedManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *DeviceTypeRequest) SetManufacturer(v NestedManufacturerRequest) {
	o.Manufacturer = v
}

// GetDefaultPlatform returns the DefaultPlatform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTypeRequest) GetDefaultPlatform() NestedPlatformRequest {
	if o == nil || IsNil(o.DefaultPlatform.Get()) {
		var ret NestedPlatformRequest
		return ret
	}
	return *o.DefaultPlatform.Get()
}

// GetDefaultPlatformOk returns a tuple with the DefaultPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTypeRequest) GetDefaultPlatformOk() (*NestedPlatformRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPlatform.Get(), o.DefaultPlatform.IsSet()
}

// HasDefaultPlatform returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasDefaultPlatform() bool {
	if o != nil && o.DefaultPlatform.IsSet() {
		return true
	}

	return false
}

// SetDefaultPlatform gets a reference to the given NullableNestedPlatformRequest and assigns it to the DefaultPlatform field.
func (o *DeviceTypeRequest) SetDefaultPlatform(v NestedPlatformRequest) {
	o.DefaultPlatform.Set(&v)
}

// SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil
func (o *DeviceTypeRequest) SetDefaultPlatformNil() {
	o.DefaultPlatform.Set(nil)
}

// UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
func (o *DeviceTypeRequest) UnsetDefaultPlatform() {
	o.DefaultPlatform.Unset()
}

// GetModel returns the Model field value
func (o *DeviceTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeviceTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *DeviceTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DeviceTypeRequest) SetSlug(v string) {
	o.Slug = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *DeviceTypeRequest) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetUHeight() float64 {
	if o == nil || IsNil(o.UHeight) {
		var ret float64
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetUHeightOk() (*float64, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given float64 and assigns it to the UHeight field.
func (o *DeviceTypeRequest) SetUHeight(v float64) {
	o.UHeight = &v
}

// GetExcludeFromUtilization returns the ExcludeFromUtilization field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetExcludeFromUtilization() bool {
	if o == nil || IsNil(o.ExcludeFromUtilization) {
		var ret bool
		return ret
	}
	return *o.ExcludeFromUtilization
}

// GetExcludeFromUtilizationOk returns a tuple with the ExcludeFromUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetExcludeFromUtilizationOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeFromUtilization) {
		return nil, false
	}
	return o.ExcludeFromUtilization, true
}

// HasExcludeFromUtilization returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasExcludeFromUtilization() bool {
	if o != nil && !IsNil(o.ExcludeFromUtilization) {
		return true
	}

	return false
}

// SetExcludeFromUtilization gets a reference to the given bool and assigns it to the ExcludeFromUtilization field.
func (o *DeviceTypeRequest) SetExcludeFromUtilization(v bool) {
	o.ExcludeFromUtilization = &v
}

// GetIsFullDepth returns the IsFullDepth field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetIsFullDepth() bool {
	if o == nil || IsNil(o.IsFullDepth) {
		var ret bool
		return ret
	}
	return *o.IsFullDepth
}

// GetIsFullDepthOk returns a tuple with the IsFullDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetIsFullDepthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullDepth) {
		return nil, false
	}
	return o.IsFullDepth, true
}

// HasIsFullDepth returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasIsFullDepth() bool {
	if o != nil && !IsNil(o.IsFullDepth) {
		return true
	}

	return false
}

// SetIsFullDepth gets a reference to the given bool and assigns it to the IsFullDepth field.
func (o *DeviceTypeRequest) SetIsFullDepth(v bool) {
	o.IsFullDepth = &v
}

// GetSubdeviceRole returns the SubdeviceRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTypeRequest) GetSubdeviceRole() DeviceTypeRequestSubdeviceRole {
	if o == nil || IsNil(o.SubdeviceRole.Get()) {
		var ret DeviceTypeRequestSubdeviceRole
		return ret
	}
	return *o.SubdeviceRole.Get()
}

// GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTypeRequest) GetSubdeviceRoleOk() (*DeviceTypeRequestSubdeviceRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubdeviceRole.Get(), o.SubdeviceRole.IsSet()
}

// HasSubdeviceRole returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasSubdeviceRole() bool {
	if o != nil && o.SubdeviceRole.IsSet() {
		return true
	}

	return false
}

// SetSubdeviceRole gets a reference to the given NullableDeviceTypeRequestSubdeviceRole and assigns it to the SubdeviceRole field.
func (o *DeviceTypeRequest) SetSubdeviceRole(v DeviceTypeRequestSubdeviceRole) {
	o.SubdeviceRole.Set(&v)
}

// SetSubdeviceRoleNil sets the value for SubdeviceRole to be an explicit nil
func (o *DeviceTypeRequest) SetSubdeviceRoleNil() {
	o.SubdeviceRole.Set(nil)
}

// UnsetSubdeviceRole ensures that no value is present for SubdeviceRole, not even an explicit nil
func (o *DeviceTypeRequest) UnsetSubdeviceRole() {
	o.SubdeviceRole.Unset()
}

// GetAirflow returns the Airflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTypeRequest) GetAirflow() DeviceTypeRequestAirflow {
	if o == nil || IsNil(o.Airflow.Get()) {
		var ret DeviceTypeRequestAirflow
		return ret
	}
	return *o.Airflow.Get()
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTypeRequest) GetAirflowOk() (*DeviceTypeRequestAirflow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Airflow.Get(), o.Airflow.IsSet()
}

// HasAirflow returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasAirflow() bool {
	if o != nil && o.Airflow.IsSet() {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given NullableDeviceTypeRequestAirflow and assigns it to the Airflow field.
func (o *DeviceTypeRequest) SetAirflow(v DeviceTypeRequestAirflow) {
	o.Airflow.Set(&v)
}

// SetAirflowNil sets the value for Airflow to be an explicit nil
func (o *DeviceTypeRequest) SetAirflowNil() {
	o.Airflow.Set(nil)
}

// UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
func (o *DeviceTypeRequest) UnsetAirflow() {
	o.Airflow.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTypeRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTypeRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *DeviceTypeRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *DeviceTypeRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *DeviceTypeRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTypeRequest) GetWeightUnit() DeviceTypeRequestWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeRequestWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTypeRequest) GetWeightUnitOk() (*DeviceTypeRequestWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeRequestWeightUnit and assigns it to the WeightUnit field.
func (o *DeviceTypeRequest) SetWeightUnit(v DeviceTypeRequestWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *DeviceTypeRequest) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *DeviceTypeRequest) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetFrontImage returns the FrontImage field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetFrontImage() *os.File {
	if o == nil || IsNil(o.FrontImage) {
		var ret *os.File
		return ret
	}
	return *o.FrontImage
}

// GetFrontImageOk returns a tuple with the FrontImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetFrontImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.FrontImage) {
		return nil, false
	}
	return o.FrontImage, true
}

// HasFrontImage returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasFrontImage() bool {
	if o != nil && !IsNil(o.FrontImage) {
		return true
	}

	return false
}

// SetFrontImage gets a reference to the given *os.File and assigns it to the FrontImage field.
func (o *DeviceTypeRequest) SetFrontImage(v *os.File) {
	o.FrontImage = &v
}

// GetRearImage returns the RearImage field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetRearImage() *os.File {
	if o == nil || IsNil(o.RearImage) {
		var ret *os.File
		return ret
	}
	return *o.RearImage
}

// GetRearImageOk returns a tuple with the RearImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetRearImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.RearImage) {
		return nil, false
	}
	return o.RearImage, true
}

// HasRearImage returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasRearImage() bool {
	if o != nil && !IsNil(o.RearImage) {
		return true
	}

	return false
}

// SetRearImage gets a reference to the given *os.File and assigns it to the RearImage field.
func (o *DeviceTypeRequest) SetRearImage(v *os.File) {
	o.RearImage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *DeviceTypeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *DeviceTypeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o DeviceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	if o.DefaultPlatform.IsSet() {
		toSerialize["default_platform"] = o.DefaultPlatform.Get()
	}
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.ExcludeFromUtilization) {
		toSerialize["exclude_from_utilization"] = o.ExcludeFromUtilization
	}
	if !IsNil(o.IsFullDepth) {
		toSerialize["is_full_depth"] = o.IsFullDepth
	}
	if o.SubdeviceRole.IsSet() {
		toSerialize["subdevice_role"] = o.SubdeviceRole.Get()
	}
	if o.Airflow.IsSet() {
		toSerialize["airflow"] = o.Airflow.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if !IsNil(o.FrontImage) {
		toSerialize["front_image"] = o.FrontImage
	}
	if !IsNil(o.RearImage) {
		toSerialize["rear_image"] = o.RearImage
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *DeviceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"manufacturer",
		"model",
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

	varDeviceTypeRequest := _DeviceTypeRequest{}

	err = json.Unmarshal(data, &varDeviceTypeRequest)

	if err != nil {
		return err
	}

	*o = DeviceTypeRequest(varDeviceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "default_platform")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "part_number")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "exclude_from_utilization")
		delete(additionalProperties, "is_full_depth")
		delete(additionalProperties, "subdevice_role")
		delete(additionalProperties, "airflow")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "front_image")
		delete(additionalProperties, "rear_image")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTypeRequest struct {
	value *DeviceTypeRequest
	isSet bool
}

func (v NullableDeviceTypeRequest) Get() *DeviceTypeRequest {
	return v.value
}

func (v *NullableDeviceTypeRequest) Set(val *DeviceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeRequest(val *DeviceTypeRequest) *NullableDeviceTypeRequest {
	return &NullableDeviceTypeRequest{value: val, isSet: true}
}

func (v NullableDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
