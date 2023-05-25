/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListPricing
type ProjectListV30DataListPricing string

// List of project_list_v3.0_data_list_pricing
const (
	PRICING_CPC_ProjectListV30DataListPricing  ProjectListV30DataListPricing = "PRICING_CPC"
	PRICING_CPM_ProjectListV30DataListPricing  ProjectListV30DataListPricing = "PRICING_CPM"
	PRICING_OCPC_ProjectListV30DataListPricing ProjectListV30DataListPricing = "PRICING_OCPC"
	PRICING_OCPM_ProjectListV30DataListPricing ProjectListV30DataListPricing = "PRICING_OCPM"
)

// All allowed values of ProjectListV30DataListPricing enum
var AllowedProjectListV30DataListPricingEnumValues = []ProjectListV30DataListPricing{
	"PRICING_CPC",
	"PRICING_CPM",
	"PRICING_OCPC",
	"PRICING_OCPM",
}

// NewProjectListV30DataListPricingFromValue returns a pointer to a valid ProjectListV30DataListPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListPricingFromValue(v string) (*ProjectListV30DataListPricing, error) {
	ev := ProjectListV30DataListPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListPricing: valid values are %v", v, AllowedProjectListV30DataListPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListPricing) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_pricing value
func (v ProjectListV30DataListPricing) Ptr() *ProjectListV30DataListPricing {
	return &v
}
