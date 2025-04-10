/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdDetailV10ResponseDataDeliverySetting
type QianchuanUniPromotionAdDetailV10ResponseDataDeliverySetting struct {
	//
	Budget     *float64                                                       `json:"budget,omitempty"`
	BudgetMode *QianchuanUniPromotionAdDetailV10DataDeliverySettingBudgetMode `json:"budget_mode,omitempty"`
	//
	DailyDeliveryTime  *float64                                                               `json:"daily_delivery_time,omitempty"`
	DeepBidType        *QianchuanUniPromotionAdDetailV10DataDeliverySettingDeepBidType        `json:"deep_bid_type,omitempty"`
	DeepExternalAction *QianchuanUniPromotionAdDetailV10DataDeliverySettingDeepExternalAction `json:"deep_external_action,omitempty"`
	//
	EndTime          *string                                                              `json:"end_time,omitempty"`
	ExternalAction   *QianchuanUniPromotionAdDetailV10DataDeliverySettingExternalAction   `json:"external_action,omitempty"`
	LiveScheduleType *QianchuanUniPromotionAdDetailV10DataDeliverySettingLiveScheduleType `json:"live_schedule_type,omitempty"`
	PricingType      *QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType      `json:"pricing_type,omitempty"`
	QcpxMode         *QianchuanUniPromotionAdDetailV10DataDeliverySettingQcpxMode         `json:"qcpx_mode,omitempty"`
	//
	Roi2Goal     *float64                                                         `json:"roi2_goal,omitempty"`
	SmartBidType *QianchuanUniPromotionAdDetailV10DataDeliverySettingSmartBidType `json:"smart_bid_type,omitempty"`
	//
	StartTime         *string                                                               `json:"start_time,omitempty"`
	VideoScheduleType *QianchuanUniPromotionAdDetailV10DataDeliverySettingVideoScheduleType `json:"video_schedule_type,omitempty"`
}
