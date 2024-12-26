/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2Career
type ToolsBidSuggestV2Career string

// List of tools_bid_suggest_v2_career
const (
	TEACHER_ToolsBidSuggestV2Career         ToolsBidSuggestV2Career = "TEACHER"
	CIVIL_SERVANTS_ToolsBidSuggestV2Career  ToolsBidSuggestV2Career = "CIVIL_SERVANTS"
	FINANCIAL_ToolsBidSuggestV2Career       ToolsBidSuggestV2Career = "FINANCIAL"
	MEDICAL_STAFF_ToolsBidSuggestV2Career   ToolsBidSuggestV2Career = "MEDICAL_STAFF"
	COLLEGE_STUDENT_ToolsBidSuggestV2Career ToolsBidSuggestV2Career = "COLLEGE_STUDENT"
	IT_ToolsBidSuggestV2Career              ToolsBidSuggestV2Career = "IT"
)

// Ptr returns reference to tools_bid_suggest_v2_career value
func (v ToolsBidSuggestV2Career) Ptr() *ToolsBidSuggestV2Career {
	return &v
}
