/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupDeleteV30ResponseData
type BudgetGroupDeleteV30ResponseData struct {
	// 删除成功的预算组id
	BudgetGroupIds []int64 `json:"budget_group_ids,omitempty"`
	// 删除失败的预算组
	Errors []*BudgetGroupDeleteV30ResponseDataErrorsInner `json:"errors,omitempty"`
}
