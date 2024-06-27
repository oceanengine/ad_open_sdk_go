/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCopyV2Request struct for ToolsSiteCopyV2Request
type ToolsSiteCopyV2Request struct {
	// 广告主id，数字范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserId int64 `json:"advertiser_id"`
	// 站点id列表, min_size=1, max_size=20
	SiteIds []string `json:"site_ids"`
}
