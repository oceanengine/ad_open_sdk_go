/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerGameAnchor 配置平台（1:不限,2:安卓,3:iOS）
type NativeAnchorGetDetailV30ResponseDataListInnerGameAnchor struct {
	//
	AnchorImageMode *int64 `json:"anchor_image_mode,omitempty"`
	//
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	//
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	//
	AppImages []*NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorAppImagesInner `json:"app_images,omitempty"`
	//
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	//
	AppTags []string `json:"app_tags,omitempty"`
	//
	GameBonus *bool `json:"game_bonus,omitempty"`
	//
	GameCharatoristic *string `json:"game_charatoristic,omitempty"`
	//
	GameDescription *string `json:"game_description,omitempty"`
	//
	GamePackageList []*NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInner `json:"game_package_list,omitempty"`
	GameType        *NativeAnchorGetDetailV30DataListGameAnchorGameType                            `json:"game_type,omitempty"`
	//
	GuideText *string `json:"guide_text,omitempty"`
	//
	HeadImageList []*NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorHeadImageListInner `json:"head_image_list,omitempty"`
	//
	IconImages []*NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorIconImagesInner `json:"icon_images,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	IosAnchorTitle *string `json:"ios_anchor_title,omitempty"`
	//
	IosDownloadUrl *string `json:"ios_download_url,omitempty"`
	//
	OtherDescription *string `json:"other_description,omitempty"`
	//
	PathParam *string `json:"path_param,omitempty"`
	// 配置平台（1:不限,2:安卓,3:iOS）
	PlatformType *int64 `json:"platform_type,omitempty"`
}