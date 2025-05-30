/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandPolicyListV30AudienceInfoAges
type BrandPolicyListV30AudienceInfoAges string

// List of brand_policy_list_v3.0_audience_info_ages
const (
	ABOVE50_BrandPolicyListV30AudienceInfoAges      BrandPolicyListV30AudienceInfoAges = "ABOVE50"
	BETWEEN18_23_BrandPolicyListV30AudienceInfoAges BrandPolicyListV30AudienceInfoAges = "BETWEEN18_23"
	BETWEEN24_30_BrandPolicyListV30AudienceInfoAges BrandPolicyListV30AudienceInfoAges = "BETWEEN24_30"
	BETWEEN31_40_BrandPolicyListV30AudienceInfoAges BrandPolicyListV30AudienceInfoAges = "BETWEEN31_40"
	BETWEEN41_49_BrandPolicyListV30AudienceInfoAges BrandPolicyListV30AudienceInfoAges = "BETWEEN41_49"
)

// Ptr returns reference to brand_policy_list_v3.0_audience_info_ages value
func (v BrandPolicyListV30AudienceInfoAges) Ptr() *BrandPolicyListV30AudienceInfoAges {
	return &v
}
