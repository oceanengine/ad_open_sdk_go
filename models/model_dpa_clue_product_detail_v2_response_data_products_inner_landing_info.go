/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductDetailV2ResponseDataProductsInnerLandingInfo 落地页信息
type DpaClueProductDetailV2ResponseDataProductsInnerLandingInfo struct {
	// 落地页
	TargetUrl *string `json:"target_url,omitempty"`
	// 安卓落地页
	TargetUrlAndroidApp *string `json:"target_url_android_app,omitempty"`
	// iOS落地页
	TargetUrlIosApp *string `json:"target_url_ios_app,omitempty"`
	// H5落地页
	TargetUrlMobile *string `json:"target_url_mobile,omitempty"`
	// 落地页universal link
	TargetUrlUniversalLink *string `json:"target_url_universal_link,omitempty"`
}
