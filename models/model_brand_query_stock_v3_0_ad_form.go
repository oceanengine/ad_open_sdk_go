/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandQueryStockV30AdForm
type BrandQueryStockV30AdForm string

// List of brand_query_stock_v3.0_ad_form
const (
	CONTENT_RECOMMEND_BrandQueryStockV30AdForm BrandQueryStockV30AdForm = "CONTENT_RECOMMEND"
	CONTENT_SERVICE_BrandQueryStockV30AdForm   BrandQueryStockV30AdForm = "CONTENT_SERVICE"
)

// All allowed values of BrandQueryStockV30AdForm enum
var AllowedBrandQueryStockV30AdFormEnumValues = []BrandQueryStockV30AdForm{
	"CONTENT_RECOMMEND",
	"CONTENT_SERVICE",
}

// NewBrandQueryStockV30AdFormFromValue returns a pointer to a valid BrandQueryStockV30AdForm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AdFormFromValue(v string) (*BrandQueryStockV30AdForm, error) {
	ev := BrandQueryStockV30AdForm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AdForm: valid values are %v", v, AllowedBrandQueryStockV30AdFormEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AdForm) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AdFormEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_ad_form value
func (v BrandQueryStockV30AdForm) Ptr() *BrandQueryStockV30AdForm {
	return &v
}
