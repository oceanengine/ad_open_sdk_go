/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2DataFormValidateType
type ClueCouponDetailV2DataFormValidateType string

// List of clue_coupon_detail_v2_data_form_validate_type
const (
	AUTO_VERIFICATION_ClueCouponDetailV2DataFormValidateType ClueCouponDetailV2DataFormValidateType = "AUTO_VERIFICATION"
	CLUE_PRIORITY_ClueCouponDetailV2DataFormValidateType     ClueCouponDetailV2DataFormValidateType = "CLUE_PRIORITY"
	ALL_VERIFICATION_ClueCouponDetailV2DataFormValidateType  ClueCouponDetailV2DataFormValidateType = "ALL_VERIFICATION"
	VALIDITY_PRIORITY_ClueCouponDetailV2DataFormValidateType ClueCouponDetailV2DataFormValidateType = "VALIDITY_PRIORITY"
	NONE_VERIFICATION_ClueCouponDetailV2DataFormValidateType ClueCouponDetailV2DataFormValidateType = "NONE_VERIFICATION"
)

// Ptr returns reference to clue_coupon_detail_v2_data_form_validate_type value
func (v ClueCouponDetailV2DataFormValidateType) Ptr() *ClueCouponDetailV2DataFormValidateType {
	return &v
}
