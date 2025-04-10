/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderUpdateV30AudienceInfoAges
type BrandOrderUpdateV30AudienceInfoAges string

// List of brand_order_update_v3.0_audience_info_ages
const (
	ABOVE50_BrandOrderUpdateV30AudienceInfoAges      BrandOrderUpdateV30AudienceInfoAges = "ABOVE50"
	BETWEEN18_23_BrandOrderUpdateV30AudienceInfoAges BrandOrderUpdateV30AudienceInfoAges = "BETWEEN18_23"
	BETWEEN24_30_BrandOrderUpdateV30AudienceInfoAges BrandOrderUpdateV30AudienceInfoAges = "BETWEEN24_30"
	BETWEEN31_40_BrandOrderUpdateV30AudienceInfoAges BrandOrderUpdateV30AudienceInfoAges = "BETWEEN31_40"
	BETWEEN41_49_BrandOrderUpdateV30AudienceInfoAges BrandOrderUpdateV30AudienceInfoAges = "BETWEEN41_49"
	UNLIMITED_BrandOrderUpdateV30AudienceInfoAges    BrandOrderUpdateV30AudienceInfoAges = "UNLIMITED"
)

// Ptr returns reference to brand_order_update_v3.0_audience_info_ages value
func (v BrandOrderUpdateV30AudienceInfoAges) Ptr() *BrandOrderUpdateV30AudienceInfoAges {
	return &v
}
