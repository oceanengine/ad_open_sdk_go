/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialDownloadCreateTaskV2Request struct for FileRebateMaterialDownloadCreateTaskV2Request
type FileRebateMaterialDownloadCreateTaskV2Request struct {
	// 代理商帐户Id
	AgentId            int64                                                    `json:"agent_id"`
	Filtering          *FileRebateMaterialDownloadCreateTaskV2RequestFiltering  `json:"filtering,omitempty"`
	MaterialPolicyType FileRebateMaterialDownloadCreateTaskV2MaterialPolicyType `json:"material_policy_type"`
	// 业绩消耗年月 年月 格式：yyyy-MM
	Period string `json:"period"`
}
