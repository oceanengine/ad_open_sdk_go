/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoCardImageInfo 图片信息
type BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoCardImageInfo struct {
	// 动图时长
	D *int64 `json:"d,omitempty"`
	// 格式
	Format *string `json:"format,omitempty"`
	//
	Hash *string `json:"hash,omitempty"`
	// 高
	Height *int64 `json:"height,omitempty"`
	// 图片URI
	WebUri *string `json:"web_uri,omitempty"`
	// 图片URL
	WebUriShowUrl *string `json:"web_uri_show_url,omitempty"`
	// 宽
	Width *int64 `json:"width,omitempty"`
}
