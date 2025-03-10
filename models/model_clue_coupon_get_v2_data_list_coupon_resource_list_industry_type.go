/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2DataListCouponResourceListIndustryType
type ClueCouponGetV2DataListCouponResourceListIndustryType string

// List of clue_coupon_get_v2_data_list_coupon_resource_list_industry_type
const (
	FOOD_ClueCouponGetV2DataListCouponResourceListIndustryType          ClueCouponGetV2DataListCouponResourceListIndustryType = "FOOD"
	GAME_ClueCouponGetV2DataListCouponResourceListIndustryType          ClueCouponGetV2DataListCouponResourceListIndustryType = "GAME"
	FINANCIAL_ClueCouponGetV2DataListCouponResourceListIndustryType     ClueCouponGetV2DataListCouponResourceListIndustryType = "FINANCIAL"
	TICKET_ClueCouponGetV2DataListCouponResourceListIndustryType        ClueCouponGetV2DataListCouponResourceListIndustryType = "TICKET"
	OTHER_ClueCouponGetV2DataListCouponResourceListIndustryType         ClueCouponGetV2DataListCouponResourceListIndustryType = "OTHER"
	ENTERTAINMENT_ClueCouponGetV2DataListCouponResourceListIndustryType ClueCouponGetV2DataListCouponResourceListIndustryType = "ENTERTAINMENT"
)

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_resource_list_industry_type value
func (v ClueCouponGetV2DataListCouponResourceListIndustryType) Ptr() *ClueCouponGetV2DataListCouponResourceListIndustryType {
	return &v
}
