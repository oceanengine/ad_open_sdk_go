/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletTransactionDetailGetV30DataResultsRemitterType
type SharedWalletTransactionDetailGetV30DataResultsRemitterType string

// List of shared_wallet_transaction_detail_get_v3.0_data_results_remitter_type
const (
	AGENT_SharedWalletTransactionDetailGetV30DataResultsRemitterType         SharedWalletTransactionDetailGetV30DataResultsRemitterType = "AGENT"
	SHARED_WALLET_SharedWalletTransactionDetailGetV30DataResultsRemitterType SharedWalletTransactionDetailGetV30DataResultsRemitterType = "SHARED_WALLET"
)

// Ptr returns reference to shared_wallet_transaction_detail_get_v3.0_data_results_remitter_type value
func (v SharedWalletTransactionDetailGetV30DataResultsRemitterType) Ptr() *SharedWalletTransactionDetailGetV30DataResultsRemitterType {
	return &v
}