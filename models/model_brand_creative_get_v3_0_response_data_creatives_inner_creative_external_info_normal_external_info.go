/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeExternalInfoNormalExternalInfo
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeExternalInfoNormalExternalInfo struct {
	// 落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 小程序链接
	MicroAppInfo *string `json:"micro_app_info,omitempty"`
	// 直达链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 直达文案
	OpenUrlText *string `json:"open_url_text,omitempty"`
}
