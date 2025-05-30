/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandToolQueryStockBalanceV30DeliveryInfoDateInfoInner struct for BrandToolQueryStockBalanceV30DeliveryInfoDateInfoInner
type BrandToolQueryStockBalanceV30DeliveryInfoDateInfoInner struct {
	// 投放日期，格式：yyyy-MM-dd，例：2023-09-08
	Date *string `json:"date,omitempty"`
	// 分时段信息，格式：[\"09:00\",\"24:00\"]
	Period []string `json:"period,omitempty"`
}
