/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseFlowCategoryGetV10ResponseData
type EnterpriseFlowCategoryGetV10ResponseData struct {
	//
	List []*EnterpriseFlowCategoryGetV10ResponseDataListInner `json:"list,omitempty"`
	//
	Offset *int64 `json:"offset,omitempty"`
	//
	TotalCount   *int64                                                `json:"total_count,omitempty"`
	TotalMetrics *EnterpriseFlowCategoryGetV10ResponseDataTotalMetrics `json:"total_metrics,omitempty"`
	TotalRatio   *EnterpriseFlowCategoryGetV10ResponseDataTotalRatio   `json:"total_ratio,omitempty"`
}
