/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30Request struct for PromotionCreateV30Request
type PromotionCreateV30Request struct {
	// 七日内留存天数，单位：天，取值范围(0,7)，仅支持最多2位小数。仅当满足以下条件时可设置： - landing_type = APP 应用推广 - ad_type = ALL 通投 - delivery_mode = MANUAL  手动投放 - external_action = AD_CONVERT_TYPE_ACTIVE 优化目标=激活 - deep_external_action = AD_CONVERT_TYPE_RETENTION_DAYS深度优化目标 = 留存天数 - delivery_setting.deep_bid_type = AD_CONVERT_TYPE_RETENTION_DAYS深度优化方式 = 留存天数 - delivery_range.inventory_catalog = MANUAL  广告位大类 = 首选媒体 - inventory_type = INVENTORY_UNION_SLOT  投放位置 只选择穿山甲
	Var7dRetention   *float64                            `json:"7d_retention,omitempty"`
	AdDownloadStatus *PromotionCreateV30AdDownloadStatus `json:"ad_download_status,omitempty"`
	//
	AdvertiserId      int64                                `json:"advertiser_id"`
	AutoExtendTraffic *PromotionCreateV30AutoExtendTraffic `json:"auto_extend_traffic,omitempty"`
	//
	Bid       *float64                            `json:"bid,omitempty"`
	BrandInfo *PromotionCreateV30RequestBrandInfo `json:"brand_info,omitempty"`
	//
	Budget     *float64                      `json:"budget,omitempty"`
	BudgetMode *PromotionCreateV30BudgetMode `json:"budget_mode,omitempty"`
	//
	ConfigId *int64 `json:"config_id,omitempty"`
	//
	CpaBid                     *float64                                      `json:"cpa_bid,omitempty"`
	CreativeAutoGenerateSwitch *PromotionCreateV30CreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DeepCpabid       *float64                            `json:"deep_cpabid,omitempty"`
	IsCommentDisable *PromotionCreateV30IsCommentDisable `json:"is_comment_disable,omitempty"`
	//
	Keywords      []*PromotionCreateV30RequestKeywordsInner `json:"keywords,omitempty"`
	MaterialsType *PromotionCreateV30MaterialsType          `json:"materials_type,omitempty"`
	//
	Name          string                                  `json:"name"`
	NativeSetting *PromotionCreateV30RequestNativeSetting `json:"native_setting,omitempty"`
	Operation     *PromotionCreateV30Operation            `json:"operation,omitempty"`
	//
	ProjectId          int64                                       `json:"project_id"`
	PromotionMaterials PromotionCreateV30RequestPromotionMaterials `json:"promotion_materials"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ShopMultiRoiGoals []*PromotionCreateV30RequestShopMultiRoiGoalsInner `json:"shop_multi_roi_goals,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	UnionBidRatio *float64 `json:"union_bid_ratio,omitempty"`
}
