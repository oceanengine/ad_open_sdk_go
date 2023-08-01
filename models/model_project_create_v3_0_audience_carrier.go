/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceCarrier
type ProjectCreateV30AudienceCarrier string

// List of project_create_v3.0_audience_carrier
const (
	MOBILE_ProjectCreateV30AudienceCarrier ProjectCreateV30AudienceCarrier = "MOBILE"
	TELCOM_ProjectCreateV30AudienceCarrier ProjectCreateV30AudienceCarrier = "TELCOM"
	UNICOM_ProjectCreateV30AudienceCarrier ProjectCreateV30AudienceCarrier = "UNICOM"
)

// All allowed values of ProjectCreateV30AudienceCarrier enum
var AllowedProjectCreateV30AudienceCarrierEnumValues = []ProjectCreateV30AudienceCarrier{
	"MOBILE",
	"TELCOM",
	"UNICOM",
}

// NewProjectCreateV30AudienceCarrierFromValue returns a pointer to a valid ProjectCreateV30AudienceCarrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceCarrierFromValue(v string) (*ProjectCreateV30AudienceCarrier, error) {
	ev := ProjectCreateV30AudienceCarrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceCarrier: valid values are %v", v, AllowedProjectCreateV30AudienceCarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceCarrier) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceCarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_carrier value
func (v ProjectCreateV30AudienceCarrier) Ptr() *ProjectCreateV30AudienceCarrier {
	return &v
}
