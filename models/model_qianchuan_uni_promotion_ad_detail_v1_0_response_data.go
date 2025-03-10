/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdDetailV10ResponseData
type QianchuanUniPromotionAdDetailV10ResponseData struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	// 抖音号id
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	CreateTime      *string                                                      `json:"create_time,omitempty"`
	CreativeSetting *QianchuanUniPromotionAdDetailV10ResponseDataCreativeSetting `json:"creative_setting,omitempty"`
	DeliverySetting *QianchuanUniPromotionAdDetailV10ResponseDataDeliverySetting `json:"delivery_setting,omitempty"`
	MarketingGoal   *QianchuanUniPromotionAdDetailV10DataMarketingGoal           `json:"marketing_goal,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	MultiProductCreativeList []*QianchuanUniPromotionAdDetailV10ResponseDataMultiProductCreativeListInner `json:"multi_product_creative_list,omitempty"`
	//
	Name      *string                                        `json:"name,omitempty"`
	OptStatus *QianchuanUniPromotionAdDetailV10DataOptStatus `json:"opt_status,omitempty"`
	//
	ProductInfos                  []*QianchuanUniPromotionAdDetailV10ResponseDataProductInfosInner           `json:"product_infos,omitempty"`
	ProgrammaticCreativeMediaList *QianchuanUniPromotionAdDetailV10ResponseDataProgrammaticCreativeMediaList `json:"programmatic_creative_media_list,omitempty"`
	//
	RoomInfo []*QianchuanUniPromotionAdDetailV10ResponseDataRoomInfoInner `json:"room_info,omitempty"`
	Status   *QianchuanUniPromotionAdDetailV10DataStatus                  `json:"status,omitempty"`
}
