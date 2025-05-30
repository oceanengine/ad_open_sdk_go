/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityOpenMaterialAuditV30ResponseData
type SecurityOpenMaterialAuditV30ResponseData struct {
	// 每次送审自动生成的对象id
	ObjectId int64 `json:"object_id"`
	// 送审是否成功
	Result bool `json:"result"`
	// 链路task_id，每次送审生成一个唯一值，贯穿全流程
	TaskId int64 `json:"task_id"`
}
