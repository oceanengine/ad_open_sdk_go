/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdStatusUpdateV10Request struct for QianchuanUniPromotionAdStatusUpdateV10Request
type QianchuanUniPromotionAdStatusUpdateV10Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64                                           `json:"advertiser_id"`
	OptStatus    QianchuanUniPromotionAdStatusUpdateV10OptStatus `json:"opt_status"`
}
