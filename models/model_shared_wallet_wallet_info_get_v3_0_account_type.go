/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletInfoGetV30AccountType
type SharedWalletWalletInfoGetV30AccountType string

// List of shared_wallet_wallet_info_get_v3.0_account_type
const (
	AD_SharedWalletWalletInfoGetV30AccountType        SharedWalletWalletInfoGetV30AccountType = "AD"
	AGENT_SharedWalletWalletInfoGetV30AccountType     SharedWalletWalletInfoGetV30AccountType = "AGENT"
	LOCAL_SharedWalletWalletInfoGetV30AccountType     SharedWalletWalletInfoGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletWalletInfoGetV30AccountType SharedWalletWalletInfoGetV30AccountType = "QIANCHUAN"
)

// Ptr returns reference to shared_wallet_wallet_info_get_v3.0_account_type value
func (v SharedWalletWalletInfoGetV30AccountType) Ptr() *SharedWalletWalletInfoGetV30AccountType {
	return &v
}
