/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30RequestCreativeInfo 创意维度信息
type AdlabGroupUpdateV30RequestCreativeInfo struct {
	// 创意标签（最多20个标签,且每个标签长度不超过10个字符）
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// 应用名，当ad_type= ANDROID或者IOS时必填 ，且必须与计划维度传入的应用名称一致
	AppName *string `json:"app_name,omitempty"`
	// 是否开启自动派生创意
	CreativeAutoGenerateSwitch *int64 `json:"creative_auto_generate_switch,omitempty"`
	// 图片素材 最多使用50张图片，当通投场景下只选择图片不选择视频时至少选择5张图片，如果视频和图片都选择则要保证一个横版视频+一个竖版视频+一张大图横图。非通投场景至少传入1张图片或1个视频
	ImageMaterials []*AdlabGroupUpdateV30RequestCreativeInfoImageMaterialsInner `json:"image_materials,omitempty"`
	// 创意分类-三级行业，填写三级行业ID
	IndustryV3 *int64 `json:"industry_v3,omitempty"`
	// 是否开启评论
	IsCommentDisable *int64 `json:"is_comment_disable,omitempty"`
	// 是否跟随素材
	IsFollowMaterial *int64 `json:"is_follow_material,omitempty"`
	// 搭配试玩素材url 可使用搭配试玩场景： 推广目的（landing_type）为APP，投放范围（delivery_range）为UNION穿山甲，且union_video_type = REWARDED_VIDEO激励视频
	PlayableUrl *string `json:"playable_url,omitempty"`
	// 广告来源，2-10个字符，当推广目的为APP时字符限制为1-20，当推广目的为应用下载且download_type为EXTERNAL_URL时必填
	Source *string `json:"source,omitempty"`
	// 标题素材 最多50个，当广告位选择通投智选时最少3个，其余场景下最少1个
	TitleMaterials []*AdlabGroupUpdateV30RequestCreativeInfoTitleMaterialsInner `json:"title_materials,omitempty"`
	// 视频素材 最多使用50个视频，当通投场景下只选择视频不选择图片时至少选择3个视频，如果视频和图片都选择则要保证一个横版视频+一个竖版视频+一张大图横图。非通投场景至少传入1张图片或1个视频
	VideoMaterials []*AdlabGroupUpdateV30RequestCreativeInfoVideoMaterialsInner `json:"video_materials,omitempty"`
}
