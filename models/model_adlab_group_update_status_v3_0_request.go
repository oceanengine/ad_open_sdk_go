/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateStatusV30Request struct for AdlabGroupUpdateStatusV30Request
type AdlabGroupUpdateStatusV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目信息
	Data []*AdlabGroupUpdateStatusV30RequestDataInner `json:"data"`
}
