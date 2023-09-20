/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameCreateV30RequestAnchorList
type ToolsWechatGameCreateV30RequestAnchorList struct {
	//
	GuideText string `json:"guide_text"`
	//
	HeaderImageUrl string `json:"header_image_url"`
	//
	IconImageUrl string `json:"icon_image_url"`
	//
	ImagesHorizontalUrl []string `json:"images_horizontal_url"`
	//
	ImagesVerticalUrl []string `json:"images_vertical_url"`
	//
	Introduction string `json:"introduction"`
	//
	Labels []string `json:"labels"`
}
