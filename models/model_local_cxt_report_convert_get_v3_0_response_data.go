/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalCxtReportConvertGetV30ResponseData
type LocalCxtReportConvertGetV30ResponseData struct {
	//
	AdPoiVerificationUvCnt *int64 `json:"ad_poi_verification_uv_cnt,omitempty"`
	//
	AdPredictPoiRecommendUvCntHigh *int64 `json:"ad_predict_poi_recommend_uv_cnt_high,omitempty"`
	//
	AdPredictPoiRecommendUvCntLow *int64 `json:"ad_predict_poi_recommend_uv_cnt_low,omitempty"`
	//
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	//
	PoiRecommendCount *int64 `json:"poi_recommend_count,omitempty"`
	//
	ShowCnt *int64 `json:"show_cnt,omitempty"`
	//
	VideoOtoPayOrderCount *int64 `json:"video_oto_pay_order_count,omitempty"`
}
