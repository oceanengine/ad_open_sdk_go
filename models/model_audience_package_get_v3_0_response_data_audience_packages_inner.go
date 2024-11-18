/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV30ResponseDataAudiencePackagesInner struct for AudiencePackageGetV30ResponseDataAudiencePackagesInner
type AudiencePackageGetV30ResponseDataAudiencePackagesInner struct {
	//
	Ac       []*AudiencePackageGetV30DataAudiencePackagesAc                  `json:"ac,omitempty"`
	Action   *AudiencePackageGetV30ResponseDataAudiencePackagesInnerAction   `json:"action,omitempty"`
	ActionV2 *AudiencePackageGetV30ResponseDataAudiencePackagesInnerActionV2 `json:"action_v2,omitempty"`
	//
	AdTagV2 []int64                                          `json:"ad_tag_v2,omitempty"`
	AdType  *AudiencePackageGetV30DataAudiencePackagesAdType `json:"ad_type,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Age        []*AudiencePackageGetV30DataAudiencePackagesAge      `json:"age,omitempty"`
	AndroidOsv *AudiencePackageGetV30DataAudiencePackagesAndroidOsv `json:"android_osv,omitempty"`
	AppType    *AudiencePackageGetV30DataAudiencePackagesAppType    `json:"app_type,omitempty"`
	//
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	//
	AutoExtendTargets []*AudiencePackageGetV30DataAudiencePackagesAutoExtendTargets   `json:"auto_extend_targets,omitempty"`
	AwemeFan          *AudiencePackageGetV30ResponseDataAudiencePackagesInnerAwemeFan `json:"aweme_fan,omitempty"`
	//
	AwemeFanTarget []int64 `json:"aweme_fan_target,omitempty"`
	//
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	//
	AwemeFansNumbersTags []int64 `json:"aweme_fans_numbers_tags,omitempty"`
	//
	Carrier               []*AudiencePackageGetV30DataAudiencePackagesCarrier             `json:"carrier,omitempty"`
	CarrierRegionOptimize *AudiencePackageGetV30DataAudiencePackagesCarrierRegionOptimize `json:"carrier_region_optimize,omitempty"`
	//
	City                  []int64                                                         `json:"city,omitempty"`
	ConvertedTimeDuration *AudiencePackageGetV30DataAudiencePackagesConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	DeliveryRange         *AudiencePackageGetV30DataAudiencePackagesDeliveryRange         `json:"delivery_range,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	DeviceBrand []*AudiencePackageGetV30DataAudiencePackagesDeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType []*AudiencePackageGetV30DataAudiencePackagesDeviceType `json:"device_type,omitempty"`
	District   *AudiencePackageGetV30DataAudiencePackagesDistrict     `json:"district,omitempty"`
	//
	ExcludeFlowPackage        []int64                                                             `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive *AudiencePackageGetV30DataAudiencePackagesFilterAwemeAbnormalActive `json:"filter_aweme_abnormal_active,omitempty"`
	FilterAwemeFansCount      *AudiencePackageGetV30DataAudiencePackagesFilterAwemeFansCount      `json:"filter_aweme_fans_count,omitempty"`
	//
	FilterEvent            []*AudiencePackageGetV30DataAudiencePackagesFilterEvent          `json:"filter_event,omitempty"`
	FilterOwnAwemeFansVTwo *AudiencePackageGetV30DataAudiencePackagesFilterOwnAwemeFansVTwo `json:"filter_own_aweme_fans_v_two,omitempty"`
	//
	FlowPackage []int64                                          `json:"flow_package,omitempty"`
	Gender      *AudiencePackageGetV30DataAudiencePackagesGender `json:"gender,omitempty"`
	//
	Geolocation     []*AudiencePackageGetV30ResponseDataAudiencePackagesInnerGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted *AudiencePackageGetV30DataAudiencePackagesHideIfConverted                 `json:"hide_if_converted,omitempty"`
	HideIfExists    *AudiencePackageGetV30DataAudiencePackagesHideIfExists                    `json:"hide_if_exists,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64                                               `json:"interest_words,omitempty"`
	IosOsv        *AudiencePackageGetV30DataAudiencePackagesIosOsv      `json:"ios_osv,omitempty"`
	LandingType   *AudiencePackageGetV30DataAudiencePackagesLandingType `json:"landing_type,omitempty"`
	//
	LaunchPrice   []int64                                                 `json:"launch_price,omitempty"`
	LocationType  *AudiencePackageGetV30DataAudiencePackagesLocationType  `json:"location_type,omitempty"`
	MarketingGoal *AudiencePackageGetV30DataAudiencePackagesMarketingGoal `json:"marketing_goal,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	NewAppCategory *string `json:"new_app_category,omitempty"`
	//
	OwnAwemeNumber *int64 `json:"own_aweme_number,omitempty"`
	//
	OwnAwemeNumberTags []int64 `json:"own_aweme_number_tags,omitempty"`
	//
	Platform []*AudiencePackageGetV30DataAudiencePackagesPlatform `json:"platform,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                                          `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *AudiencePackageGetV30DataAudiencePackagesSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
