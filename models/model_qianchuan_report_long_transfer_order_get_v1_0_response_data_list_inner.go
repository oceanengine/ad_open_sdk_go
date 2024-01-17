/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportLongTransferOrderGetV10ResponseDataListInner struct for QianchuanReportLongTransferOrderGetV10ResponseDataListInner
type QianchuanReportLongTransferOrderGetV10ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdName *string                                              `json:"ad_name,omitempty"`
	IsPay  *QianchuanReportLongTransferOrderGetV10DataListIsPay `json:"is_pay,omitempty"`
	//
	OrderAmount     *float64                                                       `json:"order_amount,omitempty"`
	OrderFlowSource *QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource `json:"order_flow_source,omitempty"`
	//
	OrderId *int64 `json:"order_id,omitempty"`
	//
	OrderPlaceTime *string                                                `json:"order_place_time,omitempty"`
	PayType        *QianchuanReportLongTransferOrderGetV10DataListPayType `json:"pay_type,omitempty"`
}
