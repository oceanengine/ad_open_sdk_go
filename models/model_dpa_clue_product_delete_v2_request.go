/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductDeleteV2Request struct for DpaClueProductDeleteV2Request
type DpaClueProductDeleteV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 类目ID
	CategoryId        int64                                           `json:"category_id"`
	StoreIdAndOuterId *DpaClueProductDeleteV2RequestStoreIdAndOuterId `json:"store_id_and_outer_id,omitempty"`
}
