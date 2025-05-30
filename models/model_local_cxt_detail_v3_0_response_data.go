/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalCxtDetailV30ResponseData
type LocalCxtDetailV30ResponseData struct {
	//
	AdId     *int64                                 `json:"ad_id,omitempty"`
	Audience *LocalCxtDetailV30ResponseDataAudience `json:"audience,omitempty"`
	//
	Budget *int64 `json:"budget,omitempty"`
	//
	CreateTime     *string                              `json:"create_time,omitempty"`
	ExternalAction *LocalCxtDetailV30DataExternalAction `json:"external_action,omitempty"`
	//
	HighBudgetRate *int64 `json:"high_budget_rate,omitempty"`
	//
	IsSetPeakBudget *bool `json:"is_set_peak_budget,omitempty"`
	//
	ModifyTime *string                         `json:"modify_time,omitempty"`
	OptStatus  *LocalCxtDetailV30DataOptStatus `json:"opt_status,omitempty"`
	//
	PeakHolidays []string `json:"peak_holidays,omitempty"`
	//
	PeakWeekDays []string `json:"peak_week_days,omitempty"`
	//
	PoiList []int64                      `json:"poi_list,omitempty"`
	Status  *LocalCxtDetailV30DataStatus `json:"status,omitempty"`
}
