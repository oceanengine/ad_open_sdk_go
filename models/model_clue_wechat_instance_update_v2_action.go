/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceUpdateV2Action
type ClueWechatInstanceUpdateV2Action string

// List of clue_wechat_instance_update_v2_action
const (
	ENABLE_ClueWechatInstanceUpdateV2Action  ClueWechatInstanceUpdateV2Action = "ENABLE"
	DISABLE_ClueWechatInstanceUpdateV2Action ClueWechatInstanceUpdateV2Action = "DISABLE"
	DELETE_ClueWechatInstanceUpdateV2Action  ClueWechatInstanceUpdateV2Action = "DELETE"
	ADD_ClueWechatInstanceUpdateV2Action     ClueWechatInstanceUpdateV2Action = "ADD"
)

// Ptr returns reference to clue_wechat_instance_update_v2_action value
func (v ClueWechatInstanceUpdateV2Action) Ptr() *ClueWechatInstanceUpdateV2Action {
	return &v
}
