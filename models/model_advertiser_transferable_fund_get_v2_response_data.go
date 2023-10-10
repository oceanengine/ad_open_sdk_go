/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserTransferableFundGetV2ResponseData
type AdvertiserTransferableFundGetV2ResponseData struct {
	//
	BidCreditValid *float64 `json:"bid_credit_valid,omitempty"`
	//
	BidPrepayValid *float64 `json:"bid_prepay_valid,omitempty"`
	//
	BrandCreditValid *float64 `json:"brand_credit_valid,omitempty"`
	//
	BrandPrepayValid *float64 `json:"brand_prepay_valid,omitempty"`
	//
	CashTransferBalance *float64 `json:"cash_transfer_balance,omitempty"`
	//
	DepositAmount *float64 `json:"deposit_amount,omitempty"`
	//
	GrantValid *float64 `json:"grant_valid,omitempty"`
	//
	TotalTransferBalance *float64 `json:"total_transfer_balance,omitempty"`
	//
	UniversalCreditValid *float64 `json:"universal_credit_valid,omitempty"`
	//
	UniversalPrepayValid *float64 `json:"universal_prepay_valid,omitempty"`
}
