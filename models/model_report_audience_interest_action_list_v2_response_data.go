/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceInterestActionListV2ResponseData
type ReportAudienceInterestActionListV2ResponseData struct {
	//
	List     []*ReportAudienceAwemeListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *AgentAdvertiserSelectV2ResponseDataPageInfo      `json:"page_info,omitempty"`
}