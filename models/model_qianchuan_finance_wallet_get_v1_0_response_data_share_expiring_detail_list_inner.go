/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFinanceWalletGetV10ResponseDataShareExpiringDetailListInner struct for QianchuanFinanceWalletGetV10ResponseDataShareExpiringDetailListInner
type QianchuanFinanceWalletGetV10ResponseDataShareExpiringDetailListInner struct {
	// 金额
	Amount   *int64                                                           `json:"amount,omitempty"`
	Category *QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory `json:"category,omitempty"`
	// 到期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`
}
