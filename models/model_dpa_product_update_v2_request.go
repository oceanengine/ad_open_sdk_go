/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductUpdateV2Request struct for DpaProductUpdateV2Request
type DpaProductUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	PlatformId int64 `json:"platform_id"`
	//
	ProductId   int64                                `json:"product_id"`
	ProductInfo DpaProductUpdateV2RequestProductInfo `json:"product_info"`
}
