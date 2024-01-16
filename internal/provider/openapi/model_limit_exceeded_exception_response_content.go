/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LimitExceededExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitExceededExceptionResponseContent{}

// LimitExceededExceptionResponseContent The client is sending more than the allowed number of requests per unit of time.
type LimitExceededExceptionResponseContent struct {
	Message *string `json:"message,omitempty"`
}

// NewLimitExceededExceptionResponseContent instantiates a new LimitExceededExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitExceededExceptionResponseContent() *LimitExceededExceptionResponseContent {
	this := LimitExceededExceptionResponseContent{}
	return &this
}

// NewLimitExceededExceptionResponseContentWithDefaults instantiates a new LimitExceededExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitExceededExceptionResponseContentWithDefaults() *LimitExceededExceptionResponseContent {
	this := LimitExceededExceptionResponseContent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LimitExceededExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitExceededExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LimitExceededExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LimitExceededExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o LimitExceededExceptionResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitExceededExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableLimitExceededExceptionResponseContent struct {
	value *LimitExceededExceptionResponseContent
	isSet bool
}

func (v NullableLimitExceededExceptionResponseContent) Get() *LimitExceededExceptionResponseContent {
	return v.value
}

func (v *NullableLimitExceededExceptionResponseContent) Set(val *LimitExceededExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitExceededExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitExceededExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitExceededExceptionResponseContent(val *LimitExceededExceptionResponseContent) *NullableLimitExceededExceptionResponseContent {
	return &NullableLimitExceededExceptionResponseContent{value: val, isSet: true}
}

func (v NullableLimitExceededExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitExceededExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

