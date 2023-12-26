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

// ProjectCreateV30AudienceRegionRecommend
type ProjectCreateV30AudienceRegionRecommend string

// List of project_create_v3.0_audience_region_recommend
const (
	OFF_ProjectCreateV30AudienceRegionRecommend ProjectCreateV30AudienceRegionRecommend = "OFF"
	ON_ProjectCreateV30AudienceRegionRecommend  ProjectCreateV30AudienceRegionRecommend = "ON"
)

// All allowed values of ProjectCreateV30AudienceRegionRecommend enum
var AllowedProjectCreateV30AudienceRegionRecommendEnumValues = []ProjectCreateV30AudienceRegionRecommend{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceRegionRecommendFromValue returns a pointer to a valid ProjectCreateV30AudienceRegionRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceRegionRecommendFromValue(v string) (*ProjectCreateV30AudienceRegionRecommend, error) {
	ev := ProjectCreateV30AudienceRegionRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceRegionRecommend: valid values are %v", v, AllowedProjectCreateV30AudienceRegionRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceRegionRecommend) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceRegionRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_region_recommend value
func (v ProjectCreateV30AudienceRegionRecommend) Ptr() *ProjectCreateV30AudienceRegionRecommend {
	return &v
}
