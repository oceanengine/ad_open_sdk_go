/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostStatusGetV10ResponseDataRaiseDetailsInner struct for QianchuanToolsSmartBoostAdBoostStatusGetV10ResponseDataRaiseDetailsInner
type QianchuanToolsSmartBoostAdBoostStatusGetV10ResponseDataRaiseDetailsInner struct {
	// 计划ID。
	AdId *int64 `json:"ad_id,omitempty"`
	// 预算。
	Budget *float64 `json:"budget,omitempty"`
	// 起量时长。
	RaiseDuration *float64 `json:"raise_duration,omitempty"`
	// 起量时间。
	RaiseTime *string                                                            `json:"raise_time,omitempty"`
	Status    *QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus `json:"status,omitempty"`
}
