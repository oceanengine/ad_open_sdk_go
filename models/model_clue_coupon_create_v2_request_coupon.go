/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCreateV2RequestCoupon
type ClueCouponCreateV2RequestCoupon struct {
	//
	DeliverEnd *string `json:"deliver_end"`
	//
	DeliverStart  *string                                     `json:"deliver_start"`
	GlobalLimit   *ClueCouponCreateV2RequestCouponGlobalLimit `json:"global_limit,omitempty"`
	NeedPhone     *ClueCouponCreateV2CouponNeedPhone          `json:"need_phone,omitempty"`
	NeedSmsVerify *ClueCouponCreateV2CouponNeedSmsVerify      `json:"need_sms_verify,omitempty"`
	//
	ResourceList []*ClueCouponCreateV2RequestCouponResourceListInner `json:"resource_list"`
	//
	Title     string                                    `json:"title"`
	UserLimit *ClueCouponCreateV2RequestCouponUserLimit `json:"user_limit,omitempty"`
}
