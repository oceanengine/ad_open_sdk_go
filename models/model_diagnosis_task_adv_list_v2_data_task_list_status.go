/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAdvListV2DataTaskListStatus
type DiagnosisTaskAdvListV2DataTaskListStatus string

// List of diagnosis_task_adv_list_v2_data_task_list_status
const (
	FAILED_DiagnosisTaskAdvListV2DataTaskListStatus  DiagnosisTaskAdvListV2DataTaskListStatus = "FAILED"
	PENDING_DiagnosisTaskAdvListV2DataTaskListStatus DiagnosisTaskAdvListV2DataTaskListStatus = "PENDING"
	SUCCESS_DiagnosisTaskAdvListV2DataTaskListStatus DiagnosisTaskAdvListV2DataTaskListStatus = "SUCCESS"
)

// Ptr returns reference to diagnosis_task_adv_list_v2_data_task_list_status value
func (v DiagnosisTaskAdvListV2DataTaskListStatus) Ptr() *DiagnosisTaskAdvListV2DataTaskListStatus {
	return &v
}
