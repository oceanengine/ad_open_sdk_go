/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexGetV2Filtering
type ToolsRubeexGetV2Filtering struct {
	//
	AuthorIds []int64 `json:"author_ids,omitempty"`
	//
	EndTime *string                           `json:"end_time,omitempty"`
	IsOwner *ToolsRubeexGetV2FilteringIsOwner `json:"is_owner,omitempty"`
	Order   *ToolsRubeexGetV2FilteringOrder   `json:"order,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
	//
	ProjectLifecycle []*ToolsRubeexGetV2FilteringProjectLifecycle `json:"project_lifecycle,omitempty"`
	//
	ProjectName *string `json:"project_name,omitempty"`
	//
	ProjectVersionLifecycle []*ToolsRubeexGetV2FilteringProjectVersionLifecycle `json:"project_version_lifecycle,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
