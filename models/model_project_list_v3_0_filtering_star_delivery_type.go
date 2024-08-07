/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringStarDeliveryType
type ProjectListV30FilteringStarDeliveryType string

// List of project_list_v3.0_filtering_star_delivery_type
const (
	NOT_STAR_DELIVERY_ProjectListV30FilteringStarDeliveryType ProjectListV30FilteringStarDeliveryType = "NOT_STAR_DELIVERY"
	STAR_DELIVERY_ProjectListV30FilteringStarDeliveryType     ProjectListV30FilteringStarDeliveryType = "STAR_DELIVERY"
)

// All allowed values of ProjectListV30FilteringStarDeliveryType enum
var AllowedProjectListV30FilteringStarDeliveryTypeEnumValues = []ProjectListV30FilteringStarDeliveryType{
	"NOT_STAR_DELIVERY",
	"STAR_DELIVERY",
}

// NewProjectListV30FilteringStarDeliveryTypeFromValue returns a pointer to a valid ProjectListV30FilteringStarDeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringStarDeliveryTypeFromValue(v string) (*ProjectListV30FilteringStarDeliveryType, error) {
	ev := ProjectListV30FilteringStarDeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringStarDeliveryType: valid values are %v", v, AllowedProjectListV30FilteringStarDeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringStarDeliveryType) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringStarDeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_star_delivery_type value
func (v ProjectListV30FilteringStarDeliveryType) Ptr() *ProjectListV30FilteringStarDeliveryType {
	return &v
}
