/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterial
type QianchuanUniPromotionAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterial struct {
	CoverImage *QianchuanUniPromotionAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterialCoverImage `json:"cover_image,omitempty"`
	//
	MaterialId  *int64                                                                                        `json:"material_id,omitempty"`
	Source      *QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource      `json:"source,omitempty"`
	StarTraffic *QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialStarTraffic `json:"star_traffic,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	VideoDuration *int64 `json:"video_duration,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
