/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAudioAdV2Request struct for FileAudioAdV2Request
type FileAudioAdV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AudioFile **FormFileInfo `json:"audio_file,omitempty"`
	// 图片的md5值(用于服务端校验) upload_type为UPLOAD_BY_FILE
	AudioSignature *string `json:"audio_signature,omitempty"`
	//
	AudioUrl   *string                 `json:"audio_url,omitempty"`
	UploadType FileAudioAdV2UploadType `json:"upload_type"`
}
