/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeUniPromotionOrderDetailV10ResponseDataCouponInfosInner struct for QianchuanAwemeUniPromotionOrderDetailV10ResponseDataCouponInfosInner
type QianchuanAwemeUniPromotionOrderDetailV10ResponseDataCouponInfosInner struct {
	// 单位元。多张券券，就是券的总金额
	CouponTotalAmount *float64 `json:"coupon_total_amount,omitempty"`
	// 优惠券类型
	CouponType *int64 `json:"coupon_type,omitempty"`
}
