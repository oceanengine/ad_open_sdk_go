/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2ResponseData
type ClueCouponDetailV2ResponseData struct {
	//
	ActivityId   *int64                                `json:"activity_id,omitempty"`
	ActivityType *ClueCouponDetailV2DataActivityType   `json:"activity_type,omitempty"`
	Coupon       *ClueCouponDetailV2ResponseDataCoupon `json:"coupon,omitempty"`
	Form         *ClueCouponDetailV2ResponseDataForm   `json:"form,omitempty"`
}
