/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInnerPromotionCardMaterial
type QianchuanCreativeGetV10ResponseDataListInnerPromotionCardMaterial struct {
	//
	ActionButton            *string                                                                      `json:"action_button,omitempty"`
	ButtonSmartOptimization *QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization `json:"button_smart_optimization,omitempty"`
	//
	ComponentId *int64 `json:"component_id,omitempty"`
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	SellingPoints []string `json:"selling_points,omitempty"`
	//
	Title *string `json:"title,omitempty"`
}
