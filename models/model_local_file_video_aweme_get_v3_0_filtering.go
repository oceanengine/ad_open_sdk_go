/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoAwemeGetV30Filtering
type LocalFileVideoAwemeGetV30Filtering struct {
	AnchorInfo LocalFileVideoAwemeGetV30FilteringAnchorInfo `json:"anchor_info"`
	// 抖音号ids筛选,当筛选anchor_types=All不限的时候必传
	AwemeIds []string `json:"aweme_ids,omitempty"`
	// 根据视频发布时间进行过滤的截止时间，与start_time搭配使用，格式：\"yyyy-MM-dd HH:mm:ss\"
	EndTime *string `json:"end_time,omitempty"`
	// 主页视频ids筛选，一次最大长度：10
	ItemIds    []int64                                       `json:"item_ids,omitempty"`
	ItemStatus *LocalFileVideoAwemeGetV30FilteringItemStatus `json:"item_status,omitempty"`
	// 根据视频发布时间进行过滤的起始时间，与end_time搭配使用，格式：\"yyyy-MM-dd HH:mm:ss\"
	StartTime *string `json:"start_time,omitempty"`
}