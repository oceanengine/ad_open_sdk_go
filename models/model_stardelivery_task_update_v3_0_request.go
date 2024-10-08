/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskUpdateV30Request struct for StardeliveryTaskUpdateV30Request
type StardeliveryTaskUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AuthorSubmitFrequency int32 `json:"author_submit_frequency"`
	//
	AuthorTaskName string `json:"author_task_name"`
	//
	StarTaskId                    int64                                                         `json:"star_task_id"`
	StarTaskMaterialsRequirements StardeliveryTaskUpdateV30RequestStarTaskMaterialsRequirements `json:"star_task_materials_requirements"`
	//
	StarTaskName string `json:"star_task_name"`
	//
	TaskAvatarId    string                                          `json:"task_avatar_id"`
	TaskContactInfo StardeliveryTaskUpdateV30RequestTaskContactInfo `json:"task_contact_info"`
	//
	TaskHeadImageId string `json:"task_head_image_id"`
	//
	TaskProductIntroduction string `json:"task_product_introduction"`
	//
	TaskProductName string `json:"task_product_name"`
}
