/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdConvertQueryV2ResponseDataListInner struct for ToolsAdConvertQueryV2ResponseDataListInner
type ToolsAdConvertQueryV2ResponseDataListInner struct {
	ConvertType *ToolsAdConvertQueryV2DataListConvertType `json:"convert_type,omitempty"`
	//
	Disabled *bool `json:"disabled,omitempty"`
	//
	ExternalActions []*ToolsAdConvertQueryV2ResponseDataListInnerExternalActionsInner `json:"external_actions,omitempty"`
}
