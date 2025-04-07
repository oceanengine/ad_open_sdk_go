/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetsDetailReadV2DataListAuditStatus
type DpaAssetsDetailReadV2DataListAuditStatus string

// List of dpa_assets_detail_read_v2_data_list_audit_status
const (
	ASSETS_AUDIT_DENY_DpaAssetsDetailReadV2DataListAuditStatus      DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_AUDIT_DENY"
	ASSETS_AUDIT_SUCCESS_DpaAssetsDetailReadV2DataListAuditStatus   DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_AUDIT_SUCCESS"
	ASSETS_TO_SUBMIT_AUDIT_DpaAssetsDetailReadV2DataListAuditStatus DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_TO_SUBMIT_AUDIT"
	ASSETS_PENDING_AUDIT_DpaAssetsDetailReadV2DataListAuditStatus   DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_PENDING_AUDIT"
)

// Ptr returns reference to dpa_assets_detail_read_v2_data_list_audit_status value
func (v DpaAssetsDetailReadV2DataListAuditStatus) Ptr() *DpaAssetsDetailReadV2DataListAuditStatus {
	return &v
}
