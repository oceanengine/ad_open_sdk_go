/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30RequestAudience
type ProjectUpdateV30RequestAudience struct {
	//
	Ac []*ProjectUpdateV30AudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64 `json:"action_categories,omitempty"`
	//
	ActionDays *int64 `json:"action_days,omitempty"`
	//
	ActionScene []*ProjectUpdateV30AudienceActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	Age        []string                            `json:"age,omitempty"`
	AndroidOsv *ProjectUpdateV30AudienceAndroidOsv `json:"android_osv,omitempty"`
	//
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	//
	AutoExtendTargets []*ProjectUpdateV30AudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*ProjectUpdateV30AudienceAwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64                                    `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope  *ProjectUpdateV30AudienceAwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	//
	Carrier []*ProjectUpdateV30AudienceCarrier `json:"carrier,omitempty"`
	//
	City                  []int64                                        `json:"city,omitempty"`
	ConvertedTimeDuration *ProjectUpdateV30AudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	//
	DeviceBrand []*ProjectUpdateV30AudienceDeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType []*ProjectUpdateV30AudienceDeviceType `json:"device_type,omitempty"`
	District   *ProjectUpdateV30AudienceDistrict     `json:"district,omitempty"`
	DpaCity    *ProjectUpdateV30AudienceDpaCity      `json:"dpa_city,omitempty"`
	DpaLbs     *ProjectUpdateV30AudienceDpaLbs       `json:"dpa_lbs,omitempty"`
	//
	ExcludeFlowPackage        []int64                                            `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive *ProjectUpdateV30AudienceFilterAwemeAbnormalActive `json:"filter_aweme_abnormal_active,omitempty"`
	FilterAwemeFansCount      *ProjectUpdateV30AudienceFilterAwemeFansCount      `json:"filter_aweme_fans_count,omitempty"`
	FilterOwnAwemeFans        *ProjectUpdateV30AudienceFilterOwnAwemeFans        `json:"filter_own_aweme_fans,omitempty"`
	//
	FlowPackage []int64                         `json:"flow_package,omitempty"`
	Gender      *ProjectUpdateV30AudienceGender `json:"gender,omitempty"`
	//
	Geolocation        []*ProjectCreateV30RequestAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted    *ProjectUpdateV30AudienceHideIfConverted           `json:"hide_if_converted,omitempty"`
	HideIfExists       *ProjectUpdateV30AudienceHideIfExists              `json:"hide_if_exists,omitempty"`
	InterestActionMode *ProjectUpdateV30AudienceInterestActionMode        `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64                         `json:"interest_words,omitempty"`
	IosOsv        *ProjectUpdateV30AudienceIosOsv `json:"ios_osv,omitempty"`
	//
	LaunchPrice  []int64                               `json:"launch_price,omitempty"`
	LocationType *ProjectUpdateV30AudienceLocationType `json:"location_type,omitempty"`
	//
	Platform        []*ProjectUpdateV30AudiencePlatform      `json:"platform,omitempty"`
	RegionRecommend *ProjectUpdateV30AudienceRegionRecommend `json:"region_recommend,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                         `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *ProjectUpdateV30AudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
