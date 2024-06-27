/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfoAdSyncConfDouPlus 产出物推DOU+配置信息（仅需advertiser_id和dou_plus_uid）
type StarDemandCreateAssignV2RequestDemandInfoAdSyncConfDouPlus struct {
	// 是否用于视频素材投放（头像跳落地页）：2=是；1=否（用于原视频投放） 默认推送素材（为2，头像跳落地页）。内容服务为2不可改
	AdSync *int32 `json:"ad_sync,omitempty"`
	// 是否投放原生视频 0或不填 = 否；1 = 是
	AdSyncOrigin *int64 `json:"ad_sync_origin,omitempty"`
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 是否自动投放 0或不填 = 否；1 = 是
	AutoSync *int64 `json:"auto_sync,omitempty"`
	// 推DOU+时可以使用此字段完成抖音UID至广告主ID的转换。若存在advertiser_id则忽略该参数。
	DouPlusUid *int64 `json:"dou_plus_uid,omitempty"`
	// 推广产品链接
	ProductLink *string `json:"product_link,omitempty"`
	// 推广产品图片文件名（材料上传接口中使用到的文件名）
	ProductPics []string `json:"product_pics,omitempty"`
	// 投放原生视频的投放时间（单位：天） 大于0的整数
	SyncDuration *int64 `json:"sync_duration,omitempty"`
}
