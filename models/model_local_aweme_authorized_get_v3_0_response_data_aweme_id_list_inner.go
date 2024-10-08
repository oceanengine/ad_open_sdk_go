/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalAwemeAuthorizedGetV30ResponseDataAwemeIdListInner struct for LocalAwemeAuthorizedGetV30ResponseDataAwemeIdListInner
type LocalAwemeAuthorizedGetV30ResponseDataAwemeIdListInner struct {
	AuthType *LocalAwemeAuthorizedGetV30DataAwemeIdListAuthType `json:"auth_type,omitempty"`
	// 抖音头像
	AwemeAvatar *string `json:"aweme_avatar,omitempty"`
	// 该抖音号是否有全域推广计划投放
	AwemeHasUniProm *bool `json:"aweme_has_uni_prom,omitempty"`
	// 抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音号名称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 是否能创建roi2广告，margoal=2时返回
	CanCreateRoi2Ad *bool `json:"can_create_roi2_ad,omitempty"`
}
