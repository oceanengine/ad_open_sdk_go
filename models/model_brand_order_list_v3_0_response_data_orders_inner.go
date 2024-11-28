/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30ResponseDataOrdersInner struct for BrandOrderListV30ResponseDataOrdersInner
type BrandOrderListV30ResponseDataOrdersInner struct {
	AdForm *BrandOrderListV30DataOrdersAdForm `json:"ad_form,omitempty"`
	// 计划信息
	AdInfos []*BrandOrderListV30ResponseDataOrdersInnerAdInfosInner `json:"ad_infos,omitempty"`
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 端
	AppOrigin    []*BrandOrderListV30DataOrdersAppOrigin               `json:"app_origin,omitempty"`
	AudienceInfo *BrandOrderListV30ResponseDataOrdersInnerAudienceInfo `json:"audience_info,omitempty"`
	AuditStatus  *BrandOrderListV30DataOrdersAuditStatus               `json:"audit_status,omitempty"`
	BudgetInfo   *BrandOrderListV30ResponseDataOrdersInnerBudgetInfo   `json:"budget_info,omitempty"`
	// 广告组ID
	CampaignId *int64                               `json:"campaign_id,omitempty"`
	Classify   *BrandOrderListV30DataOrdersClassify `json:"classify,omitempty"`
	// 预订单创建时间
	CreateTime            *string                                                        `json:"create_time,omitempty"`
	FrequencyInfo         *BrandOrderListV30ResponseDataOrdersInnerFrequencyInfo         `json:"frequency_info,omitempty"`
	GdSendType            *BrandOrderListV30DataOrdersGdSendType                         `json:"gd_send_type,omitempty"`
	MagazinePriceInfo     *BrandOrderListV30ResponseDataOrdersInnerMagazinePriceInfo     `json:"magazine_price_info,omitempty"`
	MerchantIntentionInfo *BrandOrderListV30ResponseDataOrdersInnerMerchantIntentionInfo `json:"merchant_intention_info,omitempty"`
	// 预订单名称
	Name *string `json:"name,omitempty"`
	// 预订单ID
	OrderId *int64 `json:"order_id,omitempty"`
	// 出价方式 3：CPT  6：GD
	PricingType *int64 `json:"pricing_type,omitempty"`
	// 产品类型
	ProType *int64 `json:"pro_type,omitempty"`
	// 政策
	Promptions []*BrandOrderListV30ResponseDataOrdersInnerPromptionsInner `json:"promptions,omitempty"`
	// 备注
	Remark               *string                                               `json:"remark,omitempty"`
	ScheduleInfo         *BrandOrderListV30ResponseDataOrdersInnerScheduleInfo `json:"schedule_info,omitempty"`
	Status               *BrandOrderListV30DataOrdersStatus                    `json:"status,omitempty"`
	StockIncreasePackage *BrandOrderListV30DataOrdersStockIncreasePackage      `json:"stock_increase_package,omitempty"`
}
