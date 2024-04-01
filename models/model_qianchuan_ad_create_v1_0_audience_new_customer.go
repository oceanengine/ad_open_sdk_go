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

// QianchuanAdCreateV10AudienceNewCustomer
type QianchuanAdCreateV10AudienceNewCustomer string

// List of qianchuan_ad_create_v1.0_audience_new_customer
const (
	NONE_QianchuanAdCreateV10AudienceNewCustomer          QianchuanAdCreateV10AudienceNewCustomer = "NONE"
	NO_BUY_QianchuanAdCreateV10AudienceNewCustomer        QianchuanAdCreateV10AudienceNewCustomer = "NO_BUY"
	NO_BUY_BRAND_QianchuanAdCreateV10AudienceNewCustomer  QianchuanAdCreateV10AudienceNewCustomer = "NO_BUY_BRAND"
	NO_BUY_DOUYIN_QianchuanAdCreateV10AudienceNewCustomer QianchuanAdCreateV10AudienceNewCustomer = "NO_BUY_DOUYIN"
)

// All allowed values of QianchuanAdCreateV10AudienceNewCustomer enum
var AllowedQianchuanAdCreateV10AudienceNewCustomerEnumValues = []QianchuanAdCreateV10AudienceNewCustomer{
	"NONE",
	"NO_BUY",
	"NO_BUY_BRAND",
	"NO_BUY_DOUYIN",
}

// NewQianchuanAdCreateV10AudienceNewCustomerFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceNewCustomer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceNewCustomerFromValue(v string) (*QianchuanAdCreateV10AudienceNewCustomer, error) {
	ev := QianchuanAdCreateV10AudienceNewCustomer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceNewCustomer: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceNewCustomerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceNewCustomer) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceNewCustomerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_new_customer value
func (v QianchuanAdCreateV10AudienceNewCustomer) Ptr() *QianchuanAdCreateV10AudienceNewCustomer {
	return &v
}
