/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileImageDeleteV10ResponseData
type QianchuanFileImageDeleteV10ResponseData struct {
	// 操作失败的image_id列表，不在此列表内的素材表示删除成功
	FailImageIds []string `json:"fail_image_ids"`
}
