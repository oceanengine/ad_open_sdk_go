/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentUpdateCommonComponentV2RequestCommonComponentAnchorComponent 小程序等需要在组件管理管理的anchor
type StarComponentUpdateCommonComponentV2RequestCommonComponentAnchorComponent struct {
	// 组件自定义标题
	AnchorTitle string `json:"anchor_title"`
	// 星图组件id
	Id int64 `json:"id"`
	// 传小程序id
	MicroAppId string `json:"micro_app_id"`
	// 传小程序名称
	MicroAppName string `json:"micro_app_name"`
	// 起始页
	StartPath string `json:"start_path"`
}
