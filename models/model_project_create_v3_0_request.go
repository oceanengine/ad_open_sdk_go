/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30Request struct for ProjectCreateV30Request
type ProjectCreateV30Request struct {
	AdType ProjectCreateV30AdType `json:"ad_type"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AdvertisingVolume         *int64                                     `json:"advertising_volume,omitempty"`
	AigcDynamicCreativeSwitch *ProjectCreateV30AigcDynamicCreativeSwitch `json:"aigc_dynamic_creative_switch,omitempty"`
	//
	AppName           *string                                 `json:"app_name,omitempty"`
	AppPromotionType  *ProjectCreateV30AppPromotionType       `json:"app_promotion_type,omitempty"`
	AssetType         *ProjectCreateV30AssetType              `json:"asset_type,omitempty"`
	Audience          *ProjectCreateV30RequestAudience        `json:"audience,omitempty"`
	AudienceExtend    *ProjectCreateV30AudienceExtend         `json:"audience_extend,omitempty"`
	AutoExtendTraffic *ProjectCreateV30AutoExtendTraffic      `json:"auto_extend_traffic,omitempty"`
	BlueFlowPackage   *ProjectCreateV30RequestBlueFlowPackage `json:"blue_flow_package,omitempty"`
	//
	BudgetGroupId   *int64                                 `json:"budget_group_id,omitempty"`
	Classify        *ProjectCreateV30Classify              `json:"classify,omitempty"`
	DeliveryMode    *ProjectCreateV30DeliveryMode          `json:"delivery_mode,omitempty"`
	DeliveryRange   ProjectCreateV30RequestDeliveryRange   `json:"delivery_range"`
	DeliverySetting ProjectCreateV30RequestDeliverySetting `json:"delivery_setting"`
	DeliveryType    *ProjectCreateV30DeliveryType          `json:"delivery_type,omitempty"`
	DownloadMode    *ProjectCreateV30DownloadMode          `json:"download_mode,omitempty"`
	DownloadType    *ProjectCreateV30DownloadType          `json:"download_type,omitempty"`
	//
	DownloadUrl *string                    `json:"download_url,omitempty"`
	DpaAdtype   *ProjectCreateV30DpaAdtype `json:"dpa_adtype,omitempty"`
	// 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。获取商品库元信息-商品广告-商业开放平台
	DpaProductTarget []*ProjectCreateV30RequestDpaProductTargetInner `json:"dpa_product_target,omitempty"`
	//
	ItemNumber *string `json:"item_number,omitempty"`
	// 关键词列表
	Keywords            []*ProjectCreateV30RequestKeywordsInner `json:"keywords,omitempty"`
	LandingPageStayTime *ProjectCreateV30LandingPageStayTime    `json:"landing_page_stay_time,omitempty"`
	LandingType         ProjectCreateV30LandingType             `json:"landing_type"`
	LaunchType          *ProjectCreateV30LaunchType             `json:"launch_type,omitempty"`
	MarketingGoal       ProjectCreateV30MarketingGoal           `json:"marketing_goal"`
	//
	MicroAppInstanceId *int64                              `json:"micro_app_instance_id,omitempty"`
	MicroPromotionType *ProjectCreateV30MicroPromotionType `json:"micro_promotion_type,omitempty"`
	MultiAssetType     *ProjectCreateV30MultiAssetType     `json:"multi_asset_type,omitempty"`
	//
	Name          string                                `json:"name"`
	NativeSetting *ProjectCreateV30RequestNativeSetting `json:"native_setting,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlField *string `json:"open_url_field,omitempty"`
	//
	OpenUrlParams *string                              `json:"open_url_params,omitempty"`
	OpenUrlType   *ProjectCreateV30OpenUrlType         `json:"open_url_type,omitempty"`
	Operation     *ProjectCreateV30Operation           `json:"operation,omitempty"`
	OptimizeGoal  *ProjectCreateV30RequestOptimizeGoal `json:"optimize_goal,omitempty"`
	PromotionType *ProjectCreateV30PromotionType       `json:"promotion_type,omitempty"`
	//
	QuickAppId     *int64                                 `json:"quick_app_id,omitempty"`
	RelatedProduct *ProjectCreateV30RequestRelatedProduct `json:"related_product,omitempty"`
	// 出价系数
	SearchBidRatio                 *float64                                        `json:"search_bid_ratio,omitempty"`
	StarAutoDeliverySwitch         *ProjectCreateV30StarAutoDeliverySwitch         `json:"star_auto_delivery_switch,omitempty"`
	StarAutoMaterialAdditionSwitch *ProjectCreateV30StarAutoMaterialAdditionSwitch `json:"star_auto_material_addition_switch,omitempty"`
	// 星广联投二期任务id
	StarTaskId *int64 `json:"star_task_id,omitempty"`
	//
	SubscribeUrl    *string                                 `json:"subscribe_url,omitempty"`
	TrackUrlSetting *ProjectCreateV30RequestTrackUrlSetting `json:"track_url_setting,omitempty"`
	//
	UlinkUrl           *string                             `json:"ulink_url,omitempty"`
	UlinkUrlType       *ProjectCreateV30UlinkUrlType       `json:"ulink_url_type,omitempty"`
	ValueOptimizedType *ProjectCreateV30ValueOptimizedType `json:"value_optimized_type,omitempty"`
}
