/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus
type StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus string

// List of star_vas_get_boost_group_list_v2_data_boost_group_infos_task_infos_audit_status
const (
	AUDITING_StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus     StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus = "AUDITING"
	AUDIT_FAILED_StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus = "AUDIT_FAILED"
	AUDIT_PASS_StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus   StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus = "AUDIT_PASS"
)

// Ptr returns reference to star_vas_get_boost_group_list_v2_data_boost_group_infos_task_infos_audit_status value
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus) Ptr() *StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus {
	return &v
}
