/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsVideoMaterialListInner struct for PromotionCreateV30RequestPromotionMaterialsVideoMaterialListInner
type PromotionCreateV30RequestPromotionMaterialsVideoMaterialListInner struct {
	// 引导视频id，白名单功能，仅游戏广告主投放奖励关卡广告时必填
	GuideVideoId *string                                                        `json:"guide_video_id,omitempty"`
	ImageMode    PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode `json:"image_mode"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	VideoCoverId      *string                                                                 `json:"video_cover_id,omitempty"`
	VideoHpVisibility *PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility `json:"video_hp_visibility,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	// 自定义商品库视频模板ID，创建DPA创意时可传入，传入后该素材下image_info与video_info不生效，长度限制1，从【获取 DPA 商品库视频模板】接口中获取
	VideoTaskIds      []string                                                                `json:"video_task_ids,omitempty"`
	VideoTemplateType *PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType `json:"video_template_type,omitempty"`
}
