/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CustomerCenterAccountListV30ResponseDataAccountsInner struct for CustomerCenterAccountListV30ResponseDataAccountsInner
type CustomerCenterAccountListV30ResponseDataAccountsInner struct {
	// 账户ID，对应账户类型参考account_type
	AccountId *int64 `json:"account_id,omitempty"`
	// 账户名称
	AccountName string `json:"account_name"`
	// 账户类型，枚举值： - `AD` 巨量广告 - `QIANCHUAN` 巨量千川 - `LOCAL` 巨量本地推 - `DOU+` DOU+
	AccountType *string `json:"account_type,omitempty"`
}
