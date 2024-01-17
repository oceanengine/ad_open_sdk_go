/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInnerDetailAuditInfoInner struct for StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInnerDetailAuditInfoInner
type StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInnerDetailAuditInfoInner struct {
	// 审核拒绝原因
	AuditBanReason *string `json:"audit_ban_reason,omitempty"`
	// 审核渠道 1 = 抖音内容审核 7 = 广告素材审核
	AuditSource *int64 `json:"audit_source,omitempty"`
	// 审核状态 0 = 审核不通过 99 = 审核中 1 = 审核通过 3 = 作者删除 4 = 跳过审核
	AuditStatus *int64 `json:"audit_status,omitempty"`
	// 审核时间  unix timestamp
	AuditTime *int64 `json:"audit_time,omitempty"`
}
