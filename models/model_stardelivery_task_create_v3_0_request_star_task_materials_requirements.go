/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskCreateV30RequestStarTaskMaterialsRequirements 制作要求
type StardeliveryTaskCreateV30RequestStarTaskMaterialsRequirements struct {
	// 补充说明，1-3000个字符
	AdditionalInformation *string `json:"additional_information,omitempty"`
	// 口播要求，1-80个字符，例如口播介绍指定文字、品牌、slogan、商品卖点，是否添加口播字幕等
	OnCameraRequirement string `json:"on_camera_requirement"`
	// 其他要求，不同要求用英文逗号隔开，最多6条，每条1-80个字符
	OtherRequirements []string `json:"other_requirements,omitempty"`
	// 参考素材uri，请提供需要达人在作品中配合添加的内容素材，您可上传产品图、素材图、要求示例或指引，图片最多9张，格式为 JPG、PNG 等
	SampleMaterialIds []string `json:"sample_material_ids"`
	// 示例视频，最多3个，需要输入抖音视频链接并确保该视频为在线状态 （如 https://www.douyin.com/video/xxx）
	SampleVideoUrls []string `json:"sample_video_urls,omitempty"`
	// 标题提及用户，要@ 的抖音号，抖音用户个人主页抖音昵称下方可找到
	TitleMentionsAwemeId *string `json:"title_mentions_aweme_id,omitempty"`
	// 标题要求，1-80个字符，例如标题包含商品名称、商品性能/特点关键词等；若需提及账号或话题，请在其他要求参数中填写
	TitleRequirement string `json:"title_requirement"`
	// 标题指定话题，多个话题用英文逗号隔开，最多3条 您需调用「获取星广任务标题指定话题」接口查询可使用的话题id
	TitleSpecifiesTopicIds []int64 `json:"title_specifies_topic_ids,omitempty"`
	// 镜头要求，1-80个字符，例如产品实物出镜、近景特写、达人出镜、镜头时长、指定画面等 例如：需产品实物出镜，并给到带logo的产品近景特写，特写时长不低于3s，达人需出镜亲自安利
	VoRequirement string `json:"vo_requirement"`
}
