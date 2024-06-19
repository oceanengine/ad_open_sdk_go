/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectUpdateV30ResponseDataErrorListInner struct for ToolsPrivativeWordProjectUpdateV30ResponseDataErrorListInner
type ToolsPrivativeWordProjectUpdateV30ResponseDataErrorListInner struct {
	//
	ErrorMessage *string `json:"error_message,omitempty"`
	//
	FailPhraseWords []*ToolsPrivativeWordProjectUpdateV30ResponseDataErrorListInnerFailPhraseWordsInner `json:"fail_phrase_words,omitempty"`
	//
	FailPreciseWords []*ToolsPrivativeWordProjectUpdateV30ResponseDataErrorListInnerFailPreciseWordsInner `json:"fail_precise_words,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
}
