/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaListVideoMaterialInner struct for QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaListVideoMaterialInner
type QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaListVideoMaterialInner struct {
	//
	AwemeItemId *int64                                                                             `json:"aweme_item_id,omitempty"`
	ImageMode   *QianchuanUniAwemeAdUpdateV10ProgrammaticCreativeMediaListVideoMaterialImageMode   `json:"image_mode,omitempty"`
	StarTraffic *QianchuanUniAwemeAdUpdateV10ProgrammaticCreativeMediaListVideoMaterialStarTraffic `json:"star_traffic,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
