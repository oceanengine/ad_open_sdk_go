/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ServeMarketOrderGetV10DataOrderListSkuType
type ServeMarketOrderGetV10DataOrderListSkuType string

// List of serve_market_order_get_v1.0_data_order_list_sku_type
const (
	FREE_ServeMarketOrderGetV10DataOrderListSkuType ServeMarketOrderGetV10DataOrderListSkuType = "FREE"
	TRY_ServeMarketOrderGetV10DataOrderListSkuType  ServeMarketOrderGetV10DataOrderListSkuType = "TRY"
	PAY_ServeMarketOrderGetV10DataOrderListSkuType  ServeMarketOrderGetV10DataOrderListSkuType = "PAY"
)

// All allowed values of ServeMarketOrderGetV10DataOrderListSkuType enum
var AllowedServeMarketOrderGetV10DataOrderListSkuTypeEnumValues = []ServeMarketOrderGetV10DataOrderListSkuType{
	"FREE",
	"TRY",
	"PAY",
}

// NewServeMarketOrderGetV10DataOrderListSkuTypeFromValue returns a pointer to a valid ServeMarketOrderGetV10DataOrderListSkuType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServeMarketOrderGetV10DataOrderListSkuTypeFromValue(v string) (*ServeMarketOrderGetV10DataOrderListSkuType, error) {
	ev := ServeMarketOrderGetV10DataOrderListSkuType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServeMarketOrderGetV10DataOrderListSkuType: valid values are %v", v, AllowedServeMarketOrderGetV10DataOrderListSkuTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServeMarketOrderGetV10DataOrderListSkuType) IsValid() bool {
	for _, existing := range AllowedServeMarketOrderGetV10DataOrderListSkuTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to serve_market_order_get_v1.0_data_order_list_sku_type value
func (v ServeMarketOrderGetV10DataOrderListSkuType) Ptr() *ServeMarketOrderGetV10DataOrderListSkuType {
	return &v
}
