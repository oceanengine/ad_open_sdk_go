/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2ResponseData
type ReportAdvertiserGetV2ResponseData struct {
	//
	List     []*ReportAdvertiserGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportAdvertiserGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
