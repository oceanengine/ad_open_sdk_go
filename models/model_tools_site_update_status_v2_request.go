/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteUpdateStatusV2Request struct for ToolsSiteUpdateStatusV2Request
type ToolsSiteUpdateStatusV2Request struct {
	// 广告主id， 传的这个advertiser_id的数字的范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserId int64 `json:"advertiser_id"`
	// 橙子建站站点id列表 ：1 <= 长度 <= 20
	SiteIds []string `json:"site_ids"`
	// 站点操作状态
	Status string `json:"status"`
}
