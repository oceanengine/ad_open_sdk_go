/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAvailableGetV10Filter
type QianchuanProductAvailableGetV10Filter struct {
	// 是否托管
	IsAdlab        *bool                                                `json:"is_adlab,omitempty"`
	MarketingScene *QianchuanProductAvailableGetV10FilterMarketingScene `json:"marketing_scene,omitempty"`
	//
	ProductIds []int64 `json:"product_ids,omitempty"`
	//
	ProductName *string                                   `json:"product_name,omitempty"`
	Tab         *QianchuanProductAvailableGetV10FilterTab `json:"tab,omitempty"`
}
