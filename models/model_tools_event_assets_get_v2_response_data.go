/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAssetsGetV2ResponseData
type ToolsEventAssetsGetV2ResponseData struct {
	// 应用数据
	App []*ToolsEventAssetsGetV2ResponseDataAppInner `json:"app,omitempty"`
	// 三方数据集合
	LandingPages []*ToolsEventAssetsGetV2ResponseDataLandingPagesInner `json:"landing_pages,omitempty"`
	// 字节小程序快应用资产
	MiniProgram []*ToolsEventAssetsGetV2ResponseDataMiniProgramInner `json:"mini_program,omitempty"`
	PageInfo    *ToolsEventAssetsGetV2ResponseDataPageInfo           `json:"page_info,omitempty"`
	// 快应用数据
	QuickApp []*ToolsEventAssetsGetV2ResponseDataQuickAppInner `json:"quick_app,omitempty"`
}
