/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerTrackUrlGetV2ResponseData
type EventManagerTrackUrlGetV2ResponseData struct {
	PageInfo *EventManagerTrackUrlGetV2ResponseDataPageInfo `json:"page_info,omitempty"`
	// 监测链接组信息
	TrackUrlGroups []*EventManagerTrackUrlGetV2ResponseDataTrackUrlGroupsInner `json:"track_url_groups,omitempty"`
}
