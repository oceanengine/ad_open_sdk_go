/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAvailableGetV10Filter
type QianchuanProductAvailableGetV10Filter struct {
	MarketingScene *QianchuanProductAvailableGetV10FilterMarketingScene `json:"marketing_scene,omitempty"`
	//
	ProductIds []int64 `json:"product_ids,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
}
