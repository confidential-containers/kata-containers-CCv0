/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// KernelConfig struct for KernelConfig
type KernelConfig struct {
	Path string `json:"path"`
}

// NewKernelConfig instantiates a new KernelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKernelConfig(path string) *KernelConfig {
	this := KernelConfig{}
	this.Path = path
	return &this
}

// NewKernelConfigWithDefaults instantiates a new KernelConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKernelConfigWithDefaults() *KernelConfig {
	this := KernelConfig{}
	return &this
}

// GetPath returns the Path field value
func (o *KernelConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *KernelConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *KernelConfig) SetPath(v string) {
	o.Path = v
}

func (o KernelConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableKernelConfig struct {
	value *KernelConfig
	isSet bool
}

func (v NullableKernelConfig) Get() *KernelConfig {
	return v.value
}

func (v *NullableKernelConfig) Set(val *KernelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKernelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKernelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKernelConfig(val *KernelConfig) *NullableKernelConfig {
	return &NullableKernelConfig{value: val, isSet: true}
}

func (v NullableKernelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKernelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
