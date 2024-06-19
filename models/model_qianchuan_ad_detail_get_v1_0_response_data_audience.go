/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataAudience
type QianchuanAdDetailGetV10ResponseDataAudience struct {
	//
	Ac []*QianchuanAdDetailGetV10DataAudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64 `json:"action_categories,omitempty"`
	//
	ActionDays *int64 `json:"action_days,omitempty"`
	//
	ActionScene []*QianchuanAdDetailGetV10DataAudienceActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	Age          []*QianchuanAdDetailGetV10DataAudienceAge        `json:"age,omitempty"`
	AudienceMode *QianchuanAdDetailGetV10DataAudienceAudienceMode `json:"audience_mode,omitempty"`
	//
	AutoExtendEnabled *int64 `json:"auto_extend_enabled,omitempty"`
	//
	AutoExtendTargets []*QianchuanAdDetailGetV10DataAudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors     []*QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors   `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanBehaviorsDays *QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days,omitempty"`
	//
	AwemeFanCategories []int64 `json:"aweme_fan_categories,omitempty"`
	//
	City     []int64                                      `json:"city,omitempty"`
	District *QianchuanAdDetailGetV10DataAudienceDistrict `json:"district,omitempty"`
	//
	DistrictType        *bool                                                   `json:"district_type,omitempty"`
	ElectricFenceRegion *QianchuanAdDetailGetV10DataAudienceElectricFenceRegion `json:"electric_fence_region,omitempty"`
	//
	ExcludeLimitedRegion *int64                                     `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAdDetailGetV10DataAudienceGender `json:"gender,omitempty"`
	//
	InactiveRetargetingTags []*QianchuanAdDetailGetV10ResponseDataAudienceInactiveRetargetingTagsInner `json:"inactive_retargeting_tags,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64 `json:"interest_words,omitempty"`
	//
	LivePlatformTags []string                                         `json:"live_platform_tags,omitempty"`
	LocationType     *QianchuanAdDetailGetV10DataAudienceLocationType `json:"location_type,omitempty"`
	NewCustomer      *QianchuanAdDetailGetV10DataAudienceNewCustomer  `json:"new_customer,omitempty"`
	//
	OrientationId *int64 `json:"orientation_id,omitempty"`
	//
	Platform []*QianchuanAdDetailGetV10DataAudiencePlatform `json:"platform,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                                 `json:"retargeting_tags_include,omitempty"`
	SearchExtended         *QianchuanAdDetailGetV10DataAudienceSearchExtended      `json:"search_extended,omitempty"`
	SmartInterestAction    *QianchuanAdDetailGetV10DataAudienceSmartInterestAction `json:"smart_interest_action,omitempty"`
}
