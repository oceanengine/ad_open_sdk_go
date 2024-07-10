/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostItemGroupDetailV2ResponseDataOrderStatInfoInnerOrderInfo 指派单信息
type StarVasGetBoostItemGroupDetailV2ResponseDataOrderStatInfoInnerOrderInfo struct {
	// 任务助推审核状态
	AuditStatus int64 `json:"audit_status"`
	// 达人ID
	AuthorId int64 `json:"author_id"`
	// 达人名称
	AuthorName string `json:"author_name"`
	// 视频名称
	ItemName string `json:"item_name"`
	// 指派单ID
	OrderId int64 `json:"order_id"`
}
