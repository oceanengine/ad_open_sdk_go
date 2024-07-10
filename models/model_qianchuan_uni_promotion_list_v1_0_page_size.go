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

// QianchuanUniPromotionListV10PageSize
type QianchuanUniPromotionListV10PageSize int64

// List of qianchuan_uni_promotion_list_v1.0_page_size
const (
	Enum_10_QianchuanUniPromotionListV10PageSize  QianchuanUniPromotionListV10PageSize = 10
	Enum_20_QianchuanUniPromotionListV10PageSize  QianchuanUniPromotionListV10PageSize = 20
	Enum_50_QianchuanUniPromotionListV10PageSize  QianchuanUniPromotionListV10PageSize = 50
	Enum_100_QianchuanUniPromotionListV10PageSize QianchuanUniPromotionListV10PageSize = 100
	Enum_200_QianchuanUniPromotionListV10PageSize QianchuanUniPromotionListV10PageSize = 200
)

// All allowed values of QianchuanUniPromotionListV10PageSize enum
var AllowedQianchuanUniPromotionListV10PageSizeEnumValues = []QianchuanUniPromotionListV10PageSize{
	10,
	20,
	50,
	100,
	200,
}

// NewQianchuanUniPromotionListV10PageSizeFromValue returns a pointer to a valid QianchuanUniPromotionListV10PageSize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10PageSizeFromValue(v int64) (*QianchuanUniPromotionListV10PageSize, error) {
	ev := QianchuanUniPromotionListV10PageSize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10PageSize: valid values are %v", v, AllowedQianchuanUniPromotionListV10PageSizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10PageSize) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10PageSizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_page_size value
func (v QianchuanUniPromotionListV10PageSize) Ptr() *QianchuanUniPromotionListV10PageSize {
	return &v
}
