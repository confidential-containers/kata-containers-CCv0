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

// VmResize struct for VmResize
type VmResize struct {
	DesiredVcpus *int32 `json:"desired_vcpus,omitempty"`
	// desired memory ram in bytes
	DesiredRam *int64 `json:"desired_ram,omitempty"`
	// desired balloon size in bytes
	DesiredBalloon *int64 `json:"desired_balloon,omitempty"`
}

// NewVmResize instantiates a new VmResize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmResize() *VmResize {
	this := VmResize{}
	return &this
}

// NewVmResizeWithDefaults instantiates a new VmResize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmResizeWithDefaults() *VmResize {
	this := VmResize{}
	return &this
}

// GetDesiredVcpus returns the DesiredVcpus field value if set, zero value otherwise.
func (o *VmResize) GetDesiredVcpus() int32 {
	if o == nil || o.DesiredVcpus == nil {
		var ret int32
		return ret
	}
	return *o.DesiredVcpus
}

// GetDesiredVcpusOk returns a tuple with the DesiredVcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmResize) GetDesiredVcpusOk() (*int32, bool) {
	if o == nil || o.DesiredVcpus == nil {
		return nil, false
	}
	return o.DesiredVcpus, true
}

// HasDesiredVcpus returns a boolean if a field has been set.
func (o *VmResize) HasDesiredVcpus() bool {
	if o != nil && o.DesiredVcpus != nil {
		return true
	}

	return false
}

// SetDesiredVcpus gets a reference to the given int32 and assigns it to the DesiredVcpus field.
func (o *VmResize) SetDesiredVcpus(v int32) {
	o.DesiredVcpus = &v
}

// GetDesiredRam returns the DesiredRam field value if set, zero value otherwise.
func (o *VmResize) GetDesiredRam() int64 {
	if o == nil || o.DesiredRam == nil {
		var ret int64
		return ret
	}
	return *o.DesiredRam
}

// GetDesiredRamOk returns a tuple with the DesiredRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmResize) GetDesiredRamOk() (*int64, bool) {
	if o == nil || o.DesiredRam == nil {
		return nil, false
	}
	return o.DesiredRam, true
}

// HasDesiredRam returns a boolean if a field has been set.
func (o *VmResize) HasDesiredRam() bool {
	if o != nil && o.DesiredRam != nil {
		return true
	}

	return false
}

// SetDesiredRam gets a reference to the given int64 and assigns it to the DesiredRam field.
func (o *VmResize) SetDesiredRam(v int64) {
	o.DesiredRam = &v
}

// GetDesiredBalloon returns the DesiredBalloon field value if set, zero value otherwise.
func (o *VmResize) GetDesiredBalloon() int64 {
	if o == nil || o.DesiredBalloon == nil {
		var ret int64
		return ret
	}
	return *o.DesiredBalloon
}

// GetDesiredBalloonOk returns a tuple with the DesiredBalloon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmResize) GetDesiredBalloonOk() (*int64, bool) {
	if o == nil || o.DesiredBalloon == nil {
		return nil, false
	}
	return o.DesiredBalloon, true
}

// HasDesiredBalloon returns a boolean if a field has been set.
func (o *VmResize) HasDesiredBalloon() bool {
	if o != nil && o.DesiredBalloon != nil {
		return true
	}

	return false
}

// SetDesiredBalloon gets a reference to the given int64 and assigns it to the DesiredBalloon field.
func (o *VmResize) SetDesiredBalloon(v int64) {
	o.DesiredBalloon = &v
}

func (o VmResize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DesiredVcpus != nil {
		toSerialize["desired_vcpus"] = o.DesiredVcpus
	}
	if o.DesiredRam != nil {
		toSerialize["desired_ram"] = o.DesiredRam
	}
	if o.DesiredBalloon != nil {
		toSerialize["desired_balloon"] = o.DesiredBalloon
	}
	return json.Marshal(toSerialize)
}

type NullableVmResize struct {
	value *VmResize
	isSet bool
}

func (v NullableVmResize) Get() *VmResize {
	return v.value
}

func (v *NullableVmResize) Set(val *VmResize) {
	v.value = val
	v.isSet = true
}

func (v NullableVmResize) IsSet() bool {
	return v.isSet
}

func (v *NullableVmResize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmResize(val *VmResize) *NullableVmResize {
	return &NullableVmResize{value: val, isSet: true}
}

func (v NullableVmResize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmResize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
