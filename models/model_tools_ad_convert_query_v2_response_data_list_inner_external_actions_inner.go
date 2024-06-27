/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdConvertQueryV2ResponseDataListInnerExternalActionsInner struct for ToolsAdConvertQueryV2ResponseDataListInnerExternalActionsInner
type ToolsAdConvertQueryV2ResponseDataListInnerExternalActionsInner struct {
	//
	ActionTrackUrl *string `json:"action_track_url,omitempty"`
	//
	Belong          []*ToolsAdConvertQueryV2DataListExternalActionsBelong        `json:"belong,omitempty"`
	ConvertDataType *ToolsAdConvertQueryV2DataListExternalActionsConvertDataType `json:"convert_data_type,omitempty"`
	//
	ConvertId *int64 `json:"convert_id,omitempty"`
	//
	DeepExternalActions []*ToolsAdConvertQueryV2ResponseDataListInnerExternalActionsInnerDeepExternalActionsInner `json:"deep_external_actions,omitempty"`
	//
	Disabled       *bool                                                       `json:"disabled,omitempty"`
	ExternalAction *ToolsAdConvertQueryV2DataListExternalActionsExternalAction `json:"external_action,omitempty"`
	//
	ExternalActionName *string `json:"external_action_name,omitempty"`
	//
	ExternalActions []*ToolsAdConvertQueryV2DataListExternalActionsExternalActions `json:"external_actions,omitempty"`
	//
	ExternalName *string                                             `json:"external_name,omitempty"`
	Source       *ToolsAdConvertQueryV2DataListExternalActionsSource `json:"source,omitempty"`
}
