/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerProductInfo
type QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerProductInfo struct {
	//
	CategoryName *string `json:"category_name,omitempty"`
	//
	DiscountPrice *float64 `json:"discount_price,omitempty"`
	//
	MarketPrice *float64 `json:"market_price,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	ProductImg []string `json:"product_img,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
	//
	ProductRate *float64 `json:"product_rate,omitempty"`
	//
	Recommendation *string `json:"recommendation,omitempty"`
	//
	SaleTime *int64 `json:"sale_time,omitempty"`
}
