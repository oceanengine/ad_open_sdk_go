/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfoNetServiceAnchor 网服下载锚点，当anchor_type=APP_INTERNET_SERVICE时必填
type NativeAnchorUpdateV30RequestAnchorInfoNetServiceAnchor struct {
	// APP图片比例，100：尺寸为2：1的横图，200：尺寸为3：5的竖图
	AnchorImageMode *int64 `json:"anchor_image_mode,omitempty"`
	// 安卓锚点入口标题字段
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	// 安卓下载链接，net_service_type为微信小程序场景下不用传入
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	// APP详情，1～100
	AppDescription *string `json:"app_description,omitempty"`
	// APP图片，图片个数 3～8
	AppImages []*NativeAnchorUpdateV30RequestAnchorInfoNetServiceAnchorAppImagesInner `json:"app_images,omitempty"`
	// app调起链接
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	// APP标签列表，每个标签长度：1～4；标签个数：1～3
	AppTags []string `json:"app_tags,omitempty"`
	// 引导文案，长度：1～15
	GuideText *string `json:"guide_text,omitempty"`
	// 锚点头部图片list，推荐尺寸为2：1的横图
	HeadImageList []*NativeAnchorUpdateV30RequestAnchorInfoNetServiceAnchorHeadImageListInner `json:"head_image_list,omitempty"`
	// 小程序图片list，当前锚点类型且net_service_type为'MICRO_APP' 必填，比例为1：1，最多一个
	IconImages []*NativeAnchorUpdateV30RequestAnchorInfoNetServiceAnchorIconImagesInner `json:"icon_images,omitempty"`
	// 微信小程序ID，当前锚点类型且net_service_type为MICRO_APP必填
	InstanceId *int64 `json:"instance_id,omitempty"`
	// iOS 锚点入口标题字段
	IosAnchorTitle *string `json:"ios_anchor_title,omitempty"`
	// iOS下载链接，net_service_type为微信小程序场景下不用传入
	IosDownloadUrl *string                                                        `json:"ios_download_url,omitempty"`
	NetServiceType *NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType `json:"net_service_type,omitempty"`
	// 小说章节预览，最多9999个字
	NovelChapter *string `json:"novel_chapter,omitempty"`
	// 微信小游戏/小程序路径参数
	PathParam *string `json:"path_param,omitempty"`
	// 配置平台，net_service_type为微信小程序场景下不用传入（1:不限,2:安卓,3:iOS）不限：安卓下载链接和iOS下载链接必填；安卓：安卓下载链接必填，iOS下载链接不填写；iOS：iOS下载链接必填
	PlatformType *int64 `json:"platform_type,omitempty"`
	// 快应用ID，当前锚点类型且net_service_type为QUICK_APP必填
	QuickAppId *int64 `json:"quick_app_id,omitempty"`
	// 微信外跳链接
	WechatExternalUrl *string `json:"wechat_external_url,omitempty"`
	// 微信包id
	WechatPackageId *int64 `json:"wechat_package_id,omitempty"`
}
