/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileMediaGetV2FilteringMediaName
type FileMediaGetV2FilteringMediaName string

// List of file_media_get_v2_filtering_media_name
const (
	AWEME_FileMediaGetV2FilteringMediaName FileMediaGetV2FilteringMediaName = "AWEME"
)

// All allowed values of FileMediaGetV2FilteringMediaName enum
var AllowedFileMediaGetV2FilteringMediaNameEnumValues = []FileMediaGetV2FilteringMediaName{
	"AWEME",
}

// NewFileMediaGetV2FilteringMediaNameFromValue returns a pointer to a valid FileMediaGetV2FilteringMediaName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMediaGetV2FilteringMediaNameFromValue(v string) (*FileMediaGetV2FilteringMediaName, error) {
	ev := FileMediaGetV2FilteringMediaName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMediaGetV2FilteringMediaName: valid values are %v", v, AllowedFileMediaGetV2FilteringMediaNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMediaGetV2FilteringMediaName) IsValid() bool {
	for _, existing := range AllowedFileMediaGetV2FilteringMediaNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_media_get_v2_filtering_media_name value
func (v FileMediaGetV2FilteringMediaName) Ptr() *FileMediaGetV2FilteringMediaName {
	return &v
}
