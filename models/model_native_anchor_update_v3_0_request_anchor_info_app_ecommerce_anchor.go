/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchor 电商下载锚点，当anchor_type=APP_SHOP: 应用下载-电商锚点时必填
type NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchor struct {
	// 活动信息文案，长度：1～20
	ActivityInfo *string `json:"activity_info,omitempty"`
	// 安卓锚点入口标题字段，长度：1～12
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	// 安卓包名
	AndroidPkgName *string                                                          `json:"android_pkg_name,omitempty"`
	AppIcon        *NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorAppIcon `json:"app_icon,omitempty"`
	// 应用名称，长度：1～6
	AppName *string `json:"app_name,omitempty"`
	// 应用吊起链接
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	// 详情信息文案，长度：1～20
	DetailInfo *string `json:"detail_info,omitempty"`
	// APP下载引导文案，长度：1～12
	DownloadGuideText *string `json:"download_guide_text,omitempty"`
	// 跳转链接类型枚举，1：橙子建站，橙子建站落地页设置字段必填；2：第三方落地页，第三方落地页设置字段必填
	ExternalType *int64 `json:"external_type,omitempty"`
	// ios 锚点入口标题字段，长度：1～12
	IosAnchorTitle          *string                                                                          `json:"ios_anchor_title,omitempty"`
	OfficialActiBannerImage *NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorOfficialActiBannerImage `json:"official_acti_banner_image,omitempty"`
	// 官方活动描述详情，长度：1～15
	OfficialActiText *string                                                                 `json:"official_acti_text,omitempty"`
	OrangeSiteInfo   *NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorOrangeSiteInfo `json:"orange_site_info,omitempty"`
	ProductImage     *NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorProductImage   `json:"product_image,omitempty"`
	// 商品名称，长度1～10
	ProductName *string `json:"product_name,omitempty"`
	// 商品价格（整数，且单位：元）
	ProductPrice *float64 `json:"product_price,omitempty"`
	// 商品标签列表，每个标签长度：1～4，标签个数：1～3
	ProductTags []string `json:"product_tags,omitempty"`
	// 服务信息文案，长度：1～20
	ServiceInfo   *string                                                                `json:"service_info,omitempty"`
	ThirdSiteInfo *NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorThirdSiteInfo `json:"third_site_info,omitempty"`
}
