/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AccountFundGetV30ResponseDataListInner struct for AccountFundGetV30ResponseDataListInner
type AccountFundGetV30ResponseDataListInner struct {
	//
	AccountId *int64 `json:"account_id,omitempty"`
	//
	Balance *int64 `json:"balance,omitempty"`
	//
	Cash *int64 `json:"cash,omitempty"`
	//
	CommonGrant *int64 `json:"common_grant,omitempty"`
	//
	CompensationGrant *int64 `json:"compensation_grant,omitempty"`
	//
	CompensationValidGrant *int64 `json:"compensation_valid_grant,omitempty"`
	//
	DefaultGrant *int64 `json:"default_grant,omitempty"`
	//
	Grant *int64 `json:"grant,omitempty"`
	//
	ReturnGoodsAbs *int64 `json:"return_goods_abs,omitempty"`
	//
	SearchGrant *int64 `json:"search_grant,omitempty"`
	//
	UnionGrant *int64 `json:"union_grant,omitempty"`
	//
	ValidBalance *int64 `json:"valid_balance,omitempty"`
	//
	ValidCash *int64 `json:"valid_cash,omitempty"`
	//
	ValidGrant *int64 `json:"valid_grant,omitempty"`
	//
	ValidReturnGoodsAbs *int64 `json:"valid_return_goods_abs,omitempty"`
	//
	WalletId *int64 `json:"wallet_id,omitempty"`
	//
	WalletTotalBalanceValid *int64 `json:"wallet_total_balance_valid,omitempty"`
}
