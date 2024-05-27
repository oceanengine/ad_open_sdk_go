/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseSetV30RequestRaiseInfoInnerAppointedTime 指定投放时间，当is_effective_now为FALSE时填写有效
type ToolsPromotionRaiseSetV30RequestRaiseInfoInnerAppointedTime struct {
	// 起量时间， 重复周期不传时，格式为yyyy-mm-dd HH:MM 传重复周期时，格式为HH:MM
	RaiseTime *string `json:"raise_time,omitempty"`
	// 重复周期，不传则不重复，仅生效一次，所有枚举传入则每天重复
	RepeatedDay []*ToolsPromotionRaiseSetV30RaiseInfoAppointedTimeRepeatedDay `json:"repeated_day,omitempty"`
}
