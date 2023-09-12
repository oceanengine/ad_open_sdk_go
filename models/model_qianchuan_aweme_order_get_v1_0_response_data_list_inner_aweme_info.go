/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderGetV10ResponseDataListInnerAwemeInfo 素材作者信息
type QianchuanAwemeOrderGetV10ResponseDataListInnerAwemeInfo struct {
	// 抖音号头像
	AwemeAvatar *string `json:"aweme_avatar,omitempty"`
	// 抖音uid
	AwemeId *int64 `json:"aweme_id,omitempty"`
	// 抖音昵称
	AwemeName *string `json:"aweme_name,omitempty"`
	//  抖音号
	AwemeShowId *string `json:"aweme_show_id,omitempty"`
}
