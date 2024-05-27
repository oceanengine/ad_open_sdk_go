/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAgentV2Request struct for FileVideoAgentV2Request
type FileVideoAgentV2Request struct {
	// 代理商id
	AgentId int64 `json:"agent_id"`
	// 视频上传名称
	FileName string `json:"file_name"`
	// 标注是否为AIGC生成,默认false
	IsAigc *bool `json:"is_aigc,omitempty"`
	// 标注是否允许授权(素材搬运)，目前只支持true
	IsNeedAuth bool                        `json:"is_need_auth"`
	UploadType *FileVideoAgentV2UploadType `json:"upload_type,omitempty"`
	// 文件数据
	VideoFile **FormFileInfo `json:"video_file,omitempty"`
	// 视频的md5值(用于服务端校验) upload_type为UPLOAD_BY_File必填
	VideoSignature *string `json:"video_signature,omitempty"`
	// 视频URL地址，upload_type为UPLOAD_BY_URL必传 只支持连山tos地址
	VideoUrl *string `json:"video_url,omitempty"`
}
