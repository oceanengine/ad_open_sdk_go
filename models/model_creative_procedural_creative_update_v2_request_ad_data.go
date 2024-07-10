/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestAdData
type CreativeProceduralCreativeUpdateV2RequestAdData struct {
	//
	AdCategory       *int64                                                    `json:"ad_category,omitempty"`
	AdDownloadStatus *CreativeProceduralCreativeUpdateV2AdDataAdDownloadStatus `json:"ad_download_status,omitempty"`
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	//
	AnchorId          *string                                                    `json:"anchor_id,omitempty"`
	AnchorRelatedType *CreativeProceduralCreativeUpdateV2AdDataAnchorRelatedType `json:"anchor_related_type,omitempty"`
	AnchorType        *CreativeProceduralCreativeUpdateV2AdDataAnchorType        `json:"anchor_type,omitempty"`
	//
	AppName                    *string                                                             `json:"app_name,omitempty"`
	CreativeAutoGenerateSwitch *CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DpaExternalUrlField *string `json:"dpa_external_url_field,omitempty"`
	//
	DynamicCreativeSwitch []*CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch `json:"dynamic_creative_switch,omitempty"`
	EnableSmartSource     *CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource       `json:"enable_smart_source,omitempty"`
	//
	ExternalUrl *string `json:"external_url,omitempty"`
	//
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	//
	IesCoreUserId    *string                                                         `json:"ies_core_user_id,omitempty"`
	IsCommentDisable *CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable       `json:"is_comment_disable,omitempty"`
	IsFeedAndFavSee  *CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee        `json:"is_feed_and_fav_see,omitempty"`
	IsPresentedVideo *CreativeProceduralCreativeUpdateV2AdDataIsPresentedVideo       `json:"is_presented_video,omitempty"`
	MiniProgramInfo  *CreativeProceduralCreativeUpdateV2RequestAdDataMiniProgramInfo `json:"mini_program_info,omitempty"`
	//
	OpenUrl    *string                                             `json:"open_url,omitempty"`
	ParamsType *CreativeProceduralCreativeUpdateV2AdDataParamsType `json:"params_type,omitempty"`
	//
	PlayableUrl   *string                                                `json:"playable_url,omitempty"`
	PriorityTrial *CreativeProceduralCreativeUpdateV2AdDataPriorityTrial `json:"priority_trial,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	SubLinkIdList []int64 `json:"sub_link_id_list,omitempty"`
	//
	Supplements []*CreativeProceduralCreativeUpdateV2RequestAdDataSupplementsInner `json:"supplements,omitempty"`
	//
	ThirdIndustryId *int64 `json:"third_industry_id,omitempty"`
	//
	WebUrl *string `json:"web_url,omitempty"`
}
