/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectUpdateV30RequestProjectListInner struct for ToolsPrivativeWordProjectUpdateV30RequestProjectListInner
type ToolsPrivativeWordProjectUpdateV30RequestProjectListInner struct {
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
}
