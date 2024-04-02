/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdvertiserTypeGetV10DataListShopBusinessType
type QianchuanAdvertiserTypeGetV10DataListShopBusinessType string

// List of qianchuan_advertiser_type_get_v1.0_data_list_shop_business_type
const (
	UNKOWN_QianchuanAdvertiserTypeGetV10DataListShopBusinessType      QianchuanAdvertiserTypeGetV10DataListShopBusinessType = "UNKOWN"
	SHOP_RETAIL_QianchuanAdvertiserTypeGetV10DataListShopBusinessType QianchuanAdvertiserTypeGetV10DataListShopBusinessType = "SHOP_RETAIL"
)

// All allowed values of QianchuanAdvertiserTypeGetV10DataListShopBusinessType enum
var AllowedQianchuanAdvertiserTypeGetV10DataListShopBusinessTypeEnumValues = []QianchuanAdvertiserTypeGetV10DataListShopBusinessType{
	"UNKOWN",
	"SHOP_RETAIL",
}

// NewQianchuanAdvertiserTypeGetV10DataListShopBusinessTypeFromValue returns a pointer to a valid QianchuanAdvertiserTypeGetV10DataListShopBusinessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdvertiserTypeGetV10DataListShopBusinessTypeFromValue(v string) (*QianchuanAdvertiserTypeGetV10DataListShopBusinessType, error) {
	ev := QianchuanAdvertiserTypeGetV10DataListShopBusinessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdvertiserTypeGetV10DataListShopBusinessType: valid values are %v", v, AllowedQianchuanAdvertiserTypeGetV10DataListShopBusinessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdvertiserTypeGetV10DataListShopBusinessType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdvertiserTypeGetV10DataListShopBusinessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_advertiser_type_get_v1.0_data_list_shop_business_type value
func (v QianchuanAdvertiserTypeGetV10DataListShopBusinessType) Ptr() *QianchuanAdvertiserTypeGetV10DataListShopBusinessType {
	return &v
}