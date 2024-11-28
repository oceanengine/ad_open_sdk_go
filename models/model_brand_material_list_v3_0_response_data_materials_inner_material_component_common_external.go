/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternal 普通落地页链接组件
type BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternal struct {
	CloseLoopInfo *BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternalCloseLoopInfo `json:"close_loop_info,omitempty"`
	// 落地页标题
	ExternalTitle *string `json:"external_title,omitempty"`
	// 落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 是否下载意图 0：否  1：是
	ExternalUrlIsDownload *int64 `json:"external_url_is_download,omitempty"`
	// 是否关闭评论 0：否  1：是
	IsCommentDisable *int64 `json:"is_comment_disable,omitempty"`
	// 是否以webview打开 0：否  1：是
	IsWebView    *int64                                                                                     `json:"is_web_view,omitempty"`
	MicroAppInfo *BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternalMicroAppInfo `json:"micro_app_info,omitempty"`
	// 小程序链接类型 0：默认(抖音) 1：微信
	MicroAppOpenUrlType *int64 `json:"micro_app_open_url_type,omitempty"`
	// 直达链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 直达链接文案
	OpenUrlText        *string                                                                                          `json:"open_url_text,omitempty"`
	WechatMicroAppInfo *BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternalWechatMicroAppInfo `json:"wechat_micro_app_info,omitempty"`
}
