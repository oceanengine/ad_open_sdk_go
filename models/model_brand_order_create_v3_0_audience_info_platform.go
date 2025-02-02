/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderCreateV30AudienceInfoPlatform
type BrandOrderCreateV30AudienceInfoPlatform string

// List of brand_order_create_v3.0_audience_info_platform
const (
	ANDROID_BrandOrderCreateV30AudienceInfoPlatform   BrandOrderCreateV30AudienceInfoPlatform = "ANDROID"
	IOS_BrandOrderCreateV30AudienceInfoPlatform       BrandOrderCreateV30AudienceInfoPlatform = "IOS"
	UNLIMITED_BrandOrderCreateV30AudienceInfoPlatform BrandOrderCreateV30AudienceInfoPlatform = "UNLIMITED"
)

// Ptr returns reference to brand_order_create_v3.0_audience_info_platform value
func (v BrandOrderCreateV30AudienceInfoPlatform) Ptr() *BrandOrderCreateV30AudienceInfoPlatform {
	return &v
}
