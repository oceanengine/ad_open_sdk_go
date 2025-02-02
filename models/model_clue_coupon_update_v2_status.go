/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponUpdateV2Status
type ClueCouponUpdateV2Status string

// List of clue_coupon_update_v2_status
const (
	AUDIT_FAIL_ClueCouponUpdateV2Status  ClueCouponUpdateV2Status = "AUDIT_FAIL"
	NORMAL_ClueCouponUpdateV2Status      ClueCouponUpdateV2Status = "NORMAL"
	AUDIT_DOING_ClueCouponUpdateV2Status ClueCouponUpdateV2Status = "AUDIT_DOING"
	PAUSE_ClueCouponUpdateV2Status       ClueCouponUpdateV2Status = "PAUSE"
	UNAUDITED_ClueCouponUpdateV2Status   ClueCouponUpdateV2Status = "UNAUDITED"
	DELETED_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "DELETED"
	OFFLINE_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "OFFLINE"
)

// Ptr returns reference to clue_coupon_update_v2_status value
func (v ClueCouponUpdateV2Status) Ptr() *ClueCouponUpdateV2Status {
	return &v
}
