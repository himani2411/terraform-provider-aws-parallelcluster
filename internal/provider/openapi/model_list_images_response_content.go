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

// checks if the ListImagesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImagesResponseContent{}

// ListImagesResponseContent struct for ListImagesResponseContent
type ListImagesResponseContent struct {
	// Token to use for paginated requests.
	NextToken *string `json:"nextToken,omitempty"`
	Images []ImageInfoSummary `json:"images"`
}

type _ListImagesResponseContent ListImagesResponseContent

// NewListImagesResponseContent instantiates a new ListImagesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImagesResponseContent(images []ImageInfoSummary) *ListImagesResponseContent {
	this := ListImagesResponseContent{}
	this.Images = images
	return &this
}

// NewListImagesResponseContentWithDefaults instantiates a new ListImagesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImagesResponseContentWithDefaults() *ListImagesResponseContent {
	this := ListImagesResponseContent{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListImagesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImagesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListImagesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListImagesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetImages returns the Images field value
func (o *ListImagesResponseContent) GetImages() []ImageInfoSummary {
	if o == nil {
		var ret []ImageInfoSummary
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *ListImagesResponseContent) GetImagesOk() ([]ImageInfoSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *ListImagesResponseContent) SetImages(v []ImageInfoSummary) {
	o.Images = v
}

func (o ListImagesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListImagesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["images"] = o.Images
	return toSerialize, nil
}

func (o *ListImagesResponseContent) UnmarshalJSON(data []byte) (err error) {
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

	varListImagesResponseContent := _ListImagesResponseContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListImagesResponseContent)

	if err != nil {
		return err
	}

	*o = ListImagesResponseContent(varListImagesResponseContent)

	return err
}

type NullableListImagesResponseContent struct {
	value *ListImagesResponseContent
	isSet bool
}

func (v NullableListImagesResponseContent) Get() *ListImagesResponseContent {
	return v.value
}

func (v *NullableListImagesResponseContent) Set(val *ListImagesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListImagesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListImagesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImagesResponseContent(val *ListImagesResponseContent) *NullableListImagesResponseContent {
	return &NullableListImagesResponseContent{value: val, isSet: true}
}

func (v NullableListImagesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListImagesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


