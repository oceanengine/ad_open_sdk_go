/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestAudience
type QianchuanAdUpdateV10RequestAudience struct {
	//
	Ac []*QianchuanAdUpdateV10AudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64                                 `json:"action_categories,omitempty"`
	ActionDays       *QianchuanAdUpdateV10AudienceActionDays `json:"action_days,omitempty"`
	//
	ActionScene []*QianchuanAdUpdateV10AudienceActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	Age               []*QianchuanAdUpdateV10AudienceAge             `json:"age,omitempty"`
	AudienceMode      *QianchuanAdUpdateV10AudienceAudienceMode      `json:"audience_mode,omitempty"`
	AutoExtendEnabled *QianchuanAdUpdateV10AudienceAutoExtendEnabled `json:"auto_extend_enabled,omitempty"`
	//
	AutoExtendTargets []*QianchuanAdUpdateV10AudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors     []*QianchuanAdUpdateV10AudienceAwemeFanBehaviors   `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanBehaviorsDays *QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days,omitempty"`
	//
	AwemeFanCategories []int64 `json:"aweme_fan_categories,omitempty"`
	//
	City     []int64                               `json:"city,omitempty"`
	District *QianchuanAdUpdateV10AudienceDistrict `json:"district,omitempty"`
	//
	DistrictType         *bool                                             `json:"district_type,omitempty"`
	ExcludeLimitedRegion *QianchuanAdUpdateV10AudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAdUpdateV10AudienceGender               `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64 `json:"interest_words,omitempty"`
	//
	LivePlatformTags []*QianchuanAdUpdateV10AudienceLivePlatformTags `json:"live_platform_tags,omitempty"`
	LocationType     *QianchuanAdUpdateV10AudienceLocationType       `json:"location_type,omitempty"`
	NewCustomer      *QianchuanAdUpdateV10AudienceNewCustomer        `json:"new_customer,omitempty"`
	//
	OrientationId *int64 `json:"orientation_id,omitempty"`
	//
	Platform []*QianchuanAdUpdateV10AudiencePlatform `json:"platform,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                          `json:"retargeting_tags_include,omitempty"`
	SmartInterestAction    *QianchuanAdUpdateV10AudienceSmartInterestAction `json:"smart_interest_action,omitempty"`
}
