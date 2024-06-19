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

// ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType
type ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType string

// List of tools_app_management_android_basic_package_update_v2_payment_type
const (
	FREE_ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType                      ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType = "FREE"
	PRODUCT_ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType                   ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType = "PRODUCT"
	TRIAL_AND_LIMIT_FUNCTIONS_ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType = "TRIAL_AND_LIMIT_FUNCTIONS"
	TRIAL_AND_PURCHASE_ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType        ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType = "TRIAL_AND_PURCHASE"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType enum
var AllowedToolsAppManagementAndroidBasicPackageUpdateV2PaymentTypeEnumValues = []ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType{
	"FREE",
	"PRODUCT",
	"TRIAL_AND_LIMIT_FUNCTIONS",
	"TRIAL_AND_PURCHASE",
}

// NewToolsAppManagementAndroidBasicPackageUpdateV2PaymentTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageUpdateV2PaymentTypeFromValue(v string) (*ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType, error) {
	ev := ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageUpdateV2PaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageUpdateV2PaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_update_v2_payment_type value
func (v ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType) Ptr() *ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType {
	return &v
}
