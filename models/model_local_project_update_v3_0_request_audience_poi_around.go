/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectUpdateV30RequestAudiencePoiAround 门店附近设置
type LocalProjectUpdateV30RequestAudiencePoiAround struct {
	//
	PoiAroundIds    []int64                                               `json:"poi_around_ids,omitempty"`
	PoiAroundRadius LocalProjectUpdateV30AudiencePoiAroundPoiAroundRadius `json:"poi_around_radius"`
}
