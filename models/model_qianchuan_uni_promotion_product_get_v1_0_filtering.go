/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionProductGetV10Filtering
type QianchuanUniPromotionProductGetV10Filtering struct {
	//
	ProductIds []int64 `json:"product_ids,omitempty"`
	//
	ProductName *string                                         `json:"product_name,omitempty"`
	Tab         *QianchuanUniPromotionProductGetV10FilteringTab `json:"tab,omitempty"`
}
