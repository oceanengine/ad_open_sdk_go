/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCreateV2CouponResourceListResourceType
type ClueCouponCreateV2CouponResourceListResourceType string

// List of clue_coupon_create_v2_coupon_resource_list_resource_type
const (
	FULL_ClueCouponCreateV2CouponResourceListResourceType     ClueCouponCreateV2CouponResourceListResourceType = "FULL"
	COMMON_ClueCouponCreateV2CouponResourceListResourceType   ClueCouponCreateV2CouponResourceListResourceType = "COMMON"
	DISCOUNT_ClueCouponCreateV2CouponResourceListResourceType ClueCouponCreateV2CouponResourceListResourceType = "DISCOUNT"
	GAME_ClueCouponCreateV2CouponResourceListResourceType     ClueCouponCreateV2CouponResourceListResourceType = "GAME"
	PHYSICAL_ClueCouponCreateV2CouponResourceListResourceType ClueCouponCreateV2CouponResourceListResourceType = "PHYSICAL"
)

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_resource_type value
func (v ClueCouponCreateV2CouponResourceListResourceType) Ptr() *ClueCouponCreateV2CouponResourceListResourceType {
	return &v
}
