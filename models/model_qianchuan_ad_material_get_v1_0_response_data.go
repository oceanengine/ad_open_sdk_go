/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10ResponseData
type QianchuanAdMaterialGetV10ResponseData struct {
	// 返回的素材信息列表
	AdMaterialInfos []*QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInner `json:"ad_material_infos"`
	PageInfo        *QianchuanAdMaterialGetV10ResponseDataPageInfo               `json:"page_info,omitempty"`
}
