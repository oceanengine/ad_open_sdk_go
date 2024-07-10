/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoPauseV2Request struct for FileVideoPauseV2Request
type FileVideoPauseV2Request struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 待清理素材id列表
	MaterialIds []int64 `json:"material_ids,omitempty"`
}
