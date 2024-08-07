/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletTransactionDetailGetV30ResponseDataResultsInner struct for SharedWalletTransactionDetailGetV30ResponseDataResultsInner
type SharedWalletTransactionDetailGetV30ResponseDataResultsInner struct {
	// 总金额(单位元)
	Amount  *float64                                               `json:"amount,omitempty"`
	BizType *SharedWalletTransactionDetailGetV30DataResultsBizType `json:"biz_type,omitempty"`
	// 交易时间,精确到秒
	BusinessTime *string `json:"business_time,omitempty"`
	// 授信金额(单位元)
	CreditAmount *float64 `json:"credit_amount,omitempty"`
	// 收款钱包ID
	Payee *int64 `json:"payee,omitempty"`
	// 收款钱包名称
	PayeeName *string `json:"payee_name,omitempty"`
	// 预付金额(单位元)
	PrepayAmount *float64 `json:"prepay_amount,omitempty"`
	// 付款钱包ID
	Remitter *int64 `json:"remitter,omitempty"`
	// 付款钱包名称
	RemitterName *string `json:"remitter_name,omitempty"`
	// 交易流水id
	TransactionSeq *int64 `json:"transaction_seq,omitempty"`
}
