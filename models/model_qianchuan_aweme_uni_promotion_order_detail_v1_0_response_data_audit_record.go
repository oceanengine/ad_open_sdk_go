/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeUniPromotionOrderDetailV10ResponseDataAuditRecord 订单状态的详细补充（审核等）
type QianchuanAwemeUniPromotionOrderDetailV10ResponseDataAuditRecord struct {
	// 详细的审核建议和拒绝理由
	DetailDescList []string `json:"detail_desc_list,omitempty"`
	// 详细的状态描述（二级标题）
	StatusAttachDesp *string `json:"status_attach_desp,omitempty"`
}
