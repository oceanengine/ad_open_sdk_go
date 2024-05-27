/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeGetV2ResponseDataListInner struct for ToolsClueLifeGetV2ResponseDataListInner
type ToolsClueLifeGetV2ResponseDataListInner struct {
	ActionType *ToolsClueLifeGetV2DataListActionType `json:"action_type,omitempty"`
	// 详细地址
	Address *string `json:"address,omitempty"`
	// 广告主名
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	// 年龄
	Age              *int64                                      `json:"age,omitempty"`
	AllocationStatus *ToolsClueLifeGetV2DataListAllocationStatus `json:"allocation_status,omitempty"`
	// 用户填写城市
	CityName *string `json:"city_name,omitempty"`
	// 线索ID
	ClueId *string `json:"clue_id,omitempty"`
	// 所属人姓名
	ClueOwnerName *string                             `json:"clue_owner_name,omitempty"`
	ClueType      *ToolsClueLifeGetV2DataListClueType `json:"clue_type,omitempty"`
	// 内容ID
	ContentId *string `json:"content_id,omitempty"`
	// 用户填写区县
	CountyName *string `json:"county_name,omitempty"`
	// 线索创建时间，如：2020-04-29 00:00:00
	CreateTimeDetail   *string                                       `json:"create_time_detail,omitempty"`
	EffectiveState     *ToolsClueLifeGetV2DataListEffectiveState     `json:"effective_state,omitempty"`
	EffectiveStateName *ToolsClueLifeGetV2DataListEffectiveStateName `json:"effective_state_name,omitempty"`
	FlowType           *ToolsClueLifeGetV2DataListFlowType           `json:"flow_type,omitempty"`
	// 跟进账户ID
	FollowLifeAccountId *string `json:"follow_life_account_id,omitempty"`
	// 跟进账户名称
	FollowLifeAccountName *string                                          `json:"follow_life_account_name,omitempty"`
	FollowLifeAccountType *ToolsClueLifeGetV2DataListFollowLifeAccountType `json:"follow_life_account_type,omitempty"`
	FollowStateName       *ToolsClueLifeGetV2DataListFollowStateName       `json:"follow_state_name,omitempty"`
	Gender                *ToolsClueLifeGetV2DataListGender                `json:"gender,omitempty"`
	LeadsPage             *ToolsClueLifeGetV2DataListLeadsPage             `json:"leads_page,omitempty"`
	// 广告主ID
	LocalAccountId *string `json:"local_account_id,omitempty"`
	// 线索修改时间，如：2020-04-29 00:00:00
	ModifyTime *string `json:"modify_time,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty"`
	// 订单ID
	OrderId *int64 `json:"order_id,omitempty"`
	// 广告ID
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 广告名称
	PromotionName *string `json:"promotion_name,omitempty"`
	// 用户填写省份
	ProvinceName *string `json:"province_name,omitempty"`
	// 商家备注
	Remark *string `json:"remark,omitempty"`
	// 商家表单自定义的字段信息，及其他线索相关信息
	RemarkDict *string `json:"remark_dict,omitempty"`
	// 当前线索对应广告的请求id
	ReqId *string `json:"req_id,omitempty"`
	// 线索被打上的系统标签，是一个标签项的数组
	SystemTags []string `json:"system_tags,omitempty"`
	// 线索被打上的人工标签，是一个标签项的数组，包括自定义标签和行业标签
	Tags []string `json:"tags,omitempty"`
	// 客户留资手机号。当团购订单退款后，不可获取明文手机号。
	Telephone *string `json:"telephone,omitempty"`
	// 线索工具ID
	ToolId *string `json:"tool_id,omitempty"`
}
