/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30RequestCreativeInfoVideoMaterialsInner struct for AdlabGroupUpdateV30RequestCreativeInfoVideoMaterialsInner
type AdlabGroupUpdateV30RequestCreativeInfoVideoMaterialsInner struct {
	// 视频封面信息
	ImageInfo []*AdlabGroupUpdateV30RequestCreativeInfoVideoMaterialsInnerImageInfoInner `json:"image_info"`
	ImageMode AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode                     `json:"image_mode"`
	VideoInfo AdlabGroupUpdateV30RequestCreativeInfoVideoMaterialsInnerVideoInfo         `json:"video_info"`
}
