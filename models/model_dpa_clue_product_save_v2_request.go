/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductSaveV2Request struct for DpaClueProductSaveV2Request
type DpaClueProductSaveV2Request struct {
	// 广告主ID
	AdvertiserId int64                              `json:"advertiser_id"`
	Product      DpaClueProductSaveV2RequestProduct `json:"product"`
}
