/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2DataListActivityType
type ClueCouponGetV2DataListActivityType string

// List of clue_coupon_get_v2_data_list_activity_type
const (
	DIRECT_NEED_PHONE_ClueCouponGetV2DataListActivityType     ClueCouponGetV2DataListActivityType = "DIRECT_NEED_PHONE"
	DIRECT_NOT_NEED_PHONE_ClueCouponGetV2DataListActivityType ClueCouponGetV2DataListActivityType = "DIRECT_NOT_NEED_PHONE"
)

// Ptr returns reference to clue_coupon_get_v2_data_list_activity_type value
func (v ClueCouponGetV2DataListActivityType) Ptr() *ClueCouponGetV2DataListActivityType {
	return &v
}
