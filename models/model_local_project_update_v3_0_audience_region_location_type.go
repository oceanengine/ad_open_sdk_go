/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectUpdateV30AudienceRegionLocationType
type LocalProjectUpdateV30AudienceRegionLocationType string

// List of local_project_update_v3.0_audience_region_location_type
const (
	ALL_LocalProjectUpdateV30AudienceRegionLocationType     LocalProjectUpdateV30AudienceRegionLocationType = "ALL"
	CURRENT_LocalProjectUpdateV30AudienceRegionLocationType LocalProjectUpdateV30AudienceRegionLocationType = "CURRENT"
	HOME_LocalProjectUpdateV30AudienceRegionLocationType    LocalProjectUpdateV30AudienceRegionLocationType = "HOME"
	TRAVEL_LocalProjectUpdateV30AudienceRegionLocationType  LocalProjectUpdateV30AudienceRegionLocationType = "TRAVEL"
)

// Ptr returns reference to local_project_update_v3.0_audience_region_location_type value
func (v LocalProjectUpdateV30AudienceRegionLocationType) Ptr() *LocalProjectUpdateV30AudienceRegionLocationType {
	return &v
}