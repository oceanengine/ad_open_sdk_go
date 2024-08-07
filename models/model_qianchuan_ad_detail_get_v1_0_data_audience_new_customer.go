/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudienceNewCustomer
type QianchuanAdDetailGetV10DataAudienceNewCustomer string

// List of qianchuan_ad_detail_get_v1.0_data_audience_new_customer
const (
	NONE_QianchuanAdDetailGetV10DataAudienceNewCustomer          QianchuanAdDetailGetV10DataAudienceNewCustomer = "NONE"
	NO_BUY_QianchuanAdDetailGetV10DataAudienceNewCustomer        QianchuanAdDetailGetV10DataAudienceNewCustomer = "NO_BUY"
	NO_BUY_BRAND_QianchuanAdDetailGetV10DataAudienceNewCustomer  QianchuanAdDetailGetV10DataAudienceNewCustomer = "NO_BUY_BRAND"
	NO_BUY_DOUYIN_QianchuanAdDetailGetV10DataAudienceNewCustomer QianchuanAdDetailGetV10DataAudienceNewCustomer = "NO_BUY_DOUYIN"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceNewCustomer enum
var AllowedQianchuanAdDetailGetV10DataAudienceNewCustomerEnumValues = []QianchuanAdDetailGetV10DataAudienceNewCustomer{
	"NONE",
	"NO_BUY",
	"NO_BUY_BRAND",
	"NO_BUY_DOUYIN",
}

// NewQianchuanAdDetailGetV10DataAudienceNewCustomerFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceNewCustomer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceNewCustomerFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceNewCustomer, error) {
	ev := QianchuanAdDetailGetV10DataAudienceNewCustomer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceNewCustomer: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceNewCustomerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceNewCustomer) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceNewCustomerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_new_customer value
func (v QianchuanAdDetailGetV10DataAudienceNewCustomer) Ptr() *QianchuanAdDetailGetV10DataAudienceNewCustomer {
	return &v
}
