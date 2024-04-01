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

// QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend
type QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend int64

// List of qianchuan_file_video_aweme_get_v1.0_data_video_list_is_recommend
const (
	Enum_1_QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend = 1
	Enum_0_QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend = 0
)

// All allowed values of QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend enum
var AllowedQianchuanFileVideoAwemeGetV10DataVideoListIsRecommendEnumValues = []QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend{
	1,
	0,
}

// NewQianchuanFileVideoAwemeGetV10DataVideoListIsRecommendFromValue returns a pointer to a valid QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanFileVideoAwemeGetV10DataVideoListIsRecommendFromValue(v int64) (*QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend, error) {
	ev := QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend: valid values are %v", v, AllowedQianchuanFileVideoAwemeGetV10DataVideoListIsRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend) IsValid() bool {
	for _, existing := range AllowedQianchuanFileVideoAwemeGetV10DataVideoListIsRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_file_video_aweme_get_v1.0_data_video_list_is_recommend value
func (v QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend) Ptr() *QianchuanFileVideoAwemeGetV10DataVideoListIsRecommend {
	return &v
}
