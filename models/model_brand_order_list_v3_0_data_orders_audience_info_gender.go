/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30DataOrdersAudienceInfoGender
type BrandOrderListV30DataOrdersAudienceInfoGender string

// List of brand_order_list_v3.0_data_orders_audience_info_gender
const (
	FEMALE_BrandOrderListV30DataOrdersAudienceInfoGender    BrandOrderListV30DataOrdersAudienceInfoGender = "FEMALE"
	MALE_BrandOrderListV30DataOrdersAudienceInfoGender      BrandOrderListV30DataOrdersAudienceInfoGender = "MALE"
	UNLIMITED_BrandOrderListV30DataOrdersAudienceInfoGender BrandOrderListV30DataOrdersAudienceInfoGender = "UNLIMITED"
)

// Ptr returns reference to brand_order_list_v3.0_data_orders_audience_info_gender value
func (v BrandOrderListV30DataOrdersAudienceInfoGender) Ptr() *BrandOrderListV30DataOrdersAudienceInfoGender {
	return &v
}