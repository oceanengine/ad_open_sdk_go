/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareStatsDataV10ResponseData
type QianchuanProductAnalyseCompareStatsDataV10ResponseData struct {
	CompeteProductData *QianchuanProductAnalyseCompareStatsDataV10ResponseDataCompeteProductData `json:"compete_product_data,omitempty"`
	OwnProductData     *QianchuanProductAnalyseCompareStatsDataV10ResponseDataOwnProductData     `json:"own_product_data,omitempty"`
	SimilarProductData *QianchuanProductAnalyseCompareStatsDataV10ResponseDataSimilarProductData `json:"similar_product_data,omitempty"`
}
