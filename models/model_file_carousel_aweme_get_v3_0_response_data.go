/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileCarouselAwemeGetV30ResponseData
type FileCarouselAwemeGetV30ResponseData struct {
	//
	AwemeCarouselList []*FileCarouselAwemeGetV30ResponseDataAwemeCarouselListInner `json:"aweme_carousel_list,omitempty"`
	CursorInfo        *FileCarouselAwemeGetV30ResponseDataCursorInfo               `json:"cursor_info,omitempty"`
}
