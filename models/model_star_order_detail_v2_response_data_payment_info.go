/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderDetailV2ResponseDataPaymentInfo 资金信息
type StarOrderDetailV2ResponseDataPaymentInfo struct {
	// 已扣除金额
	DeductAmount *int64 `json:"deduct_amount,omitempty"`
	// 服务费
	PlatformFee *int64 `json:"platform_fee,omitempty"`
	// 已退还金额
	RefundAmount *int64 `json:"refund_amount,omitempty"`
	// 任务金额
	TaskCost *int64 `json:"task_cost,omitempty"`
	// 任务总金额
	Total *int64 `json:"total,omitempty"`
	// 已付金额（元，下同）
	TotalPaid *int64 `json:"total_paid,omitempty"`
}
