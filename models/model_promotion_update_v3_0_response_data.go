/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30ResponseData
type PromotionUpdateV30ResponseData struct {
	//
	ErrorList []*PromotionUpdateV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
