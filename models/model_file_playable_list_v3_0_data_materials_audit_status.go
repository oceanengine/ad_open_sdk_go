/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FilePlayableListV30DataMaterialsAuditStatus
type FilePlayableListV30DataMaterialsAuditStatus string

// List of file_playable_list_v3.0_data_materials_audit_status
const (
	STATUS_CONFIRMED_FilePlayableListV30DataMaterialsAuditStatus    FilePlayableListV30DataMaterialsAuditStatus = "STATUS_CONFIRMED"
	STATUS_CONFIRMING_FilePlayableListV30DataMaterialsAuditStatus   FilePlayableListV30DataMaterialsAuditStatus = "STATUS_CONFIRMING"
	STATUS_CONFIRM_FAIL_FilePlayableListV30DataMaterialsAuditStatus FilePlayableListV30DataMaterialsAuditStatus = "STATUS_CONFIRM_FAIL"
)

// Ptr returns reference to file_playable_list_v3.0_data_materials_audit_status value
func (v FilePlayableListV30DataMaterialsAuditStatus) Ptr() *FilePlayableListV30DataMaterialsAuditStatus {
	return &v
}
