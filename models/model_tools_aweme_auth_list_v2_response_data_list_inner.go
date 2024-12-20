/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthListV2ResponseDataListInner struct for ToolsAwemeAuthListV2ResponseDataListInner
type ToolsAwemeAuthListV2ResponseDataListInner struct {
	AuthStatus *ToolsAwemeAuthListV2DataListAuthStatus `json:"auth_status,omitempty"`
	AuthType   *ToolsAwemeAuthListV2DataListAuthType   `json:"auth_type,omitempty"`
	// 抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音账号名称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 授权结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 备注
	Note      *string                                `json:"note,omitempty"`
	ShareType *ToolsAwemeAuthListV2DataListShareType `json:"share_type,omitempty"`
	// 授权开始时间
	StartTime *string                                             `json:"start_time,omitempty"`
	SubStatus *ToolsAwemeAuthListV2DataListSubStatus              `json:"sub_status,omitempty"`
	VideoInfo *ToolsAwemeAuthListV2ResponseDataListInnerVideoInfo `json:"video_info,omitempty"`
}
