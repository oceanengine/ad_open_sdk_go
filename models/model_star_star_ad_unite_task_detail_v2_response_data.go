/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarStarAdUniteTaskDetailV2ResponseData
type StarStarAdUniteTaskDetailV2ResponseData struct {
	AuditStatus StarStarAdUniteTaskDetailV2DataAuditStatus `json:"audit_status"`
	// 任务预算，单位：元*1000，建议获取后除以1000展示为元单位
	Budget int64 `json:"budget"`
	// 任务id
	DemandId     int64                                       `json:"demand_id"`
	EvaluateType StarStarAdUniteTaskDetailV2DataEvaluateType `json:"evaluate_type"`
	// 客户的星图id
	StarId int64 `json:"star_id"`
	//
	StatInfo []*StarStarAdUniteTaskDetailV2ResponseDataStatInfoInner `json:"stat_info"`
	Status   StarStarAdUniteTaskDetailV2DataStatus                   `json:"status"`
	// 投放时段(没返回就是全时段投放)
	WeekSchedule []*[]int64 `json:"week_schedule,omitempty"`
}
