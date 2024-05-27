/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAudioGetV2Filtering
type FileAudioGetV2Filtering struct {
	// 按音频id过滤，list长度 1-100 注意audio_ids，material_ids，audio_signatures三者只能选一个过滤，同时传入会报错
	AudioIds []string `json:"audio_ids,omitempty"`
	// 根据音频上传时间过滤，与start_time搭配使用，格式：yyyy-mm-dd hh:mm:ss
	EndTime *string `json:"end_time,omitempty"`
	// 按音频素材id（全局唯一）过滤，list长度1-100 注意audio_ids，material_ids，audio_signatures三者只能选一个过滤，同时传入会报错
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 按音频md5过滤，list长度1-100 注意audio_ids，material_ids，audio_signatures三者只能选一个过滤，同时传入会报错
	Signatures []string `json:"signatures,omitempty"`
	// 音频素材来源 枚举值大小写敏感，请严格按照定义的名称传参 AD_SITE  AD后台本地上传 OPEN_API 开放平台上传 TTS_TEXT_TO_AUDIO文本转音频
	Sources []*FileAudioGetV2FilteringSources `json:"sources,omitempty"`
	// 根据音频上传时间过滤，与end_time搭配使用，格式：yyyy-mm-dd hh:mm:ss
	StartTime *string `json:"start_time,omitempty"`
}
