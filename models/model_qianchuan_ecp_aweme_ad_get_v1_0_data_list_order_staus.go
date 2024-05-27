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

// QianchuanEcpAwemeAdGetV10DataListOrderStaus
type QianchuanEcpAwemeAdGetV10DataListOrderStaus string

// List of qianchuan_ecp_aweme_ad_get_v1.0_data_list_order_staus
const (
	AUDITING_QianchuanEcpAwemeAdGetV10DataListOrderStaus     QianchuanEcpAwemeAdGetV10DataListOrderStaus = "AUDITING"
	BOOK_QianchuanEcpAwemeAdGetV10DataListOrderStaus         QianchuanEcpAwemeAdGetV10DataListOrderStaus = "BOOK"
	CREATING_QianchuanEcpAwemeAdGetV10DataListOrderStaus     QianchuanEcpAwemeAdGetV10DataListOrderStaus = "CREATING"
	DELIVERING_QianchuanEcpAwemeAdGetV10DataListOrderStaus   QianchuanEcpAwemeAdGetV10DataListOrderStaus = "DELIVERING"
	FAILED_QianchuanEcpAwemeAdGetV10DataListOrderStaus       QianchuanEcpAwemeAdGetV10DataListOrderStaus = "FAILED"
	FINISHED_QianchuanEcpAwemeAdGetV10DataListOrderStaus     QianchuanEcpAwemeAdGetV10DataListOrderStaus = "FINISHED"
	OVER_QianchuanEcpAwemeAdGetV10DataListOrderStaus         QianchuanEcpAwemeAdGetV10DataListOrderStaus = "OVER"
	REJECT_QianchuanEcpAwemeAdGetV10DataListOrderStaus       QianchuanEcpAwemeAdGetV10DataListOrderStaus = "REJECT"
	UNDELIVERING_QianchuanEcpAwemeAdGetV10DataListOrderStaus QianchuanEcpAwemeAdGetV10DataListOrderStaus = "UNDELIVERING"
	UNPAID_QianchuanEcpAwemeAdGetV10DataListOrderStaus       QianchuanEcpAwemeAdGetV10DataListOrderStaus = "UNPAID"
)

// All allowed values of QianchuanEcpAwemeAdGetV10DataListOrderStaus enum
var AllowedQianchuanEcpAwemeAdGetV10DataListOrderStausEnumValues = []QianchuanEcpAwemeAdGetV10DataListOrderStaus{
	"AUDITING",
	"BOOK",
	"CREATING",
	"DELIVERING",
	"FAILED",
	"FINISHED",
	"OVER",
	"REJECT",
	"UNDELIVERING",
	"UNPAID",
}

// NewQianchuanEcpAwemeAdGetV10DataListOrderStausFromValue returns a pointer to a valid QianchuanEcpAwemeAdGetV10DataListOrderStaus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEcpAwemeAdGetV10DataListOrderStausFromValue(v string) (*QianchuanEcpAwemeAdGetV10DataListOrderStaus, error) {
	ev := QianchuanEcpAwemeAdGetV10DataListOrderStaus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEcpAwemeAdGetV10DataListOrderStaus: valid values are %v", v, AllowedQianchuanEcpAwemeAdGetV10DataListOrderStausEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEcpAwemeAdGetV10DataListOrderStaus) IsValid() bool {
	for _, existing := range AllowedQianchuanEcpAwemeAdGetV10DataListOrderStausEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ecp_aweme_ad_get_v1.0_data_list_order_staus value
func (v QianchuanEcpAwemeAdGetV10DataListOrderStaus) Ptr() *QianchuanEcpAwemeAdGetV10DataListOrderStaus {
	return &v
}
