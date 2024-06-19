/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAdV2Request struct for FileVideoAdV2Request
type FileVideoAdV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 素材的文件名，可自定义素材名，不传择默认取文件名，最长255个字符 注：若同一素材已进行上传，重新上传不会改名
	Filename *string `json:"filename,omitempty"`
	// 视频素材是否为AIGC生成
	IsAigc *bool `json:"is_aigc,omitempty"`
	// 上传视频场景是否是引导视频
	IsGuideVideo *bool `json:"is_guide_video,omitempty"`
	//
	Labels     []string                 `json:"labels,omitempty"`
	UploadType *FileVideoAdV2UploadType `json:"upload_type,omitempty"`
	// 视频文件 允许格式：mp4、mpeg、3gp、avi（10s超时限制） upload_type为UPLOAD_BY_File必填
	VideoFile **FormFileInfo `json:"video_file,omitempty"`
	// 视频的md5值(用于服务端校验) upload_type为UPLOAD_BY_File必填
	VideoSignature *string `json:"video_signature,omitempty"`
	// 视频url地址，如http://xxx.xxx，upload_type为UPLOAD_BY_URL必填
	VideoUrl *string `json:"video_url,omitempty"`
}
