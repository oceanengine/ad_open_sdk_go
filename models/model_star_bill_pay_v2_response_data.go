/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarBillPayV2ResponseData
type StarBillPayV2ResponseData struct {
	// 实付金额（单位元，下同）
	ActualPay *int64 `json:"actual_pay,omitempty"`
	// 付款后账户余额
	Balance *int64 `json:"balance,omitempty"`
	// 精确实付金额（单位元，下同）
	PreciseActualPay *float64 `json:"precise_actual_pay,omitempty"`
	// 精确付款后账户余额
	PreciseBalance *float64 `json:"precise_balance,omitempty"`
}
