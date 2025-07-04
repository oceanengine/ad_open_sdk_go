/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsNoBidSuggestBidV2FilteringFlowControlMode
type ToolsNoBidSuggestBidV2FilteringFlowControlMode string

// List of tools_no_bid_suggest_bid_v2_filtering_flow_control_mode
const (
	FLOW_CONTROL_MODE_FAST_ToolsNoBidSuggestBidV2FilteringFlowControlMode       ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_FAST"
	FLOW_CONTROL_MODE_HOURLY_ToolsNoBidSuggestBidV2FilteringFlowControlMode     ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
	FLOW_CONTROL_MODE_SMOOTH_ToolsNoBidSuggestBidV2FilteringFlowControlMode     ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_TWO_PHASES_ToolsNoBidSuggestBidV2FilteringFlowControlMode ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
	FLOW_CONTROL_MODE_BALANCE_ToolsNoBidSuggestBidV2FilteringFlowControlMode    ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
)

// Ptr returns reference to tools_no_bid_suggest_bid_v2_filtering_flow_control_mode value
func (v ToolsNoBidSuggestBidV2FilteringFlowControlMode) Ptr() *ToolsNoBidSuggestBidV2FilteringFlowControlMode {
	return &v
}
