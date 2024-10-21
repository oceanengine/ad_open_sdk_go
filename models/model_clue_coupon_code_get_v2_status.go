/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCodeGetV2Status
type ClueCouponCodeGetV2Status string

// List of clue_coupon_code_get_v2_status
const (
	USED_ClueCouponCodeGetV2Status      ClueCouponCodeGetV2Status = "USED"
	EXPIRED_ClueCouponCodeGetV2Status   ClueCouponCodeGetV2Status = "EXPIRED"
	INVALID_ClueCouponCodeGetV2Status   ClueCouponCodeGetV2Status = "INVALID"
	VALID_ClueCouponCodeGetV2Status     ClueCouponCodeGetV2Status = "VALID"
	ABANDONED_ClueCouponCodeGetV2Status ClueCouponCodeGetV2Status = "ABANDONED"
)

// Ptr returns reference to clue_coupon_code_get_v2_status value
func (v ClueCouponCodeGetV2Status) Ptr() *ClueCouponCodeGetV2Status {
	return &v
}
