/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCreateV2Request struct for ClueCouponCreateV2Request
type ClueCouponCreateV2Request struct {
	ActivityType ClueCouponCreateV2ActivityType `json:"activity_type"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	BindFormId *int64                           `json:"bind_form_id,omitempty"`
	Coupon     *ClueCouponCreateV2RequestCoupon `json:"coupon,omitempty"`
}
