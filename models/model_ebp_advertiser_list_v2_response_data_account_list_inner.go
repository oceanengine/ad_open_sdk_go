/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EbpAdvertiserListV2ResponseDataAccountListInner struct for EbpAdvertiserListV2ResponseDataAccountListInner
type EbpAdvertiserListV2ResponseDataAccountListInner struct {
	// 客户id
	AccountId *int64 `json:"account_id,omitempty"`
	//
	AccountName *string                                        `json:"account_name,omitempty"`
	AccountType *EbpAdvertiserListV2DataAccountListAccountType `json:"account_type,omitempty"`
}
