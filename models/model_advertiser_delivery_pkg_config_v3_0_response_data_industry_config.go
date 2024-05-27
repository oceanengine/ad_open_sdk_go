/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfig 资质规则
type AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfig struct {
	// 资质规则id
	ConfigId int64 `json:"config_id"`
	// 第一级到第三级行业id
	IndustryIds []int64 `json:"industry_ids"`
	// 第一级到第三级行业名称
	IndustryNames  []string                                                       `json:"industry_names"`
	IndustryStatus AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus `json:"industry_status"`
	// 必填资质模块配置
	Necessaries []*AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigNecessariesInner `json:"necessaries,omitempty"`
	// 选填资质模块配置
	Unnecessaries []*AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInner `json:"unnecessaries,omitempty"`
}
