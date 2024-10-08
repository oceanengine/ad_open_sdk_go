/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasCreateBoostItemGroupV2Request struct for StarVasCreateBoostItemGroupV2Request
type StarVasCreateBoostItemGroupV2Request struct {
	// 手动出价，乘百取整（如：0.1 -> 10）。仅在boost_action为LINK_ACTION传参。如果不传，则为自动出价。可出价范围0.1~1000
	BidAmount   *int64                                   `json:"bid_amount,omitempty"`
	BoostAction StarVasCreateBoostItemGroupV2BoostAction `json:"boost_action"`
	// 助推金额，单位元。范围为1千~1亿，且为100的倍数
	BoostAmount int64 `json:"boost_amount"`
	// 自定义投放时长，单位时，仅且必须在boost_action为LINK_ACTION传参。时长范围为2天~7天。
	BoostHours        *int32                                                 `json:"boost_hours,omitempty"`
	BoostType         StarVasCreateBoostItemGroupV2BoostType                 `json:"boost_type"`
	CustomAudienceTag *StarVasCreateBoostItemGroupV2RequestCustomAudienceTag `json:"custom_audience_tag,omitempty"`
	// 加热组名称
	Name string `json:"name"`
	// 关联指派单，指派单必须在当前同一主体账户下
	OrderIds []int64 `json:"order_ids"`
	// 关联人群包，仅boost_type为CUSTOM_PACKAGE时传参
	PackId *int64 `json:"pack_id,omitempty"`
	// 客户ID
	StarId int64 `json:"star_id"`
}
