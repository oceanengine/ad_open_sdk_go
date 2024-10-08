/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostGroupListV2ResponseDataBoostGroupInfosInner struct for StarVasGetBoostGroupListV2ResponseDataBoostGroupInfosInner
type StarVasGetBoostGroupListV2ResponseDataBoostGroupInfosInner struct {
	// 手动出价，乘百取整（如：0.1 -> 10）。如果为空则没有
	BidAmount   *int64                                                   `json:"bid_amount,omitempty"`
	BidType     *StarVasGetBoostGroupListV2DataBoostGroupInfosBidType    `json:"bid_type,omitempty"`
	BoostAction StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction `json:"boost_action"`
	// 助推预算，单位元
	BoostAmount int64 `json:"boost_amount"`
	// 助推消耗，单位千分之一元
	BoostCost int64 `json:"boost_cost"`
	// 自定义投放时长，单位时。如果为空则没有
	BoostHours *int64                                                 `json:"boost_hours,omitempty"`
	BoostType  StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType `json:"boost_type"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 关联人群包，仅boost_type为CUSTOM_PACKAGE有效
	PackId *int64                                              `json:"pack_id,omitempty"`
	Status StarVasGetBoostGroupListV2DataBoostGroupInfosStatus `json:"status"`
	// 助推组ID
	TaskId int64 `json:"task_id"`
	// 助推的指派单信息
	TaskInfos []*StarVasGetBoostGroupListV2ResponseDataBoostGroupInfosInnerTaskInfosInner `json:"task_infos"`
	// 助推组名称
	TaskName string `json:"task_name"`
}
