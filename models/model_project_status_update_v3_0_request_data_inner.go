/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectStatusUpdateV30RequestDataInner struct for ProjectStatusUpdateV30RequestDataInner
type ProjectStatusUpdateV30RequestDataInner struct {
	OptStatus ProjectStatusUpdateV30DataOptStatus `json:"opt_status"`
	// 项目ID
	ProjectId int64 `json:"project_id"`
}
