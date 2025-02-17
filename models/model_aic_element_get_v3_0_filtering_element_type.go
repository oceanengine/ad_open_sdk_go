/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicElementGetV30FilteringElementType
type AicElementGetV30FilteringElementType string

// List of aic_element_get_v3.0_filtering_element_type
const (
	FOLDER_AicElementGetV30FilteringElementType AicElementGetV30FilteringElementType = "FOLDER"
	FONT_AicElementGetV30FilteringElementType   AicElementGetV30FilteringElementType = "FONT"
	IMAGE_AicElementGetV30FilteringElementType  AicElementGetV30FilteringElementType = "IMAGE"
	MUSIC_AicElementGetV30FilteringElementType  AicElementGetV30FilteringElementType = "MUSIC"
	VIDEO_AicElementGetV30FilteringElementType  AicElementGetV30FilteringElementType = "VIDEO"
	VOICE_AicElementGetV30FilteringElementType  AicElementGetV30FilteringElementType = "VOICE"
)

// Ptr returns reference to aic_element_get_v3.0_filtering_element_type value
func (v AicElementGetV30FilteringElementType) Ptr() *AicElementGetV30FilteringElementType {
	return &v
}
