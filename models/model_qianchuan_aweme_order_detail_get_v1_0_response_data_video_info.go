/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataVideoInfo 视频信息
type QianchuanAwemeOrderDetailGetV10ResponseDataVideoInfo struct {
	// 视频封面
	AwemeItemCover *string `json:"aweme_item_cover,omitempty"`
	// 视频id
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	// 视频标题
	AwemeItemTitle *string                                               `json:"aweme_item_title,omitempty"`
	ItemType       *QianchuanAwemeOrderDetailGetV10DataVideoInfoItemType `json:"item_type,omitempty"`
}
