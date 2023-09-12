/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPreAuditGetV2Filter
type ToolsPreAuditGetV2Filter struct {
	// 前置预审物料 - 文本物料：用户上传文本 - 图片物料：用户上传图片链接或图片ID - 视频物料：用户上传的视频云ID - 落地页物料： 用户上传落地页URL
	Contents     []string                              `json:"contents,omitempty"`
	MaterialType *ToolsPreAuditGetV2FilterMaterialType `json:"material_type,omitempty"`
	// 前置预审ID列表
	PreAuditIds []int64 `json:"pre_audit_ids,omitempty"`
}
