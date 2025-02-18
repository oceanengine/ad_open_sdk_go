/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2DataCouponResourceListCodeType
type ClueCouponDetailV2DataCouponResourceListCodeType string

// List of clue_coupon_detail_v2_data_coupon_resource_list_code_type
const (
	PLATFORM_ClueCouponDetailV2DataCouponResourceListCodeType ClueCouponDetailV2DataCouponResourceListCodeType = "PLATFORM"
	COMMON_ClueCouponDetailV2DataCouponResourceListCodeType   ClueCouponDetailV2DataCouponResourceListCodeType = "COMMON"
	MERCHANT_ClueCouponDetailV2DataCouponResourceListCodeType ClueCouponDetailV2DataCouponResourceListCodeType = "MERCHANT"
	API_ClueCouponDetailV2DataCouponResourceListCodeType      ClueCouponDetailV2DataCouponResourceListCodeType = "API"
)

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_resource_list_code_type value
func (v ClueCouponDetailV2DataCouponResourceListCodeType) Ptr() *ClueCouponDetailV2DataCouponResourceListCodeType {
	return &v
}
