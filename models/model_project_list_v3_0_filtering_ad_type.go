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

// ProjectListV30FilteringAdType
type ProjectListV30FilteringAdType string

// List of project_list_v3.0_filtering_ad_type
const (
	ALL_ProjectListV30FilteringAdType    ProjectListV30FilteringAdType = "ALL"
	SEARCH_ProjectListV30FilteringAdType ProjectListV30FilteringAdType = "SEARCH"
)

// All allowed values of ProjectListV30FilteringAdType enum
var AllowedProjectListV30FilteringAdTypeEnumValues = []ProjectListV30FilteringAdType{
	"ALL",
	"SEARCH",
}

// NewProjectListV30FilteringAdTypeFromValue returns a pointer to a valid ProjectListV30FilteringAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringAdTypeFromValue(v string) (*ProjectListV30FilteringAdType, error) {
	ev := ProjectListV30FilteringAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringAdType: valid values are %v", v, AllowedProjectListV30FilteringAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringAdType) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_ad_type value
func (v ProjectListV30FilteringAdType) Ptr() *ProjectListV30FilteringAdType {
	return &v
}
