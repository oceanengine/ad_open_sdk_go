/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceHideIfConverted
type ProjectListV30DataListAudienceHideIfConverted string

// List of project_list_v3.0_data_list_audience_hide_if_converted
const (
	ADVERTISER_ProjectListV30DataListAudienceHideIfConverted   ProjectListV30DataListAudienceHideIfConverted = "ADVERTISER"
	APP_ProjectListV30DataListAudienceHideIfConverted          ProjectListV30DataListAudienceHideIfConverted = "APP"
	CUSTOMER_ProjectListV30DataListAudienceHideIfConverted     ProjectListV30DataListAudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_ProjectListV30DataListAudienceHideIfConverted   ProjectListV30DataListAudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_ProjectListV30DataListAudienceHideIfConverted ProjectListV30DataListAudienceHideIfConverted = "ORGANIZATION"
	PROJECT_ProjectListV30DataListAudienceHideIfConverted      ProjectListV30DataListAudienceHideIfConverted = "PROJECT"
	PROMOTION_ProjectListV30DataListAudienceHideIfConverted    ProjectListV30DataListAudienceHideIfConverted = "PROMOTION"
)

// All allowed values of ProjectListV30DataListAudienceHideIfConverted enum
var AllowedProjectListV30DataListAudienceHideIfConvertedEnumValues = []ProjectListV30DataListAudienceHideIfConverted{
	"ADVERTISER",
	"APP",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
	"PROJECT",
	"PROMOTION",
}

// NewProjectListV30DataListAudienceHideIfConvertedFromValue returns a pointer to a valid ProjectListV30DataListAudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceHideIfConvertedFromValue(v string) (*ProjectListV30DataListAudienceHideIfConverted, error) {
	ev := ProjectListV30DataListAudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceHideIfConverted: valid values are %v", v, AllowedProjectListV30DataListAudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_hide_if_converted value
func (v ProjectListV30DataListAudienceHideIfConverted) Ptr() *ProjectListV30DataListAudienceHideIfConverted {
	return &v
}
