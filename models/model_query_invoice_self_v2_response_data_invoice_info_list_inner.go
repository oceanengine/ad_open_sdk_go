/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceSelfV2ResponseDataInvoiceInfoListInner struct for QueryInvoiceSelfV2ResponseDataInvoiceInfoListInner
type QueryInvoiceSelfV2ResponseDataInvoiceInfoListInner struct {
	// 作废金额
	AbandonedAmount *float64 `json:"abandoned_amount,omitempty"`
	// 实际开票金额
	ActualInvoiceAmount *float64 `json:"actual_invoice_amount,omitempty"`
	// 开票申请金额
	ApplyAmount *float64 `json:"apply_amount,omitempty"`
	// 合同名称
	ContractName *string `json:"contract_name,omitempty"`
	// 合同编号
	ContractSerial *string `json:"contract_serial,omitempty"`
	// 签约主体
	ContractSubjectName *string `json:"contract_subject_name,omitempty"`
	// 客户ID
	CustomerId *int64 `json:"customer_id,omitempty"`
	// 客户名称
	CustomerName *string `json:"customer_name,omitempty"`
	// 开票客户名称
	CustomerSubjectName *string `json:"customer_subject_name,omitempty"`
	// 是否差额开票
	DifferenceInvoiceName *string `json:"difference_invoice_name,omitempty"`
	// 开票金额
	InvoiceAmount *float64 `json:"invoice_amount,omitempty"`
	// 发票代码
	InvoiceCodes []string `json:"invoice_codes,omitempty"`
	// 开票单ID
	InvoiceId int64 `json:"invoice_id"`
	// 发票号码
	InvoiceNos []string `json:"invoice_nos,omitempty"`
	// 开票类型名称
	InvoiceObjectName *string `json:"invoice_object_name,omitempty"`
	// 开票项目名称
	InvoiceProjectList []string `json:"invoice_project_list,omitempty"`
	// 开票单编号
	InvoiceSerial string `json:"invoice_serial"`
	// 开票状态 CANCELLED(0, \"已作废\"),NO_INVOICED(1, \"未开票\"),PART_INVOICED(2, \"部分开票\"),INVOICED(3, \"已开票\"),NO_NEED_INVOICE(4, \"无需开票\");
	InvoiceStatus *int64 `json:"invoice_status,omitempty"`
	// 开票状态名称
	InvoiceStatusName *string `json:"invoice_status_name,omitempty"`
	// 发票类型
	InvoiceType *int64 `json:"invoice_type,omitempty"`
	// 发票类型名称
	InvoiceTypeName *string `json:"invoice_type_name,omitempty"`
	// 发票开具时间 yyyy-MM-dd HH:mm:ss
	IssueTime *string `json:"issue_time,omitempty"`
	// 业务线
	PlatformName *string `json:"platform_name,omitempty"`
	// 红冲金额
	RevertAmount *float64 `json:"revert_amount,omitempty"`
	// 红冲冻结金额
	RevertFrozenAmount *float64 `json:"revert_frozen_amount,omitempty"`
	// 红冲状态  NORMAL(1, \"未红冲\"),REVERTED(2, \"全额红冲\"),RED(3, \"红票\"),PART_REVERTED(4, \"部分红冲\");
	RevertStatus *int64 `json:"revert_status,omitempty"`
	// 红冲状态名称
	RevertStatusName *string `json:"revert_status_name,omitempty"`
	// 提交时间 yyyy-MM-dd HH:mm:ss
	SubmitTime *string `json:"submit_time,omitempty"`
	// 提交人名称
	SubmitterName *string `json:"submitter_name,omitempty"`
}