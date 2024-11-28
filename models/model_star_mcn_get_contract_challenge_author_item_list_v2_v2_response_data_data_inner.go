/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInner struct for StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInner
type StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInner struct {
	//
	AndroidConvertCount *int64 `json:"android_convert_count,omitempty"`
	// 达人已发放收益
	AuthorAmount *int64 `json:"author_amount,omitempty"`
	//
	AuthorId *int64 `json:"author_id,omitempty"`
	//
	AuthorName *string `json:"author_name,omitempty"`
	// 计费明细数据
	BillDetail *string `json:"bill_detail,omitempty"`
	// 用户参与的渠道ID
	ChannelId *string `json:"channel_id,omitempty"`
	// 线索量
	ClueCnt *int64 `json:"clue_cnt,omitempty"`
	// 作品封面url（有效期1小时）
	CoverUrl *string `json:"cover_url,omitempty"`
	//
	DemandId int64 `json:"demand_id"`
	//
	DemandName string `json:"demand_name"`
	//
	DouyinId *string `json:"douyin_id,omitempty"`
	// 累计广告收益
	IaaCommercialIncomeHour *int64 `json:"iaa_commercial_income_hour,omitempty"`
	// item的iaa收益
	IaaIncome *int64 `json:"iaa_income,omitempty"`
	// 累计订单金额
	IapOrderIncomeHour *int64 `json:"iap_order_income_hour,omitempty"`
	// 安装完成量
	InstallFinishCnt *int64 `json:"install_finish_cnt,omitempty"`
	//
	IosConvertCount *int64 `json:"ios_convert_count,omitempty"`
	//
	IsOwn bool `json:"is_own"`
	// 视频是否有违规
	IsViolated *int32 `json:"is_violated,omitempty"`
	//
	ItemClickRate *int64 `json:"item_click_rate,omitempty"`
	//
	ItemClickRateDouble *float64 `json:"item_click_rate_double,omitempty"`
	//
	ItemComment *int64 `json:"item_comment,omitempty"`
	//
	ItemFeelGood *int64 `json:"item_feel_good,omitempty"`
	//
	ItemForward *int64 `json:"item_forward,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	ItemInfoDailyList []*StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerItemInfoDailyListInner `json:"item_info_daily_list,omitempty"`
	//
	ItemInteractRate *float64 `json:"item_interact_rate,omitempty"`
	// 作品链接
	ItemUrl *string `json:"item_url,omitempty"`
	//
	ItemView       *int64                                                                            `json:"item_view,omitempty"`
	LiveEffectData *StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerLiveEffectData `json:"live_effect_data,omitempty"`
	// MCN 抽成金额
	McnAmount *int64 `json:"mcn_amount,omitempty"`
	// 视频公开状态，1：公开，0：非公开
	OpenState *int64 `json:"open_state,omitempty"`
	// 审核状态
	OverallAuditStatus *int64 `json:"overall_audit_status,omitempty"`
	// 短视频播放量，或直播看播人次
	Play *int64 `json:"play,omitempty"`
	// 发布时间
	PublishTime *int64 `json:"publish_time,omitempty"`
	// 撮合中介已发放佣金(和mcn_amount类似)，单位为厘（元*1000）
	ServiceProviderAmount *int64 `json:"service_provider_amount,omitempty"`
	// 结算时间
	SettleTime *string `json:"settle_time,omitempty"`
	// MCN和达人可获得部分的预估流水，单位为厘（元*1000）
	SharePrice *int64 `json:"share_price,omitempty"`
	//
	SharePriceHour *int64 `json:"share_price_hour,omitempty"`
	// 达人预估可收益金额（非已发放金额），单位为厘（元*1000），包含MCN分佣部分
	SharePriceTotal *int64 `json:"share_price_total,omitempty"`
	//
	SharePriceTotalHour *int64 `json:"share_price_total_hour,omitempty"`
	// 视频有效播放量
	ShareVv         *int64                                                                             `json:"share_vv,omitempty"`
	SmallGameSettle *StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerSmallGameSettle `json:"small_game_settle,omitempty"`
	// 相关性审核，-1 未通过 0 审核中 1 通过
	Uncorrelated *int64 `json:"uncorrelated,omitempty"`
	// 作品时长（毫秒）
	VideoDurationMs *int64 `json:"video_duration_ms,omitempty"`
	// 投稿来源渠道（视频参与来源渠道）
	VideoParticipateProviderChannel *string `json:"video_participate_provider_channel,omitempty"`
}
