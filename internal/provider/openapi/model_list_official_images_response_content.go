/*
ParallelCluster

ParallelCluster API

API version: 3.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListOfficialImagesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfficialImagesResponseContent{}

// ListOfficialImagesResponseContent struct for ListOfficialImagesResponseContent
type ListOfficialImagesResponseContent struct {
	Images []AmiInfo `json:"images"`
}

type _ListOfficialImagesResponseContent ListOfficialImagesResponseContent

// NewListOfficialImagesResponseContent instantiates a new ListOfficialImagesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfficialImagesResponseContent(images []AmiInfo) *ListOfficialImagesResponseContent {
	this := ListOfficialImagesResponseContent{}
	this.Images = images
	return &this
}

// NewListOfficialImagesResponseContentWithDefaults instantiates a new ListOfficialImagesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfficialImagesResponseContentWithDefaults() *ListOfficialImagesResponseContent {
	this := ListOfficialImagesResponseContent{}
	return &this
}

// GetImages returns the Images field value
func (o *ListOfficialImagesResponseContent) GetImages() []AmiInfo {
	if o == nil {
		var ret []AmiInfo
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *ListOfficialImagesResponseContent) GetImagesOk() ([]AmiInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *ListOfficialImagesResponseContent) SetImages(v []AmiInfo) {
	o.Images = v
}

func (o ListOfficialImagesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOfficialImagesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["images"] = o.Images
	return toSerialize, nil
}

func (o *ListOfficialImagesResponseContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"images",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListOfficialImagesResponseContent := _ListOfficialImagesResponseContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOfficialImagesResponseContent)

	if err != nil {
		return err
	}

	*o = ListOfficialImagesResponseContent(varListOfficialImagesResponseContent)

	return err
}

type NullableListOfficialImagesResponseContent struct {
	value *ListOfficialImagesResponseContent
	isSet bool
}

func (v NullableListOfficialImagesResponseContent) Get() *ListOfficialImagesResponseContent {
	return v.value
}

func (v *NullableListOfficialImagesResponseContent) Set(val *ListOfficialImagesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfficialImagesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfficialImagesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfficialImagesResponseContent(val *ListOfficialImagesResponseContent) *NullableListOfficialImagesResponseContent {
	return &NullableListOfficialImagesResponseContent{value: val, isSet: true}
}

func (v NullableListOfficialImagesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOfficialImagesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


