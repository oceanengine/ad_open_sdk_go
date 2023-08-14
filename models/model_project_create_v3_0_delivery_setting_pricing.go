/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingPricing
type ProjectCreateV30DeliverySettingPricing string

// List of project_create_v3.0_delivery_setting_pricing
const (
	PRICING_CPC_ProjectCreateV30DeliverySettingPricing  ProjectCreateV30DeliverySettingPricing = "PRICING_CPC"
	PRICING_CPM_ProjectCreateV30DeliverySettingPricing  ProjectCreateV30DeliverySettingPricing = "PRICING_CPM"
	PRICING_OCPC_ProjectCreateV30DeliverySettingPricing ProjectCreateV30DeliverySettingPricing = "PRICING_OCPC"
	PRICING_OCPM_ProjectCreateV30DeliverySettingPricing ProjectCreateV30DeliverySettingPricing = "PRICING_OCPM"
)

// All allowed values of ProjectCreateV30DeliverySettingPricing enum
var AllowedProjectCreateV30DeliverySettingPricingEnumValues = []ProjectCreateV30DeliverySettingPricing{
	"PRICING_CPC",
	"PRICING_CPM",
	"PRICING_OCPC",
	"PRICING_OCPM",
}

// NewProjectCreateV30DeliverySettingPricingFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingPricingFromValue(v string) (*ProjectCreateV30DeliverySettingPricing, error) {
	ev := ProjectCreateV30DeliverySettingPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingPricing: valid values are %v", v, AllowedProjectCreateV30DeliverySettingPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingPricing) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_pricing value
func (v ProjectCreateV30DeliverySettingPricing) Ptr() *ProjectCreateV30DeliverySettingPricing {
	return &v
}
