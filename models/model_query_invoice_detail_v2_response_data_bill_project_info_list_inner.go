/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceDetailV2ResponseDataBillProjectInfoListInner struct for QueryInvoiceDetailV2ResponseDataBillProjectInfoListInner
type QueryInvoiceDetailV2ResponseDataBillProjectInfoListInner struct {
	// 开票项目金额
	InvoiceProjectAmount *float64 `json:"invoice_project_amount,omitempty"`
	// 开票项目名称
	InvoiceProjectName *string `json:"invoice_project_name,omitempty"`
	// 开具日期
	IssueDate *string `json:"issue_date,omitempty"`
	// 打印备注
	Remark *string `json:"remark,omitempty"`
}
