/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundTransactionGetV2DataListTransactionType
type AdvertiserFundTransactionGetV2DataListTransactionType string

// List of advertiser_fund_transaction_get_v2_data_list_transaction_type
const (
	RECHARGE_AdvertiserFundTransactionGetV2DataListTransactionType AdvertiserFundTransactionGetV2DataListTransactionType = "RECHARGE"
	TRANSFER_AdvertiserFundTransactionGetV2DataListTransactionType AdvertiserFundTransactionGetV2DataListTransactionType = "TRANSFER"
)

// Ptr returns reference to advertiser_fund_transaction_get_v2_data_list_transaction_type value
func (v AdvertiserFundTransactionGetV2DataListTransactionType) Ptr() *AdvertiserFundTransactionGetV2DataListTransactionType {
	return &v
}
