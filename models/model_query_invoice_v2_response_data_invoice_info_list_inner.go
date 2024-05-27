/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceV2ResponseDataInvoiceInfoListInner struct for QueryInvoiceV2ResponseDataInvoiceInfoListInner
type QueryInvoiceV2ResponseDataInvoiceInfoListInner struct {
	//
	AbandonedAmount *float64 `json:"abandoned_amount,omitempty"`
	//
	ActualInvoiceAmount *float64 `json:"actual_invoice_amount,omitempty"`
	//
	ApplyAmount *float64 `json:"apply_amount,omitempty"`
	//
	InvoiceAmount *float64 `json:"invoice_amount,omitempty"`
	//
	InvoiceCodes []string `json:"invoice_codes,omitempty"`
	//
	InvoiceId int64 `json:"invoice_id"`
	//
	InvoiceNos []string `json:"invoice_nos,omitempty"`
	//
	InvoiceSerial string `json:"invoice_serial"`
	//
	InvoiceStatus *int64 `json:"invoice_status,omitempty"`
	//
	InvoiceStatusName *string `json:"invoice_status_name,omitempty"`
	//
	IssueTime *string `json:"issue_time,omitempty"`
	//
	RevertStatus *int64 `json:"revert_status,omitempty"`
	//
	RevertStatusName *string `json:"revert_status_name,omitempty"`
}
