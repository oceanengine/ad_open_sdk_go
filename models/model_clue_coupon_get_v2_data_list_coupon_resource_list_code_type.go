/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2DataListCouponResourceListCodeType
type ClueCouponGetV2DataListCouponResourceListCodeType string

// List of clue_coupon_get_v2_data_list_coupon_resource_list_code_type
const (
	API_ClueCouponGetV2DataListCouponResourceListCodeType      ClueCouponGetV2DataListCouponResourceListCodeType = "API"
	PLATFORM_ClueCouponGetV2DataListCouponResourceListCodeType ClueCouponGetV2DataListCouponResourceListCodeType = "PLATFORM"
	MERCHANT_ClueCouponGetV2DataListCouponResourceListCodeType ClueCouponGetV2DataListCouponResourceListCodeType = "MERCHANT"
	COMMON_ClueCouponGetV2DataListCouponResourceListCodeType   ClueCouponGetV2DataListCouponResourceListCodeType = "COMMON"
)

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_resource_list_code_type value
func (v ClueCouponGetV2DataListCouponResourceListCodeType) Ptr() *ClueCouponGetV2DataListCouponResourceListCodeType {
	return &v
}
