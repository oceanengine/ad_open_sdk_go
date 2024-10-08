/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductStatusBatchUpdateV2Request struct for DpaProductStatusBatchUpdateV2Request
type DpaProductStatusBatchUpdateV2Request struct {
	// 广告主ID
	AdvertiserId int64                                  `json:"advertiser_id"`
	OptStatus    DpaProductStatusBatchUpdateV2OptStatus `json:"opt_status"`
	// 商品库ID
	PlatformId int64 `json:"platform_id"`
	// 商品ID列表
	ProductIds []int64 `json:"product_ids"`
}
