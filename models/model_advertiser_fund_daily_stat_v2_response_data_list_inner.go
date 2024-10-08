/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundDailyStatV2ResponseDataListInner struct for AdvertiserFundDailyStatV2ResponseDataListInner
type AdvertiserFundDailyStatV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Balance *float64 `json:"balance,omitempty"`
	//
	CashCost *float64 `json:"cash_cost,omitempty"`
	//
	CompanyWalletCost *float64 `json:"company_wallet_cost,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	Date *string `json:"date,omitempty"`
	//
	Frozen *float64 `json:"frozen,omitempty"`
	//
	GrantBalance *float64 `json:"grant_balance,omitempty"`
	//
	Income *float64 `json:"income,omitempty"`
	//
	NonGrantBalance *float64 `json:"non_grant_balance,omitempty"`
	//
	RewardCost *float64 `json:"reward_cost,omitempty"`
	//
	SharedWalletCost *float64 `json:"shared_wallet_cost,omitempty"`
	//
	TransferIn *float64 `json:"transfer_in,omitempty"`
	//
	TransferOut *float64 `json:"transfer_out,omitempty"`
}
