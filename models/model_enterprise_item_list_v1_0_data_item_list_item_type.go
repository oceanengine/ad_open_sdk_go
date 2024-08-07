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

// EnterpriseItemListV10DataItemListItemType
type EnterpriseItemListV10DataItemListItemType string

// List of enterprise_item_list_v1.0_data_item_list_item_type
const (
	ITEM_AD_EnterpriseItemListV10DataItemListItemType      EnterpriseItemListV10DataItemListItemType = "ITEM_AD"
	ITEM_CONTENT_EnterpriseItemListV10DataItemListItemType EnterpriseItemListV10DataItemListItemType = "ITEM_CONTENT"
)

// All allowed values of EnterpriseItemListV10DataItemListItemType enum
var AllowedEnterpriseItemListV10DataItemListItemTypeEnumValues = []EnterpriseItemListV10DataItemListItemType{
	"ITEM_AD",
	"ITEM_CONTENT",
}

// NewEnterpriseItemListV10DataItemListItemTypeFromValue returns a pointer to a valid EnterpriseItemListV10DataItemListItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseItemListV10DataItemListItemTypeFromValue(v string) (*EnterpriseItemListV10DataItemListItemType, error) {
	ev := EnterpriseItemListV10DataItemListItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseItemListV10DataItemListItemType: valid values are %v", v, AllowedEnterpriseItemListV10DataItemListItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseItemListV10DataItemListItemType) IsValid() bool {
	for _, existing := range AllowedEnterpriseItemListV10DataItemListItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_item_list_v1.0_data_item_list_item_type value
func (v EnterpriseItemListV10DataItemListItemType) Ptr() *EnterpriseItemListV10DataItemListItemType {
	return &v
}
