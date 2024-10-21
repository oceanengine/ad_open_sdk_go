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
	USED_ClueCouponCodeGetV2DataListStatus      ClueCouponCodeGetV2DataListStatus = "USED"
	EXPIRED_ClueCouponCodeGetV2DataListStatus   ClueCouponCodeGetV2DataListStatus = "EXPIRED"
	INVALID_ClueCouponCodeGetV2DataListStatus   ClueCouponCodeGetV2DataListStatus = "INVALID"
	VALID_ClueCouponCodeGetV2DataListStatus     ClueCouponCodeGetV2DataListStatus = "VALID"
	ABANDONED_ClueCouponCodeGetV2DataListStatus ClueCouponCodeGetV2DataListStatus = "ABANDONED"
)

// Ptr returns reference to clue_coupon_code_get_v2_data_list_status value
func (v ClueCouponCodeGetV2DataListStatus) Ptr() *ClueCouponCodeGetV2DataListStatus {
	return &v
}
