/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30ResponseData
type BudgetGroupListV30ResponseData struct {
	//
	BudgetGroups []*BudgetGroupListV30ResponseDataBudgetGroupsInner `json:"budget_groups,omitempty"`
	PageInfo     *BudgetGroupListV30ResponseDataPageInfo            `json:"page_info,omitempty"`
}
