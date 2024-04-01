/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliverySettingSearchContinueDelivery
type ProjectListV30DataListDeliverySettingSearchContinueDelivery string

// List of project_list_v3.0_data_list_delivery_setting_search_continue_delivery
const (
	OFF_ProjectListV30DataListDeliverySettingSearchContinueDelivery ProjectListV30DataListDeliverySettingSearchContinueDelivery = "OFF"
	ON_ProjectListV30DataListDeliverySettingSearchContinueDelivery  ProjectListV30DataListDeliverySettingSearchContinueDelivery = "ON"
)

// All allowed values of ProjectListV30DataListDeliverySettingSearchContinueDelivery enum
var AllowedProjectListV30DataListDeliverySettingSearchContinueDeliveryEnumValues = []ProjectListV30DataListDeliverySettingSearchContinueDelivery{
	"OFF",
	"ON",
}

// NewProjectListV30DataListDeliverySettingSearchContinueDeliveryFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingSearchContinueDelivery
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingSearchContinueDeliveryFromValue(v string) (*ProjectListV30DataListDeliverySettingSearchContinueDelivery, error) {
	ev := ProjectListV30DataListDeliverySettingSearchContinueDelivery(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingSearchContinueDelivery: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingSearchContinueDeliveryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingSearchContinueDelivery) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingSearchContinueDeliveryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_search_continue_delivery value
func (v ProjectListV30DataListDeliverySettingSearchContinueDelivery) Ptr() *ProjectListV30DataListDeliverySettingSearchContinueDelivery {
	return &v
}
