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

// ProjectCreateV30LandingPageStayTime
type ProjectCreateV30LandingPageStayTime int64

// List of project_create_v3.0_landing_page_stay_time
const (
	Enum_1000_ProjectCreateV30LandingPageStayTime  ProjectCreateV30LandingPageStayTime = 1000
	Enum_12000_ProjectCreateV30LandingPageStayTime ProjectCreateV30LandingPageStayTime = 12000
	Enum_20000_ProjectCreateV30LandingPageStayTime ProjectCreateV30LandingPageStayTime = 20000
	Enum_30000_ProjectCreateV30LandingPageStayTime ProjectCreateV30LandingPageStayTime = 30000
	Enum_40000_ProjectCreateV30LandingPageStayTime ProjectCreateV30LandingPageStayTime = 40000
	Enum_5000_ProjectCreateV30LandingPageStayTime  ProjectCreateV30LandingPageStayTime = 5000
	Enum_50000_ProjectCreateV30LandingPageStayTime ProjectCreateV30LandingPageStayTime = 50000
	Enum_60000_ProjectCreateV30LandingPageStayTime ProjectCreateV30LandingPageStayTime = 60000
)

// All allowed values of ProjectCreateV30LandingPageStayTime enum
var AllowedProjectCreateV30LandingPageStayTimeEnumValues = []ProjectCreateV30LandingPageStayTime{
	1000,
	12000,
	20000,
	30000,
	40000,
	5000,
	50000,
	60000,
}

// NewProjectCreateV30LandingPageStayTimeFromValue returns a pointer to a valid ProjectCreateV30LandingPageStayTime
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30LandingPageStayTimeFromValue(v int64) (*ProjectCreateV30LandingPageStayTime, error) {
	ev := ProjectCreateV30LandingPageStayTime(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30LandingPageStayTime: valid values are %v", v, AllowedProjectCreateV30LandingPageStayTimeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30LandingPageStayTime) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30LandingPageStayTimeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_landing_page_stay_time value
func (v ProjectCreateV30LandingPageStayTime) Ptr() *ProjectCreateV30LandingPageStayTime {
	return &v
}
