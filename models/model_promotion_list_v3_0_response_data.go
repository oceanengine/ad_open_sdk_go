/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseData
type PromotionListV30ResponseData struct {
	//
	List     []*PromotionListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *PromotionListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
