/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasCancelBoostItemGroupV2Request struct for StarVasCancelBoostItemGroupV2Request
type StarVasCancelBoostItemGroupV2Request struct {
	// 客户ID
	StarId int64 `json:"star_id"`
	// 助推组ID
	TaskId int64 `json:"task_id"`
}
