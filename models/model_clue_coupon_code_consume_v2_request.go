/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCodeConsumeV2Request struct for ClueCouponCodeConsumeV2Request
type ClueCouponCodeConsumeV2Request struct {
	//
	ActivityId *int64 `json:"activity_id,omitempty"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Code *string `json:"code,omitempty"`
	//
	CodeId *string `json:"code_id,omitempty"`
	//
	CouponId *int64                                 `json:"coupon_id,omitempty"`
	Employee ClueCouponCodeConsumeV2RequestEmployee `json:"employee"`
	//
	Extra map[string]interface{} `json:"extra,omitempty"`
	//
	ResourceId *int64 `json:"resource_id,omitempty"`
}
