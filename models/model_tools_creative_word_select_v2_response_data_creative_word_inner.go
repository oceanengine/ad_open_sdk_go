/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCreativeWordSelectV2ResponseDataCreativeWordInner struct for ToolsCreativeWordSelectV2ResponseDataCreativeWordInner
type ToolsCreativeWordSelectV2ResponseDataCreativeWordInner struct {
	ContentType *ToolsCreativeWordSelectV2DataCreativeWordContentType `json:"content_type,omitempty"`
	//
	CreativeWordId *int64 `json:"creative_word_id,omitempty"`
	//
	DefaultWord *string `json:"default_word,omitempty"`
	//
	MaxWordLen *int32 `json:"max_word_len,omitempty"`
	//
	MinWordLen *int32 `json:"min_word_len,omitempty"`
	//
	Name   *string                                          `json:"name,omitempty"`
	Status *ToolsCreativeWordSelectV2DataCreativeWordStatus `json:"status,omitempty"`
	//
	UserRate *float64 `json:"user_rate,omitempty"`
	//
	Words []string `json:"words,omitempty"`
}
