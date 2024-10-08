/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10Filter
type EnterpriseCommentListGetV10Filter struct {
	//
	AdId []int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId []int64 `json:"advertiser_id,omitempty"`
	//
	CampaignId []int64 `json:"campaign_id,omitempty"`
	//
	Content *string `json:"content,omitempty"`
	//
	CreativeId []int64                                   `json:"creative_id,omitempty"`
	Emotion    *EnterpriseCommentListGetV10FilterEmotion `json:"emotion,omitempty"`
	//
	ItemId   []int64                                    `json:"item_id,omitempty"`
	ItemType *EnterpriseCommentListGetV10FilterItemType `json:"item_type,omitempty"`
	Level    *EnterpriseCommentListGetV10FilterLevel    `json:"level,omitempty"`
	//
	MaterialId           *int64                                                 `json:"material_id,omitempty"`
	ReplyStatusByEDouyin *EnterpriseCommentListGetV10FilterReplyStatusByEDouyin `json:"reply_status_by_e_douyin,omitempty"`
	Source               *EnterpriseCommentListGetV10FilterSource               `json:"source,omitempty"`
}
