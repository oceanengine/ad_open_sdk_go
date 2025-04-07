/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderUpdateV30RequestAudienceInfo 定向
type BrandOrderUpdateV30RequestAudienceInfo struct {
	Ac *BrandOrderUpdateV30AudienceInfoAc `json:"ac,omitempty"`
	// 兴趣
	ActionCategories []int64 `json:"action_categories,omitempty"`
	// 年龄
	Ages            []*BrandOrderUpdateV30AudienceInfoAges                 `json:"ages,omitempty"`
	DistrictInfo    *BrandOrderUpdateV30RequestAudienceInfoDistrictInfo    `json:"district_info,omitempty"`
	Gender          *BrandOrderUpdateV30AudienceInfoGender                 `json:"gender,omitempty"`
	Platform        *BrandOrderUpdateV30AudienceInfoPlatform               `json:"platform,omitempty"`
	RetargetingInfo *BrandOrderUpdateV30RequestAudienceInfoRetargetingInfo `json:"retargeting_info,omitempty"`
}
