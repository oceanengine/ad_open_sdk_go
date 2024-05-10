/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeListInnerVideoMaterial 视频素材信息
type CreativeDetailGetV30ResponseDataCreativeListInnerVideoMaterial struct {
	// 抖音短视频ID
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	// 自定义商品库视频模板ID
	DpaVideoTaskIds      []string                                                                 `json:"dpa_video_task_ids,omitempty"`
	DpaVideoTemplateType *CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType   `json:"dpa_video_template_type,omitempty"`
	ImageInfo            *CreativeDetailGetV30ResponseDataCreativeListInnerVideoMaterialImageInfo `json:"image_info,omitempty"`
	VideoInfo            *CreativeDetailGetV30ResponseDataCreativeListInnerVideoMaterialVideoInfo `json:"video_info,omitempty"`
}
