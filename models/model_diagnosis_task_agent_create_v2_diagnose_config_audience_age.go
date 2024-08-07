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

// DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge
type DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge string

// List of diagnosis_task_agent_create_v2_diagnose_config_audience_age
const (
	ALL_DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge        DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge = "ALL"
	Enum_18_23_DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge = "18-23"
	Enum_24_30_DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge = "24-30"
	Enum_31_40_DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge = "31-40"
	Enum_41_49_DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge = "41-49"
	Enum_50_DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge    DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge = "50"
)

// All allowed values of DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge enum
var AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAgeEnumValues = []DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge{
	"ALL",
	"18-23",
	"24-30",
	"31-40",
	"41-49",
	"50",
}

// NewDiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAgeFromValue returns a pointer to a valid DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAgeFromValue(v string) (*DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge, error) {
	ev := DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge: valid values are %v", v, AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_create_v2_diagnose_config_audience_age value
func (v DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge) Ptr() *DiagnosisTaskAgentCreateV2DiagnoseConfigAudienceAge {
	return &v
}
