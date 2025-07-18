/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2FlowControlMode
type ToolsBidSuggestV2FlowControlMode string

// List of tools_bid_suggest_v2_flow_control_mode
const (
	FLOW_CONTROL_MODE_SMOOTH_ToolsBidSuggestV2FlowControlMode     ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_BALANCE_ToolsBidSuggestV2FlowControlMode    ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
	FLOW_CONTROL_MODE_FAST_ToolsBidSuggestV2FlowControlMode       ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_FAST"
	FLOW_CONTROL_MODE_TWO_PHASES_ToolsBidSuggestV2FlowControlMode ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
	FLOW_CONTROL_MODE_HOURLY_ToolsBidSuggestV2FlowControlMode     ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
)

// Ptr returns reference to tools_bid_suggest_v2_flow_control_mode value
func (v ToolsBidSuggestV2FlowControlMode) Ptr() *ToolsBidSuggestV2FlowControlMode {
	return &v
}
