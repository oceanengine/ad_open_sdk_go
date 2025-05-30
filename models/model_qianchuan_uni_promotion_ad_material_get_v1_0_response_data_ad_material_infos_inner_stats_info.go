/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdMaterialGetV10ResponseDataAdMaterialInfosInnerStatsInfo
type QianchuanUniPromotionAdMaterialGetV10ResponseDataAdMaterialInfosInnerStatsInfo struct {
	// 整体点击次数
	ProductClickCountForRoi2 *int64 `json:"product_click_count_for_roi2,omitempty"`
	// 整体转化率
	ProductConvertRateForRoi2 *float64 `json:"product_convert_rate_for_roi2,omitempty"`
	// 整体点击率
	ProductCvrRateForRoi2 *float64 `json:"product_cvr_rate_for_roi2,omitempty"`
	// 整体展示次数
	ProductShowCountForRoi2 *int64 `json:"product_show_count_for_roi2,omitempty"`
	// 整体消耗，单位元
	StatCostForRoi2 *float64 `json:"stat_cost_for_roi2,omitempty"`
	// 整体成交订单成本，单位元
	TotalCostPerPayOrderForRoi2 *float64 `json:"total_cost_per_pay_order_for_roi2,omitempty"`
	// 整体成交订单数
	TotalPayOrderCountForRoi2 *int64 `json:"total_pay_order_count_for_roi2,omitempty"`
	// 整体成交智能优惠券金额，单位元
	TotalPayOrderCouponAmountForRoi2 *float64 `json:"total_pay_order_coupon_amount_for_roi2,omitempty"`
	// 整体成交金额，单位元
	TotalPayOrderGmvForRoi2 *float64 `json:"total_pay_order_gmv_for_roi2,omitempty"`
	// 整体支付ROI
	TotalPrepayAndPayOrderRoi2 *float64 `json:"total_prepay_and_pay_order_roi2,omitempty"`
	// 整体未完结预售订单预估金额，单位元
	TotalUnfinishedEstimateOrderGmvForRoi2 *float64 `json:"total_unfinished_estimate_order_gmv_for_roi2,omitempty"`
}
