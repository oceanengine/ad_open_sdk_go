/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandPolicyListV30AudienceInfoDistrictInfo 地域信息
type BrandPolicyListV30AudienceInfoDistrictInfo struct {
	// 选择的城市
	City           []int64                                                   `json:"city,omitempty"`
	CitySelectType *BrandPolicyListV30AudienceInfoDistrictInfoCitySelectType `json:"city_select_type,omitempty"`
	DistrictType   *BrandPolicyListV30AudienceInfoDistrictInfoDistrictType   `json:"district_type,omitempty"`
	LocationType   *BrandPolicyListV30AudienceInfoDistrictInfoLocationType   `json:"location_type,omitempty"`
}
