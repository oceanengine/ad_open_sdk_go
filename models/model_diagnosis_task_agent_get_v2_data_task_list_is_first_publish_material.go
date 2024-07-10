/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial
type DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial string

// List of diagnosis_task_agent_get_v2_data_task_list_is_first_publish_material
const (
	NO_DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial      DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial = "NO"
	UNKNOWN_DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial = "UNKNOWN"
	YES_DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial     DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial = "YES"
)

// All allowed values of DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial enum
var AllowedDiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterialEnumValues = []DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial{
	"NO",
	"UNKNOWN",
	"YES",
}

// NewDiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterialFromValue returns a pointer to a valid DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterialFromValue(v string) (*DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial, error) {
	ev := DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial: valid values are %v", v, AllowedDiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_get_v2_data_task_list_is_first_publish_material value
func (v DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial) Ptr() *DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial {
	return &v
}
