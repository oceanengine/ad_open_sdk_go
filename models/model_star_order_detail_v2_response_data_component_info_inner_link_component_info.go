/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderDetailV2ResponseDataComponentInfoInnerLinkComponentInfo
type StarOrderDetailV2ResponseDataComponentInfoInnerLinkComponentInfo struct {
	// 审核失败原因
	AuditRejectReason *string `json:"audit_reject_reason,omitempty"`
	// 组件ID
	ComponentId *int64 `json:"component_id,omitempty"`
	// 安卓下载链接
	DownloadUrlAndroid *string `json:"download_url_android,omitempty"`
	// iOS app 下载链接
	DownloadUrlIos *string `json:"download_url_ios,omitempty"`
	// 安卓落地页链接
	LandingUrlAndroid *string `json:"landing_url_android,omitempty"`
	// iOS落地页链接
	LandingUrlIos *string `json:"landing_url_ios,omitempty"`
	// 常规组件类别 1 = 落地页 2 = 短视频小程序 3 = 原生落地页 4 = 直播落地页 5 = 直播小程序 6 = 直播下载组件 7 = 直播引流组件
	LinkType *int64 `json:"link_type,omitempty"`
	// 组件状态 0 = 待送审 1 = 有效 2 =审核中 3 = 审核失败
	Status *int64 `json:"status,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
}
