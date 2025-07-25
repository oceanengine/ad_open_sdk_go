/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30DeliverySettingBidType
type ProjectCreateV30DeliverySettingBidType string

// List of project_create_v3.0_delivery_setting_bid_type
const (
	CUSTOM_ProjectCreateV30DeliverySettingBidType        ProjectCreateV30DeliverySettingBidType = "CUSTOM"
	NO_BID_ProjectCreateV30DeliverySettingBidType        ProjectCreateV30DeliverySettingBidType = "NO_BID"
	OPTIMAL_COST_ProjectCreateV30DeliverySettingBidType  ProjectCreateV30DeliverySettingBidType = "OPTIMAL_COST"
	UPPER_CONTROL_ProjectCreateV30DeliverySettingBidType ProjectCreateV30DeliverySettingBidType = "UPPER_CONTROL"
)

// Ptr returns reference to project_create_v3.0_delivery_setting_bid_type value
func (v ProjectCreateV30DeliverySettingBidType) Ptr() *ProjectCreateV30DeliverySettingBidType {
	return &v
}
