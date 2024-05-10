/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerTrackUrlCreateV2Request struct for EventManagerTrackUrlCreateV2Request
type EventManagerTrackUrlCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 资产ID
	AssetsId int64 `json:"assets_id"`
	// 应用下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// 监测链接组信息，IOS和安卓应用可绑定多组监测链接
	TrackUrlGroups []*EventManagerTrackUrlCreateV2RequestTrackUrlGroupsInner `json:"track_url_groups"`
}
