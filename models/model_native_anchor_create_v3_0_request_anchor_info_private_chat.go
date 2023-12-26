/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoPrivateChat 私信锚点，当anchor_type=PRIVATE_CHAT必填
type NativeAnchorCreateV30RequestAnchorInfoPrivateChat struct {
	Button NativeAnchorCreateV30AnchorInfoPrivateChatButton `json:"button"`
	// 咨询引导文案，如私信获取一对一解答，不超过9个字
	ChatGuide string `json:"chat_guide"`
	// 工具名称，例如直播间私信咨询/在线咨询等，不超过20个字
	ToolsName string `json:"tools_name"`
}
