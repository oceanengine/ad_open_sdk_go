/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerBudgetInfoQuantityInfoInner struct for BrandAdGetV30ResponseDataAdsInnerBudgetInfoQuantityInfoInner
type BrandAdGetV30ResponseDataAdsInnerBudgetInfoQuantityInfoInner struct {
	// 上传商圈Id
	BusinessId *int64 `json:"business_id,omitempty"`
	// 城市
	City *int64 `json:"city,omitempty"`
	// 日期,格式\"2020-01-01\"
	Date        *string                                                                  `json:"date,omitempty"`
	Geolocation *BrandAdGetV30ResponseDataAdsInnerBudgetInfoQuantityInfoInnerGeolocation `json:"geolocation,omitempty"`
	// 预定量
	Quantity *int64 `json:"quantity,omitempty"`
	// 预定类型
	QuantityType *int64 `json:"quantity_type,omitempty"`
}
