/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestCreativeInfoVideoMaterialsInner struct for AdlabGroupCreateV30RequestCreativeInfoVideoMaterialsInner
type AdlabGroupCreateV30RequestCreativeInfoVideoMaterialsInner struct {
	// 视频封面信息
	ImageInfo []*AdlabGroupCreateV30RequestCreativeInfoVideoMaterialsInnerImageInfoInner `json:"image_info"`
	ImageMode AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode                     `json:"image_mode"`
	VideoInfo AdlabGroupCreateV30RequestCreativeInfoVideoMaterialsInnerVideoInfo         `json:"video_info"`
}
