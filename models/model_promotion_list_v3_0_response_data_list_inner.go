/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInner struct for PromotionListV30ResponseDataListInner
type PromotionListV30ResponseDataListInner struct {
	// 七日内留存天数
	Var7dRetention   *float64                                  `json:"7d_retention,omitempty"`
	AdDownloadStatus *PromotionListV30DataListAdDownloadStatus `json:"ad_download_status,omitempty"`
	//
	AdvertiserId      *int64                                     `json:"advertiser_id,omitempty"`
	AutoExtendTraffic *PromotionListV30DataListAutoExtendTraffic `json:"auto_extend_traffic,omitempty"`
	//
	Bid             *float64                                              `json:"bid,omitempty"`
	BlueFlowPackage *PromotionListV30ResponseDataListInnerBlueFlowPackage `json:"blue_flow_package,omitempty"`
	BrandInfo       *PromotionListV30ResponseDataListInnerBrandInfo       `json:"brand_info,omitempty"`
	//
	Budget     *float64                            `json:"budget,omitempty"`
	BudgetMode *PromotionListV30DataListBudgetMode `json:"budget_mode,omitempty"`
	//
	ConfigId *int64 `json:"config_id,omitempty"`
	//
	CpaBid                     *float64                                            `json:"cpa_bid,omitempty"`
	CreativeAutoGenerateSwitch *PromotionListV30DataListCreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DeepCpabid       *float64                                  `json:"deep_cpabid,omitempty"`
	HasCarryMaterial *PromotionListV30DataListHasCarryMaterial `json:"has_carry_material,omitempty"`
	IsCommentDisable *PromotionListV30DataListIsCommentDisable `json:"is_comment_disable,omitempty"`
	// 搜索广告关键词
	Keywords          []*PromotionListV30ResponseDataListInnerKeywordsInner   `json:"keywords,omitempty"`
	LearningPhase     *PromotionListV30DataListLearningPhase                  `json:"learning_phase,omitempty"`
	MaterialScoreInfo *PromotionListV30ResponseDataListInnerMaterialScoreInfo `json:"material_score_info,omitempty"`
	MaterialsType     *PromotionListV30DataListMaterialsType                  `json:"materials_type,omitempty"`
	NativeSetting     *PromotionListV30ResponseDataListInnerNativeSetting     `json:"native_setting,omitempty"`
	OptStatus         *PromotionListV30DataListOptStatus                      `json:"opt_status,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
	//
	PromotionCreateTime *string `json:"promotion_create_time,omitempty"`
	//
	PromotionId        *int64                                                   `json:"promotion_id,omitempty"`
	PromotionMaterials *PromotionListV30ResponseDataListInnerPromotionMaterials `json:"promotion_materials,omitempty"`
	//
	PromotionModifyTime *string `json:"promotion_modify_time,omitempty"`
	//
	PromotionName *string `json:"promotion_name,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleTime *string `json:"schedule_time,omitempty"`
	//
	ShopMultiRoiGoals []*PromotionListV30ResponseDataListInnerShopMultiRoiGoalsInner `json:"shop_multi_roi_goals,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	StarTaskId  *int64                               `json:"star_task_id,omitempty"`
	Status      *PromotionListV30DataListStatus      `json:"status,omitempty"`
	StatusFirst *PromotionListV30DataListStatusFirst `json:"status_first,omitempty"`
	//
	StatusSecond []*PromotionListV30DataListStatusSecond `json:"status_second,omitempty"`
	//
	UnionBid *float64 `json:"union_bid,omitempty"`
	//
	UnionBidRatio *float64 `json:"union_bid_ratio,omitempty"`
	//
	UnionCpaBid *float64 `json:"union_cpa_bid,omitempty"`
	//
	UnionDeepCpabid *float64 `json:"union_deep_cpabid,omitempty"`
	//
	UnionRoiGoal *float64 `json:"union_roi_goal,omitempty"`
}
