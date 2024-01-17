/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfo 广告计划维度信息
type AdlabGroupDetailV30ResponseDataDataAdInfo struct {
	AdTarget *AdlabGroupDetailV30DataDataAdInfoAdTarget         `json:"ad_target,omitempty"`
	AppType  *AdlabGroupDetailV30DataDataAdInfoAppType          `json:"app_type,omitempty"`
	Audience *AdlabGroupDetailV30ResponseDataDataAdInfoAudience `json:"audience,omitempty"`
	// 自定义转化目标
	ConvertId *int64 `json:"convert_id,omitempty"`
	// 目标转化出价/预期成本
	CpaBid      *float64                                      `json:"cpa_bid,omitempty"`
	DeepBidType *AdlabGroupDetailV30DataDataAdInfoDeepBidType `json:"deep_bid_type,omitempty"`
	// 深度优化出价
	DeepCpaBid         *float64                                                `json:"deep_cpa_bid,omitempty"`
	DeepExternalAction *AdlabGroupDetailV30DataDataAdInfoDeepExternalAction    `json:"deep_external_action,omitempty"`
	DeliveryRange      *AdlabGroupDetailV30ResponseDataDataAdInfoDeliveryRange `json:"delivery_range,omitempty"`
	DownloadMode       *AdlabGroupDetailV30DataDataAdInfoDownloadMode          `json:"download_mode,omitempty"`
	DownloadType       *AdlabGroupDetailV30DataDataAdInfoDownloadType          `json:"download_type,omitempty"`
	// 下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// 投放结束时间
	EndTime        *string                                          `json:"end_time,omitempty"`
	ExternalAction *AdlabGroupDetailV30DataDataAdInfoExternalAction `json:"external_action,omitempty"`
	// 直达链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 应用包名
	Package *string                                   `json:"package,omitempty"`
	Pricing *AdlabGroupDetailV30DataDataAdInfoPricing `json:"pricing,omitempty"`
	// 商品信息
	ProductInfo []*AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInner `json:"product_info,omitempty"`
	// 深度转化ROI系数
	RoiGoal      *float64                                       `json:"roi_goal,omitempty"`
	ScheduleType *AdlabGroupDetailV30DataDataAdInfoScheduleType `json:"schedule_type,omitempty"`
	// 投放起始时间
	StartTime       *string                                                   `json:"start_time,omitempty"`
	TrackUrlSetting *AdlabGroupDetailV30ResponseDataDataAdInfoTrackUrlSetting `json:"track_url_setting,omitempty"`
	// 投放时段
	WeekSchedule []*[]int64 `json:"week_schedule,omitempty"`
}
