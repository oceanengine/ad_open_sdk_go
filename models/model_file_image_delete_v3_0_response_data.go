/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageDeleteV30ResponseData
type FileImageDeleteV30ResponseData struct {
	// 操作失败的图片ID列表，不在此列表内的素材表示删除成功
	FailImageIds []string `json:"fail_image_ids,omitempty"`
}
