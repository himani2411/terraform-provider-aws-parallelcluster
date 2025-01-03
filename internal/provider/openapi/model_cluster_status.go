/*
ParallelCluster

ParallelCluster API

API version: 3.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ClusterStatus the model 'ClusterStatus'
type ClusterStatus string

// List of ClusterStatus
const (
	CLUSTERSTATUS_CREATE_IN_PROGRESS ClusterStatus = "CREATE_IN_PROGRESS"
	CLUSTERSTATUS_CREATE_FAILED ClusterStatus = "CREATE_FAILED"
	CLUSTERSTATUS_CREATE_COMPLETE ClusterStatus = "CREATE_COMPLETE"
	CLUSTERSTATUS_DELETE_IN_PROGRESS ClusterStatus = "DELETE_IN_PROGRESS"
	CLUSTERSTATUS_DELETE_FAILED ClusterStatus = "DELETE_FAILED"
	CLUSTERSTATUS_DELETE_COMPLETE ClusterStatus = "DELETE_COMPLETE"
	CLUSTERSTATUS_UPDATE_IN_PROGRESS ClusterStatus = "UPDATE_IN_PROGRESS"
	CLUSTERSTATUS_UPDATE_COMPLETE ClusterStatus = "UPDATE_COMPLETE"
	CLUSTERSTATUS_UPDATE_FAILED ClusterStatus = "UPDATE_FAILED"
)

// All allowed values of ClusterStatus enum
var AllowedClusterStatusEnumValues = []ClusterStatus{
	"CREATE_IN_PROGRESS",
	"CREATE_FAILED",
	"CREATE_COMPLETE",
	"DELETE_IN_PROGRESS",
	"DELETE_FAILED",
	"DELETE_COMPLETE",
	"UPDATE_IN_PROGRESS",
	"UPDATE_COMPLETE",
	"UPDATE_FAILED",
}

func (v *ClusterStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterStatus(value)
	for _, existing := range AllowedClusterStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterStatus", value)
}

// NewClusterStatusFromValue returns a pointer to a valid ClusterStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterStatusFromValue(v string) (*ClusterStatus, error) {
	ev := ClusterStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterStatus: valid values are %v", v, AllowedClusterStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterStatus) IsValid() bool {
	for _, existing := range AllowedClusterStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterStatus value
func (v ClusterStatus) Ptr() *ClusterStatus {
	return &v
}

type NullableClusterStatus struct {
	value *ClusterStatus
	isSet bool
}

func (v NullableClusterStatus) Get() *ClusterStatus {
	return v.value
}

func (v *NullableClusterStatus) Set(val *ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus(val *ClusterStatus) *NullableClusterStatus {
	return &NullableClusterStatus{value: val, isSet: true}
}

func (v NullableClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

