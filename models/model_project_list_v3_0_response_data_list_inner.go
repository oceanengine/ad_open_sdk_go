/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInner struct for ProjectListV30ResponseDataListInner
type ProjectListV30ResponseDataListInner struct {
	AdType *ProjectListV30DataListAdType `json:"ad_type,omitempty"`
	//
	AdvertiserId              *int64                                           `json:"advertiser_id,omitempty"`
	AigcDynamicCreativeSwitch *ProjectListV30DataListAigcDynamicCreativeSwitch `json:"aigc_dynamic_creative_switch,omitempty"`
	//
	AppName          *string                                      `json:"app_name,omitempty"`
	AppPromotionType *ProjectListV30DataListAppPromotionType      `json:"app_promotion_type,omitempty"`
	AssetType        *ProjectListV30DataListAssetType             `json:"asset_type,omitempty"`
	Audience         *ProjectListV30ResponseDataListInnerAudience `json:"audience,omitempty"`
	//
	AudienceExtend  *string                                             `json:"audience_extend,omitempty"`
	BlueFlowPackage *ProjectListV30ResponseDataListInnerBlueFlowPackage `json:"blue_flow_package,omitempty"`
	//
	BudgetGroupId   *int64                                              `json:"budget_group_id,omitempty"`
	DeliveryMedium  *ProjectListV30DataListDeliveryMedium               `json:"delivery_medium,omitempty"`
	DeliveryMode    *ProjectListV30DataListDeliveryMode                 `json:"delivery_mode,omitempty"`
	DeliveryProduct *ProjectListV30DataListDeliveryProduct              `json:"delivery_product,omitempty"`
	DeliveryRange   *ProjectListV30ResponseDataListInnerDeliveryRange   `json:"delivery_range,omitempty"`
	DeliverySetting *ProjectListV30ResponseDataListInnerDeliverySetting `json:"delivery_setting,omitempty"`
	DeliveryType    *ProjectListV30DataListDeliveryType                 `json:"delivery_type,omitempty"`
	DownloadMode    *ProjectListV30DataListDownloadMode                 `json:"download_mode,omitempty"`
	DownloadType    *ProjectListV30DataListDownloadType                 `json:"download_type,omitempty"`
	//
	DownloadUrl *string                          `json:"download_url,omitempty"`
	DpaAdtype   *ProjectListV30DataListDpaAdtype `json:"dpa_adtype,omitempty"`
	//
	DpaCategories []int64 `json:"dpa_categories,omitempty"`
	//
	DpaProductTarget []*ProjectListV30ResponseDataListInnerDpaProductTargetInner `json:"dpa_product_target,omitempty"`
	//
	FeedDeliverySearch *string `json:"feed_delivery_search,omitempty"`
	//
	IfNewcustomerdelivery  *bool                                                      `json:"if_newcustomerdelivery,omitempty"`
	InternalAdvertiserInfo *ProjectListV30ResponseDataListInnerInternalAdvertiserInfo `json:"internal_advertiser_info,omitempty"`
	// 关键词
	Keywords            []*ProjectListV30ResponseDataListInnerKeywordsInner `json:"keywords,omitempty"`
	LandingPageStayTime *ProjectListV30DataListLandingPageStayTime          `json:"landing_page_stay_time,omitempty"`
	LandingType         *ProjectListV30DataListLandingType                  `json:"landing_type,omitempty"`
	LaunchType          *ProjectListV30DataListLaunchType                   `json:"launch_type,omitempty"`
	MarketingGoal       *ProjectListV30DataListMarketingGoal                `json:"marketing_goal,omitempty"`
	// 对应「创建字节小程序」接口获取的instance_id
	MicroAppInstanceId  *int64                                     `json:"micro_app_instance_id,omitempty"`
	MicroPromotionType  *ProjectListV30DataListMicroPromotionType  `json:"micro_promotion_type,omitempty"`
	MultiAssetType      *ProjectListV30DataListMultiAssetType      `json:"multi_asset_type,omitempty"`
	MultiDeliveryMedium *ProjectListV30DataListMultiDeliveryMedium `json:"multi_delivery_medium,omitempty"`
	//
	Name          *string                                           `json:"name,omitempty"`
	NativeSetting *ProjectListV30ResponseDataListInnerNativeSetting `json:"native_setting,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlField *string `json:"open_url_field,omitempty"`
	//
	OpenUrlParams *string                                          `json:"open_url_params,omitempty"`
	OpenUrlType   *ProjectListV30DataListOpenUrlType               `json:"open_url_type,omitempty"`
	OptStatus     *ProjectListV30DataListOptStatus                 `json:"opt_status,omitempty"`
	OptimizeGoal  *ProjectListV30ResponseDataListInnerOptimizeGoal `json:"optimize_goal,omitempty"`
	//
	PackageName *string                        `json:"package_name,omitempty"`
	Pricing     *ProjectListV30DataListPricing `json:"pricing,omitempty"`
	//
	ProjectCreateTime *string `json:"project_create_time,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
	//
	ProjectModifyTime *string                              `json:"project_modify_time,omitempty"`
	PromotionType     *ProjectListV30DataListPromotionType `json:"promotion_type,omitempty"`
	// 对应「查询快应用信息」接口获取的quick_app_id
	QuickAppId     *int64                                             `json:"quick_app_id,omitempty"`
	RelatedProduct *ProjectListV30ResponseDataListInnerRelatedProduct `json:"related_product,omitempty"`
	//
	SearchBidRatio                 *float64                                              `json:"search_bid_ratio,omitempty"`
	StarAutoMaterialAdditionSwitch *ProjectListV30DataListStarAutoMaterialAdditionSwitch `json:"star_auto_material_addition_switch,omitempty"`
	//
	StarTaskId      *int64                                 `json:"star_task_id,omitempty"`
	StarTaskVersion *ProjectListV30DataListStarTaskVersion `json:"star_task_version,omitempty"`
	Status          *ProjectListV30DataListStatus          `json:"status,omitempty"`
	StatusFirst     *ProjectListV30DataListStatusFirst     `json:"status_first,omitempty"`
	//
	StatusSecond []*ProjectListV30DataListStatusSecond `json:"status_second,omitempty"`
	//
	SubscribeUrl    *string                                             `json:"subscribe_url,omitempty"`
	TrackUrlSetting *ProjectListV30ResponseDataListInnerTrackUrlSetting `json:"track_url_setting,omitempty"`
	//
	UlinkUrl           *string                                   `json:"ulink_url,omitempty"`
	ValueOptimizedType *ProjectListV30DataListValueOptimizedType `json:"value_optimized_type,omitempty"`
}
