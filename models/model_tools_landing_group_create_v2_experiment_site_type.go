/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupCreateV2ExperimentSiteType
type ToolsLandingGroupCreateV2ExperimentSiteType string

// List of tools_landing_group_create_v2_experiment_site_type
const (
	OWN_ToolsLandingGroupCreateV2ExperimentSiteType         ToolsLandingGroupCreateV2ExperimentSiteType = "OWN"
	THIRD_PARTY_ToolsLandingGroupCreateV2ExperimentSiteType ToolsLandingGroupCreateV2ExperimentSiteType = "THIRD_PARTY"
)

// Ptr returns reference to tools_landing_group_create_v2_experiment_site_type value
func (v ToolsLandingGroupCreateV2ExperimentSiteType) Ptr() *ToolsLandingGroupCreateV2ExperimentSiteType {
	return &v
}
