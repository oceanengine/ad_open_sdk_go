/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMaterialPushV30ResponseDataSuccessResultsInner struct for AicMaterialPushV30ResponseDataSuccessResultsInner
type AicMaterialPushV30ResponseDataSuccessResultsInner struct {
	// 推送成功的广告主
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
	// 素材id
	MaterialId *string `json:"material_id,omitempty"`
	// 推送成功的视频id
	VideoId *string `json:"video_id,omitempty"`
}