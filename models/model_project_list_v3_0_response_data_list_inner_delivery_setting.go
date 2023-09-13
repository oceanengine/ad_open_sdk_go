/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInnerDeliverySetting
type ProjectListV30ResponseDataListInnerDeliverySetting struct {
	//
	Bid      *float64                                       `json:"bid,omitempty"`
	BidSpeed *ProjectListV30DataListDeliverySettingBidSpeed `json:"bid_speed,omitempty"`
	//
	BidType *string `json:"bid_type,omitempty"`
	//
	Budget               *float64                                                   `json:"budget,omitempty"`
	BudgetMode           *ProjectListV30DataListDeliverySettingBudgetMode           `json:"budget_mode,omitempty"`
	BudgetOptimizeSwitch *ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch `json:"budget_optimize_switch,omitempty"`
	//
	CpaBid      *float64                                          `json:"cpa_bid,omitempty"`
	DeepBidType *ProjectListV30DataListDeliverySettingDeepBidType `json:"deep_bid_type,omitempty"`
	//
	DeepCpabid *float64 `json:"deep_cpabid,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	FilterNightSwitch *string `json:"filter_night_switch,omitempty"`
	//
	FirstRoiGoal  *float64                                            `json:"first_roi_goal,omitempty"`
	ProjectCustom *ProjectListV30DataListDeliverySettingProjectCustom `json:"project_custom,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleTime *string                                            `json:"schedule_time,omitempty"`
	ScheduleType *ProjectListV30DataListDeliverySettingScheduleType `json:"schedule_type,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
