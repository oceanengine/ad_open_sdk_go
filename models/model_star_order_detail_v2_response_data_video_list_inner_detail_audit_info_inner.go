/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderDetailV2ResponseDataVideoListInnerDetailAuditInfoInner struct for StarOrderDetailV2ResponseDataVideoListInnerDetailAuditInfoInner
type StarOrderDetailV2ResponseDataVideoListInnerDetailAuditInfoInner struct {
	// 审核拒绝原因
	AuditBanReason *string `json:"audit_ban_reason,omitempty"`
	// 审核渠道 1 = 抖音内容审核 7 = 广告素材审核
	AuditSource *string `json:"audit_source,omitempty"`
	// 审核状态 0 = 审核不通过 99 = 审核中 1 = 审核通过 3 = 作者删除 4 = 跳过审核
	AuditStatus *string `json:"audit_status,omitempty"`
	// 审核时间  unix timestamp
	AuditTime *string `json:"audit_time,omitempty"`
}
