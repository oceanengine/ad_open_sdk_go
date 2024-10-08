/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBlueFlowPackageListV30ResponseDataBlueFlowPackagesInner struct for ToolsBlueFlowPackageListV30ResponseDataBlueFlowPackagesInner
type ToolsBlueFlowPackageListV30ResponseDataBlueFlowPackagesInner struct {
	//
	BlueFlowKeyword []*ToolsBlueFlowPackageListV30ResponseDataBlueFlowPackagesInnerBlueFlowKeywordInner `json:"blue_flow_keyword,omitempty"`
	//
	BlueFlowPackageId *int64 `json:"blue_flow_package_id,omitempty"`
	//
	BlueFlowPackageName *string `json:"blue_flow_package_name,omitempty"`
	//
	EstimatedCostRange []int64 `json:"estimated_cost_range,omitempty"`
	//
	EstimatedReachTrafficRange []int64 `json:"estimated_reach_traffic_range,omitempty"`
	//
	FlowCoverageRatio *string `json:"flow_coverage_ratio,omitempty"`
	//
	SuggestedBidRange []int64 `json:"suggested_bid_range,omitempty"`
}
