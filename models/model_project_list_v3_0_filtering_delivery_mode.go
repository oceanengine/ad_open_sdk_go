/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30FilteringDeliveryMode
type ProjectListV30FilteringDeliveryMode string

// List of project_list_v3.0_filtering_delivery_mode
const (
	MANUAL_ProjectListV30FilteringDeliveryMode     ProjectListV30FilteringDeliveryMode = "MANUAL"
	PROCEDURAL_ProjectListV30FilteringDeliveryMode ProjectListV30FilteringDeliveryMode = "PROCEDURAL"
)

// Ptr returns reference to project_list_v3.0_filtering_delivery_mode value
func (v ProjectListV30FilteringDeliveryMode) Ptr() *ProjectListV30FilteringDeliveryMode {
	return &v
}
