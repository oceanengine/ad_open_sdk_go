/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityAuditResultsV30DataStatus
type SecurityAuditResultsV30DataStatus string

// List of security_audit_results_v3.0_data_status
const (
	APPROVE_SecurityAuditResultsV30DataStatus  SecurityAuditResultsV30DataStatus = "APPROVE"
	REJECT_SecurityAuditResultsV30DataStatus   SecurityAuditResultsV30DataStatus = "REJECT"
	AUDITING_SecurityAuditResultsV30DataStatus SecurityAuditResultsV30DataStatus = "AUDITING"
)

// Ptr returns reference to security_audit_results_v3.0_data_status value
func (v SecurityAuditResultsV30DataStatus) Ptr() *SecurityAuditResultsV30DataStatus {
	return &v
}
