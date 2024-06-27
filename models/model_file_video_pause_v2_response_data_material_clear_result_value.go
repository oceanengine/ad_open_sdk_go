/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoPauseV2ResponseDataMaterialClearResultValue struct for FileVideoPauseV2ResponseDataMaterialClearResultValue
type FileVideoPauseV2ResponseDataMaterialClearResultValue struct {
	// 清理失败的使用该素材的创意ID
	CreativeFailureList []int64 `json:"creative_failure_list,omitempty"`
	// 清理成功的使用该素材的创意ID
	CreativeSuccessList []int64 `json:"creative_success_list,omitempty"`
	// 素材是否被清理（存在内部错误导致素材没有被清理）
	IsCleared bool `json:"is_cleared"`
	// 素材id
	MaterialId int64 `json:"material_id"`
	// 素材未被清理原因
	Msg *string `json:"msg,omitempty"`
	// 清理失败的使用该素材的广告ID
	PromotionFailureList []int64 `json:"promotion_failure_list,omitempty"`
	// 清理成功的使用该素材的广告ID
	PromotionSuccessList []int64 `json:"promotion_success_list,omitempty"`
}
