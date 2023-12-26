/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataDeliverySetting
type QianchuanAdDetailGetV10ResponseDataDeliverySetting struct {
	//
	AllowQcpx *bool `json:"allow_qcpx,omitempty"`
	//
	AutoManageStrategyCmd *int64 `json:"auto_manage_strategy_cmd,omitempty"`
	//
	Budget     *float64                                              `json:"budget,omitempty"`
	BudgetMode *QianchuanAdDetailGetV10DataDeliverySettingBudgetMode `json:"budget_mode,omitempty"`
	//
	CpaBid      *float64                                               `json:"cpa_bid,omitempty"`
	DeepBidType *QianchuanAdDetailGetV10DataDeliverySettingDeepBidType `json:"deep_bid_type,omitempty"`
	//
	DeepCpaBid           *float64                                                        `json:"deep_cpa_bid,omitempty"`
	DeepExternalAction   *QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction   `json:"deep_external_action,omitempty"`
	EnableAutoPause      *QianchuanAdDetailGetV10DataDeliverySettingEnableAutoPause      `json:"enable_auto_pause,omitempty"`
	EnableFollowMaterial *QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial `json:"enable_follow_material,omitempty"`
	//
	EndTime          *string                                                     `json:"end_time,omitempty"`
	ExternalAction   *QianchuanAdDetailGetV10DataDeliverySettingExternalAction   `json:"external_action,omitempty"`
	LiveScheduleType *QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType `json:"live_schedule_type,omitempty"`
	//
	ProductNewOpen *bool                                               `json:"product_new_open,omitempty"`
	QcpxMode       *QianchuanAdDetailGetV10DataDeliverySettingQcpxMode `json:"qcpx_mode,omitempty"`
	//
	ReviveBudget *float64 `json:"revive_budget,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleFixedRange *int64 `json:"schedule_fixed_range,omitempty"`
	//
	ScheduleTime *string                                                 `json:"schedule_time,omitempty"`
	SmartBidType *QianchuanAdDetailGetV10DataDeliverySettingSmartBidType `json:"smart_bid_type,omitempty"`
	//
	StartTime         *string                                                      `json:"start_time,omitempty"`
	VideoScheduleType *QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType `json:"video_schedule_type,omitempty"`
}
