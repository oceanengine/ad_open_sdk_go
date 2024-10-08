/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAdvListV2Results
type DiagnosisTaskAdvListV2Results string

// List of diagnosis_task_adv_list_v2_results
const (
	AD_HIGH_QUALITY_MATERIAL_DiagnosisTaskAdvListV2Results      DiagnosisTaskAdvListV2Results = "AD_HIGH_QUALITY_MATERIAL"
	ECP_HIGH_QUALITY_MATERIAL_DiagnosisTaskAdvListV2Results     DiagnosisTaskAdvListV2Results = "ECP_HIGH_QUALITY_MATERIAL"
	FIRST_PUBLISH_MATERIAL_DiagnosisTaskAdvListV2Results        DiagnosisTaskAdvListV2Results = "FIRST_PUBLISH_MATERIAL"
	INEFFICIENT_MATERIAL_DiagnosisTaskAdvListV2Results          DiagnosisTaskAdvListV2Results = "INEFFICIENT_MATERIAL"
	NON_AD_HIGH_QUALITY_MATERIAL_DiagnosisTaskAdvListV2Results  DiagnosisTaskAdvListV2Results = "NON_AD_HIGH_QUALITY_MATERIAL"
	NON_ECP_HIGH_QUALITY_MATERIAL_DiagnosisTaskAdvListV2Results DiagnosisTaskAdvListV2Results = "NON_ECP_HIGH_QUALITY_MATERIAL"
	NON_FIRST_PUBLISH_MATERIAL_DiagnosisTaskAdvListV2Results    DiagnosisTaskAdvListV2Results = "NON_FIRST_PUBLISH_MATERIAL"
	NON_INEFFICIENT_MATERIAL_DiagnosisTaskAdvListV2Results      DiagnosisTaskAdvListV2Results = "NON_INEFFICIENT_MATERIAL"
)

// Ptr returns reference to diagnosis_task_adv_list_v2_results value
func (v DiagnosisTaskAdvListV2Results) Ptr() *DiagnosisTaskAdvListV2Results {
	return &v
}
