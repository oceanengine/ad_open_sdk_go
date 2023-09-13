/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaDetailGetV2Filtering
type DpaDetailGetV2Filtering struct {
	//
	DpaCategories []int64                             `json:"dpa_categories,omitempty"`
	IsRecommend   *DpaDetailGetV2FilteringIsRecommend `json:"is_recommend,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	ProductName *string                        `json:"product_name,omitempty"`
	Status      *DpaDetailGetV2FilteringStatus `json:"status,omitempty"`
}
