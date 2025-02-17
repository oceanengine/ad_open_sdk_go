/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerWechatGame 微信小游戏组件描述
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerWechatGame struct {
	// 微信小游戏的路径参数
	GamePath *string `json:"game_path,omitempty"`
	// 微信小游戏组件ID
	InstanceId *int64 `json:"instance_id,omitempty"`
	// 简介，长度不超过40
	Introduction *string `json:"introduction,omitempty"`
	// 标签，个数不超过2，字数不超过5
	Items []string `json:"items,omitempty"`
	// logo链接
	Logo *string `json:"logo,omitempty"`
}
