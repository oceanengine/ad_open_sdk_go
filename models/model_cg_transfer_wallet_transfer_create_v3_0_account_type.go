/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferCreateV30AccountType
type CgTransferWalletTransferCreateV30AccountType string

// List of cg_transfer_wallet_transfer_create_v3.0_account_type
const (
	AD_CgTransferWalletTransferCreateV30AccountType        CgTransferWalletTransferCreateV30AccountType = "AD"
	AGENT_CgTransferWalletTransferCreateV30AccountType     CgTransferWalletTransferCreateV30AccountType = "AGENT"
	LOCAL_CgTransferWalletTransferCreateV30AccountType     CgTransferWalletTransferCreateV30AccountType = "LOCAL"
	QIANCHUAN_CgTransferWalletTransferCreateV30AccountType CgTransferWalletTransferCreateV30AccountType = "QIANCHUAN"
)

// Ptr returns reference to cg_transfer_wallet_transfer_create_v3.0_account_type value
func (v CgTransferWalletTransferCreateV30AccountType) Ptr() *CgTransferWalletTransferCreateV30AccountType {
	return &v
}
