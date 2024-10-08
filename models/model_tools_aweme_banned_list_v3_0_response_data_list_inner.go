/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeBannedListV30ResponseDataListInner struct for ToolsAwemeBannedListV30ResponseDataListInner
type ToolsAwemeBannedListV30ResponseDataListInner struct {
	// 抖音昵称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 抖音用户id
	AwemeUserId *string                                    `json:"aweme_user_id,omitempty"`
	BannedType  *ToolsAwemeBannedListV30DataListBannedType `json:"banned_type,omitempty"`
	// 昵称关键词
	NicknameKeyword *string `json:"nickname_keyword,omitempty"`
}
