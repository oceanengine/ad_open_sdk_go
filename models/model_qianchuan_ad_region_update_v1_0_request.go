/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRegionUpdateV10Request struct for QianchuanAdRegionUpdateV10Request
type QianchuanAdRegionUpdateV10Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	City                 []int64                                         `json:"city,omitempty"`
	District             *QianchuanAdRegionUpdateV10District             `json:"district,omitempty"`
	ElectricFenceRegion  *QianchuanAdRegionUpdateV10ElectricFenceRegion  `json:"electric_fence_region,omitempty"`
	ExcludeLimitedRegion *QianchuanAdRegionUpdateV10ExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	LocationType         *QianchuanAdRegionUpdateV10LocationType         `json:"location_type,omitempty"`
}
