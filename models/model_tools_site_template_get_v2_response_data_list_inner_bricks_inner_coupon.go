/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerCoupon 发券组件描述
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerCoupon struct {
	// 活动ID，当`coupon`不为空时，有值。用户可以通过[【获取卡券列表】]（https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710639364108#item-link-%E6%A6%82%E5%BF%B5%E4%BB%8B%E7%BB%8D）接口或[【青鸟线索通平台】]（https://clue.oceanengine.com/）获取活动ID
	ActivityId int64 `json:"activity_id"`
}
