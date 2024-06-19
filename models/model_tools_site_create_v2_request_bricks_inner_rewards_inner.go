/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCreateV2RequestBricksInnerRewardsInner struct for ToolsSiteCreateV2RequestBricksInnerRewardsInner
type ToolsSiteCreateV2RequestBricksInnerRewardsInner struct {
	// 奖品图片
	Image *string `json:"image,omitempty"`
	// 中奖页奖品名称
	PopupText *string `json:"popup_text,omitempty"`
	// 中奖概率,所有奖品中奖概率和为100
	Rate *float64 `json:"rate,omitempty"`
	// 奖品名称
	Text *string `json:"text,omitempty"`
}
