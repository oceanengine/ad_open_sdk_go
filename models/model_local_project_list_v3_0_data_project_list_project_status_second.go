/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectListV30DataProjectListProjectStatusSecond
type LocalProjectListV30DataProjectListProjectStatusSecond string

// List of local_project_list_v3.0_data_project_list_project_status_second
const (
	PROJECT_STATUS_BUDGET_EXCEED_LocalProjectListV30DataProjectListProjectStatusSecond LocalProjectListV30DataProjectListProjectStatusSecond = "PROJECT_STATUS_BUDGET_EXCEED"
	PROJECT_STATUS_DISABLE_LocalProjectListV30DataProjectListProjectStatusSecond       LocalProjectListV30DataProjectListProjectStatusSecond = "PROJECT_STATUS_DISABLE"
	PROJECT_STATUS_NOT_SCHEDULE_LocalProjectListV30DataProjectListProjectStatusSecond  LocalProjectListV30DataProjectListProjectStatusSecond = "PROJECT_STATUS_NOT_SCHEDULE"
	PROJECT_STATUS_NOT_START_LocalProjectListV30DataProjectListProjectStatusSecond     LocalProjectListV30DataProjectListProjectStatusSecond = "PROJECT_STATUS_NOT_START"
)

// Ptr returns reference to local_project_list_v3.0_data_project_list_project_status_second value
func (v LocalProjectListV30DataProjectListProjectStatusSecond) Ptr() *LocalProjectListV30DataProjectListProjectStatusSecond {
	return &v
}
