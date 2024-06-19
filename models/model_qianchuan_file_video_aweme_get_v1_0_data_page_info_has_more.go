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

// QianchuanFileVideoAwemeGetV10DataPageInfoHasMore
type QianchuanFileVideoAwemeGetV10DataPageInfoHasMore int64

// List of qianchuan_file_video_aweme_get_v1.0_data_page_info_has_more
const (
	Enum_1_QianchuanFileVideoAwemeGetV10DataPageInfoHasMore QianchuanFileVideoAwemeGetV10DataPageInfoHasMore = 1
	Enum_0_QianchuanFileVideoAwemeGetV10DataPageInfoHasMore QianchuanFileVideoAwemeGetV10DataPageInfoHasMore = 0
)

// All allowed values of QianchuanFileVideoAwemeGetV10DataPageInfoHasMore enum
var AllowedQianchuanFileVideoAwemeGetV10DataPageInfoHasMoreEnumValues = []QianchuanFileVideoAwemeGetV10DataPageInfoHasMore{
	1,
	0,
}

// NewQianchuanFileVideoAwemeGetV10DataPageInfoHasMoreFromValue returns a pointer to a valid QianchuanFileVideoAwemeGetV10DataPageInfoHasMore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanFileVideoAwemeGetV10DataPageInfoHasMoreFromValue(v int64) (*QianchuanFileVideoAwemeGetV10DataPageInfoHasMore, error) {
	ev := QianchuanFileVideoAwemeGetV10DataPageInfoHasMore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanFileVideoAwemeGetV10DataPageInfoHasMore: valid values are %v", v, AllowedQianchuanFileVideoAwemeGetV10DataPageInfoHasMoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanFileVideoAwemeGetV10DataPageInfoHasMore) IsValid() bool {
	for _, existing := range AllowedQianchuanFileVideoAwemeGetV10DataPageInfoHasMoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_file_video_aweme_get_v1.0_data_page_info_has_more value
func (v QianchuanFileVideoAwemeGetV10DataPageInfoHasMore) Ptr() *QianchuanFileVideoAwemeGetV10DataPageInfoHasMore {
	return &v
}
