/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAdvGetV2DataTaskListStatus
type DiagnosisTaskAdvGetV2DataTaskListStatus string

// List of diagnosis_task_adv_get_v2_data_task_list_status
const (
	FAILED_DiagnosisTaskAdvGetV2DataTaskListStatus  DiagnosisTaskAdvGetV2DataTaskListStatus = "FAILED"
	PENDING_DiagnosisTaskAdvGetV2DataTaskListStatus DiagnosisTaskAdvGetV2DataTaskListStatus = "PENDING"
	SUCCESS_DiagnosisTaskAdvGetV2DataTaskListStatus DiagnosisTaskAdvGetV2DataTaskListStatus = "SUCCESS"
)

// Ptr returns reference to diagnosis_task_adv_get_v2_data_task_list_status value
func (v DiagnosisTaskAdvGetV2DataTaskListStatus) Ptr() *DiagnosisTaskAdvGetV2DataTaskListStatus {
	return &v
}
