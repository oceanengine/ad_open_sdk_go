/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanVideoGetV10DataListImageMode
type QianchuanVideoGetV10DataListImageMode string

// List of qianchuan_video_get_v1.0_data_list_image_mode
const (
	LARGE_QianchuanVideoGetV10DataListImageMode          QianchuanVideoGetV10DataListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanVideoGetV10DataListImageMode QianchuanVideoGetV10DataListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanVideoGetV10DataListImageMode          QianchuanVideoGetV10DataListImageMode = "SMALL"
	UNION_SPLASH_QianchuanVideoGetV10DataListImageMode   QianchuanVideoGetV10DataListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanVideoGetV10DataListImageMode    QianchuanVideoGetV10DataListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanVideoGetV10DataListImageMode QianchuanVideoGetV10DataListImageMode = "VIDEO_VERTICAL"
)

// Ptr returns reference to qianchuan_video_get_v1.0_data_list_image_mode value
func (v QianchuanVideoGetV10DataListImageMode) Ptr() *QianchuanVideoGetV10DataListImageMode {
	return &v
}
