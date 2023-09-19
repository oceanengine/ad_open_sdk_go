/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdminInfoV2ResponseDataDistrictsInnerSubDistrictsInnerSubDistrictsInner struct for ToolsAdminInfoV2ResponseDataDistrictsInnerSubDistrictsInnerSubDistrictsInner
type ToolsAdminInfoV2ResponseDataDistrictsInnerSubDistrictsInnerSubDistrictsInner struct {
	// 中国大陆行政区域编码
	Code *string `json:"code,omitempty"`
	// 港澳台、国外行政区域编码
	GeonameId *int64                                                      `json:"geoname_id,omitempty"`
	Level     *ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel `json:"level,omitempty"`
	// 行政区域名称
	Name *string `json:"name,omitempty"`
	// 子行政层级信息
	SubDistricts []*ToolsAdminInfoV2ResponseDataDistrictsInnerSubDistrictsInnerSubDistrictsInnerSubDistrictsInner `json:"sub_districts,omitempty"`
}
