/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10ResponseDataAdListInnerStatsInfo
type QianchuanUniPromotionListV10ResponseDataAdListInnerStatsInfo struct {
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	TotalCostPerPayOrderForRoi2 *float64 `json:"total_cost_per_pay_order_for_roi2,omitempty"`
	//
	TotalPayOrderCountForRoi2 *float64 `json:"total_pay_order_count_for_roi2,omitempty"`
	//
	TotalPayOrderGmvForRoi2 *int64 `json:"total_pay_order_gmv_for_roi2,omitempty"`
	//
	TotalPrepayAndPayOrderRoi2 *float64 `json:"total_prepay_and_pay_order_roi2,omitempty"`
	//
	TotalPrepayOrderCountForRoi2 *int64 `json:"total_prepay_order_count_for_roi2,omitempty"`
	//
	TotalPrepayOrderGmvForRoi2 *float64 `json:"total_prepay_order_gmv_for_roi2,omitempty"`
}
