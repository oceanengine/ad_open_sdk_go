/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeProductAvailableGetV10DataProductListChannelType
type QianchuanAwemeProductAvailableGetV10DataProductListChannelType string

// List of qianchuan_aweme_product_available_get_v1.0_data_product_list_channel_type
const (
	SHOP_SELL_QianchuanAwemeProductAvailableGetV10DataProductListChannelType QianchuanAwemeProductAvailableGetV10DataProductListChannelType = "SHOP_SELL"
	STAR_SELL_QianchuanAwemeProductAvailableGetV10DataProductListChannelType QianchuanAwemeProductAvailableGetV10DataProductListChannelType = "STAR_SELL"
)

// All allowed values of QianchuanAwemeProductAvailableGetV10DataProductListChannelType enum
var AllowedQianchuanAwemeProductAvailableGetV10DataProductListChannelTypeEnumValues = []QianchuanAwemeProductAvailableGetV10DataProductListChannelType{
	"SHOP_SELL",
	"STAR_SELL",
}

// NewQianchuanAwemeProductAvailableGetV10DataProductListChannelTypeFromValue returns a pointer to a valid QianchuanAwemeProductAvailableGetV10DataProductListChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeProductAvailableGetV10DataProductListChannelTypeFromValue(v string) (*QianchuanAwemeProductAvailableGetV10DataProductListChannelType, error) {
	ev := QianchuanAwemeProductAvailableGetV10DataProductListChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeProductAvailableGetV10DataProductListChannelType: valid values are %v", v, AllowedQianchuanAwemeProductAvailableGetV10DataProductListChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeProductAvailableGetV10DataProductListChannelType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeProductAvailableGetV10DataProductListChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_product_available_get_v1.0_data_product_list_channel_type value
func (v QianchuanAwemeProductAvailableGetV10DataProductListChannelType) Ptr() *QianchuanAwemeProductAvailableGetV10DataProductListChannelType {
	return &v
}
