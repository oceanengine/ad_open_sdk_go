/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandQueryStockV30GdSendType
type BrandQueryStockV30GdSendType string

// List of brand_query_stock_v3.0_gd_send_type
const (
	CTR_BrandQueryStockV30GdSendType         BrandQueryStockV30GdSendType = "CTR"
	FANS_INCR_BrandQueryStockV30GdSendType   BrandQueryStockV30GdSendType = "FANS_INCR"
	FORM_BrandQueryStockV30GdSendType        BrandQueryStockV30GdSendType = "FORM"
	HOISTING_BrandQueryStockV30GdSendType    BrandQueryStockV30GdSendType = "HOISTING"
	INTERACTIVE_BrandQueryStockV30GdSendType BrandQueryStockV30GdSendType = "INTERACTIVE"
	PLANT_GRASS_BrandQueryStockV30GdSendType BrandQueryStockV30GdSendType = "PLANT_GRASS"
	SEQUENCE_BrandQueryStockV30GdSendType    BrandQueryStockV30GdSendType = "SEQUENCE"
)

// All allowed values of BrandQueryStockV30GdSendType enum
var AllowedBrandQueryStockV30GdSendTypeEnumValues = []BrandQueryStockV30GdSendType{
	"CTR",
	"FANS_INCR",
	"FORM",
	"HOISTING",
	"INTERACTIVE",
	"PLANT_GRASS",
	"SEQUENCE",
}

// NewBrandQueryStockV30GdSendTypeFromValue returns a pointer to a valid BrandQueryStockV30GdSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30GdSendTypeFromValue(v string) (*BrandQueryStockV30GdSendType, error) {
	ev := BrandQueryStockV30GdSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30GdSendType: valid values are %v", v, AllowedBrandQueryStockV30GdSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30GdSendType) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30GdSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_gd_send_type value
func (v BrandQueryStockV30GdSendType) Ptr() *BrandQueryStockV30GdSendType {
	return &v
}
