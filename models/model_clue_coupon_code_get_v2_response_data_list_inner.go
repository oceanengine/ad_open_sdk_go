/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCodeGetV2ResponseDataListInner struct for ClueCouponCodeGetV2ResponseDataListInner
type ClueCouponCodeGetV2ResponseDataListInner struct {
	//
	ActivityId *int64 `json:"activity_id,omitempty"`
	//
	Code *string `json:"code,omitempty"`
	//
	CodeId *string `json:"code_id,omitempty"`
	//
	CouponId *int64 `json:"coupon_id,omitempty"`
	//
	CreateTime **string `json:"create_time,omitempty"`
	//
	ResourceId *int64 `json:"resource_id,omitempty"`
	//
	ResourceTitle *string                            `json:"resource_title,omitempty"`
	Status        *ClueCouponCodeGetV2DataListStatus `json:"status,omitempty"`
	//
	UpdateTime **string `json:"update_time,omitempty"`
	//
	ValidEnd **string `json:"valid_end,omitempty"`
	//
	ValidStart **string `json:"valid_start,omitempty"`
}
