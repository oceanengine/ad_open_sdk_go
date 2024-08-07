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

// ProjectUpdateV30AudienceAwemeFanTimeScope
type ProjectUpdateV30AudienceAwemeFanTimeScope string

// List of project_update_v3.0_audience_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_ProjectUpdateV30AudienceAwemeFanTimeScope ProjectUpdateV30AudienceAwemeFanTimeScope = "FIFTEEN_DAYS"
	SIXTY_DAYS_ProjectUpdateV30AudienceAwemeFanTimeScope   ProjectUpdateV30AudienceAwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_ProjectUpdateV30AudienceAwemeFanTimeScope  ProjectUpdateV30AudienceAwemeFanTimeScope = "THIRTY_DAYS"
)

// All allowed values of ProjectUpdateV30AudienceAwemeFanTimeScope enum
var AllowedProjectUpdateV30AudienceAwemeFanTimeScopeEnumValues = []ProjectUpdateV30AudienceAwemeFanTimeScope{
	"FIFTEEN_DAYS",
	"SIXTY_DAYS",
	"THIRTY_DAYS",
}

// NewProjectUpdateV30AudienceAwemeFanTimeScopeFromValue returns a pointer to a valid ProjectUpdateV30AudienceAwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceAwemeFanTimeScopeFromValue(v string) (*ProjectUpdateV30AudienceAwemeFanTimeScope, error) {
	ev := ProjectUpdateV30AudienceAwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceAwemeFanTimeScope: valid values are %v", v, AllowedProjectUpdateV30AudienceAwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceAwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceAwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_aweme_fan_time_scope value
func (v ProjectUpdateV30AudienceAwemeFanTimeScope) Ptr() *ProjectUpdateV30AudienceAwemeFanTimeScope {
	return &v
}
