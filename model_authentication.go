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

// Authentication the model 'Authentication'
type Authentication string

// List of Authentication
const (
	AUTHENTICATION_EMPTY       Authentication = ""
	AUTHENTICATION_HMAC_MD5    Authentication = "hmac-md5"
	AUTHENTICATION_HMAC_SHA1   Authentication = "hmac-sha1"
	AUTHENTICATION_HMAC_SHA256 Authentication = "hmac-sha256"
	AUTHENTICATION_HMAC_SHA384 Authentication = "hmac-sha384"
	AUTHENTICATION_HMAC_SHA512 Authentication = "hmac-sha512"
)

// All allowed values of Authentication enum
var AllowedAuthenticationEnumValues = []Authentication{
	"",
	"hmac-md5",
	"hmac-sha1",
	"hmac-sha256",
	"hmac-sha384",
	"hmac-sha512",
}

func (v *Authentication) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Authentication(value)
	for _, existing := range AllowedAuthenticationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Authentication", value)
}

// NewAuthenticationFromValue returns a pointer to a valid Authentication
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationFromValue(v string) (*Authentication, error) {
	ev := Authentication(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Authentication: valid values are %v", v, AllowedAuthenticationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Authentication) IsValid() bool {
	for _, existing := range AllowedAuthenticationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Authentication value
func (v Authentication) Ptr() *Authentication {
	return &v
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
