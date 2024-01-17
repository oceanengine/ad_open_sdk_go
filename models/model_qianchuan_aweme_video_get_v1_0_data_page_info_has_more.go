/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeVideoGetV10DataPageInfoHasMore
type QianchuanAwemeVideoGetV10DataPageInfoHasMore int64

// List of qianchuan_aweme_video_get_v1.0_data_page_info_has_more
const (
	Enum_0_QianchuanAwemeVideoGetV10DataPageInfoHasMore QianchuanAwemeVideoGetV10DataPageInfoHasMore = 0
	Enum_1_QianchuanAwemeVideoGetV10DataPageInfoHasMore QianchuanAwemeVideoGetV10DataPageInfoHasMore = 1
)

// All allowed values of QianchuanAwemeVideoGetV10DataPageInfoHasMore enum
var AllowedQianchuanAwemeVideoGetV10DataPageInfoHasMoreEnumValues = []QianchuanAwemeVideoGetV10DataPageInfoHasMore{
	0,
	1,
}

// NewQianchuanAwemeVideoGetV10DataPageInfoHasMoreFromValue returns a pointer to a valid QianchuanAwemeVideoGetV10DataPageInfoHasMore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeVideoGetV10DataPageInfoHasMoreFromValue(v int64) (*QianchuanAwemeVideoGetV10DataPageInfoHasMore, error) {
	ev := QianchuanAwemeVideoGetV10DataPageInfoHasMore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeVideoGetV10DataPageInfoHasMore: valid values are %v", v, AllowedQianchuanAwemeVideoGetV10DataPageInfoHasMoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeVideoGetV10DataPageInfoHasMore) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeVideoGetV10DataPageInfoHasMoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_video_get_v1.0_data_page_info_has_more value
func (v QianchuanAwemeVideoGetV10DataPageInfoHasMore) Ptr() *QianchuanAwemeVideoGetV10DataPageInfoHasMore {
	return &v
}
