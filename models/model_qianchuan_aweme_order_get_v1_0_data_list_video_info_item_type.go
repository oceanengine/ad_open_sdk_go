/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderGetV10DataListVideoInfoItemType
type QianchuanAwemeOrderGetV10DataListVideoInfoItemType int64

// List of qianchuan_aweme_order_get_v1.0_data_list_video_info_item_type
const (
	Enum_0_QianchuanAwemeOrderGetV10DataListVideoInfoItemType QianchuanAwemeOrderGetV10DataListVideoInfoItemType = 0
	Enum_1_QianchuanAwemeOrderGetV10DataListVideoInfoItemType QianchuanAwemeOrderGetV10DataListVideoInfoItemType = 1
)

// All allowed values of QianchuanAwemeOrderGetV10DataListVideoInfoItemType enum
var AllowedQianchuanAwemeOrderGetV10DataListVideoInfoItemTypeEnumValues = []QianchuanAwemeOrderGetV10DataListVideoInfoItemType{
	0,
	1,
}

// NewQianchuanAwemeOrderGetV10DataListVideoInfoItemTypeFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10DataListVideoInfoItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10DataListVideoInfoItemTypeFromValue(v int64) (*QianchuanAwemeOrderGetV10DataListVideoInfoItemType, error) {
	ev := QianchuanAwemeOrderGetV10DataListVideoInfoItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10DataListVideoInfoItemType: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10DataListVideoInfoItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10DataListVideoInfoItemType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10DataListVideoInfoItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_data_list_video_info_item_type value
func (v QianchuanAwemeOrderGetV10DataListVideoInfoItemType) Ptr() *QianchuanAwemeOrderGetV10DataListVideoInfoItemType {
	return &v
}
