/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerNetServiceAnchor
type NativeAnchorGetDetailV30ResponseDataListInnerNetServiceAnchor struct {
	//
	AnchorImageMode *int64 `json:"anchor_image_mode,omitempty"`
	//
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	//
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	//
	AppDescription *string `json:"app_description,omitempty"`
	//
	AppImages []*NativeAnchorGetDetailV30ResponseDataListInnerNetServiceAnchorAppImagesInner `json:"app_images,omitempty"`
	//
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	//
	AppTags []string `json:"app_tags,omitempty"`
	//
	GuideText *string `json:"guide_text,omitempty"`
	//
	HeadImageList []*NativeAnchorGetDetailV30ResponseDataListInnerNetServiceAnchorHeadImageListInner `json:"head_image_list,omitempty"`
	//
	IconImages []*NativeAnchorGetDetailV30ResponseDataListInnerNetServiceAnchorIconImagesInner `json:"icon_images,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	IosAnchorTitle *string `json:"ios_anchor_title,omitempty"`
	//
	IosDownloadUrl *string                                                         `json:"ios_download_url,omitempty"`
	NetServiceType *NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType `json:"net_service_type,omitempty"`
	//
	NovelChapter *string `json:"novel_chapter,omitempty"`
	//
	PathParam *string `json:"path_param,omitempty"`
	// 配置平台（1:不限,2:安卓,3:iOS）
	PlatformType *int64 `json:"platform_type,omitempty"`
	//
	QuickAppId *int64 `json:"quick_app_id,omitempty"`
	//
	WechatExternalUrl *string `json:"wechat_external_url,omitempty"`
	//
	WechatPackageId *int64 `json:"wechat_package_id,omitempty"`
}
