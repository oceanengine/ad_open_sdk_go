/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderQuotaGetV10ResponseData
type QianchuanAwemeOrderQuotaGetV10ResponseData struct {
	//
	MaxCost *float64 `json:"max_cost,omitempty"`
	//
	QuotaGear *int64 `json:"quota_gear,omitempty"`
	//
	SumQuota           *int64                                                        `json:"sum_quota,omitempty"`
	TerminateQuotaInfo *QianchuanAwemeOrderQuotaGetV10ResponseDataTerminateQuotaInfo `json:"terminate_quota_info,omitempty"`
	//
	UsedQuota *int64 `json:"used_quota,omitempty"`
}
