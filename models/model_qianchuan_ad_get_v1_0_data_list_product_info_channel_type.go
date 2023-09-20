/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListProductInfoChannelType
type QianchuanAdGetV10DataListProductInfoChannelType string

// List of qianchuan_ad_get_v1.0_data_list_product_info_channel_type
const (
	SHOP_SELL_QianchuanAdGetV10DataListProductInfoChannelType QianchuanAdGetV10DataListProductInfoChannelType = "SHOP_SELL"
	STAR_SELL_QianchuanAdGetV10DataListProductInfoChannelType QianchuanAdGetV10DataListProductInfoChannelType = "STAR_SELL"
)

// All allowed values of QianchuanAdGetV10DataListProductInfoChannelType enum
var AllowedQianchuanAdGetV10DataListProductInfoChannelTypeEnumValues = []QianchuanAdGetV10DataListProductInfoChannelType{
	"SHOP_SELL",
	"STAR_SELL",
}

// NewQianchuanAdGetV10DataListProductInfoChannelTypeFromValue returns a pointer to a valid QianchuanAdGetV10DataListProductInfoChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListProductInfoChannelTypeFromValue(v string) (*QianchuanAdGetV10DataListProductInfoChannelType, error) {
	ev := QianchuanAdGetV10DataListProductInfoChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListProductInfoChannelType: valid values are %v", v, AllowedQianchuanAdGetV10DataListProductInfoChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListProductInfoChannelType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListProductInfoChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_product_info_channel_type value
func (v QianchuanAdGetV10DataListProductInfoChannelType) Ptr() *QianchuanAdGetV10DataListProductInfoChannelType {
	return &v
}
