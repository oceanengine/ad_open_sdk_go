/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionEasyKeepDeliverySwitchUpdateV30Request struct for PromotionEasyKeepDeliverySwitchUpdateV30Request
type PromotionEasyKeepDeliverySwitchUpdateV30Request struct {
	//
	AdvertiserId       int64                                                      `json:"advertiser_id"`
	KeepDeliverySwitch PromotionEasyKeepDeliverySwitchUpdateV30KeepDeliverySwitch `json:"keep_delivery_switch"`
	//
	ProjectIds []int64 `json:"project_ids"`
}
