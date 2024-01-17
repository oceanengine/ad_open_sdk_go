/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInner struct for BrandAdGetV30ResponseDataAdsInner
type BrandAdGetV30ResponseDataAdsInner struct {
	// 计划ID
	AdId *string `json:"ad_id,omitempty"`
	// 计划状态
	AdStatus *int64 `json:"ad_status,omitempty"`
	// 流量位置信息
	AdStockInfos []*BrandAdGetV30ResponseDataAdsInnerAdStockInfosInner `json:"ad_stock_infos,omitempty"`
	// 广告主ID
	AdvertiserId *int64                                         `json:"advertiser_id,omitempty"`
	AudienceInfo *BrandAdGetV30ResponseDataAdsInnerAudienceInfo `json:"audience_info,omitempty"`
	BudgetInfo   *BrandAdGetV30ResponseDataAdsInnerBudgetInfo   `json:"budget_info,omitempty"`
	// 资源立项编号
	BudgetNo *string `json:"budget_no,omitempty"`
	// 广告组ID
	CampaignId *string                       `json:"campaign_id,omitempty"`
	Classify   *BrandAdGetV30DataAdsClassify `json:"classify,omitempty"`
	// 计划创建时间
	CreateTime              *string                                         `json:"create_time,omitempty"`
	EnableMerchantIntention *BrandAdGetV30DataAdsEnableMerchantIntention    `json:"enable_merchant_intention,omitempty"`
	FrequencyInfo           *BrandAdGetV30ResponseDataAdsInnerFrequencyInfo `json:"frequency_info,omitempty"`
	// 推广目的（枚举值）
	LandingType       *int64                                              `json:"landing_type,omitempty"`
	MagazinePriceInfo *BrandAdGetV30ResponseDataAdsInnerMagazinePriceInfo `json:"magazine_price_info,omitempty"`
	// 招商意向单ID
	MerchantIntentionId *int64 `json:"merchant_intention_id,omitempty"`
	// 招商意向单名称
	MerchantIntentionName *string `json:"merchant_intention_name,omitempty"`
	// 招商意向单编号
	MerchantIntentionNo *string `json:"merchant_intention_no,omitempty"`
	// 广告计划名称
	Name *string `json:"name,omitempty"`
	// 政策信息
	Promotions []*BrandAdGetV30ResponseDataAdsInnerPromotionsInner `json:"promotions,omitempty"`
	// 备注
	Remark       *string                                        `json:"remark,omitempty"`
	ScheduleInfo *BrandAdGetV30ResponseDataAdsInnerScheduleInfo `json:"schedule_info,omitempty"`
	// 商品ID
	SpuId *int64 `json:"spu_id,omitempty"`
}
