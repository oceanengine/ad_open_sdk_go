/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanMaterialAdGetV10ResponseDataAdListInner struct for QianchuanMaterialAdGetV10ResponseDataAdListInner
type QianchuanMaterialAdGetV10ResponseDataAdListInner struct {
	// 广告主id
	AdId *int64 `json:"ad_id,omitempty"`
	// 计划名称
	AdName  *string                                                  `json:"ad_name,omitempty"`
	Metrics *QianchuanMaterialAdGetV10ResponseDataAdListInnerMetrics `json:"metrics,omitempty"`
}
