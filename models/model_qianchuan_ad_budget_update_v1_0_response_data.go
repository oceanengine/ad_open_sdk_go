/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdBudgetUpdateV10ResponseData
type QianchuanAdBudgetUpdateV10ResponseData struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	Errors []*QianchuanAdBudgetUpdateV10ResponseDataErrorsInner `json:"errors,omitempty"`
}
