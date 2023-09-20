/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageGetV2ResponseDataListInner struct for FileImageGetV2ResponseDataListInner
type FileImageGetV2ResponseDataListInner struct {
	// 素材是否是aigc生成
	Aigc *bool `json:"aigc,omitempty"`
	// 素材的上传时间，格式：\"yyyy-mm-dd HH:MM:SS\"
	CreateTime *string `json:"create_time,omitempty"`
	// 素材的文件名
	Filename *string `json:"filename,omitempty"`
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
	// 图片预览地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL” 链接仅做预览使用，预览链接有效期为1小时
	Url *string `json:"url,omitempty"`
	// 图片宽度
	Width *int64 `json:"width,omitempty"`
}
