/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandToolQueryPublishPriceV30FrequencyInfo
type BrandToolQueryPublishPriceV30FrequencyInfo struct {
	// 频控信息
	Frequency       *int64                                                     `json:"frequency,omitempty"`
	FrequencyStatus *BrandToolQueryPublishPriceV30FrequencyInfoFrequencyStatus `json:"frequency_status,omitempty"`
}
