/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30DataListStatusFirst
type ProjectListV30DataListStatusFirst string

// List of project_list_v3.0_data_list_status_first
const (
	PROJECT_STATUS_DELETE_ProjectListV30DataListStatusFirst  ProjectListV30DataListStatusFirst = "PROJECT_STATUS_DELETE"
	PROJECT_STATUS_DISABLE_ProjectListV30DataListStatusFirst ProjectListV30DataListStatusFirst = "PROJECT_STATUS_DISABLE"
	PROJECT_STATUS_DONE_ProjectListV30DataListStatusFirst    ProjectListV30DataListStatusFirst = "PROJECT_STATUS_DONE"
	PROJECT_STATUS_ENABLE_ProjectListV30DataListStatusFirst  ProjectListV30DataListStatusFirst = "PROJECT_STATUS_ENABLE"
)

// Ptr returns reference to project_list_v3.0_data_list_status_first value
func (v ProjectListV30DataListStatusFirst) Ptr() *ProjectListV30DataListStatusFirst {
	return &v
}
