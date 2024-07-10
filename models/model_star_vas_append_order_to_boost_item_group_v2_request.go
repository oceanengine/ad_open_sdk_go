/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasAppendOrderToBoostItemGroupV2Request struct for StarVasAppendOrderToBoostItemGroupV2Request
type StarVasAppendOrderToBoostItemGroupV2Request struct {
	// 助推金额，单位元。若为null，则不追加
	BoostAmount *int64 `json:"boost_amount,omitempty"`
	// 关联指派单。若为null，则不追加
	OrderIds []int64 `json:"order_ids,omitempty"`
	// 客户ID
	StarId int64 `json:"star_id"`
	// 助推组ID
	TaskId int64 `json:"task_id"`
}
