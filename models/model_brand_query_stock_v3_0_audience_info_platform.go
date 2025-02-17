/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryStockV30AudienceInfoPlatform
type BrandQueryStockV30AudienceInfoPlatform string

// List of brand_query_stock_v3.0_audience_info_platform
const (
	ANDROID_BrandQueryStockV30AudienceInfoPlatform   BrandQueryStockV30AudienceInfoPlatform = "ANDROID"
	IOS_BrandQueryStockV30AudienceInfoPlatform       BrandQueryStockV30AudienceInfoPlatform = "IOS"
	UNLIMITED_BrandQueryStockV30AudienceInfoPlatform BrandQueryStockV30AudienceInfoPlatform = "UNLIMITED"
)

// Ptr returns reference to brand_query_stock_v3.0_audience_info_platform value
func (v BrandQueryStockV30AudienceInfoPlatform) Ptr() *BrandQueryStockV30AudienceInfoPlatform {
	return &v
}
