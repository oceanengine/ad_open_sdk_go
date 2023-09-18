/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanProductAvailableGetV10DataProductListTags
type QianchuanProductAvailableGetV10DataProductListTags string

// List of qianchuan_product_available_get_v1.0_data_product_list_tags
const (
	Enum_1_QianchuanProductAvailableGetV10DataProductListTags QianchuanProductAvailableGetV10DataProductListTags = "1"
	Enum_0_QianchuanProductAvailableGetV10DataProductListTags QianchuanProductAvailableGetV10DataProductListTags = "0"
)

// All allowed values of QianchuanProductAvailableGetV10DataProductListTags enum
var AllowedQianchuanProductAvailableGetV10DataProductListTagsEnumValues = []QianchuanProductAvailableGetV10DataProductListTags{
	"1",
	"0",
}

// NewQianchuanProductAvailableGetV10DataProductListTagsFromValue returns a pointer to a valid QianchuanProductAvailableGetV10DataProductListTags
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAvailableGetV10DataProductListTagsFromValue(v string) (*QianchuanProductAvailableGetV10DataProductListTags, error) {
	ev := QianchuanProductAvailableGetV10DataProductListTags(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAvailableGetV10DataProductListTags: valid values are %v", v, AllowedQianchuanProductAvailableGetV10DataProductListTagsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAvailableGetV10DataProductListTags) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAvailableGetV10DataProductListTagsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_available_get_v1.0_data_product_list_tags value
func (v QianchuanProductAvailableGetV10DataProductListTags) Ptr() *QianchuanProductAvailableGetV10DataProductListTags {
	return &v
}