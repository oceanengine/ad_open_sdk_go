/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskAuthorVideoAuditV30OptStatus
type StardeliveryTaskAuthorVideoAuditV30OptStatus string

// List of stardelivery_task_author_video_audit_v3.0_opt_status
const (
	PASS_StardeliveryTaskAuthorVideoAuditV30OptStatus   StardeliveryTaskAuthorVideoAuditV30OptStatus = "PASS"
	REJECT_StardeliveryTaskAuthorVideoAuditV30OptStatus StardeliveryTaskAuthorVideoAuditV30OptStatus = "REJECT"
)

// Ptr returns reference to stardelivery_task_author_video_audit_v3.0_opt_status value
func (v StardeliveryTaskAuthorVideoAuditV30OptStatus) Ptr() *StardeliveryTaskAuthorVideoAuditV30OptStatus {
	return &v
}