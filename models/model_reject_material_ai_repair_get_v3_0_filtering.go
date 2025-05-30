/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// RejectMaterialAiRepairGetV30Filtering
type RejectMaterialAiRepairGetV30Filtering struct {
	//
	AiRepairIds []int64 `json:"ai_repair_ids,omitempty"`
	//
	MaterialIds []int64 `json:"material_ids,omitempty"`
	//
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
	//
	RepairEndTime *string `json:"repair_end_time,omitempty"`
	//
	RepairStartTime *string `json:"repair_start_time,omitempty"`
}
