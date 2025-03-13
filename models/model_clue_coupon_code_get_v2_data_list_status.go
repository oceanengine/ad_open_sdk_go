/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCodeGetV2DataListStatus
type ClueCouponCodeGetV2DataListStatus string

// List of clue_coupon_code_get_v2_data_list_status
const (
	VALID_ClueCouponCodeGetV2DataListStatus     ClueCouponCodeGetV2DataListStatus = "VALID"
	EXPIRED_ClueCouponCodeGetV2DataListStatus   ClueCouponCodeGetV2DataListStatus = "EXPIRED"
	USED_ClueCouponCodeGetV2DataListStatus      ClueCouponCodeGetV2DataListStatus = "USED"
	ABANDONED_ClueCouponCodeGetV2DataListStatus ClueCouponCodeGetV2DataListStatus = "ABANDONED"
	INVALID_ClueCouponCodeGetV2DataListStatus   ClueCouponCodeGetV2DataListStatus = "INVALID"
)

// Ptr returns reference to clue_coupon_code_get_v2_data_list_status value
func (v ClueCouponCodeGetV2DataListStatus) Ptr() *ClueCouponCodeGetV2DataListStatus {
	return &v
}
