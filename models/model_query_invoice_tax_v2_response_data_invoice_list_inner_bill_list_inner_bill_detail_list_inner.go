/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInnerBillDetailListInner struct for QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInnerBillDetailListInner
type QueryInvoiceTaxV2ResponseDataInvoiceListInnerBillListInnerBillDetailListInner struct {
	//
	Amount *float64 `json:"amount,omitempty"`
	//
	CommodifyServiceName *string `json:"commodify_service_name,omitempty"`
	//
	SpecificationsModels *string `json:"specifications_models,omitempty"`
	//
	TaxAmount *float64 `json:"tax_amount,omitempty"`
	//
	TaxRate *float64 `json:"tax_rate,omitempty"`
	//
	TotalPriceTax *float64 `json:"total_price_tax,omitempty"`
	//
	Unit *string `json:"unit,omitempty"`
}
