/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundGrantTransactionGetV2DataListTransactionType
type AdvertiserFundGrantTransactionGetV2DataListTransactionType string

// List of advertiser_fund_grant_transaction_get_v2_data_list_transaction_type
const (
	RECHARGE_AdvertiserFundGrantTransactionGetV2DataListTransactionType AdvertiserFundGrantTransactionGetV2DataListTransactionType = "RECHARGE"
	TRANSFER_AdvertiserFundGrantTransactionGetV2DataListTransactionType AdvertiserFundGrantTransactionGetV2DataListTransactionType = "TRANSFER"
)

// Ptr returns reference to advertiser_fund_grant_transaction_get_v2_data_list_transaction_type value
func (v AdvertiserFundGrantTransactionGetV2DataListTransactionType) Ptr() *AdvertiserFundGrantTransactionGetV2DataListTransactionType {
	return &v
}
