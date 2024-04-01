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

// QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags
type QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags string

// List of qianchuan_tools_estimate_audience_v1.0_audience_live_platform_tags
const (
	ABNORMAL_ACTIVE_QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags = "ABNORMAL_ACTIVE"
	AWEME_FANS_QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags      QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags = "AWEME_FANS"
	LARGE_FANSCOUNT_QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags = "LARGE_FANSCOUNT"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceLivePlatformTagsEnumValues = []QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags{
	"ABNORMAL_ACTIVE",
	"AWEME_FANS",
	"LARGE_FANSCOUNT",
}

// NewQianchuanToolsEstimateAudienceV10AudienceLivePlatformTagsFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceLivePlatformTagsFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceLivePlatformTagsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceLivePlatformTagsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_live_platform_tags value
func (v QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags) Ptr() *QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags {
	return &v
}
