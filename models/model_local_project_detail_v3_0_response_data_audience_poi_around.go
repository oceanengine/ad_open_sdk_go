/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectDetailV30ResponseDataAudiencePoiAround 门店附近定向设置
type LocalProjectDetailV30ResponseDataAudiencePoiAround struct {
	// 定向门店ids
	PoiAroundIds    []int64                                                    `json:"poi_around_ids,omitempty"`
	PoiAroundRadius *LocalProjectDetailV30DataAudiencePoiAroundPoiAroundRadius `json:"poi_around_radius,omitempty"`
}
