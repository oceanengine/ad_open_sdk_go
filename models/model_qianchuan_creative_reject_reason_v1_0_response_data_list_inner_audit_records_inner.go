/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeRejectReasonV10ResponseDataListInnerAuditRecordsInner struct for QianchuanCreativeRejectReasonV10ResponseDataListInnerAuditRecordsInner
type QianchuanCreativeRejectReasonV10ResponseDataListInnerAuditRecordsInner struct {
	AuditPlatform *QianchuanCreativeRejectReasonV10DataListAuditRecordsAuditPlatform `json:"audit_platform,omitempty"`
	//
	Content *string `json:"content,omitempty"`
	//
	Desc *string `json:"desc,omitempty"`
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	RejectReason []string `json:"reject_reason,omitempty"`
	//
	Suggestion []string `json:"suggestion,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
