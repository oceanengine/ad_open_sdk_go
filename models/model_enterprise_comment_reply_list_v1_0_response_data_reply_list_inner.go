/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentReplyListV10ResponseDataReplyListInner struct for EnterpriseCommentReplyListV10ResponseDataReplyListInner
type EnterpriseCommentReplyListV10ResponseDataReplyListInner struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	IsEDouyinId *bool `json:"is_e_douyin_id,omitempty"`
	//
	ReplyAwemeName *string `json:"reply_aweme_name,omitempty"`
	//
	ReplyId *int64 `json:"reply_id,omitempty"`
	//
	ReplyOpenId *string                                                `json:"reply_open_id,omitempty"`
	ReplyStatus *EnterpriseCommentReplyListV10DataReplyListReplyStatus `json:"reply_status,omitempty"`
	//
	ReplyText *string `json:"reply_text,omitempty"`
}
