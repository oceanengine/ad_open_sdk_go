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

// EnterpriseOperationLogGetV10DataListBusinessPageOperationType
type EnterpriseOperationLogGetV10DataListBusinessPageOperationType string

// List of enterprise_operation_log_get_v1.0_data_list_business_page_operation_type
const (
	VIDEO_EnterpriseOperationLogGetV10DataListBusinessPageOperationType EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "VIDEO"
	LIVE_EnterpriseOperationLogGetV10DataListBusinessPageOperationType  EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "LIVE"
	AD_EnterpriseOperationLogGetV10DataListBusinessPageOperationType    EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "AD"
	DOU_EnterpriseOperationLogGetV10DataListBusinessPageOperationType   EnterpriseOperationLogGetV10DataListBusinessPageOperationType = "DOU"
)

// All allowed values of EnterpriseOperationLogGetV10DataListBusinessPageOperationType enum
var AllowedEnterpriseOperationLogGetV10DataListBusinessPageOperationTypeEnumValues = []EnterpriseOperationLogGetV10DataListBusinessPageOperationType{
	"VIDEO",
	"LIVE",
	"AD",
	"DOU",
}

// NewEnterpriseOperationLogGetV10DataListBusinessPageOperationTypeFromValue returns a pointer to a valid EnterpriseOperationLogGetV10DataListBusinessPageOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseOperationLogGetV10DataListBusinessPageOperationTypeFromValue(v string) (*EnterpriseOperationLogGetV10DataListBusinessPageOperationType, error) {
	ev := EnterpriseOperationLogGetV10DataListBusinessPageOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseOperationLogGetV10DataListBusinessPageOperationType: valid values are %v", v, AllowedEnterpriseOperationLogGetV10DataListBusinessPageOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseOperationLogGetV10DataListBusinessPageOperationType) IsValid() bool {
	for _, existing := range AllowedEnterpriseOperationLogGetV10DataListBusinessPageOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_operation_log_get_v1.0_data_list_business_page_operation_type value
func (v EnterpriseOperationLogGetV10DataListBusinessPageOperationType) Ptr() *EnterpriseOperationLogGetV10DataListBusinessPageOperationType {
	return &v
}
