/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataDataInner struct for ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataDataInner
type ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataDataInner struct {
	// 广告ID
	PromotionId *int64                                                                   `json:"promotion_id,omitempty"`
	RaiseInfo   *ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataDataInnerRaiseInfo `json:"raise_info,omitempty"`
	Status      *ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus                 `json:"status,omitempty"`
}
