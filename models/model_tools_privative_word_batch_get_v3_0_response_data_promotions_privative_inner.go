/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordBatchGetV30ResponseDataPromotionsPrivativeInner struct for ToolsPrivativeWordBatchGetV30ResponseDataPromotionsPrivativeInner
type ToolsPrivativeWordBatchGetV30ResponseDataPromotionsPrivativeInner struct {
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
