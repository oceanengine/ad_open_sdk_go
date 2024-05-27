/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseItemListV10ResponseDataItemListInner struct for EnterpriseItemListV10ResponseDataItemListInner
type EnterpriseItemListV10ResponseDataItemListInner struct {
	//
	CommentCount *string `json:"comment_count,omitempty"`
	//
	ItemCoverUrl *string `json:"item_cover_url,omitempty"`
	//
	ItemCreateTime *string `json:"item_create_time,omitempty"`
	//
	ItemDuration *int64 `json:"item_duration,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	ItemTitle *string                                    `json:"item_title,omitempty"`
	ItemType  *EnterpriseItemListV10DataItemListItemType `json:"item_type,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
}
