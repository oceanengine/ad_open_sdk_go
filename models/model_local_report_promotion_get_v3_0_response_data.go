/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportPromotionGetV30ResponseData
type LocalReportPromotionGetV30ResponseData struct {
	PageInfo *LocalReportPromotionGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	PromotionList []*LocalReportPromotionGetV30ResponseDataPromotionListInner `json:"promotion_list,omitempty"`
}
