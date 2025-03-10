/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2Gender
type AudiencePackageUpdateV2Gender string

// List of audience_package_update_v2_gender
const (
	GENDER_FEMALE_AudiencePackageUpdateV2Gender    AudiencePackageUpdateV2Gender = "GENDER_FEMALE"
	GENDER_MALE_AudiencePackageUpdateV2Gender      AudiencePackageUpdateV2Gender = "GENDER_MALE"
	NONE_AudiencePackageUpdateV2Gender             AudiencePackageUpdateV2Gender = "NONE"
	GENDER_UNLIMITED_AudiencePackageUpdateV2Gender AudiencePackageUpdateV2Gender = "GENDER_UNLIMITED"
)

// Ptr returns reference to audience_package_update_v2_gender value
func (v AudiencePackageUpdateV2Gender) Ptr() *AudiencePackageUpdateV2Gender {
	return &v
}
