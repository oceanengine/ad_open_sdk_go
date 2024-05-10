/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestAudience
type QianchuanAdCreateV10RequestAudience struct {
	//
	Ac []*QianchuanAdCreateV10AudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64                                 `json:"action_categories,omitempty"`
	ActionDays       *QianchuanAdCreateV10AudienceActionDays `json:"action_days,omitempty"`
	//
	ActionScene []*QianchuanAdCreateV10AudienceActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	Age               []*QianchuanAdCreateV10AudienceAge             `json:"age,omitempty"`
	AudienceMode      *QianchuanAdCreateV10AudienceAudienceMode      `json:"audience_mode,omitempty"`
	AutoExtendEnabled *QianchuanAdCreateV10AudienceAutoExtendEnabled `json:"auto_extend_enabled,omitempty"`
	//
	AutoExtendTargets []*QianchuanAdCreateV10AudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors     []*QianchuanAdCreateV10AudienceAwemeFanBehaviors   `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanBehaviorsDays *QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days,omitempty"`
	//
	AwemeFanCategories []int64 `json:"aweme_fan_categories,omitempty"`
	//
	City     []int64                               `json:"city,omitempty"`
	District *QianchuanAdCreateV10AudienceDistrict `json:"district,omitempty"`
	//
	DistrictType         *bool                                             `json:"district_type,omitempty"`
	ElectricFenceRegion  *QianchuanAdCreateV10AudienceElectricFenceRegion  `json:"electric_fence_region,omitempty"`
	ExcludeLimitedRegion *QianchuanAdCreateV10AudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAdCreateV10AudienceGender               `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64 `json:"interest_words,omitempty"`
	//
	LivePlatformTags []*QianchuanAdCreateV10AudienceLivePlatformTags `json:"live_platform_tags,omitempty"`
	LocationType     *QianchuanAdCreateV10AudienceLocationType       `json:"location_type,omitempty"`
	NewCustomer      *QianchuanAdCreateV10AudienceNewCustomer        `json:"new_customer,omitempty"`
	//
	OrientationId *int64 `json:"orientation_id,omitempty"`
	//
	Platform []*QianchuanAdCreateV10AudiencePlatform `json:"platform,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                          `json:"retargeting_tags_include,omitempty"`
	SmartInterestAction    *QianchuanAdCreateV10AudienceSmartInterestAction `json:"smart_interest_action,omitempty"`
}
