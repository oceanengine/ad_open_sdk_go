/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialSuggestionV10ResponseDataListInnerAuditRecords
type QianchuanAdMaterialSuggestionV10ResponseDataListInnerAuditRecords struct {
	AuditPlatform *QianchuanAdMaterialSuggestionV10DataListAuditRecordsAuditPlatform `json:"audit_platform,omitempty"`
	//
	RejectReason []string `json:"reject_reason,omitempty"`
	//
	Suggestions []string `json:"suggestions,omitempty"`
}