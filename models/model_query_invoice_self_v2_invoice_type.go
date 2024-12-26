/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceSelfV2InvoiceType
type QueryInvoiceSelfV2InvoiceType string

// List of query_invoice_self_v2_invoice_type
const (
	GENERAL_QueryInvoiceSelfV2InvoiceType       QueryInvoiceSelfV2InvoiceType = "GENERAL"
	SPECIAL_QueryInvoiceSelfV2InvoiceType       QueryInvoiceSelfV2InvoiceType = "SPECIAL"
	E_GENERAL_QueryInvoiceSelfV2InvoiceType     QueryInvoiceSelfV2InvoiceType = "E_GENERAL"
	E_SPECIAL_QueryInvoiceSelfV2InvoiceType     QueryInvoiceSelfV2InvoiceType = "E_SPECIAL"
	ALL_E_GENERAL_QueryInvoiceSelfV2InvoiceType QueryInvoiceSelfV2InvoiceType = "ALL_E_GENERAL"
	ALL_E_SPECIAL_QueryInvoiceSelfV2InvoiceType QueryInvoiceSelfV2InvoiceType = "ALL_E_SPECIAL"
)

// Ptr returns reference to query_invoice_self_v2_invoice_type value
func (v QueryInvoiceSelfV2InvoiceType) Ptr() *QueryInvoiceSelfV2InvoiceType {
	return &v
}
