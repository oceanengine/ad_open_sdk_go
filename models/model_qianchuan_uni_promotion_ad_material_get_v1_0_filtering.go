/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdMaterialGetV10Filtering
type QianchuanUniPromotionAdMaterialGetV10Filtering struct {
	//
	AnalysisType []*QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType `json:"analysis_type,omitempty"`
	// 查询结束日期
	EndDate            *string                                                           `json:"end_date,omitempty"`
	MaterialSelectType *QianchuanUniPromotionAdMaterialGetV10FilteringMaterialSelectType `json:"material_select_type,omitempty"`
	MaterialStatus     *QianchuanUniPromotionAdMaterialGetV10FilteringMaterialStatus     `json:"material_status,omitempty"`
	MaterialType       QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType        `json:"material_type"`
	//
	ProductKeyword *string `json:"product_keyword,omitempty"`
	//
	SearchKeyword *string `json:"search_keyword,omitempty"`
	// 查询起始日期
	StartDate *string                                                  `json:"start_date,omitempty"`
	VideoType *QianchuanUniPromotionAdMaterialGetV10FilteringVideoType `json:"video_type,omitempty"`
}
