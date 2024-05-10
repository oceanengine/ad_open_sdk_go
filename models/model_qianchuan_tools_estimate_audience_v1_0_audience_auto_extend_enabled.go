/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled
type QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled int64

// List of qianchuan_tools_estimate_audience_v1.0_audience_auto_extend_enabled
const (
	Enum_0_QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled = 0
	Enum_1_QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled = 1
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabledEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled{
	0,
	1,
}

// NewQianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabledFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabledFromValue(v int64) (*QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabledEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabledEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_auto_extend_enabled value
func (v QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled {
	return &v
}
