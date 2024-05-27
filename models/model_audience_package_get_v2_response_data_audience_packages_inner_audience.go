/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudience
type AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudience struct {
	//
	Ac       []*AudiencePackageGetV2DataAudiencePackagesAudienceAc                  `json:"ac,omitempty"`
	Action   *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceAction   `json:"action,omitempty"`
	ActionV2 *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceActionV2 `json:"action_v2,omitempty"`
	//
	ActivateType []*AudiencePackageGetV2DataAudiencePackagesAudienceActivateType `json:"activate_type,omitempty"`
	//
	AdTag []int64 `json:"ad_tag,omitempty"`
	//
	AdTagV2 []int64 `json:"ad_tag_v2,omitempty"`
	//
	Age []string `json:"age,omitempty"`
	//
	AlbumId            []int64                                                                          `json:"album_id,omitempty"`
	AndroidLiteVersion *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceAndroidLiteVersion `json:"android_lite_version,omitempty"`
	//
	AndroidOsv     *string                                                                      `json:"android_osv,omitempty"`
	AndroidVersion *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceAndroidVersion `json:"android_version,omitempty"`
	//
	AppCategory []string `json:"app_category,omitempty"`
	//
	AppIds []string `json:"app_ids,omitempty"`
	//
	AppLocalAudience *bool `json:"app_local_audience,omitempty"`
	//
	AppRetargetingType *int64 `json:"app_retargeting_type,omitempty"`
	//
	AppVersion []int64 `json:"app_version,omitempty"`
	//
	ArticleCategory *string `json:"article_category,omitempty"`
	//
	ArticleSource []int64 `json:"article_source,omitempty"`
	//
	ArticleType *int64 `json:"article_type,omitempty"`
	//
	AuthorIds []int64 `json:"author_ids,omitempty"`
	//
	AutoExtendTargets []string                                                               `json:"auto_extend_targets,omitempty"`
	AwemeFan          *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceAwemeFan `json:"aweme_fan,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64 `json:"aweme_fan_categories,omitempty"`
	//
	AwemeFanTarget    []int64                                                            `json:"aweme_fan_target,omitempty"`
	AwemeFanTimeScope *AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	//
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	//
	AwemeFansNumbersTags []int64 `json:"aweme_fans_numbers_tags,omitempty"`
	//
	AwemeItemIds []int64 `json:"aweme_item_ids,omitempty"`
	//
	AwemeStarUserId *int64 `json:"aweme_star_user_id,omitempty"`
	//
	BrandSafety *int64 `json:"brand_safety,omitempty"`
	//
	BusinessIds []int64 `json:"business_ids,omitempty"`
	//
	CarSeries []int64 `json:"car_series,omitempty"`
	//
	Career []*AudiencePackageGetV2DataAudiencePackagesAudienceCareer `json:"career,omitempty"`
	//
	Carrier []*AudiencePackageGetV2DataAudiencePackagesAudienceCarrier `json:"carrier,omitempty"`
	//
	Carriers []int64 `json:"carriers,omitempty"`
	//
	Channel *string `json:"channel,omitempty"`
	//
	City []int64 `json:"city,omitempty"`
	//
	Country []string `json:"country,omitempty"`
	//
	CptSequence []int64 `json:"cpt_sequence,omitempty"`
	//
	CustomAudience *int64 `json:"custom_audience,omitempty"`
	//
	DeviceBrand []*AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceId []int64 `json:"device_id,omitempty"`
	//
	DeviceModels []int64 `json:"device_models,omitempty"`
	//
	DeviceType []*AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType `json:"device_type,omitempty"`
	District   *AudiencePackageGetV2DataAudiencePackagesAudienceDistrict     `json:"district,omitempty"`
	//
	DpaLocalAudience *int64 `json:"dpa_local_audience,omitempty"`
	//
	DpaRetargetingType *int64 `json:"dpa_retargeting_type,omitempty"`
	//
	EcomGuestType *int64 `json:"ecom_guest_type,omitempty"`
	//
	EpisodeId []int64 `json:"episode_id,omitempty"`
	//
	EpisodeName *string `json:"episode_name,omitempty"`
	//
	EssayChannel []int64 `json:"essay_channel,omitempty"`
	//
	ExcludeAppPackageId *string `json:"exclude_app_package_id,omitempty"`
	//
	ExcludeCustomActions *string `json:"exclude_custom_actions,omitempty"`
	//
	ExcludeFlowPackage []int64 `json:"exclude_flow_package,omitempty"`
	//
	FactoryType []int64 `json:"factory_type,omitempty"`
	//
	FilterAwemeAbnormalActive *int64 `json:"filter_aweme_abnormal_active,omitempty"`
	//
	FilterAwemeFansCount *int64 `json:"filter_aweme_fans_count,omitempty"`
	//
	FilterOwnAwemeFans *int64 `json:"filter_own_aweme_fans,omitempty"`
	//
	FilterOwnAwemeFansVTwo *int64 `json:"filter_own_aweme_fans_v_two,omitempty"`
	//
	FlowPackage []int64 `json:"flow_package,omitempty"`
	//
	GCategory []int64 `json:"g_category,omitempty"`
	//
	GKeyword []int64 `json:"g_keyword,omitempty"`
	//
	GdTags []int64                                                 `json:"gd_tags,omitempty"`
	Gender *AudiencePackageGetV2DataAudiencePackagesAudienceGender `json:"gender,omitempty"`
	//
	Geolocation []*AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceGeolocationInner `json:"geolocation,omitempty"`
	//
	GetUp *int64 `json:"get_up,omitempty"`
	//
	GroupAdTag []string `json:"group_ad_tag,omitempty"`
	//
	GroupAttribute []int64 `json:"group_attribute,omitempty"`
	//
	GroupId []int64 `json:"group_id,omitempty"`
	//
	Humidity []int64 `json:"humidity,omitempty"`
	//
	ImageMode *int64 `json:"image_mode,omitempty"`
	//
	IncludeContextWords []string `json:"include_context_words,omitempty"`
	//
	IncludeCustomActions *string `json:"include_custom_actions,omitempty"`
	//
	IndustryWords []*AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceIndustryWordsInner `json:"industry_words,omitempty"`
	//
	InstalledApps []int64 `json:"installed_apps,omitempty"`
	//
	InterestActionMode *string `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestEntry []int64 `json:"interest_entry,omitempty"`
	//
	InterestKeywordsI18n []int64 `json:"interest_keywords_i18n,omitempty"`
	//
	InterestKeywordsLangI18n []string `json:"interest_keywords_lang_i18n,omitempty"`
	//
	InterestTags []int64 `json:"interest_tags,omitempty"`
	//
	InterestWords  []int64                                                                      `json:"interest_words,omitempty"`
	IosLiteVersion *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceIosLiteVersion `json:"ios_lite_version,omitempty"`
	//
	IosOsv     *string                                                                  `json:"ios_osv,omitempty"`
	IosVersion *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceIosVersion `json:"ios_version,omitempty"`
	//
	Keywords []string `json:"keywords,omitempty"`
	//
	Language *string `json:"language,omitempty"`
	//
	LanguageList []string `json:"language_list,omitempty"`
	//
	LaunchPrice []int64 `json:"launch_price,omitempty"`
	//
	LinkItemid *string `json:"link_itemid,omitempty"`
	//
	LinkTopicIds []int64 `json:"link_topic_ids,omitempty"`
	//
	LinkUid *string `json:"link_uid,omitempty"`
	//
	LinkVideoLabel []int64 `json:"link_video_label,omitempty"`
	//
	LkCarPrice   []int64                                                       `json:"lk_car_price,omitempty"`
	LocationType *AudiencePackageGetV2DataAudiencePackagesAudienceLocationType `json:"location_type,omitempty"`
	//
	MccMnc []string `json:"mcc_mnc,omitempty"`
	//
	MediaApps []int64 `json:"media_apps,omitempty"`
	//
	MediaId *int64 `json:"media_id,omitempty"`
	//
	MediaIds []int64 `json:"media_ids,omitempty"`
	//
	MinAppVersion *int64 `json:"min_app_version,omitempty"`
	//
	MobileBrand  []string                                                                   `json:"mobile_brand,omitempty"`
	MovieAndStar *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceMovieAndStar `json:"movie_and_star,omitempty"`
	//
	NewAppCategory *string `json:"new_app_category,omitempty"`
	//
	NewAppIds *string `json:"new_app_ids,omitempty"`
	//
	NovelCategory []int64 `json:"novel_category,omitempty"`
	//
	OwnAwemeNumber *int64 `json:"own_aweme_number,omitempty"`
	//
	OwnAwemeNumberTags []int64 `json:"own_aweme_number_tags,omitempty"`
	//
	ParticleLocations []int64 `json:"particle_locations,omitempty"`
	//
	PgcMediaIds *string `json:"pgc_media_ids,omitempty"`
	//
	Platform []*AudiencePackageGetV2DataAudiencePackagesAudiencePlatform `json:"platform,omitempty"`
	//
	Platforms []int64 `json:"platforms,omitempty"`
	//
	Product []int64 `json:"product,omitempty"`
	//
	Province []int64 `json:"province,omitempty"`
	//
	QcLimitedRegion *int64 `json:"qc_limited_region,omitempty"`
	//
	RecQuality []int64 `json:"rec_quality,omitempty"`
	//
	RefreshTimes []int64 `json:"refresh_times,omitempty"`
	//
	RegionVer *string `json:"region_ver,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTags []int64 `json:"retargeting_tags,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	ScreenSizeType []int64 `json:"screen_size_type,omitempty"`
	//
	SearchOffset []int64 `json:"search_offset,omitempty"`
	//
	Sleep *int64 `json:"sleep,omitempty"`
	//
	SpecialIds []int64 `json:"special_ids,omitempty"`
	//
	SpecificUserIds []int64 `json:"specific_user_ids,omitempty"`
	//
	StartingMode []int64 `json:"starting_mode,omitempty"`
	//
	StartingTimes []int64 `json:"starting_times,omitempty"`
	//
	SuperiorPopularityType *int64 `json:"superior_popularity_type,omitempty"`
	//
	SupportC2s *int64 `json:"support_c2s,omitempty"`
	//
	TargetUid *int64 `json:"target_uid,omitempty"`
	//
	TargetUids []int64 `json:"target_uids,omitempty"`
	//
	Temperature []int64 `json:"temperature,omitempty"`
	//
	Temperatureoptional []int64 `json:"temperatureoptional,omitempty"`
	//
	TextKeywordsIds []int64 `json:"text_keywords_ids,omitempty"`
	//
	TopicCategory []int64 `json:"topic_category,omitempty"`
	//
	UnionSlotId []int64 `json:"union_slot_id,omitempty"`
	//
	UnionSlotType *int64 `json:"union_slot_type,omitempty"`
	//
	UserType []int64 `json:"user_type,omitempty"`
	//
	VideoPatchNumber []int64 `json:"video_patch_number,omitempty"`
	//
	VideoType []int64 `json:"video_type,omitempty"`
	//
	Weather []int64 `json:"weather,omitempty"`
	//
	WholeCountry []string `json:"whole_country,omitempty"`
	//
	Wind []int64 `json:"wind,omitempty"`
	//
	WordbagIds []int64 `json:"wordbag_ids,omitempty"`
}
