/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdDetailV10DataOptStatus
type QianchuanUniPromotionAdDetailV10DataOptStatus string

// List of qianchuan_uni_promotion_ad_detail_v1.0_data_opt_status
const (
	DELETE_QianchuanUniPromotionAdDetailV10DataOptStatus         QianchuanUniPromotionAdDetailV10DataOptStatus = "DELETE"
	DISABLE_QianchuanUniPromotionAdDetailV10DataOptStatus        QianchuanUniPromotionAdDetailV10DataOptStatus = "DISABLE"
	ENABLE_QianchuanUniPromotionAdDetailV10DataOptStatus         QianchuanUniPromotionAdDetailV10DataOptStatus = "ENABLE"
	QUOTA_DISABLE_QianchuanUniPromotionAdDetailV10DataOptStatus  QianchuanUniPromotionAdDetailV10DataOptStatus = "QUOTA_DISABLE"
	ROI2_DISABLE_QianchuanUniPromotionAdDetailV10DataOptStatus   QianchuanUniPromotionAdDetailV10DataOptStatus = "ROI2_DISABLE"
	SYSTEM_DISABLE_QianchuanUniPromotionAdDetailV10DataOptStatus QianchuanUniPromotionAdDetailV10DataOptStatus = "SYSTEM_DISABLE"
)

// Ptr returns reference to qianchuan_uni_promotion_ad_detail_v1.0_data_opt_status value
func (v QianchuanUniPromotionAdDetailV10DataOptStatus) Ptr() *QianchuanUniPromotionAdDetailV10DataOptStatus {
	return &v
}
