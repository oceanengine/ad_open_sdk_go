/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfo
type StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfo struct {
	//
	AuthorTaskName string `json:"author_task_name"`
	//
	Budget int64 `json:"budget"`
	//
	EndTime int64 `json:"end_time"`
	//
	MaxUploadItemCnt       int64                                                                               `json:"max_upload_item_cnt"`
	ParticipateAuthorRange StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoParticipateAuthorRange `json:"participate_author_range"`
	//
	SampleDelivery int64 `json:"sample_delivery"`
	//
	SampleVideo    []int64                                                                     `json:"sample_video,omitempty"`
	SettlementInfo StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoSettlementInfo `json:"settlement_info"`
	//
	StartTime int64 `json:"start_time"`
	//
	TaskHeadImage *string `json:"task_head_image,omitempty"`
	//
	TaskIcon string `json:"task_icon"`
}
