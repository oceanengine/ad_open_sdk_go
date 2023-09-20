/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceNewCustomer
type QianchuanAdUpdateV10AudienceNewCustomer string

// List of qianchuan_ad_update_v1.0_audience_new_customer
const (
	NONE_QianchuanAdUpdateV10AudienceNewCustomer   QianchuanAdUpdateV10AudienceNewCustomer = "NONE"
	NO_BUY_QianchuanAdUpdateV10AudienceNewCustomer QianchuanAdUpdateV10AudienceNewCustomer = "NO_BUY"
)

// All allowed values of QianchuanAdUpdateV10AudienceNewCustomer enum
var AllowedQianchuanAdUpdateV10AudienceNewCustomerEnumValues = []QianchuanAdUpdateV10AudienceNewCustomer{
	"NONE",
	"NO_BUY",
}

// NewQianchuanAdUpdateV10AudienceNewCustomerFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceNewCustomer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceNewCustomerFromValue(v string) (*QianchuanAdUpdateV10AudienceNewCustomer, error) {
	ev := QianchuanAdUpdateV10AudienceNewCustomer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceNewCustomer: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceNewCustomerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceNewCustomer) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceNewCustomerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_new_customer value
func (v QianchuanAdUpdateV10AudienceNewCustomer) Ptr() *QianchuanAdUpdateV10AudienceNewCustomer {
	return &v
}
