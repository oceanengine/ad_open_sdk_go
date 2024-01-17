/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliverySettingBidSpeed
type ProjectListV30DataListDeliverySettingBidSpeed string

// List of project_list_v3.0_data_list_delivery_setting_bid_speed
const (
	BALANCE_ProjectListV30DataListDeliverySettingBidSpeed ProjectListV30DataListDeliverySettingBidSpeed = "BALANCE"
	FAST_ProjectListV30DataListDeliverySettingBidSpeed    ProjectListV30DataListDeliverySettingBidSpeed = "FAST"
)

// All allowed values of ProjectListV30DataListDeliverySettingBidSpeed enum
var AllowedProjectListV30DataListDeliverySettingBidSpeedEnumValues = []ProjectListV30DataListDeliverySettingBidSpeed{
	"BALANCE",
	"FAST",
}

// NewProjectListV30DataListDeliverySettingBidSpeedFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingBidSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingBidSpeedFromValue(v string) (*ProjectListV30DataListDeliverySettingBidSpeed, error) {
	ev := ProjectListV30DataListDeliverySettingBidSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingBidSpeed: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingBidSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingBidSpeed) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingBidSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_bid_speed value
func (v ProjectListV30DataListDeliverySettingBidSpeed) Ptr() *ProjectListV30DataListDeliverySettingBidSpeed {
	return &v
}
