/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCreateV2CouponResourceListIndustryType
type ClueCouponCreateV2CouponResourceListIndustryType string

// List of clue_coupon_create_v2_coupon_resource_list_industry_type
const (
	FOOD_ClueCouponCreateV2CouponResourceListIndustryType          ClueCouponCreateV2CouponResourceListIndustryType = "FOOD"
	TICKET_ClueCouponCreateV2CouponResourceListIndustryType        ClueCouponCreateV2CouponResourceListIndustryType = "TICKET"
	GAME_ClueCouponCreateV2CouponResourceListIndustryType          ClueCouponCreateV2CouponResourceListIndustryType = "GAME"
	ENTERTAINMENT_ClueCouponCreateV2CouponResourceListIndustryType ClueCouponCreateV2CouponResourceListIndustryType = "ENTERTAINMENT"
	FINANCIAL_ClueCouponCreateV2CouponResourceListIndustryType     ClueCouponCreateV2CouponResourceListIndustryType = "FINANCIAL"
	OTHER_ClueCouponCreateV2CouponResourceListIndustryType         ClueCouponCreateV2CouponResourceListIndustryType = "OTHER"
)

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_industry_type value
func (v ClueCouponCreateV2CouponResourceListIndustryType) Ptr() *ClueCouponCreateV2CouponResourceListIndustryType {
	return &v
}
