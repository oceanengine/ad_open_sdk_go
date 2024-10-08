/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareStatsDataV10ResponseDataSimilarProductData
type QianchuanProductAnalyseCompareStatsDataV10ResponseDataSimilarProductData struct {
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	DirectOrderPayCost *float64 `json:"direct_order_pay_cost,omitempty"`
	//
	DirectOrderPayCostPerOrder *float64 `json:"direct_order_pay_cost_per_order,omitempty"`
	//
	DirectOrderPayCount *int64 `json:"direct_order_pay_count,omitempty"`
	//
	DirectOrderPayGmv *float64 `json:"direct_order_pay_gmv,omitempty"`
	//
	DirectOrderPayRate *float64 `json:"direct_order_pay_rate,omitempty"`
	//
	DirectOrderPrepayAndPayRoi *float64 `json:"direct_order_prepay_and_pay_roi,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
}
