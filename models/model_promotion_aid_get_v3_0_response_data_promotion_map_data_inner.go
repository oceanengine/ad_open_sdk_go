/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAidGetV30ResponseDataPromotionMapDataInner struct for PromotionAidGetV30ResponseDataPromotionMapDataInner
type PromotionAidGetV30ResponseDataPromotionMapDataInner struct {
	//
	Ads []*PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInner `json:"ads,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
