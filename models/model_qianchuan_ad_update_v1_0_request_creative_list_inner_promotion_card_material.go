/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestCreativeListInnerPromotionCardMaterial
type QianchuanAdUpdateV10RequestCreativeListInnerPromotionCardMaterial struct {
	//
	ActionButton            *string                                                                       `json:"action_button,omitempty"`
	ButtonSmartOptimization *QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization `json:"button_smart_optimization,omitempty"`
	//
	ComponentId *int64 `json:"component_id,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	SellingPoints []string `json:"selling_points,omitempty"`
	//
	Title *string `json:"title,omitempty"`
}
