/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10ResponseData
type QianchuanUniPromotionListV10ResponseData struct {
	//
	AdList   []*QianchuanUniPromotionListV10ResponseDataAdListInner `json:"ad_list"`
	PageInfo *QianchuanUniPromotionListV10ResponseDataPageInfo      `json:"page_info,omitempty"`
}
