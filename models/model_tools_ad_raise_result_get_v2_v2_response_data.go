/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdRaiseResultGetV2V2ResponseData
type ToolsAdRaiseResultGetV2V2ResponseData struct {
	//
	AdRaiseReport []*ToolsAdRaiseResultGetV2V2ResponseDataAdRaiseReportInner `json:"ad_raise_report,omitempty"`
	//
	AdRaiseVersion *int64                                       `json:"ad_raise_version,omitempty"`
	PageInfo       *AgentAdvertiserSelectV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
