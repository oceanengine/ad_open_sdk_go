/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarStarAdUniteTaskListV2ResponseDataDemandsInner struct for StarStarAdUniteTaskListV2ResponseDataDemandsInner
type StarStarAdUniteTaskListV2ResponseDataDemandsInner struct {
	// 任务创建时间，格式：yyyy-mm-dd HH:MM:SS
	CreateTime *string `json:"create_time,omitempty"`
	// 任务id
	DemandId *int64 `json:"demand_id,omitempty"`
	// 任务名称
	DemandName *string `json:"demand_name,omitempty"`
}
