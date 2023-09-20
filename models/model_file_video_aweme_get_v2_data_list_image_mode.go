/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoAwemeGetV2DataListImageMode
type FileVideoAwemeGetV2DataListImageMode string

// List of file_video_aweme_get_v2_data_list_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_FileVideoAwemeGetV2DataListImageMode          FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_FileVideoAwemeGetV2DataListImageMode      FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_FileVideoAwemeGetV2DataListImageMode                 FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_FileVideoAwemeGetV2DataListImageMode               FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_FileVideoAwemeGetV2DataListImageMode               FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_FileVideoAwemeGetV2DataListImageMode      FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_FileVideoAwemeGetV2DataListImageMode FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_FileVideoAwemeGetV2DataListImageMode   FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_FileVideoAwemeGetV2DataListImageMode               FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_FileVideoAwemeGetV2DataListImageMode        FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_FileVideoAwemeGetV2DataListImageMode  FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_FileVideoAwemeGetV2DataListImageMode               FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_FileVideoAwemeGetV2DataListImageMode      FileVideoAwemeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_FileVideoAwemeGetV2DataListImageMode                   FileVideoAwemeGetV2DataListImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_FileVideoAwemeGetV2DataListImageMode                 FileVideoAwemeGetV2DataListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of FileVideoAwemeGetV2DataListImageMode enum
var AllowedFileVideoAwemeGetV2DataListImageModeEnumValues = []FileVideoAwemeGetV2DataListImageMode{
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewFileVideoAwemeGetV2DataListImageModeFromValue returns a pointer to a valid FileVideoAwemeGetV2DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoAwemeGetV2DataListImageModeFromValue(v string) (*FileVideoAwemeGetV2DataListImageMode, error) {
	ev := FileVideoAwemeGetV2DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoAwemeGetV2DataListImageMode: valid values are %v", v, AllowedFileVideoAwemeGetV2DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoAwemeGetV2DataListImageMode) IsValid() bool {
	for _, existing := range AllowedFileVideoAwemeGetV2DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_aweme_get_v2_data_list_image_mode value
func (v FileVideoAwemeGetV2DataListImageMode) Ptr() *FileVideoAwemeGetV2DataListImageMode {
	return &v
}
