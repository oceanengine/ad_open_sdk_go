/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageDeleteV30Request struct for FileImageDeleteV30Request
type FileImageDeleteV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 待删除的图片ID列表，长度范围：1 ~ 100
	ImageIds []string `json:"image_ids,omitempty"`
}
