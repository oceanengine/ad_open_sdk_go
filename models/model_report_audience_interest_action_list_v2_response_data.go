/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceInterestActionListV2ResponseData
type ReportAudienceInterestActionListV2ResponseData struct {
	//
	List     []*ReportAudienceInterestActionListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportAudienceInterestActionListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
