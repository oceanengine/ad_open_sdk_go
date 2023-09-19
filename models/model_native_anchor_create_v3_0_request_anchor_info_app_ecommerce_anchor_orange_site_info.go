/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoAppEcommerceAnchorOrangeSiteInfo 橙子建站落地页设置
type NativeAnchorCreateV30RequestAnchorInfoAppEcommerceAnchorOrangeSiteInfo struct {
	// 安卓端落地页链接，内部需要包含应用下载链接
	AndroidExternalUrl *string `json:"android_external_url,omitempty"`
	// iOS端落地页链接，内部需要包含应用下载链接
	IosExternalUrl *string `json:"ios_external_url,omitempty"`
}
