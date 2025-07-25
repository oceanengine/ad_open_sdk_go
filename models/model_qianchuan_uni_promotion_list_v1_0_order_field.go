/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10OrderField
type QianchuanUniPromotionListV10OrderField string

// List of qianchuan_uni_promotion_list_v1.0_order_field
const (
	CREATE_TIME_QianchuanUniPromotionListV10OrderField                                  QianchuanUniPromotionListV10OrderField = "create_time"
	STAT_COST_QianchuanUniPromotionListV10OrderField                                    QianchuanUniPromotionListV10OrderField = "stat_cost"
	TOTAL_COST_PER_PAY_ORDER_FOR_ROI2_QianchuanUniPromotionListV10OrderField            QianchuanUniPromotionListV10OrderField = "total_cost_per_pay_order_for_roi2"
	TOTAL_ECOM_PLATFORM_SUBSIDY_AMOUNT_FOR_ROI2_QianchuanUniPromotionListV10OrderField  QianchuanUniPromotionListV10OrderField = "total_ecom_platform_subsidy_amount_for_roi2"
	TOTAL_PAY_ORDER_COUNT_FOR_ROI2_QianchuanUniPromotionListV10OrderField               QianchuanUniPromotionListV10OrderField = "total_pay_order_count_for_roi2"
	TOTAL_PAY_ORDER_COUPON_AMOUNT_FOR_ROI2_QianchuanUniPromotionListV10OrderField       QianchuanUniPromotionListV10OrderField = "total_pay_order_coupon_amount_for_roi2"
	TOTAL_PAY_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10OrderField                 QianchuanUniPromotionListV10OrderField = "total_pay_order_gmv_for_roi2"
	TOTAL_PAY_ORDER_GMV_INCLUDE_COUPON_FOR_ROI2_QianchuanUniPromotionListV10OrderField  QianchuanUniPromotionListV10OrderField = "total_pay_order_gmv_include_coupon_for_roi2"
	TOTAL_PREPAY_AND_PAY_ORDER_ROI2_QianchuanUniPromotionListV10OrderField              QianchuanUniPromotionListV10OrderField = "total_prepay_and_pay_order_roi2"
	TOTAL_PREPAY_ORDER_COUNT_FOR_ROI2_QianchuanUniPromotionListV10OrderField            QianchuanUniPromotionListV10OrderField = "total_prepay_order_count_for_roi2"
	TOTAL_PREPAY_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10OrderField              QianchuanUniPromotionListV10OrderField = "total_prepay_order_gmv_for_roi2"
	TOTAL_UNFINISHED_ESTIMATE_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10OrderField QianchuanUniPromotionListV10OrderField = "total_unfinished_estimate_order_gmv_for_roi2"
)

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_order_field value
func (v QianchuanUniPromotionListV10OrderField) Ptr() *QianchuanUniPromotionListV10OrderField {
	return &v
}
