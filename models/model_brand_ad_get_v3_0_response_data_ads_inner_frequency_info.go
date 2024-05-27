/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerFrequencyInfo 频控信息
type BrandAdGetV30ResponseDataAdsInnerFrequencyInfo struct {
	// 周期频控
	Frequency       *int64                                            `json:"frequency,omitempty"`
	FrequencyStatus *BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus `json:"frequency_status,omitempty"`
	// 触达次数
	ReachTarget *int64 `json:"reach_target,omitempty"`
}
