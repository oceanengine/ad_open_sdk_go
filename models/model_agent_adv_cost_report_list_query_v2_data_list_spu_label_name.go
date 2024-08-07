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

// AgentAdvCostReportListQueryV2DataListSpuLabelName
type AgentAdvCostReportListQueryV2DataListSpuLabelName string

// List of agent_adv_cost_report_list_query_v2_data_list_spu_label_name
const (
	BRAND_CONVERSION_AgentAdvCostReportListQueryV2DataListSpuLabelName       AgentAdvCostReportListQueryV2DataListSpuLabelName = "BRAND_CONVERSION"
	BRAND_EXPOSURE_AgentAdvCostReportListQueryV2DataListSpuLabelName         AgentAdvCostReportListQueryV2DataListSpuLabelName = "BRAND_EXPOSURE"
	BRAND_GRASS_AgentAdvCostReportListQueryV2DataListSpuLabelName            AgentAdvCostReportListQueryV2DataListSpuLabelName = "BRAND_GRASS"
	CONCENTRATING_RESOURCE_AgentAdvCostReportListQueryV2DataListSpuLabelName AgentAdvCostReportListQueryV2DataListSpuLabelName = "CONCENTRATING_RESOURCE"
	STAR_RESOURCE_AgentAdvCostReportListQueryV2DataListSpuLabelName          AgentAdvCostReportListQueryV2DataListSpuLabelName = "STAR_RESOURCE"
)

// All allowed values of AgentAdvCostReportListQueryV2DataListSpuLabelName enum
var AllowedAgentAdvCostReportListQueryV2DataListSpuLabelNameEnumValues = []AgentAdvCostReportListQueryV2DataListSpuLabelName{
	"BRAND_CONVERSION",
	"BRAND_EXPOSURE",
	"BRAND_GRASS",
	"CONCENTRATING_RESOURCE",
	"STAR_RESOURCE",
}

// NewAgentAdvCostReportListQueryV2DataListSpuLabelNameFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListSpuLabelName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListSpuLabelNameFromValue(v string) (*AgentAdvCostReportListQueryV2DataListSpuLabelName, error) {
	ev := AgentAdvCostReportListQueryV2DataListSpuLabelName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListSpuLabelName: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListSpuLabelNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListSpuLabelName) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListSpuLabelNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_spu_label_name value
func (v AgentAdvCostReportListQueryV2DataListSpuLabelName) Ptr() *AgentAdvCostReportListQueryV2DataListSpuLabelName {
	return &v
}
