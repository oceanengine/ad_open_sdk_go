/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseStatusGetV30ResponseDataListInnerRaiseInfoInnerAppointTime 指定投放时间
type ToolsPromotionRaiseStatusGetV30ResponseDataListInnerRaiseInfoInnerAppointTime struct {
	// 起量时间， 重复周期不传时，格式为yyyy-mm-dd HH:MM 传重复周期时，格式为HH:MM
	RaiseTime *string `json:"raise_time,omitempty"`
	// 重复周期，不传则不重复，仅生效一次 允许值： PER_MONDAY；PER_TUESDAY；PER_WEDNESDAY；PER_THURSDAY；PER_FRIDAY：PER_SATURDAY；PER_SUNDAY；EVERY_DAY
	RepeatDay []*ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay `json:"repeat_day,omitempty"`
}
