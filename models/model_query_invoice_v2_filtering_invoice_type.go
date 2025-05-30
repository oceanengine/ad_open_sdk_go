/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceV2FilteringInvoiceType
type QueryInvoiceV2FilteringInvoiceType string

// List of query_invoice_v2_filtering_invoice_type
const (
	GENERAL_QueryInvoiceV2FilteringInvoiceType       QueryInvoiceV2FilteringInvoiceType = "GENERAL"
	SPECIAL_QueryInvoiceV2FilteringInvoiceType       QueryInvoiceV2FilteringInvoiceType = "SPECIAL"
	E_GENERAL_QueryInvoiceV2FilteringInvoiceType     QueryInvoiceV2FilteringInvoiceType = "E_GENERAL"
	E_SPECIAL_QueryInvoiceV2FilteringInvoiceType     QueryInvoiceV2FilteringInvoiceType = "E_SPECIAL"
	ALL_E_GENERAL_QueryInvoiceV2FilteringInvoiceType QueryInvoiceV2FilteringInvoiceType = "ALL_E_GENERAL"
	ALL_E_SPECIAL_QueryInvoiceV2FilteringInvoiceType QueryInvoiceV2FilteringInvoiceType = "ALL_E_SPECIAL"
)

// Ptr returns reference to query_invoice_v2_filtering_invoice_type value
func (v QueryInvoiceV2FilteringInvoiceType) Ptr() *QueryInvoiceV2FilteringInvoiceType {
	return &v
}
