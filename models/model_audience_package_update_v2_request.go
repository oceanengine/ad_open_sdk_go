/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2Request struct for AudiencePackageUpdateV2Request
type AudiencePackageUpdateV2Request struct {
	//
	Ac []*AudiencePackageUpdateV2Ac `json:"ac,omitempty"`
	//
	ActionCategories []int64                            `json:"action_categories,omitempty"`
	ActionDays       *AudiencePackageUpdateV2ActionDays `json:"action_days,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Age        []*AudiencePackageUpdateV2Age      `json:"age,omitempty"`
	AndroidOsv *AudiencePackageUpdateV2AndroidOsv `json:"android_osv,omitempty"`
	//
	AudiencePackageId int64 `json:"audience_package_id"`
	//
	AutoExtendTargets []*AudiencePackageUpdateV2AutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*AudiencePackageUpdateV2AwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64                                   `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope  *AudiencePackageUpdateV2AwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	//
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	//
	BusinessIds []int64 `json:"business_ids,omitempty"`
	//
	Carrier               []*AudiencePackageUpdateV2Carrier             `json:"carrier,omitempty"`
	CarrierRegionOptimize *AudiencePackageUpdateV2CarrierRegionOptimize `json:"carrier_region_optimize,omitempty"`
	//
	City                  []int64                                       `json:"city,omitempty"`
	ConvertedTimeDuration *AudiencePackageUpdateV2ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	DeviceBrand []*AudiencePackageUpdateV2DeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType []*AudiencePackageUpdateV2DeviceType `json:"device_type,omitempty"`
	District   *AudiencePackageUpdateV2District     `json:"district,omitempty"`
	//
	ExcludeFlowPackage        []int64                                           `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive *AudiencePackageUpdateV2FilterAwemeAbnormalActive `json:"filter_aweme_abnormal_active,omitempty"`
	//
	FilterAwemeFansCount *int64 `json:"filter_aweme_fans_count,omitempty"`
	//
	FilterEvent        []*AudiencePackageUpdateV2FilterEvent      `json:"filter_event,omitempty"`
	FilterOwnAwemeFans *AudiencePackageUpdateV2FilterOwnAwemeFans `json:"filter_own_aweme_fans,omitempty"`
	//
	FlowPackage []int64                        `json:"flow_package,omitempty"`
	Gender      *AudiencePackageUpdateV2Gender `json:"gender,omitempty"`
	//
	Geolocation        []*AudiencePackageUpdateV2RequestGeolocationInner `json:"geolocation,omitempty"`
	HarmonyOsv         *AudiencePackageUpdateV2HarmonyOsv                `json:"harmony_osv,omitempty"`
	HideIfConverted    *AudiencePackageUpdateV2HideIfConverted           `json:"hide_if_converted,omitempty"`
	HideIfExists       *AudiencePackageUpdateV2HideIfExists              `json:"hide_if_exists,omitempty"`
	InterestActionMode *AudiencePackageUpdateV2InterestActionMode        `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64                        `json:"interest_words,omitempty"`
	IosOsv        *AudiencePackageUpdateV2IosOsv `json:"ios_osv,omitempty"`
	//
	LaunchPrice  []int64                              `json:"launch_price,omitempty"`
	LocationType *AudiencePackageUpdateV2LocationType `json:"location_type,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Platform []*AudiencePackageUpdateV2Platform `json:"platform,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTags []int64 `json:"retargeting_tags,omitempty"`
	//
	RetargetingTagsExclude []int64                                        `json:"retargeting_tags_exclude,omitempty"`
	SuperiorPopularityType *AudiencePackageUpdateV2SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
