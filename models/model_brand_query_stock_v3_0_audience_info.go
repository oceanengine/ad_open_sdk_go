/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryStockV30AudienceInfo
type BrandQueryStockV30AudienceInfo struct {
	Ac *BrandQueryStockV30AudienceInfoAc `json:"ac,omitempty"`
	// 兴趣定向
	ActionCategory []int64 `json:"action_category,omitempty"`
	// 年龄定向，支持的枚举值：<br> BETWEEN18_23：18～23岁<br> BETWEEN24_30：24～30岁<br> BETWEEN31_40：31～40岁<br> BETWEEN41_49：41～49岁<br> ABOVE50：50岁及以上 UNLIMITED：不限
	Ages            []*BrandQueryStockV30AudienceInfoAges          `json:"ages,omitempty"`
	DistrictInfo    *BrandQueryStockV30AudienceInfoDistrictInfo    `json:"district_info,omitempty"`
	Gender          *BrandQueryStockV30AudienceInfoGender          `json:"gender,omitempty"`
	Platform        *BrandQueryStockV30AudienceInfoPlatform        `json:"platform,omitempty"`
	RetargetingInfo *BrandQueryStockV30AudienceInfoRetargetingInfo `json:"retargeting_info,omitempty"`
}
