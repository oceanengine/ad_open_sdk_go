/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceHideIfConverted
type ProjectCreateV30AudienceHideIfConverted string

// List of project_create_v3.0_audience_hide_if_converted
const (
	ADVERTISER_ProjectCreateV30AudienceHideIfConverted   ProjectCreateV30AudienceHideIfConverted = "ADVERTISER"
	APP_ProjectCreateV30AudienceHideIfConverted          ProjectCreateV30AudienceHideIfConverted = "APP"
	CUSTOMER_ProjectCreateV30AudienceHideIfConverted     ProjectCreateV30AudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_ProjectCreateV30AudienceHideIfConverted   ProjectCreateV30AudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_ProjectCreateV30AudienceHideIfConverted ProjectCreateV30AudienceHideIfConverted = "ORGANIZATION"
	PROJECT_ProjectCreateV30AudienceHideIfConverted      ProjectCreateV30AudienceHideIfConverted = "PROJECT"
	PROMOTION_ProjectCreateV30AudienceHideIfConverted    ProjectCreateV30AudienceHideIfConverted = "PROMOTION"
)

// All allowed values of ProjectCreateV30AudienceHideIfConverted enum
var AllowedProjectCreateV30AudienceHideIfConvertedEnumValues = []ProjectCreateV30AudienceHideIfConverted{
	"ADVERTISER",
	"APP",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
	"PROJECT",
	"PROMOTION",
}

// NewProjectCreateV30AudienceHideIfConvertedFromValue returns a pointer to a valid ProjectCreateV30AudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceHideIfConvertedFromValue(v string) (*ProjectCreateV30AudienceHideIfConverted, error) {
	ev := ProjectCreateV30AudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceHideIfConverted: valid values are %v", v, AllowedProjectCreateV30AudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_hide_if_converted value
func (v ProjectCreateV30AudienceHideIfConverted) Ptr() *ProjectCreateV30AudienceHideIfConverted {
	return &v
}
