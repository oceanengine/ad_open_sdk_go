/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandPolicyListV30AudienceInfoGender
type BrandPolicyListV30AudienceInfoGender string

// List of brand_policy_list_v3.0_audience_info_gender
const (
	FEMALE_BrandPolicyListV30AudienceInfoGender    BrandPolicyListV30AudienceInfoGender = "FEMALE"
	MALE_BrandPolicyListV30AudienceInfoGender      BrandPolicyListV30AudienceInfoGender = "MALE"
	UNLIMITED_BrandPolicyListV30AudienceInfoGender BrandPolicyListV30AudienceInfoGender = "UNLIMITED"
)

// Ptr returns reference to brand_policy_list_v3.0_audience_info_gender value
func (v BrandPolicyListV30AudienceInfoGender) Ptr() *BrandPolicyListV30AudienceInfoGender {
	return &v
}
