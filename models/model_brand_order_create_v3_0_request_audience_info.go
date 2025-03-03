/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderCreateV30RequestAudienceInfo 定向
type BrandOrderCreateV30RequestAudienceInfo struct {
	Ac *BrandOrderCreateV30AudienceInfoAc `json:"ac,omitempty"`
	// 兴趣
	ActionCategories []int64 `json:"action_categories,omitempty"`
	// 年龄
	Ages            []*BrandOrderCreateV30AudienceInfoAges                 `json:"ages,omitempty"`
	DistrictInfo    *BrandOrderCreateV30RequestAudienceInfoDistrictInfo    `json:"district_info,omitempty"`
	Gender          *BrandOrderCreateV30AudienceInfoGender                 `json:"gender,omitempty"`
	Platform        *BrandOrderCreateV30AudienceInfoPlatform               `json:"platform,omitempty"`
	RetargetingInfo *BrandOrderCreateV30RequestAudienceInfoRetargetingInfo `json:"retargeting_info,omitempty"`
}
