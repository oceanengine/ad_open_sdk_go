/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmCreateChallengeV2RequestChallengeInfo
type StarDemandOmCreateChallengeV2RequestChallengeInfo struct {
	// 最长分账周期
	AccountDivideDay int64 `json:"account_divide_day"`
	// 组件标题 最长14字
	AnchorTitle string                                                `json:"anchor_title"`
	AuthorScope StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope `json:"author_scope"`
	// 达人侧任务名称 字符串，最长17
	AuthorTaskName string `json:"author_task_name"`
	// 分佣比例 整数，范围根据任务类型限制
	CommissionRate *int64                                                   `json:"commission_rate,omitempty"`
	CommissionType StarDemandOmCreateChallengeV2ChallengeInfoCommissionType `json:"commission_type"`
	// 任务介绍 文本，最长500字
	DemandDesc *string `json:"demand_desc,omitempty"`
	// 小程序ID
	MicroAppId string `json:"micro_app_id"`
	// 任务标签 list<string> 长度固定为2  [${形式标签}, ${内容标签}]
	OmTaskTag              []string                                                                `json:"om_task_tag"`
	ParticipateAuthorRange StarDemandOmCreateChallengeV2RequestChallengeInfoParticipateAuthorRange `json:"participate_author_range"`
	// 示例视频id list<i64> 最多5个
	SampleVideo []int64 `json:"sample_video,omitempty"`
	// 小程序落地页地址 字符串，支持校验合规性
	StartPage *string `json:"start_page,omitempty"`
	// 任务头图 URL文本
	TaskHeadImage *string `json:"task_head_image,omitempty"`
	// 任务图标 URL文本
	TaskIcon string `json:"task_icon"`
}
