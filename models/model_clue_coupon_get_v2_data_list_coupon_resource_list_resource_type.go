/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2DataListCouponResourceListResourceType
type ClueCouponGetV2DataListCouponResourceListResourceType string

// List of clue_coupon_get_v2_data_list_coupon_resource_list_resource_type
const (
	GAME_ClueCouponGetV2DataListCouponResourceListResourceType     ClueCouponGetV2DataListCouponResourceListResourceType = "GAME"
	FULL_ClueCouponGetV2DataListCouponResourceListResourceType     ClueCouponGetV2DataListCouponResourceListResourceType = "FULL"
	COMMON_ClueCouponGetV2DataListCouponResourceListResourceType   ClueCouponGetV2DataListCouponResourceListResourceType = "COMMON"
	PHYSICAL_ClueCouponGetV2DataListCouponResourceListResourceType ClueCouponGetV2DataListCouponResourceListResourceType = "PHYSICAL"
	DISCOUNT_ClueCouponGetV2DataListCouponResourceListResourceType ClueCouponGetV2DataListCouponResourceListResourceType = "DISCOUNT"
)

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_resource_list_resource_type value
func (v ClueCouponGetV2DataListCouponResourceListResourceType) Ptr() *ClueCouponGetV2DataListCouponResourceListResourceType {
	return &v
}
