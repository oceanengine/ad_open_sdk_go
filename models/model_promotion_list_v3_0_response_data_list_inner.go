/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInner struct for PromotionListV30ResponseDataListInner
type PromotionListV30ResponseDataListInner struct {
	AdDownloadStatus *PromotionListV30DataListAdDownloadStatus `json:"ad_download_status,omitempty"`
	//
	AdvertiserId      *int64                                     `json:"advertiser_id,omitempty"`
	AutoExtendTraffic *PromotionListV30DataListAutoExtendTraffic `json:"auto_extend_traffic,omitempty"`
	//
	Bid       *float32                                        `json:"bid,omitempty"`
	BrandInfo *PromotionListV30ResponseDataListInnerBrandInfo `json:"brand_info,omitempty"`
	//
	Budget *float32 `json:"budget,omitempty"`
	//
	ConfigId *int64 `json:"config_id,omitempty"`
	//
	CpaBid                     *float32                                            `json:"cpa_bid,omitempty"`
	CreativeAutoGenerateSwitch *PromotionListV30DataListCreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DeepCpabid       *float32                                  `json:"deep_cpabid,omitempty"`
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
	RoiGoal *float32 `json:"roi_goal,omitempty"`
	//
	Source *string                         `json:"source,omitempty"`
	Status *PromotionListV30DataListStatus `json:"status,omitempty"`
}
