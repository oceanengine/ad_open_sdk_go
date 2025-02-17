/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAudioGetV2FilteringSources
type FileAudioGetV2FilteringSources string

// List of file_audio_get_v2_filtering_sources
const (
	AD_SITE_FileAudioGetV2FilteringSources           FileAudioGetV2FilteringSources = "AD_SITE"
	OPEN_API_FileAudioGetV2FilteringSources          FileAudioGetV2FilteringSources = "OPEN_API"
	TTS_TEXT_TO_AUDIO_FileAudioGetV2FilteringSources FileAudioGetV2FilteringSources = "TTS_TEXT_TO_AUDIO"
)

// Ptr returns reference to file_audio_get_v2_filtering_sources value
func (v FileAudioGetV2FilteringSources) Ptr() *FileAudioGetV2FilteringSources {
	return &v
}
