/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10AudienceDistrict
type QianchuanAdCreateV10AudienceDistrict string

// List of qianchuan_ad_create_v1.0_audience_district
const (
	CITY_QianchuanAdCreateV10AudienceDistrict   QianchuanAdCreateV10AudienceDistrict = "CITY"
	COUNTY_QianchuanAdCreateV10AudienceDistrict QianchuanAdCreateV10AudienceDistrict = "COUNTY"
	NONE_QianchuanAdCreateV10AudienceDistrict   QianchuanAdCreateV10AudienceDistrict = "NONE"
)

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_district value
func (v QianchuanAdCreateV10AudienceDistrict) Ptr() *QianchuanAdCreateV10AudienceDistrict {
	return &v
}
