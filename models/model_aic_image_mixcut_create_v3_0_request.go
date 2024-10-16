/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicImageMixcutCreateV30Request struct for AicImageMixcutCreateV30Request
type AicImageMixcutCreateV30Request struct {
	// BP账户体系:组织id，Agent账户体系:代理商id
	AccountId   int64                                      `json:"account_id"`
	AccountType AicImageMixcutCreateV30AccountType         `json:"account_type"`
	AudioOption *AicImageMixcutCreateV30RequestAudioOption `json:"audio_option,omitempty"`
	// 生成成片的数量，最大值5
	Count      *int64                                    `json:"count,omitempty"`
	GenElement *AicImageMixcutCreateV30RequestGenElement `json:"gen_element,omitempty"`
	// 图片元素id，用于视频生成，最多20个
	ImageElementIds []int64                                      `json:"image_element_ids"`
	ProductInfo     *AicImageMixcutCreateV30RequestProductInfo   `json:"product_info,omitempty"`
	RenderOption    *AicImageMixcutCreateV30RequestRenderOption  `json:"render_option,omitempty"`
	SubtitleStyle   *AicImageMixcutCreateV30RequestSubtitleStyle `json:"subtitle_style,omitempty"`
	// 可自定义任务名称，生成的成片均以此名称命名，最多50字，多个视频命名规则为任务名称_1-n。不传入，默认使用product_name+时间戳命名 示例:洗衣凝珠_20240809_152301
	TaskName *string `json:"task_name,omitempty"`
}
