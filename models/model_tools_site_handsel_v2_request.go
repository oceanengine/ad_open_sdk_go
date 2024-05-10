/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteHandselV2Request struct for ToolsSiteHandselV2Request
type ToolsSiteHandselV2Request struct {
	// 广告主id，范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserId int64 `json:"advertiser_id"`
	// 站点id列表, min_size=1, max_size=20
	SiteIds []string `json:"site_ids"`
	// 目标广告主id列表, min_size=1, max_size=20
	TargetAdvertiserIds []string `json:"target_advertiser_ids"`
}
