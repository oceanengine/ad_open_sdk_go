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

// QianchuanDmpAudiencesGetV10DataRetargetingTagsSource
type QianchuanDmpAudiencesGetV10DataRetargetingTagsSource string

// List of qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_source
const (
	CUSTOM_AUDIENCE_TYPE_BRAND_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource          QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_BRAND"
	CUSTOM_AUDIENCE_TYPE_DATA_SOURCE_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource    QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_DATA_SOURCE"
	CUSTOM_AUDIENCE_TYPE_DOU_PLUS_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource       QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_DOU_PLUS"
	CUSTOM_AUDIENCE_TYPE_EXTEND_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource         QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_EXTEND"
	CUSTOM_AUDIENCE_TYPE_FRIEND_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource         QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_FRIEND"
	CUSTOM_AUDIENCE_TYPE_ONE_KEY_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource        QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_ONE_KEY"
	CUSTOM_AUDIENCE_TYPE_OPERATE_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource        QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_OPERATE"
	CUSTOM_AUDIENCE_TYPE_PACK_RULE_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource      QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_PACK_RULE"
	CUSTOM_AUDIENCE_TYPE_RULE_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource           QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_RULE"
	CUSTOM_AUDIENCE_TYPE_THEME_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource          QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_THEME"
	CUSTOM_AUDIENCE_TYPE_THIRD_PARTY_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource    QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_THIRD_PARTY"
	CUSTOM_AUDIENCE_TYPE_UPLOAD_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource         QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "CUSTOM_AUDIENCE_TYPE_UPLOAD"
	FINANCECUSTOM_AUDIENCE_TYPE_FINANCE_QianchuanDmpAudiencesGetV10DataRetargetingTagsSource QianchuanDmpAudiencesGetV10DataRetargetingTagsSource = "FINANCECUSTOM_AUDIENCE_TYPE_FINANCE"
)

// All allowed values of QianchuanDmpAudiencesGetV10DataRetargetingTagsSource enum
var AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsSourceEnumValues = []QianchuanDmpAudiencesGetV10DataRetargetingTagsSource{
	"CUSTOM_AUDIENCE_TYPE_BRAND",
	"CUSTOM_AUDIENCE_TYPE_DATA_SOURCE",
	"CUSTOM_AUDIENCE_TYPE_DOU_PLUS",
	"CUSTOM_AUDIENCE_TYPE_EXTEND",
	"CUSTOM_AUDIENCE_TYPE_FRIEND",
	"CUSTOM_AUDIENCE_TYPE_ONE_KEY",
	"CUSTOM_AUDIENCE_TYPE_OPERATE",
	"CUSTOM_AUDIENCE_TYPE_PACK_RULE",
	"CUSTOM_AUDIENCE_TYPE_RULE",
	"CUSTOM_AUDIENCE_TYPE_THEME",
	"CUSTOM_AUDIENCE_TYPE_THIRD_PARTY",
	"CUSTOM_AUDIENCE_TYPE_UPLOAD",
	"FINANCECUSTOM_AUDIENCE_TYPE_FINANCE",
}

// NewQianchuanDmpAudiencesGetV10DataRetargetingTagsSourceFromValue returns a pointer to a valid QianchuanDmpAudiencesGetV10DataRetargetingTagsSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanDmpAudiencesGetV10DataRetargetingTagsSourceFromValue(v string) (*QianchuanDmpAudiencesGetV10DataRetargetingTagsSource, error) {
	ev := QianchuanDmpAudiencesGetV10DataRetargetingTagsSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanDmpAudiencesGetV10DataRetargetingTagsSource: valid values are %v", v, AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsSource) IsValid() bool {
	for _, existing := range AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_source value
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsSource) Ptr() *QianchuanDmpAudiencesGetV10DataRetargetingTagsSource {
	return &v
}
