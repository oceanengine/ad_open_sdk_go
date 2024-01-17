/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigCreateV2Request struct for CreativeAutoGenerateConfigCreateV2Request
type CreativeAutoGenerateConfigCreateV2Request struct {
	// 广告计划ID
	AdId int64 `json:"ad_id"`
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 配置ID，新建时不传入、修改时必传
	ConfigId           *int64                                               `json:"config_id,omitempty"`
	SelectTemplateMode CreativeAutoGenerateConfigCreateV2SelectTemplateMode `json:"select_template_mode"`
	// 模板列表
	Templates []*CreativeAutoGenerateConfigCreateV2RequestTemplatesInner `json:"templates,omitempty"`
}
