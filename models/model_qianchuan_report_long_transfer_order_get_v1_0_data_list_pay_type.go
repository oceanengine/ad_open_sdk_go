/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportLongTransferOrderGetV10DataListPayType
type QianchuanReportLongTransferOrderGetV10DataListPayType string

// List of qianchuan_report_long_transfer_order_get_v1.0_data_list_pay_type
const (
	DIRECT_QianchuanReportLongTransferOrderGetV10DataListPayType   QianchuanReportLongTransferOrderGetV10DataListPayType = "DIRECT"
	INDIRECT_QianchuanReportLongTransferOrderGetV10DataListPayType QianchuanReportLongTransferOrderGetV10DataListPayType = "INDIRECT"
)

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_data_list_pay_type value
func (v QianchuanReportLongTransferOrderGetV10DataListPayType) Ptr() *QianchuanReportLongTransferOrderGetV10DataListPayType {
	return &v
}
