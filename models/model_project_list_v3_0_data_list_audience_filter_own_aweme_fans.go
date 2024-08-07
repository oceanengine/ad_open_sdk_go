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

// ProjectListV30DataListAudienceFilterOwnAwemeFans
type ProjectListV30DataListAudienceFilterOwnAwemeFans string

// List of project_list_v3.0_data_list_audience_filter_own_aweme_fans
const (
	OFF_ProjectListV30DataListAudienceFilterOwnAwemeFans ProjectListV30DataListAudienceFilterOwnAwemeFans = "OFF"
	ON_ProjectListV30DataListAudienceFilterOwnAwemeFans  ProjectListV30DataListAudienceFilterOwnAwemeFans = "ON"
)

// All allowed values of ProjectListV30DataListAudienceFilterOwnAwemeFans enum
var AllowedProjectListV30DataListAudienceFilterOwnAwemeFansEnumValues = []ProjectListV30DataListAudienceFilterOwnAwemeFans{
	"OFF",
	"ON",
}

// NewProjectListV30DataListAudienceFilterOwnAwemeFansFromValue returns a pointer to a valid ProjectListV30DataListAudienceFilterOwnAwemeFans
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceFilterOwnAwemeFansFromValue(v string) (*ProjectListV30DataListAudienceFilterOwnAwemeFans, error) {
	ev := ProjectListV30DataListAudienceFilterOwnAwemeFans(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceFilterOwnAwemeFans: valid values are %v", v, AllowedProjectListV30DataListAudienceFilterOwnAwemeFansEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceFilterOwnAwemeFans) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceFilterOwnAwemeFansEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_filter_own_aweme_fans value
func (v ProjectListV30DataListAudienceFilterOwnAwemeFans) Ptr() *ProjectListV30DataListAudienceFilterOwnAwemeFans {
	return &v
}
