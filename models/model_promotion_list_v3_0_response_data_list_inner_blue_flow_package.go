/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerBlueFlowPackage
type PromotionListV30ResponseDataListInnerBlueFlowPackage struct {
	//
	BlueFlowPackageId      *int64                                                         `json:"blue_flow_package_id,omitempty"`
	BlueFlowPackageSetting *PromotionListV30DataListBlueFlowPackageBlueFlowPackageSetting `json:"blue_flow_package_setting,omitempty"`
	//
	BlueFlowSuggestion *string `json:"blue_flow_suggestion,omitempty"`
}
