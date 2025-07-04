/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2DataListCouponStatus
type ClueCouponGetV2DataListCouponStatus string

// List of clue_coupon_get_v2_data_list_coupon_status
const (
	DELETED_ClueCouponGetV2DataListCouponStatus     ClueCouponGetV2DataListCouponStatus = "DELETED"
	PAUSE_ClueCouponGetV2DataListCouponStatus       ClueCouponGetV2DataListCouponStatus = "PAUSE"
	OFFLINE_ClueCouponGetV2DataListCouponStatus     ClueCouponGetV2DataListCouponStatus = "OFFLINE"
	AUDIT_FAIL_ClueCouponGetV2DataListCouponStatus  ClueCouponGetV2DataListCouponStatus = "AUDIT_FAIL"
	NORMAL_ClueCouponGetV2DataListCouponStatus      ClueCouponGetV2DataListCouponStatus = "NORMAL"
	UNAUDITED_ClueCouponGetV2DataListCouponStatus   ClueCouponGetV2DataListCouponStatus = "UNAUDITED"
	AUDIT_DOING_ClueCouponGetV2DataListCouponStatus ClueCouponGetV2DataListCouponStatus = "AUDIT_DOING"
)

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_status value
func (v ClueCouponGetV2DataListCouponStatus) Ptr() *ClueCouponGetV2DataListCouponStatus {
	return &v
}
