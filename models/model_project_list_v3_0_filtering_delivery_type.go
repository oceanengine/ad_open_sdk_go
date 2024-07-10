/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringDeliveryType
type ProjectListV30FilteringDeliveryType string

// List of project_list_v3.0_filtering_delivery_type
const (
	DURATION_ProjectListV30FilteringDeliveryType ProjectListV30FilteringDeliveryType = "DURATION"
	NORMAL_ProjectListV30FilteringDeliveryType   ProjectListV30FilteringDeliveryType = "NORMAL"
)

// All allowed values of ProjectListV30FilteringDeliveryType enum
var AllowedProjectListV30FilteringDeliveryTypeEnumValues = []ProjectListV30FilteringDeliveryType{
	"DURATION",
	"NORMAL",
}

// NewProjectListV30FilteringDeliveryTypeFromValue returns a pointer to a valid ProjectListV30FilteringDeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringDeliveryTypeFromValue(v string) (*ProjectListV30FilteringDeliveryType, error) {
	ev := ProjectListV30FilteringDeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringDeliveryType: valid values are %v", v, AllowedProjectListV30FilteringDeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringDeliveryType) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringDeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_delivery_type value
func (v ProjectListV30FilteringDeliveryType) Ptr() *ProjectListV30FilteringDeliveryType {
	return &v
}
