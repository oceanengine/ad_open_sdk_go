/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponUpdateV2Request struct for ClueCouponUpdateV2Request
type ClueCouponUpdateV2Request struct {
	//
	ActivityId int64 `json:"activity_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	DeliverEnd *NullableTime `json:"deliver_end,omitempty"`
	//
	DeliverStart *NullableTime                         `json:"deliver_start,omitempty"`
	GlobalLimit  *ClueCouponUpdateV2RequestGlobalLimit `json:"global_limit,omitempty"`
	Status       *ClueCouponUpdateV2Status             `json:"status,omitempty"`
	UserLimit    *ClueCouponUpdateV2RequestUserLimit   `json:"user_limit,omitempty"`
}
