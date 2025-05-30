/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalChargeSubmitV30ResponseData
type LocalChargeSubmitV30ResponseData struct {
	// 支付链接，用来跳转收银台
	CashierUrl *string `json:"cashier_url,omitempty"`
	// 充值单ID，用来查询充值结果
	ChargeOrderId *int64 `json:"charge_order_id,omitempty"`
}
