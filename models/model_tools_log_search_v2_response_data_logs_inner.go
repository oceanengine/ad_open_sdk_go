/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLogSearchV2ResponseDataLogsInner struct for ToolsLogSearchV2ResponseDataLogsInner
type ToolsLogSearchV2ResponseDataLogsInner struct {
	//
	ContentLog []string `json:"content_log,omitempty"`
	//
	ContentTitle *string `json:"content_title,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	ObjectId *int64 `json:"object_id,omitempty"`
	//
	ObjectName *string `json:"object_name,omitempty"`
	//
	ObjectType *string `json:"object_type,omitempty"`
	//
	Operator *string `json:"operator,omitempty"`
	//
	OptIp *string `json:"opt_ip,omitempty"`
}
