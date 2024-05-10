/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchorProductImage 服务主图
type NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchorProductImage struct {
	// 图片高度
	Height *int64 `json:"height,omitempty"`
	// 图片uri，大小不超过10M，服务主图以小图及大图形式分别展示
	Uri *string `json:"uri,omitempty"`
	// 图片宽度
	Width *int64 `json:"width,omitempty"`
}
