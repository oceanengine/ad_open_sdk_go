/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimatedPriceGetV2ResponseData
type ToolsEstimatedPriceGetV2ResponseData struct {
	//
	CpcPrice *float64 `json:"cpc_price,omitempty"`
	//
	LowerBound *float64 `json:"lower_bound,omitempty"`
	//
	UpperBound *float64 `json:"upper_bound,omitempty"`
}
