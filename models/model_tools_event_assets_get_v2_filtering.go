/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAssetsGetV2Filtering
type ToolsEventAssetsGetV2Filtering struct {
	App         *ToolsEventAssetsGetV2FilteringApp         `json:"app,omitempty"`
	LandingPage *ToolsEventAssetsGetV2FilteringLandingPage `json:"landing_page,omitempty"`
	MiniProgram *ToolsEventAssetsGetV2FilteringMiniProgram `json:"mini_program,omitempty"`
	QuickApp    *ToolsEventAssetsGetV2FilteringQuickApp    `json:"quick_app,omitempty"`
}
