/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus
type QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus string

// List of qianchuan_ad_material_get_v1.0_data_ad_material_infos_audit_status
const (
	PASS_QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus        QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus = "PASS"
	REJECT_QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus      QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus = "REJECT"
	IN_PROGRESS_QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus = "IN_PROGRESS"
)

// Ptr returns reference to qianchuan_ad_material_get_v1.0_data_ad_material_infos_audit_status value
func (v QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus) Ptr() *QianchuanAdMaterialGetV10DataAdMaterialInfosAuditStatus {
	return &v
}
