/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30Filtering
type BudgetGroupListV30Filtering struct {
	AdType *BudgetGroupListV30FilteringAdType `json:"ad_type,omitempty"`
	// 预算组id列表
	BudgetGroupIdList []int64 `json:"budget_group_id_list,omitempty"`
	// 预算组名称模糊搜索关键字
	BudgetGroupKeyword      *string                                             `json:"budget_group_keyword,omitempty"`
	BudgetGroupStatusFirst  *BudgetGroupListV30FilteringBudgetGroupStatusFirst  `json:"budget_group_status_first,omitempty"`
	BudgetGroupStatusSecond *BudgetGroupListV30FilteringBudgetGroupStatusSecond `json:"budget_group_status_second,omitempty"`
	DeliveryType            *BudgetGroupListV30FilteringDeliveryType            `json:"delivery_type,omitempty"`
	StatisticsTime          *BudgetGroupListV30FilteringStatisticsTime          `json:"statistics_time,omitempty"`
}
