/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductCreateV2RequestProductInfoPriceInfo
type DpaProductCreateV2RequestProductInfoPriceInfo struct {
	//
	DailyMortgage *string `json:"daily_mortgage,omitempty"`
	//
	Discount *float64 `json:"discount,omitempty"`
	//
	DownPayment *string `json:"down_payment,omitempty"`
	//
	Mortgage *string `json:"mortgage,omitempty"`
	//
	Price *float64 `json:"price,omitempty"`
	//
	PriceUnit *string `json:"price_unit,omitempty"`
	//
	SalesPromotion *string `json:"sales_promotion,omitempty"`
	//
	Saving *float64 `json:"saving,omitempty"`
	//
	Value *float64 `json:"value,omitempty"`
}
