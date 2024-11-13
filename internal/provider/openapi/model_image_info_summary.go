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

// checks if the ImageInfoSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageInfoSummary{}

// ImageInfoSummary struct for ImageInfoSummary
type ImageInfoSummary struct {
	// Id of the image.
	ImageId string `json:"imageId" validate:"regexp=^[a-zA-Z][a-zA-Z0-9-]+$"`
	Ec2AmiInfo *Ec2AmiInfoSummary `json:"ec2AmiInfo,omitempty"`
	// AWS region where the image is built.
	Region string `json:"region"`
	// ParallelCluster version used to build the image.
	Version string `json:"version"`
	// ARN of the main CloudFormation stack.
	CloudformationStackArn *string `json:"cloudformationStackArn,omitempty"`
	ImageBuildStatus ImageBuildStatus `json:"imageBuildStatus"`
	CloudformationStackStatus *CloudFormationStackStatus `json:"cloudformationStackStatus,omitempty"`
}

type _ImageInfoSummary ImageInfoSummary

// NewImageInfoSummary instantiates a new ImageInfoSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageInfoSummary(imageId string, region string, version string, imageBuildStatus ImageBuildStatus) *ImageInfoSummary {
	this := ImageInfoSummary{}
	this.ImageId = imageId
	this.Region = region
	this.Version = version
	this.ImageBuildStatus = imageBuildStatus
	return &this
}

// NewImageInfoSummaryWithDefaults instantiates a new ImageInfoSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageInfoSummaryWithDefaults() *ImageInfoSummary {
	this := ImageInfoSummary{}
	return &this
}

// GetImageId returns the ImageId field value
func (o *ImageInfoSummary) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *ImageInfoSummary) SetImageId(v string) {
	o.ImageId = v
}

// GetEc2AmiInfo returns the Ec2AmiInfo field value if set, zero value otherwise.
func (o *ImageInfoSummary) GetEc2AmiInfo() Ec2AmiInfoSummary {
	if o == nil || IsNil(o.Ec2AmiInfo) {
		var ret Ec2AmiInfoSummary
		return ret
	}
	return *o.Ec2AmiInfo
}

// GetEc2AmiInfoOk returns a tuple with the Ec2AmiInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetEc2AmiInfoOk() (*Ec2AmiInfoSummary, bool) {
	if o == nil || IsNil(o.Ec2AmiInfo) {
		return nil, false
	}
	return o.Ec2AmiInfo, true
}

// HasEc2AmiInfo returns a boolean if a field has been set.
func (o *ImageInfoSummary) HasEc2AmiInfo() bool {
	if o != nil && !IsNil(o.Ec2AmiInfo) {
		return true
	}

	return false
}

// SetEc2AmiInfo gets a reference to the given Ec2AmiInfoSummary and assigns it to the Ec2AmiInfo field.
func (o *ImageInfoSummary) SetEc2AmiInfo(v Ec2AmiInfoSummary) {
	o.Ec2AmiInfo = &v
}

// GetRegion returns the Region field value
func (o *ImageInfoSummary) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ImageInfoSummary) SetRegion(v string) {
	o.Region = v
}

// GetVersion returns the Version field value
func (o *ImageInfoSummary) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ImageInfoSummary) SetVersion(v string) {
	o.Version = v
}

// GetCloudformationStackArn returns the CloudformationStackArn field value if set, zero value otherwise.
func (o *ImageInfoSummary) GetCloudformationStackArn() string {
	if o == nil || IsNil(o.CloudformationStackArn) {
		var ret string
		return ret
	}
	return *o.CloudformationStackArn
}

// GetCloudformationStackArnOk returns a tuple with the CloudformationStackArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetCloudformationStackArnOk() (*string, bool) {
	if o == nil || IsNil(o.CloudformationStackArn) {
		return nil, false
	}
	return o.CloudformationStackArn, true
}

// HasCloudformationStackArn returns a boolean if a field has been set.
func (o *ImageInfoSummary) HasCloudformationStackArn() bool {
	if o != nil && !IsNil(o.CloudformationStackArn) {
		return true
	}

	return false
}

// SetCloudformationStackArn gets a reference to the given string and assigns it to the CloudformationStackArn field.
func (o *ImageInfoSummary) SetCloudformationStackArn(v string) {
	o.CloudformationStackArn = &v
}

// GetImageBuildStatus returns the ImageBuildStatus field value
func (o *ImageInfoSummary) GetImageBuildStatus() ImageBuildStatus {
	if o == nil {
		var ret ImageBuildStatus
		return ret
	}

	return o.ImageBuildStatus
}

// GetImageBuildStatusOk returns a tuple with the ImageBuildStatus field value
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetImageBuildStatusOk() (*ImageBuildStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageBuildStatus, true
}

// SetImageBuildStatus sets field value
func (o *ImageInfoSummary) SetImageBuildStatus(v ImageBuildStatus) {
	o.ImageBuildStatus = v
}

// GetCloudformationStackStatus returns the CloudformationStackStatus field value if set, zero value otherwise.
func (o *ImageInfoSummary) GetCloudformationStackStatus() CloudFormationStackStatus {
	if o == nil || IsNil(o.CloudformationStackStatus) {
		var ret CloudFormationStackStatus
		return ret
	}
	return *o.CloudformationStackStatus
}

// GetCloudformationStackStatusOk returns a tuple with the CloudformationStackStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageInfoSummary) GetCloudformationStackStatusOk() (*CloudFormationStackStatus, bool) {
	if o == nil || IsNil(o.CloudformationStackStatus) {
		return nil, false
	}
	return o.CloudformationStackStatus, true
}

// HasCloudformationStackStatus returns a boolean if a field has been set.
func (o *ImageInfoSummary) HasCloudformationStackStatus() bool {
	if o != nil && !IsNil(o.CloudformationStackStatus) {
		return true
	}

	return false
}

// SetCloudformationStackStatus gets a reference to the given CloudFormationStackStatus and assigns it to the CloudformationStackStatus field.
func (o *ImageInfoSummary) SetCloudformationStackStatus(v CloudFormationStackStatus) {
	o.CloudformationStackStatus = &v
}

func (o ImageInfoSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageInfoSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imageId"] = o.ImageId
	if !IsNil(o.Ec2AmiInfo) {
		toSerialize["ec2AmiInfo"] = o.Ec2AmiInfo
	}
	toSerialize["region"] = o.Region
	toSerialize["version"] = o.Version
	if !IsNil(o.CloudformationStackArn) {
		toSerialize["cloudformationStackArn"] = o.CloudformationStackArn
	}
	toSerialize["imageBuildStatus"] = o.ImageBuildStatus
	if !IsNil(o.CloudformationStackStatus) {
		toSerialize["cloudformationStackStatus"] = o.CloudformationStackStatus
	}
	return toSerialize, nil
}

func (o *ImageInfoSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageId",
		"region",
		"version",
		"imageBuildStatus",
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

	varImageInfoSummary := _ImageInfoSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageInfoSummary)

	if err != nil {
		return err
	}

	*o = ImageInfoSummary(varImageInfoSummary)

	return err
}

type NullableImageInfoSummary struct {
	value *ImageInfoSummary
	isSet bool
}

func (v NullableImageInfoSummary) Get() *ImageInfoSummary {
	return v.value
}

func (v *NullableImageInfoSummary) Set(val *ImageInfoSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableImageInfoSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableImageInfoSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageInfoSummary(val *ImageInfoSummary) *NullableImageInfoSummary {
	return &NullableImageInfoSummary{value: val, isSet: true}
}

func (v NullableImageInfoSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageInfoSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


