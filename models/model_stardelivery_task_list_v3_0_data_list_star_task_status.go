/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskListV30DataListStarTaskStatus
type StardeliveryTaskListV30DataListStarTaskStatus string

// List of stardelivery_task_list_v3.0_data_list_star_task_status
const (
	BILLING_IN_PROGRESS_StardeliveryTaskListV30DataListStarTaskStatus StardeliveryTaskListV30DataListStarTaskStatus = "BILLING_IN_PROGRESS"
	CANCELLED_StardeliveryTaskListV30DataListStarTaskStatus           StardeliveryTaskListV30DataListStarTaskStatus = "CANCELLED"
	COMPLETED_StardeliveryTaskListV30DataListStarTaskStatus           StardeliveryTaskListV30DataListStarTaskStatus = "COMPLETED"
	PENDING_LAUNCH_StardeliveryTaskListV30DataListStarTaskStatus      StardeliveryTaskListV30DataListStarTaskStatus = "PENDING_LAUNCH"
	PREPARATION_StardeliveryTaskListV30DataListStarTaskStatus         StardeliveryTaskListV30DataListStarTaskStatus = "PREPARATION"
)

// Ptr returns reference to stardelivery_task_list_v3.0_data_list_star_task_status value
func (v StardeliveryTaskListV30DataListStarTaskStatus) Ptr() *StardeliveryTaskListV30DataListStarTaskStatus {
	return &v
}
