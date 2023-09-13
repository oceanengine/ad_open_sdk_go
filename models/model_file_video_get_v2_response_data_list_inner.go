/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoGetV2ResponseDataListInner struct for FileVideoGetV2ResponseDataListInner
type FileVideoGetV2ResponseDataListInner struct {
	// 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	BitRate *int64 `json:"bit_rate,omitempty"`
	// 素材的上传时间，格式：\"yyyy-mm-dd HH:MM:SS\"
	CreateTime *string `json:"create_time,omitempty"`
	// 码率，单位bps
	Duration *float64 `json:"duration,omitempty"`
	// 素材的文件名
	Filename *string `json:"filename,omitempty"`
	// 视频格式
	Format *string `json:"format,omitempty"`
	// 视频高度
	Height *int64 `json:"height,omitempty"`
	// 视频ID
	Id *string `json:"id,omitempty"`
	// 视频标签
	Labels []string `json:"labels,omitempty"`
	// 视频时长
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	OrganizationTags []string `json:"organization_tags,omitempty"`
	// 视频首帧截图，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”
	PosterUrl *string `json:"poster_url,omitempty"`
	// 视频md5值
	Signature *string `json:"signature,omitempty"`
	// 视频大小
	Size *int64 `json:"size,omitempty"`
	// 素材来源，详见【附录-素材来源】
	Source *string `json:"source,omitempty"`
	// 视频地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL” 链接仅做预览使用，预览链接有效期为1小时
	Url *string `json:"url,omitempty"`
	// 视频宽度
	Width *int64 `json:"width,omitempty"`
}
