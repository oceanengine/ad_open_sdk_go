/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestChallengeInfo
type StarDemandCreateChallengeV2RequestChallengeInfo struct {
	// 达人选择方式
	AuthorChooseType *int32 `json:"author_choose_type,omitempty"`
	//
	AuthorNoThresholdTask *bool `json:"author_no_threshold_task,omitempty"`
	// 达人侧任务名称 60字内
	AuthorTaskName string `json:"author_task_name"`
	// 总预算，单位元 正整数，最低10000
	Budget int64 `json:"budget"`
	// 投稿活动结束时间，unix时间戳（秒） 不可早于start_time
	EndTime int64 `json:"end_time"`
	//
	MarketingType *int64 `json:"marketing_type,omitempty"`
	// 单达人可投稿次数 范围1~20
	MaxUploadItemCnt         int64                                                                    `json:"max_upload_item_cnt"`
	ParticipateAuthorRange   StarDemandCreateChallengeV2RequestChallengeInfoParticipateAuthorRange    `json:"participate_author_range"`
	ParticipateMcnRange      *StarDemandCreateChallengeV2RequestChallengeInfoParticipateMcnRange      `json:"participate_mcn_range,omitempty"`
	ParticipateProviderRange *StarDemandCreateChallengeV2RequestChallengeInfoParticipateProviderRange `json:"participate_provider_range,omitempty"`
	// 是否给达人寄送样品 0：否；1：是
	SampleDelivery int64 `json:"sample_delivery"`
	// 示例视频ID（https://www.douyin.com/video/xxx 中的XXX部分） 最多3个
	SampleVideo    []int64                                                       `json:"sample_video,omitempty"`
	SettlementInfo StarDemandCreateChallengeV2RequestChallengeInfoSettlementInfo `json:"settlement_info"`
	// 投稿活动开始时间，unix时间戳（秒） 不可早于当前日期
	StartTime int64 `json:"start_time"`
	// 任务头图 通过上传材料接口上传的文件名
	TaskHeadImage *string `json:"task_head_image,omitempty"`
	// 任务图标 通过上传材料接口上传的文件名
	TaskIcon string `json:"task_icon"`
}
