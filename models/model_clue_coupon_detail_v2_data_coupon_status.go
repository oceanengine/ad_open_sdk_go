/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2DataCouponStatus
type ClueCouponDetailV2DataCouponStatus string

// List of clue_coupon_detail_v2_data_coupon_status
const (
	AUDIT_FAIL_ClueCouponDetailV2DataCouponStatus  ClueCouponDetailV2DataCouponStatus = "AUDIT_FAIL"
	AUDIT_DOING_ClueCouponDetailV2DataCouponStatus ClueCouponDetailV2DataCouponStatus = "AUDIT_DOING"
	PAUSE_ClueCouponDetailV2DataCouponStatus       ClueCouponDetailV2DataCouponStatus = "PAUSE"
	OFFLINE_ClueCouponDetailV2DataCouponStatus     ClueCouponDetailV2DataCouponStatus = "OFFLINE"
	DELETED_ClueCouponDetailV2DataCouponStatus     ClueCouponDetailV2DataCouponStatus = "DELETED"
	UNAUDITED_ClueCouponDetailV2DataCouponStatus   ClueCouponDetailV2DataCouponStatus = "UNAUDITED"
	NORMAL_ClueCouponDetailV2DataCouponStatus      ClueCouponDetailV2DataCouponStatus = "NORMAL"
)

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_status value
func (v ClueCouponDetailV2DataCouponStatus) Ptr() *ClueCouponDetailV2DataCouponStatus {
	return &v
}
