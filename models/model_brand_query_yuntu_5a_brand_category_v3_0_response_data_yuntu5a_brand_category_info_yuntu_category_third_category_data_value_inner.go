/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategoryThirdCategoryDataValueInner struct for BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategoryThirdCategoryDataValueInner
type BrandQueryYuntu5aBrandCategoryV30ResponseDataYuntu5aBrandCategoryInfoYuntuCategoryThirdCategoryDataValueInner struct {
	//
	BrandIds []int64 `json:"brand_ids,omitempty"`
	// 三级级行业分类ID
	Id *int64 `json:"id,omitempty"`
	// 三级级行业分类名称
	Value *string `json:"value,omitempty"`
}
