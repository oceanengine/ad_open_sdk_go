/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerWechatGame 微信小游戏组件描述
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerWechatGame struct {
	// 微信小游戏路径
	GamePath *string `json:"game_path,omitempty"`
	// 微信小游戏组件ID
	InstanceId int64 `json:"instance_id"`
	// 简介，长度不超过40
	Introduction *string `json:"introduction,omitempty"`
	// 标签，个数不超过2，字数不超过5
	Items []string `json:"items,omitempty"`
	// logo链接
	Logo *string `json:"logo,omitempty"`
}
