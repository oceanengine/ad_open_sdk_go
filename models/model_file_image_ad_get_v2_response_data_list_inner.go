/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageAdGetV2ResponseDataListInner struct for FileImageAdGetV2ResponseDataListInner
type FileImageAdGetV2ResponseDataListInner struct {
	// 图片格式
	Format *string `json:"format,omitempty"`
	// 图片高度
	Height *int64 `json:"height,omitempty"`
	// 图片ID
	Id *string `json:"id,omitempty"`
	// 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	MaterialId *int64 `json:"material_id,omitempty"`
	// 图片md5
	Signature *string `json:"signature,omitempty"`
	// 图片大小
	Size *int64 `json:"size,omitempty"`
	// 图片预览地址
	Url *string `json:"url,omitempty"`
	// 图片宽度
	Width *int64 `json:"width,omitempty"`
}
