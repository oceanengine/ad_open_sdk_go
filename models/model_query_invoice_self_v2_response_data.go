/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryInvoiceSelfV2ResponseData
type QueryInvoiceSelfV2ResponseData struct {
	// 开票列表
	InvoiceInfoList []*QueryInvoiceSelfV2ResponseDataInvoiceInfoListInner `json:"invoice_info_list"`
	PageInfo        *QueryInvoiceSelfV2ResponseDataPageInfo               `json:"page_info,omitempty"`
}
