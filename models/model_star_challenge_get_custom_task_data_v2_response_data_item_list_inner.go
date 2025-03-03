/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeGetCustomTaskDataV2ResponseDataItemListInner struct for StarChallengeGetCustomTaskDataV2ResponseDataItemListInner
type StarChallengeGetCustomTaskDataV2ResponseDataItemListInner struct {
	// 含作弊的Android下载数/Android激活数
	AndroidConvertCount *int32                                                               `json:"android_convert_count,omitempty"`
	AuthorInfo          *StarChallengeGetCustomTaskDataV2ResponseDataItemListInnerAuthorInfo `json:"author_info,omitempty"`
	// channel_id
	ChannelId *string `json:"channel_id,omitempty"`
	// 评论量
	Comment *int64 `json:"comment,omitempty"`
	// 组件点击量
	ComponentClickCount *int32 `json:"component_click_count,omitempty"`
	// 投稿机构ID
	ContractMcnId *int64 `json:"contract_mcn_id,omitempty"`
	// 投稿机构名称
	ContractMcnName *string `json:"contract_mcn_name,omitempty"`
	// 促活转化数（线索量/安装量等）
	ConvertCount *int64 `json:"convert_count,omitempty"`
	// 视频封面图
	CoverUrl *string `json:"cover_url,omitempty"`
	// 发布时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 去除作弊后的ios下载数/ios激活数
	InvalidAndroidConvertCount *int64 `json:"invalid_android_convert_count,omitempty"`
	// 含作弊的ios下载数/ios激活数
	InvalidIosConvertCount *int64 `json:"invalid_ios_convert_count,omitempty"`
	// 去除作弊后的Android下载数/Android激活数
	IosConvertCount *int32 `json:"ios_convert_count,omitempty"`
	// 作品id
	ItemId *int64 `json:"item_id,omitempty"`
	// 点赞量
	LikeCnt *int64 `json:"like_cnt,omitempty"`
	// acu
	LiveAcu *int64 `json:"live_acu,omitempty"`
	// 直播时长
	LiveDuration *int64 `json:"live_duration,omitempty"`
	// pcu
	LivePcu *int64 `json:"live_pcu,omitempty"`
	// 播放量
	Play *int64 `json:"play,omitempty"`
	// 预约成功次数（仅安卓）
	ReserveCnt *int64 `json:"reserve_cnt,omitempty"`
	// 预约安装完成次数
	ReserveInstallCnt *int64 `json:"reserve_install_cnt,omitempty"`
	// 奖励金额
	RewardAmount *int64 `json:"reward_amount,omitempty"`
	// 分享量
	Share *int64 `json:"share,omitempty"`
	// 有效播放量
	ShareVv *int32 `json:"share_vv,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 相关性审核（1-通过）
	Uncorrelated *int32 `json:"uncorrelated,omitempty"`
	// 链接
	Url *string `json:"url,omitempty"`
	// 视频时长
	VideoDurationMs *int64 `json:"video_duration_ms,omitempty"`
	// pv
	WatchCnt *int64 `json:"watch_cnt,omitempty"`
	// 直播人均时长
	WatchDurationAvg *float64 `json:"watch_duration_avg,omitempty"`
	// uv
	WatchUv *int64 `json:"watch_uv,omitempty"`
	// 有效评论量
	WhiteComment *int32 `json:"white_comment,omitempty"`
	// 有效点赞量
	WhiteLikeCnt *int32 `json:"white_like_cnt,omitempty"`
	// 有效分享量
	WhiteShare *int32 `json:"white_share,omitempty"`
}
