/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreateProjectInvoiceV2RequestInvoiceBillListInnerInvoiceBillProjectListInner struct for CreateProjectInvoiceV2RequestInvoiceBillListInnerInvoiceBillProjectListInner
type CreateProjectInvoiceV2RequestInvoiceBillListInnerInvoiceBillProjectListInner struct {
	// 开票项目金额
	ApplyAmount float64 `json:"apply_amount"`
	// 开票项目名称
	InvoiceProjectName string `json:"invoice_project_name"`
	// 规则型号
	Specifications *string `json:"specifications,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty"`
}
