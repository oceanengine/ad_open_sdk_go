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

// ProjectUpdateV30AudienceDistrict
type ProjectUpdateV30AudienceDistrict string

// List of project_update_v3.0_audience_district
const (
	BUSINESS_DISTRICT_ProjectUpdateV30AudienceDistrict ProjectUpdateV30AudienceDistrict = "BUSINESS_DISTRICT"
	NONE_ProjectUpdateV30AudienceDistrict              ProjectUpdateV30AudienceDistrict = "NONE"
	OVERSEA_ProjectUpdateV30AudienceDistrict           ProjectUpdateV30AudienceDistrict = "OVERSEA"
	REGION_ProjectUpdateV30AudienceDistrict            ProjectUpdateV30AudienceDistrict = "REGION"
)

// All allowed values of ProjectUpdateV30AudienceDistrict enum
var AllowedProjectUpdateV30AudienceDistrictEnumValues = []ProjectUpdateV30AudienceDistrict{
	"BUSINESS_DISTRICT",
	"NONE",
	"OVERSEA",
	"REGION",
}

// NewProjectUpdateV30AudienceDistrictFromValue returns a pointer to a valid ProjectUpdateV30AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceDistrictFromValue(v string) (*ProjectUpdateV30AudienceDistrict, error) {
	ev := ProjectUpdateV30AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceDistrict: valid values are %v", v, AllowedProjectUpdateV30AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_district value
func (v ProjectUpdateV30AudienceDistrict) Ptr() *ProjectUpdateV30AudienceDistrict {
	return &v
}
