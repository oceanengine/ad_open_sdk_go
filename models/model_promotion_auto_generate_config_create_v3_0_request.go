/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAutoGenerateConfigCreateV30Request struct for PromotionAutoGenerateConfigCreateV30Request
type PromotionAutoGenerateConfigCreateV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// promotionId
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 策略数据(列表中strategy_id需唯一, 即同一个策略（strategy）的策略配置项列表(strategy_state)，需放到同一个对象内，不可分开传递)
	StrategyData []*PromotionAutoGenerateConfigCreateV30RequestStrategyDataInner `json:"strategy_data,omitempty"`
}
