/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerMagazinePriceInfo 刊例价信息
type BrandAdGetV30ResponseDataAdsInnerMagazinePriceInfo struct {
	// 实际投放刊例
	ActualMagazinePrice *string `json:"actual_magazine_price,omitempty"`
	// 政策折扣
	Discount *string `json:"discount,omitempty"`
	// 刊例单价
	PublishPrice *string `json:"publish_price,omitempty"`
	// 政策单价
	SellPrice *string `json:"sell_price,omitempty"`
	// 刊例总价
	TotalPublishPrice *string `json:"total_publish_price,omitempty"`
	// 政策总价
	TotalSellPrice *string `json:"total_sell_price,omitempty"`
}
