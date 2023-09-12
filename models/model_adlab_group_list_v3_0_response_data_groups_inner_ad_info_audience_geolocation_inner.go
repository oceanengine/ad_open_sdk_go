/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudienceGeolocationInner struct for AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudienceGeolocationInner
type AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudienceGeolocationInner struct {
	// 纬度
	Lat *float64 `json:"lat,omitempty"`
	// 经度
	Long *float64 `json:"long,omitempty"`
	// 地点名称
	Name *string `json:"name,omitempty"`
	// 半径
	Radius *int64 `json:"radius,omitempty"`
}
