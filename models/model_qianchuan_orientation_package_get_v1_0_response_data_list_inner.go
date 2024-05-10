/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanOrientationPackageGetV10ResponseDataListInner struct for QianchuanOrientationPackageGetV10ResponseDataListInner
type QianchuanOrientationPackageGetV10ResponseDataListInner struct {
	//
	InActiveRetargetingTags []*QianchuanOrientationPackageGetV10ResponseDataListInnerInActiveRetargetingTagsInner `json:"InActive_retargeting_tags,omitempty"`
	//
	Ac []*QianchuanOrientationPackageGetV10DataListAc `json:"ac,omitempty"`
	//
	ActionCategories []int32 `json:"action_categories,omitempty"`
	//
	ActionDays *int32 `json:"action_days,omitempty"`
	//
	ActionScene []*QianchuanOrientationPackageGetV10DataListActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	AdNum *int64 `json:"ad_num,omitempty"`
	//
	Age []*QianchuanOrientationPackageGetV10DataListAge `json:"age,omitempty"`
	//
	AutoExtendEnabled *int32 `json:"auto_extend_enabled,omitempty"`
	//
	AutoExtendTargets []*QianchuanOrientationPackageGetV10DataListAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors     []*QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors   `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanBehaviorsDays *QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days,omitempty"`
	//
	AwemeFanCategories []int64 `json:"aweme_fan_categories,omitempty"`
	//
	City                []int64                                                       `json:"city,omitempty"`
	District            *QianchuanOrientationPackageGetV10DataListDistrict            `json:"district,omitempty"`
	ElectricFenceRegion *QianchuanOrientationPackageGetV10DataListElectricFenceRegion `json:"electric_fence_region,omitempty"`
	Gender              *QianchuanOrientationPackageGetV10DataListGender              `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64 `json:"interest_words,omitempty"`
	//
	LivePlatformTags []string                                               `json:"live_platform_tags,omitempty"`
	LocationType     *QianchuanOrientationPackageGetV10DataListLocationType `json:"location_type,omitempty"`
	NewCustomer      *QianchuanOrientationPackageGetV10DataListNewCustomer  `json:"new_customer,omitempty"`
	//
	OrientationId int64 `json:"orientation_id"`
	//
	OrientationInfo *string `json:"orientation_info,omitempty"`
	//
	OrientationName string `json:"orientation_name"`
	//
	Platform []*QianchuanOrientationPackageGetV10DataListPlatform `json:"platform,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                                       `json:"retargeting_tags_include,omitempty"`
	SmartInterestAction    *QianchuanOrientationPackageGetV10DataListSmartInterestAction `json:"smart_interest_action,omitempty"`
}
