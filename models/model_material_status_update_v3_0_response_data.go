/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// MaterialStatusUpdateV30ResponseData
type MaterialStatusUpdateV30ResponseData struct {
	//
	Errors []*MaterialStatusUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	MaterialIds []int64 `json:"material_ids,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
