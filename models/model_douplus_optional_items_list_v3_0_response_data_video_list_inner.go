/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOptionalItemsListV30ResponseDataVideoListInner struct for DouplusOptionalItemsListV30ResponseDataVideoListInner
type DouplusOptionalItemsListV30ResponseDataVideoListInner struct {
	//
	Height *int64 `json:"height,omitempty"`
	//
	IsHardSell *bool `json:"is_hard_sell,omitempty"`
	//
	IsImg *bool `json:"is_img,omitempty"`
	//
	ItemId *string `json:"item_id,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	Url []string `json:"url,omitempty"`
	//
	VideoCoverUrl []string `json:"video_cover_url,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
