/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdCostProtectStatusGetV2ResponseDataListInner struct for AdCostProtectStatusGetV2ResponseDataListInner
type AdCostProtectStatusGetV2ResponseDataListInner struct {
	// 计划id
	AdId   *int64                                  `json:"ad_id,omitempty"`
	Status *AdCostProtectStatusGetV2DataListStatus `json:"status,omitempty"`
}
