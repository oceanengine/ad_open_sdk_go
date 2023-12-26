/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeVideoGetV10DataVideoListIsRecommend
type QianchuanAwemeVideoGetV10DataVideoListIsRecommend int64

// List of qianchuan_aweme_video_get_v1.0_data_video_list_is_recommend
const (
	Enum_0_QianchuanAwemeVideoGetV10DataVideoListIsRecommend QianchuanAwemeVideoGetV10DataVideoListIsRecommend = 0
	Enum_1_QianchuanAwemeVideoGetV10DataVideoListIsRecommend QianchuanAwemeVideoGetV10DataVideoListIsRecommend = 1
)

// All allowed values of QianchuanAwemeVideoGetV10DataVideoListIsRecommend enum
var AllowedQianchuanAwemeVideoGetV10DataVideoListIsRecommendEnumValues = []QianchuanAwemeVideoGetV10DataVideoListIsRecommend{
	0,
	1,
}

// NewQianchuanAwemeVideoGetV10DataVideoListIsRecommendFromValue returns a pointer to a valid QianchuanAwemeVideoGetV10DataVideoListIsRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeVideoGetV10DataVideoListIsRecommendFromValue(v int64) (*QianchuanAwemeVideoGetV10DataVideoListIsRecommend, error) {
	ev := QianchuanAwemeVideoGetV10DataVideoListIsRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeVideoGetV10DataVideoListIsRecommend: valid values are %v", v, AllowedQianchuanAwemeVideoGetV10DataVideoListIsRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeVideoGetV10DataVideoListIsRecommend) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeVideoGetV10DataVideoListIsRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_video_get_v1.0_data_video_list_is_recommend value
func (v QianchuanAwemeVideoGetV10DataVideoListIsRecommend) Ptr() *QianchuanAwemeVideoGetV10DataVideoListIsRecommend {
	return &v
}
