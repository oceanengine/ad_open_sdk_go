/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeDeleteV2Request struct for ClueCouponEmployeeDeleteV2Request
type ClueCouponEmployeeDeleteV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	List []*ClueCouponEmployeeDeleteV2RequestListInner `json:"list"`
}
