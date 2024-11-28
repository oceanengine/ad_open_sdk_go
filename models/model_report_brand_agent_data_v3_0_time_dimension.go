/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandAgentDataV30TimeDimension
type ReportBrandAgentDataV30TimeDimension string

// List of report_brand_agent_data_v3.0_time_dimension
const (
	DAY_ReportBrandAgentDataV30TimeDimension   ReportBrandAgentDataV30TimeDimension = "DAY"
	MONTH_ReportBrandAgentDataV30TimeDimension ReportBrandAgentDataV30TimeDimension = "MONTH"
	WEEK_ReportBrandAgentDataV30TimeDimension  ReportBrandAgentDataV30TimeDimension = "WEEK"
)

// Ptr returns reference to report_brand_agent_data_v3.0_time_dimension value
func (v ReportBrandAgentDataV30TimeDimension) Ptr() *ReportBrandAgentDataV30TimeDimension {
	return &v
}
