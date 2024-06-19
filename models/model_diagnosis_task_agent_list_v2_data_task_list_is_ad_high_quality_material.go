/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial
type DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial string

// List of diagnosis_task_agent_list_v2_data_task_list_is_ad_high_quality_material
const (
	NO_DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial      DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial = "NO"
	UNKNOWN_DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial = "UNKNOWN"
	YES_DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial     DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial = "YES"
)

// All allowed values of DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial enum
var AllowedDiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterialEnumValues = []DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial{
	"NO",
	"UNKNOWN",
	"YES",
}

// NewDiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterialFromValue returns a pointer to a valid DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterialFromValue(v string) (*DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial, error) {
	ev := DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial: valid values are %v", v, AllowedDiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_list_v2_data_task_list_is_ad_high_quality_material value
func (v DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial) Ptr() *DiagnosisTaskAgentListV2DataTaskListIsAdHighQualityMaterial {
	return &v
}
