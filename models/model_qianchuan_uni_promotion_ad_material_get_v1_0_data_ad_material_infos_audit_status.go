/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus
type QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus string

// List of qianchuan_uni_promotion_ad_material_get_v1.0_data_ad_material_infos_audit_status
const (
	PASS_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus        QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus = "PASS"
	REJECT_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus      QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus = "REJECT"
	IN_PROGRESS_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus = "IN_PROGRESS"
)

// Ptr returns reference to qianchuan_uni_promotion_ad_material_get_v1.0_data_ad_material_infos_audit_status value
func (v QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus) Ptr() *QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosAuditStatus {
	return &v
}
