/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderGetV10OrderField
type QianchuanAwemeOrderGetV10OrderField string

// List of qianchuan_aweme_order_get_v1.0_order_field
const (
	ORDER_CREATE_TIME_QianchuanAwemeOrderGetV10OrderField        QianchuanAwemeOrderGetV10OrderField = "order_create_time"
	PAY_ORDER_AMOUNT_QianchuanAwemeOrderGetV10OrderField         QianchuanAwemeOrderGetV10OrderField = "pay_order_amount"
	PREPAY_AND_PAY_ORDER_ROI_QianchuanAwemeOrderGetV10OrderField QianchuanAwemeOrderGetV10OrderField = "prepay_and_pay_order_roi"
	STAT_COST_QianchuanAwemeOrderGetV10OrderField                QianchuanAwemeOrderGetV10OrderField = "stat_cost"
)

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_order_field value
func (v QianchuanAwemeOrderGetV10OrderField) Ptr() *QianchuanAwemeOrderGetV10OrderField {
	return &v
}
