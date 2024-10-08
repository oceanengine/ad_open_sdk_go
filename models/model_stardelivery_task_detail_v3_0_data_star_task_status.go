/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskDetailV30DataStarTaskStatus
type StardeliveryTaskDetailV30DataStarTaskStatus string

// List of stardelivery_task_detail_v3.0_data_star_task_status
const (
	BILLING_IN_PROGRESS_StardeliveryTaskDetailV30DataStarTaskStatus StardeliveryTaskDetailV30DataStarTaskStatus = "BILLING_IN_PROGRESS"
	CANCELLED_StardeliveryTaskDetailV30DataStarTaskStatus           StardeliveryTaskDetailV30DataStarTaskStatus = "CANCELLED"
	COMPLETED_StardeliveryTaskDetailV30DataStarTaskStatus           StardeliveryTaskDetailV30DataStarTaskStatus = "COMPLETED"
	PENDING_LAUNCH_StardeliveryTaskDetailV30DataStarTaskStatus      StardeliveryTaskDetailV30DataStarTaskStatus = "PENDING_LAUNCH"
	PREPARATION_StardeliveryTaskDetailV30DataStarTaskStatus         StardeliveryTaskDetailV30DataStarTaskStatus = "PREPARATION"
)

// Ptr returns reference to stardelivery_task_detail_v3.0_data_star_task_status value
func (v StardeliveryTaskDetailV30DataStarTaskStatus) Ptr() *StardeliveryTaskDetailV30DataStarTaskStatus {
	return &v
}
