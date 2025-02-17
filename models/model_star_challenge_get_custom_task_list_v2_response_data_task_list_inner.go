/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeGetCustomTaskListV2ResponseDataTaskListInner struct for StarChallengeGetCustomTaskListV2ResponseDataTaskListInner
type StarChallengeGetCustomTaskListV2ResponseDataTaskListInner struct {
	// 品牌名称
	BrandName *string `json:"brand_name,omitempty"`
	// 任务id
	ChallengeTaskId *int64 `json:"challenge_task_id,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 任务名称
	DemandName *string `json:"demand_name,omitempty"`
	// 客户id
	DemanderId *int64 `json:"demander_id,omitempty"`
	// 投稿结束时间
	EndTime *int64 `json:"end_time,omitempty"`
	// 任务头图
	HeadImageUrl *string `json:"head_image_url,omitempty"`
	// 投稿开始时间
	StartTime *int64 `json:"start_time,omitempty"`
	// 任务类型
	TaskCategory *int64 `json:"task_category,omitempty"`
	// 任务状态
	TaskStatus *int64 `json:"task_status,omitempty"`
	// 结算方式
	UniversalSettlementType *int64 `json:"universal_settlement_type,omitempty"`
}
