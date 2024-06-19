/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType
type QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType int64

// List of qianchuan_aweme_order_detail_get_v1.0_data_video_info_item_type
const (
	Enum_0_QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType = 0
	Enum_1_QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType = 1
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType enum
var AllowedQianchuanAwemeOrderDetailGetV10DataVideoInfoItemTypeEnumValues = []QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType{
	0,
	1,
}

// NewQianchuanAwemeOrderDetailGetV10DataVideoInfoItemTypeFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataVideoInfoItemTypeFromValue(v int64) (*QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataVideoInfoItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataVideoInfoItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_video_info_item_type value
func (v QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType) Ptr() *QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType {
	return &v
}
