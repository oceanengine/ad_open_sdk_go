/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderUpdateV30AudienceInfoGender
type BrandOrderUpdateV30AudienceInfoGender string

// List of brand_order_update_v3.0_audience_info_gender
const (
	FEMALE_BrandOrderUpdateV30AudienceInfoGender    BrandOrderUpdateV30AudienceInfoGender = "FEMALE"
	MALE_BrandOrderUpdateV30AudienceInfoGender      BrandOrderUpdateV30AudienceInfoGender = "MALE"
	UNLIMITED_BrandOrderUpdateV30AudienceInfoGender BrandOrderUpdateV30AudienceInfoGender = "UNLIMITED"
)

// Ptr returns reference to brand_order_update_v3.0_audience_info_gender value
func (v BrandOrderUpdateV30AudienceInfoGender) Ptr() *BrandOrderUpdateV30AudienceInfoGender {
	return &v
}
