/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportUniPromotionDimensionDataAuthorGetV10ResponseDataListInner struct for QianchuanReportUniPromotionDimensionDataAuthorGetV10ResponseDataListInner
type QianchuanReportUniPromotionDimensionDataAuthorGetV10ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	StatDatetime *string `json:"stat_datetime,omitempty"`
	//
	TotalCostPerPayOrderForRoi2 *float64 `json:"total_cost_per_pay_order_for_roi2,omitempty"`
	//
	TotalPayOrderCountForRoi2 *float64 `json:"total_pay_order_count_for_roi2,omitempty"`
	//
	TotalPayOrderCouponAmountForRoi2 *float64 `json:"total_pay_order_coupon_amount_for_roi2,omitempty"`
	//
	TotalPayOrderGmvForRoi2 *float64 `json:"total_pay_order_gmv_for_roi2,omitempty"`
	//
	TotalPrepayAndPayOrderRoi2 *float64 `json:"total_prepay_and_pay_order_roi2,omitempty"`
	//
	TotalPrepayOrderCountForRoi2 *float64 `json:"total_prepay_order_count_for_roi2,omitempty"`
	//
	TotalPrepayOrderGmvForRoi2 *float64 `json:"total_prepay_order_gmv_for_roi2,omitempty"`
	//
	TotalUnfinishedEstimateOrderGmvForRoi2 *float64 `json:"total_unfinished_estimate_order_gmv_for_roi2,omitempty"`
}
