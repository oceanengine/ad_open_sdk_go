/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdGetV10ResponseDataListInnerDeliverySetting
type QianchuanAdGetV10ResponseDataListInnerDeliverySetting struct {
	//
	AllowQcpx *bool `json:"allow_qcpx,omitempty"`
	//
	Budget     *float64                                            `json:"budget,omitempty"`
	BudgetMode *QianchuanAdGetV10DataListDeliverySettingBudgetMode `json:"budget_mode,omitempty"`
	//
	CpaBid      *float64                                             `json:"cpa_bid,omitempty"`
	DeepBidType *QianchuanAdGetV10DataListDeliverySettingDeepBidType `json:"deep_bid_type,omitempty"`
	//
	DeepCpaBid         *float64                                                    `json:"deep_cpa_bid,omitempty"`
	DeepExternalAction *QianchuanAdGetV10DataListDeliverySettingDeepExternalAction `json:"deep_external_action,omitempty"`
	//
	EndTime        *string                                                 `json:"end_time,omitempty"`
	ExternalAction *QianchuanAdGetV10DataListDeliverySettingExternalAction `json:"external_action,omitempty"`
	//
	ProductNewOpen *bool                                             `json:"product_new_open,omitempty"`
	QcpxMode       *QianchuanAdGetV10DataListDeliverySettingQcpxMode `json:"qcpx_mode,omitempty"`
	//
	ReviveBudget *float64 `json:"revive_budget,omitempty"`
	//
	RoiGoal      *float64                                              `json:"roi_goal,omitempty"`
	SmartBidType *QianchuanAdGetV10DataListDeliverySettingSmartBidType `json:"smart_bid_type,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
