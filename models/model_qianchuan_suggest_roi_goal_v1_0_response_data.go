/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanSuggestRoiGoalV10ResponseData
type QianchuanSuggestRoiGoalV10ResponseData struct {
	//
	EcpRoiGoal float64 `json:"ecp_roi_goal"`
	//
	RoiLowerBound *float64 `json:"roi_lower_bound,omitempty"`
	//
	RoiUpperBound *float64 `json:"roi_upper_bound,omitempty"`
}
