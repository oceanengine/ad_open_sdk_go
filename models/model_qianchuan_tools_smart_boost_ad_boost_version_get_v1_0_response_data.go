/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostVersionGetV10ResponseData
type QianchuanToolsSmartBoostAdBoostVersionGetV10ResponseData struct {
	// json返回值
	RaiseListInfos []*QianchuanToolsSmartBoostAdBoostVersionGetV10ResponseDataRaiseListInfosInner `json:"raise_list_infos,omitempty"`
	// 总数
	TotalCnt *int32 `json:"total_cnt,omitempty"`
}
