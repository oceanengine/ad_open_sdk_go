/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialUpdateV30RequestMaterialListInnerMaterialComponentBusinessUrl
type BrandMaterialUpdateV30RequestMaterialListInnerMaterialComponentBusinessUrl struct {
	// 落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 是否是下载链接
	ExternalUrlIsDownload *int64 `json:"external_url_is_download,omitempty"`
	// 直达链接
	OpenUrl *string `json:"open_url,omitempty"`
}
