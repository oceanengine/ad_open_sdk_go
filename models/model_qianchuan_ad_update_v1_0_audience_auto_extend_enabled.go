/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceAutoExtendEnabled
type QianchuanAdUpdateV10AudienceAutoExtendEnabled int64

// List of qianchuan_ad_update_v1.0_audience_auto_extend_enabled
const (
	Enum_0_QianchuanAdUpdateV10AudienceAutoExtendEnabled QianchuanAdUpdateV10AudienceAutoExtendEnabled = 0
	Enum_1_QianchuanAdUpdateV10AudienceAutoExtendEnabled QianchuanAdUpdateV10AudienceAutoExtendEnabled = 1
)

// All allowed values of QianchuanAdUpdateV10AudienceAutoExtendEnabled enum
var AllowedQianchuanAdUpdateV10AudienceAutoExtendEnabledEnumValues = []QianchuanAdUpdateV10AudienceAutoExtendEnabled{
	0,
	1,
}

// NewQianchuanAdUpdateV10AudienceAutoExtendEnabledFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAutoExtendEnabled
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAutoExtendEnabledFromValue(v int64) (*QianchuanAdUpdateV10AudienceAutoExtendEnabled, error) {
	ev := QianchuanAdUpdateV10AudienceAutoExtendEnabled(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAutoExtendEnabled: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAutoExtendEnabledEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAutoExtendEnabled) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAutoExtendEnabledEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_auto_extend_enabled value
func (v QianchuanAdUpdateV10AudienceAutoExtendEnabled) Ptr() *QianchuanAdUpdateV10AudienceAutoExtendEnabled {
	return &v
}
