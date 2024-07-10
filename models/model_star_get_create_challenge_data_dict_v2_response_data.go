/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarGetCreateChallengeDataDictV2ResponseData
type StarGetCreateChallengeDataDictV2ResponseData struct {
	// 观众年龄范围可选项
	AgeRange []string `json:"age_range,omitempty"`
	// 观众地域可选项
	CityLevel []string `json:"city_level,omitempty"`
	// 内容类型可选项
	ContentTags []*StarGetCreateChallengeDataDictV2ResponseDataContentTagsInner `json:"content_tags,omitempty"`
	// 观众设备品牌可选项
	DeviceBrands []string `json:"device_brands,omitempty"`
	// 达人省份可选项
	Provinces []string `json:"provinces,omitempty"`
}
