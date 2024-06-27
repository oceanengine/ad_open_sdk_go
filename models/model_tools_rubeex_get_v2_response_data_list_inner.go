/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexGetV2ResponseDataListInner struct for ToolsRubeexGetV2ResponseDataListInner
type ToolsRubeexGetV2ResponseDataListInner struct {
	//
	AuthorId *int64 `json:"author_id,omitempty"`
	//
	Cover *string `json:"cover,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	CustomFinishPlayPlayableRate *float64 `json:"custom_finish_play_playable_rate,omitempty"`
	//
	CustomStartPlayPlayableRate *float64 `json:"custom_start_play_playable_rate,omitempty"`
	//
	OriginId *int64 `json:"origin_id,omitempty"`
	//
	PlatformName *string `json:"platform_name,omitempty"`
	//
	PreviewUrl *string `json:"preview_url,omitempty"`
	//
	ProjectData *string `json:"project_data,omitempty"`
	//
	ProjectDescription *string `json:"project_description,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
	//
	ProjectLifecycle *string `json:"project_lifecycle,omitempty"`
	//
	ProjectName *string `json:"project_name,omitempty"`
	//
	ProjectStatus *string `json:"project_status,omitempty"`
	//
	SkinId *int64 `json:"skin_id,omitempty"`
	//
	TemplateId *int64 `json:"template_id,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
}
