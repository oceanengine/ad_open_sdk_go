/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderUpdateV30RequestAudienceInfoRetargetingInfo 人群包
type BrandOrderUpdateV30RequestAudienceInfoRetargetingInfo struct {
	// 人群包ID列表
	RetargetingTags []int64                                                        `json:"retargeting_tags,omitempty"`
	RetargetingType *BrandOrderUpdateV30AudienceInfoRetargetingInfoRetargetingType `json:"retargeting_type,omitempty"`
}
