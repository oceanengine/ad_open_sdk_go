/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectUpdateV30AudienceGender
type LocalProjectUpdateV30AudienceGender string

// List of local_project_update_v3.0_audience_gender
const (
	FEMALE_LocalProjectUpdateV30AudienceGender LocalProjectUpdateV30AudienceGender = "FEMALE"
	MALE_LocalProjectUpdateV30AudienceGender   LocalProjectUpdateV30AudienceGender = "MALE"
	NONE_LocalProjectUpdateV30AudienceGender   LocalProjectUpdateV30AudienceGender = "NONE"
)

// Ptr returns reference to local_project_update_v3.0_audience_gender value
func (v LocalProjectUpdateV30AudienceGender) Ptr() *LocalProjectUpdateV30AudienceGender {
	return &v
}
