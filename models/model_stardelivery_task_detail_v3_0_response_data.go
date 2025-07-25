/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskDetailV30ResponseData
type StardeliveryTaskDetailV30ResponseData struct {
	// 任务内推广的Android应用包名
	AndroidAppPackageName *string `json:"android_app_package_name,omitempty"`
	// 任务内推广的Android应用链接
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	// 任务内Android落地页，橙子落地页，仅1个
	AndroidExternalUrl *string                                                `json:"android_external_url,omitempty"`
	AudienceSettings   *StardeliveryTaskDetailV30ResponseDataAudienceSettings `json:"audience_settings,omitempty"`
	//
	AuthorSubmitFrequency *int32 `json:"author_submit_frequency,omitempty"`
	// 达人侧任务名称，1-50个字符，达人浏览任务时会看到此名称
	AuthorTaskName *string                                 `json:"author_task_name,omitempty"`
	BriefType      *StardeliveryTaskDetailV30DataBriefType `json:"brief_type,omitempty"`
	// 一级行业分类ID
	FirstIndustryId *int64 `json:"first_industry_id,omitempty"`
	// 一级行业分类Name
	FirstIndustryName *string `json:"first_industry_name,omitempty"`
	// 任务内推广的iOS应用包名
	IosAppPackageName *string `json:"ios_app_package_name,omitempty"`
	// 任务内推广的iOS应用链接
	IosDownloadUrl *string `json:"ios_download_url,omitempty"`
	// 任务内iOS落地页，橙子落地页，仅1个
	IosExternalUrl *string `json:"ios_external_url,omitempty"`
	// 任务所选锚点关联的小程序/小游戏appid
	MicroAppId *string `json:"micro_app_id,omitempty"`
	// 字节小程序链接，创建任务时填写的链接，未设置时不会返回
	MicroAppLink *string `json:"micro_app_link,omitempty"`
	// 任务所选锚点关联的小程序/小游戏名称
	MicroAppName *string `json:"micro_app_name,omitempty"`
	// 微信小程序/小游戏路径参数
	MicroAppPath                  *string                                                             `json:"micro_app_path,omitempty"`
	ProductAdditionalIntroduction *StardeliveryTaskDetailV30ResponseDataProductAdditionalIntroduction `json:"product_additional_introduction,omitempty"`
	// 产品介绍，1-1000个字符，介绍中出现数据、专利/奖项、影视片段、明星/网红、活动等信息
	ProductIntroduction *string `json:"product_introduction,omitempty"`
	// 产品名称，1-40个字符，任务推广的产品名称
	ProductName *string `json:"product_name,omitempty"`
	// 二级行业分类ID
	SecondIndustryId *int64 `json:"second_industry_id,omitempty"`
	// 二级行业分类Name
	SecondIndustryName *string `json:"second_industry_name,omitempty"`
	// 来源星图账户名称
	StarAccountName *string `json:"star_account_name,omitempty"`
	// 广告消耗分成比例
	StarAdCostDivideRatio *float64 `json:"star_ad_cost_divide_ratio,omitempty"`
	// 来源星图账户id
	StarId *int64 `json:"star_id,omitempty"`
	// 素材出价（元），每条视频最低需要付给达人的底价价格
	StarMaterialBid *float64 `json:"star_material_bid,omitempty"`
	// 基础素材费单价区间，报价区间为建议值，系统优先邀约符合报价区间的达人
	StarMaterialBidRange []float64 `json:"star_material_bid_range,omitempty"`
	// 素材一级类目id
	StarMaterialFirstType *int32 `json:"star_material_first_type,omitempty"`
	// 素材二级类目id
	StarMaterialSecondType    *int32                                                          `json:"star_material_second_type,omitempty"`
	StarTaskAlbumMicroAppInfo *StardeliveryTaskDetailV30ResponseDataStarTaskAlbumMicroAppInfo `json:"star_task_album_micro_app_info,omitempty"`
	// 原生锚点ID
	StarTaskAnchorId   *int64                                                  `json:"star_task_anchor_id,omitempty"`
	StarTaskAnchorType *StardeliveryTaskDetailV30DataStarTaskAnchorType        `json:"star_task_anchor_type,omitempty"`
	StarTaskAssetInfo  *StardeliveryTaskDetailV30ResponseDataStarTaskAssetInfo `json:"star_task_asset_info,omitempty"`
	// 任务审核拒绝原因和意见，没有则返回为空
	StarTaskBanReasonDetail *string `json:"star_task_ban_reason_detail,omitempty"`
	// 任务出价（元）
	StarTaskBid *float64 `json:"star_task_bid,omitempty"`
	// 任务预算（元），只能是10000-1000000元之间的整数
	StarTaskBudget *float64 `json:"star_task_budget,omitempty"`
	// 任务创建时间，格式YYYY-MM-DD
	StarTaskCreateTime *string `json:"star_task_create_time,omitempty"`
	// 服务商延长后的实际投稿截止时间
	StarTaskDelayedPostEndTime *string                                              `json:"star_task_delayed_post_end_time,omitempty"`
	StarTaskExternalAction     *StardeliveryTaskDetailV30DataStarTaskExternalAction `json:"star_task_external_action,omitempty"`
	// 星广联投任务ID
	StarTaskId *int64 `json:"star_task_id,omitempty"`
	// 任务内介绍文案
	StarTaskIntroductionText       *string                                                             `json:"star_task_introduction_text,omitempty"`
	StarTaskMaterialPurchaseMethod *StardeliveryTaskDetailV30DataStarTaskMaterialPurchaseMethod        `json:"star_task_material_purchase_method,omitempty"`
	StarTaskMaterialsRequirements  *StardeliveryTaskDetailV30ResponseDataStarTaskMaterialsRequirements `json:"star_task_materials_requirements,omitempty"`
	// 星广联投任务名称
	StarTaskName *string `json:"star_task_name,omitempty"`
	// 任务结束时间为投稿截止时间,格式YYYY-MM-DD
	StarTaskPostEndTime *string                                      `json:"star_task_post_end_time,omitempty"`
	StarTaskSource      *StardeliveryTaskDetailV30DataStarTaskSource `json:"star_task_source,omitempty"`
	// 任务开始时间，格式YYYY-MM-DD
	StarTaskStartTime *string                                         `json:"star_task_start_time,omitempty"`
	StarTaskStatus    *StardeliveryTaskDetailV30DataStarTaskStatus    `json:"star_task_status,omitempty"`
	StarTaskSubStatus *StardeliveryTaskDetailV30DataStarTaskSubStatus `json:"star_task_sub_status,omitempty"`
	// 任务标签
	StarTaskTags *string `json:"star_task_tags,omitempty"`
	// 任务图标uri
	TaskAvatarId *string `json:"task_avatar_id,omitempty"`
	// 任务图标url
	TaskAvatarUrl   *string                                               `json:"task_avatar_url,omitempty"`
	TaskContactInfo *StardeliveryTaskDetailV30ResponseDataTaskContactInfo `json:"task_contact_info,omitempty"`
	// 任务头图uri
	TaskHeadImageId *string `json:"task_head_image_id,omitempty"`
}
