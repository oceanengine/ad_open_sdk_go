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

// DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial
type DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial string

// List of diagnosis_task_agent_list_v2_data_task_list_is_inefficient_material
const (
	NO_DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial      DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial = "NO"
	UNKNOWN_DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial = "UNKNOWN"
	YES_DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial     DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial = "YES"
)

// All allowed values of DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial enum
var AllowedDiagnosisTaskAgentListV2DataTaskListIsInefficientMaterialEnumValues = []DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial{
	"NO",
	"UNKNOWN",
	"YES",
}

// NewDiagnosisTaskAgentListV2DataTaskListIsInefficientMaterialFromValue returns a pointer to a valid DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentListV2DataTaskListIsInefficientMaterialFromValue(v string) (*DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial, error) {
	ev := DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial: valid values are %v", v, AllowedDiagnosisTaskAgentListV2DataTaskListIsInefficientMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentListV2DataTaskListIsInefficientMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_list_v2_data_task_list_is_inefficient_material value
func (v DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial) Ptr() *DiagnosisTaskAgentListV2DataTaskListIsInefficientMaterial {
	return &v
}
