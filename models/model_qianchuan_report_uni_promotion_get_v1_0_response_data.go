/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportUniPromotionGetV10ResponseData
type QianchuanReportUniPromotionGetV10ResponseData struct {
	// 整体消耗，单位元，小数点2位
	StatCost *float64 `json:"stat_cost,omitempty"`
	// 整体成交订单成本
	TotalCostPerPayOrderForRoi2 *float64 `json:"total_cost_per_pay_order_for_roi2,omitempty"`
	// 平台补贴金额，单位元，小数点2位
	TotalEcomPlatformSubsidyAmountForRoi2 *float64 `json:"total_ecom_platform_subsidy_amount_for_roi2,omitempty"`
	// 整体成交订单数
	TotalPayOrderCountForRoi2 *float64 `json:"total_pay_order_count_for_roi2,omitempty"`
	// 整体成交智能优惠券金额，单位元，小数点2位
	TotalPayOrderCouponAmountForRoi2 *float64 `json:"total_pay_order_coupon_amount_for_roi2,omitempty"`
	// 用户实际支付金额，单位元，小数点2位
	TotalPayOrderGmvForRoi2 *float64 `json:"total_pay_order_gmv_for_roi2,omitempty"`
	// 整体成交金额，单位元，小数点2位
	TotalPayOrderGmvIncludeCouponForRoi2 *float64 `json:"total_pay_order_gmv_include_coupon_for_roi2,omitempty"`
	// 整体支付ROI
	TotalPrepayAndPayOrderRoi2 *float64 `json:"total_prepay_and_pay_order_roi2,omitempty"`
	// 整体预售订单数 注意：仅支持直播全域
	TotalPrepayOrderCountForRoi2 *float64 `json:"total_prepay_order_count_for_roi2,omitempty"`
	// 整体预售订单金额，单位元，小数点2位 注意：仅支持直播全域
	TotalPrepayOrderGmvForRoi2 *float64 `json:"total_prepay_order_gmv_for_roi2,omitempty"`
	// 整体未完结预售订单预估金额，单位元，小数点2位
	TotalUnfinishedEstimateOrderGmvForRoi2 *float64 `json:"total_unfinished_estimate_order_gmv_for_roi2,omitempty"`
}
