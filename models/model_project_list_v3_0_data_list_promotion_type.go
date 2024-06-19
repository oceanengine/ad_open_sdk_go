/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListPromotionType
type ProjectListV30DataListPromotionType string

// List of project_list_v3.0_data_list_promotion_type
const (
	AWEME_HOME_PAGE_ProjectListV30DataListPromotionType   ProjectListV30DataListPromotionType = "AWEME_HOME_PAGE"
	LANDING_PAGE_LINK_ProjectListV30DataListPromotionType ProjectListV30DataListPromotionType = "LANDING_PAGE_LINK"
)

// All allowed values of ProjectListV30DataListPromotionType enum
var AllowedProjectListV30DataListPromotionTypeEnumValues = []ProjectListV30DataListPromotionType{
	"AWEME_HOME_PAGE",
	"LANDING_PAGE_LINK",
}

// NewProjectListV30DataListPromotionTypeFromValue returns a pointer to a valid ProjectListV30DataListPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListPromotionTypeFromValue(v string) (*ProjectListV30DataListPromotionType, error) {
	ev := ProjectListV30DataListPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListPromotionType: valid values are %v", v, AllowedProjectListV30DataListPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListPromotionType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_promotion_type value
func (v ProjectListV30DataListPromotionType) Ptr() *ProjectListV30DataListPromotionType {
	return &v
}
