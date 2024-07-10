/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeMaterialMetricsV2ResponseData
type CreativeMaterialMetricsV2ResponseData struct {
	// 广告主 ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 统计数据
	Stats *map[string]CreativeMaterialMetricsV2ResponseDataStatsValue `json:"stats,omitempty"`
}
