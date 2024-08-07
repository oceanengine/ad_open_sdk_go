/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileAudioGetV2FilteringSources
type FileAudioGetV2FilteringSources string

// List of file_audio_get_v2_filtering_sources
const (
	AD_SITE_FileAudioGetV2FilteringSources           FileAudioGetV2FilteringSources = "AD_SITE"
	OPEN_API_FileAudioGetV2FilteringSources          FileAudioGetV2FilteringSources = "OPEN_API"
	TTS_TEXT_TO_AUDIO_FileAudioGetV2FilteringSources FileAudioGetV2FilteringSources = "TTS_TEXT_TO_AUDIO"
)

// All allowed values of FileAudioGetV2FilteringSources enum
var AllowedFileAudioGetV2FilteringSourcesEnumValues = []FileAudioGetV2FilteringSources{
	"AD_SITE",
	"OPEN_API",
	"TTS_TEXT_TO_AUDIO",
}

// NewFileAudioGetV2FilteringSourcesFromValue returns a pointer to a valid FileAudioGetV2FilteringSources
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileAudioGetV2FilteringSourcesFromValue(v string) (*FileAudioGetV2FilteringSources, error) {
	ev := FileAudioGetV2FilteringSources(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileAudioGetV2FilteringSources: valid values are %v", v, AllowedFileAudioGetV2FilteringSourcesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileAudioGetV2FilteringSources) IsValid() bool {
	for _, existing := range AllowedFileAudioGetV2FilteringSourcesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_audio_get_v2_filtering_sources value
func (v FileAudioGetV2FilteringSources) Ptr() *FileAudioGetV2FilteringSources {
	return &v
}
