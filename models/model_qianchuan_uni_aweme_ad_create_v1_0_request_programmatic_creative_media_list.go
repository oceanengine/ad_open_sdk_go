/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdCreateV10RequestProgrammaticCreativeMediaList
type QianchuanUniAwemeAdCreateV10RequestProgrammaticCreativeMediaList struct {
	//
	BlockVideoMaterial []*QianchuanUniAwemeAdCreateV10RequestProgrammaticCreativeMediaListBlockVideoMaterialInner `json:"block_video_material,omitempty"`
	//
	TitleMaterial []*QianchuanUniAwemeAdCreateV10RequestProgrammaticCreativeMediaListTitleMaterialInner `json:"title_material,omitempty"`
	//
	VideoMaterial []*QianchuanUniAwemeAdCreateV10RequestProgrammaticCreativeMediaListVideoMaterialInner `json:"video_material,omitempty"`
}
