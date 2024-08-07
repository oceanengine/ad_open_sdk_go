/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceLivePlatformTags
type QianchuanAdCreateV10AudienceLivePlatformTags string

// List of qianchuan_ad_create_v1.0_audience_live_platform_tags
const (
	ABNORMAL_ACTIVE_QianchuanAdCreateV10AudienceLivePlatformTags QianchuanAdCreateV10AudienceLivePlatformTags = "ABNORMAL_ACTIVE"
	AWEME_FANS_QianchuanAdCreateV10AudienceLivePlatformTags      QianchuanAdCreateV10AudienceLivePlatformTags = "AWEME_FANS"
	LARGE_FANSCOUNT_QianchuanAdCreateV10AudienceLivePlatformTags QianchuanAdCreateV10AudienceLivePlatformTags = "LARGE_FANSCOUNT"
)

// All allowed values of QianchuanAdCreateV10AudienceLivePlatformTags enum
var AllowedQianchuanAdCreateV10AudienceLivePlatformTagsEnumValues = []QianchuanAdCreateV10AudienceLivePlatformTags{
	"ABNORMAL_ACTIVE",
	"AWEME_FANS",
	"LARGE_FANSCOUNT",
}

// NewQianchuanAdCreateV10AudienceLivePlatformTagsFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceLivePlatformTags
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceLivePlatformTagsFromValue(v string) (*QianchuanAdCreateV10AudienceLivePlatformTags, error) {
	ev := QianchuanAdCreateV10AudienceLivePlatformTags(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceLivePlatformTags: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceLivePlatformTagsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceLivePlatformTags) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceLivePlatformTagsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_live_platform_tags value
func (v QianchuanAdCreateV10AudienceLivePlatformTags) Ptr() *QianchuanAdCreateV10AudienceLivePlatformTags {
	return &v
}
