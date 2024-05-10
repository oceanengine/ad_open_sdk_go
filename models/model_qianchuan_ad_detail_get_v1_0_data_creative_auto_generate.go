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

// QianchuanAdDetailGetV10DataCreativeAutoGenerate
type QianchuanAdDetailGetV10DataCreativeAutoGenerate int64

// List of qianchuan_ad_detail_get_v1.0_data_creative_auto_generate
const (
	Enum_0_QianchuanAdDetailGetV10DataCreativeAutoGenerate QianchuanAdDetailGetV10DataCreativeAutoGenerate = 0
	Enum_1_QianchuanAdDetailGetV10DataCreativeAutoGenerate QianchuanAdDetailGetV10DataCreativeAutoGenerate = 1
)

// All allowed values of QianchuanAdDetailGetV10DataCreativeAutoGenerate enum
var AllowedQianchuanAdDetailGetV10DataCreativeAutoGenerateEnumValues = []QianchuanAdDetailGetV10DataCreativeAutoGenerate{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataCreativeAutoGenerateFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataCreativeAutoGenerate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataCreativeAutoGenerateFromValue(v int64) (*QianchuanAdDetailGetV10DataCreativeAutoGenerate, error) {
	ev := QianchuanAdDetailGetV10DataCreativeAutoGenerate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataCreativeAutoGenerate: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataCreativeAutoGenerateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataCreativeAutoGenerate) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataCreativeAutoGenerateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_creative_auto_generate value
func (v QianchuanAdDetailGetV10DataCreativeAutoGenerate) Ptr() *QianchuanAdDetailGetV10DataCreativeAutoGenerate {
	return &v
}
