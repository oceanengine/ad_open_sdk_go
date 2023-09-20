/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceHideIfExists
type ProjectListV30DataListAudienceHideIfExists string

// List of project_list_v3.0_data_list_audience_hide_if_exists
const (
	FILTER_ProjectListV30DataListAudienceHideIfExists    ProjectListV30DataListAudienceHideIfExists = "FILTER"
	TARGETING_ProjectListV30DataListAudienceHideIfExists ProjectListV30DataListAudienceHideIfExists = "TARGETING"
	UNLIMITED_ProjectListV30DataListAudienceHideIfExists ProjectListV30DataListAudienceHideIfExists = "UNLIMITED"
)

// All allowed values of ProjectListV30DataListAudienceHideIfExists enum
var AllowedProjectListV30DataListAudienceHideIfExistsEnumValues = []ProjectListV30DataListAudienceHideIfExists{
	"FILTER",
	"TARGETING",
	"UNLIMITED",
}

// NewProjectListV30DataListAudienceHideIfExistsFromValue returns a pointer to a valid ProjectListV30DataListAudienceHideIfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceHideIfExistsFromValue(v string) (*ProjectListV30DataListAudienceHideIfExists, error) {
	ev := ProjectListV30DataListAudienceHideIfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceHideIfExists: valid values are %v", v, AllowedProjectListV30DataListAudienceHideIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceHideIfExists) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceHideIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_hide_if_exists value
func (v ProjectListV30DataListAudienceHideIfExists) Ptr() *ProjectListV30DataListAudienceHideIfExists {
	return &v
}
