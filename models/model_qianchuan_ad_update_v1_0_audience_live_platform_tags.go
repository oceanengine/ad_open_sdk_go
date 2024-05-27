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

// QianchuanAdUpdateV10AudienceLivePlatformTags
type QianchuanAdUpdateV10AudienceLivePlatformTags string

// List of qianchuan_ad_update_v1.0_audience_live_platform_tags
const (
	ABNORMAL_ACTIVE_QianchuanAdUpdateV10AudienceLivePlatformTags QianchuanAdUpdateV10AudienceLivePlatformTags = "ABNORMAL_ACTIVE"
	AWEME_FANS_QianchuanAdUpdateV10AudienceLivePlatformTags      QianchuanAdUpdateV10AudienceLivePlatformTags = "AWEME_FANS"
	LARGE_FANSCOUNT_QianchuanAdUpdateV10AudienceLivePlatformTags QianchuanAdUpdateV10AudienceLivePlatformTags = "LARGE_FANSCOUNT"
)

// All allowed values of QianchuanAdUpdateV10AudienceLivePlatformTags enum
var AllowedQianchuanAdUpdateV10AudienceLivePlatformTagsEnumValues = []QianchuanAdUpdateV10AudienceLivePlatformTags{
	"ABNORMAL_ACTIVE",
	"AWEME_FANS",
	"LARGE_FANSCOUNT",
}

// NewQianchuanAdUpdateV10AudienceLivePlatformTagsFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceLivePlatformTags
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceLivePlatformTagsFromValue(v string) (*QianchuanAdUpdateV10AudienceLivePlatformTags, error) {
	ev := QianchuanAdUpdateV10AudienceLivePlatformTags(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceLivePlatformTags: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceLivePlatformTagsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceLivePlatformTags) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceLivePlatformTagsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_live_platform_tags value
func (v QianchuanAdUpdateV10AudienceLivePlatformTags) Ptr() *QianchuanAdUpdateV10AudienceLivePlatformTags {
	return &v
}
