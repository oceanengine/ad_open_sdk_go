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

// ProjectListV30DataListAdType
type ProjectListV30DataListAdType string

// List of project_list_v3.0_data_list_ad_type
const (
	ALL_ProjectListV30DataListAdType    ProjectListV30DataListAdType = "ALL"
	SEARCH_ProjectListV30DataListAdType ProjectListV30DataListAdType = "SEARCH"
)

// All allowed values of ProjectListV30DataListAdType enum
var AllowedProjectListV30DataListAdTypeEnumValues = []ProjectListV30DataListAdType{
	"ALL",
	"SEARCH",
}

// NewProjectListV30DataListAdTypeFromValue returns a pointer to a valid ProjectListV30DataListAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAdTypeFromValue(v string) (*ProjectListV30DataListAdType, error) {
	ev := ProjectListV30DataListAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAdType: valid values are %v", v, AllowedProjectListV30DataListAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAdType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_ad_type value
func (v ProjectListV30DataListAdType) Ptr() *ProjectListV30DataListAdType {
	return &v
}
