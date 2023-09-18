/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentTermsBannedUpdateV30Request struct for ToolsCommentTermsBannedUpdateV30Request
type ToolsCommentTermsBannedUpdateV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 抖音号，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeId *string `json:"aweme_id,omitempty"`
	// 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv *bool `json:"is_apply_to_adv,omitempty"`
	// 更新后的屏蔽词，屏蔽词长度范围为1-20字
	NewTerm string `json:"new_term"`
	// 待更新的屏蔽词，屏蔽词长度范围为1-20字
	OriginTerm string `json:"origin_term"`
}