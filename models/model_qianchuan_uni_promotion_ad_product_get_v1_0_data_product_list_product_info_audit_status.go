/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus
type QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus string

// List of qianchuan_uni_promotion_ad_product_get_v1.0_data_product_list_product_info_audit_status
const (
	PASS_QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus        QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus = "PASS"
	REJECT_QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus      QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus = "REJECT"
	IN_PROGRESS_QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus = "IN_PROGRESS"
	NONE_QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus        QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus = "NONE"
)

// Ptr returns reference to qianchuan_uni_promotion_ad_product_get_v1.0_data_product_list_product_info_audit_status value
func (v QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus) Ptr() *QianchuanUniPromotionAdProductGetV10DataProductListProductInfoAuditStatus {
	return &v
}
