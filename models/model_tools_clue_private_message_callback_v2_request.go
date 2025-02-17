/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCluePrivateMessageCallbackV2Request struct for ToolsCluePrivateMessageCallbackV2Request
type ToolsCluePrivateMessageCallbackV2Request struct {
	Adinfo *ToolsCluePrivateMessageCallbackV2RequestAdinfo `json:"adinfo,omitempty"`
	// 授权的广告主id，可以取抖开私信拉取接口里的adv_id
	AdvertiserId int64 `json:"advertiser_id"`
	// 接待商家的抖音open id
	BOpenId string `json:"b_open_id"`
	// 留资用户的抖音open id
	COpenId string `json:"c_open_id"`
	// 抖开接口里面的开发者唯一标识
	ClientKey       string                                                  `json:"client_key"`
	ClueConvertInfo ToolsCluePrivateMessageCallbackV2RequestClueConvertInfo `json:"clue_convert_info"`
	ClueData        ToolsCluePrivateMessageCallbackV2RequestClueData        `json:"clue_data"`
	// 事件/状态变更发生的秒级时间戳
	EventTimestamp *int64                                           `json:"event_timestamp,omitempty"`
	MsgInfo        *ToolsCluePrivateMessageCallbackV2RequestMsgInfo `json:"msg_info,omitempty"`
}
