/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsOrangeSiteGetV30ResponseData
type ToolsOrangeSiteGetV30ResponseData struct {
	//
	List         []*ToolsOrangeSiteGetV30ResponseDataListInner  `json:"list,omitempty"`
	OptimizeGoal *ToolsOrangeSiteGetV30ResponseDataOptimizeGoal `json:"optimize_goal,omitempty"`
	PageInfo     *ToolsOrangeSiteGetV30ResponseDataPageInfo     `json:"page_info,omitempty"`
}
