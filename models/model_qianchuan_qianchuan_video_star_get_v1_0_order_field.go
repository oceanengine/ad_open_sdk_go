/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanQianchuanVideoStarGetV10OrderField
type QianchuanQianchuanVideoStarGetV10OrderField string

// List of qianchuan_qianchuan_video_star_get_v1.0_order_field
const (
	CREAT_TIME_QianchuanQianchuanVideoStarGetV10OrderField                      QianchuanQianchuanVideoStarGetV10OrderField = "creat_time"
	STAT_COST_QianchuanQianchuanVideoStarGetV10OrderField                       QianchuanQianchuanVideoStarGetV10OrderField = "stat_cost"
	CTR_QianchuanQianchuanVideoStarGetV10OrderField                             QianchuanQianchuanVideoStarGetV10OrderField = "ctr"
	CONVERT_RATE_QianchuanQianchuanVideoStarGetV10OrderField                    QianchuanQianchuanVideoStarGetV10OrderField = "convert_rate"
	TOTAL_PREPAY_AND_PAY_ORDER_ROI2_QianchuanQianchuanVideoStarGetV10OrderField QianchuanQianchuanVideoStarGetV10OrderField = "total_prepay_and_pay_order_roi2"
	TOTAL_PAY_ORDER_GMV_FOR_ROI2_QianchuanQianchuanVideoStarGetV10OrderField    QianchuanQianchuanVideoStarGetV10OrderField = "total_pay_order_gmv_for_roi2"
	VIEW_CNT_QianchuanQianchuanVideoStarGetV10OrderField                        QianchuanQianchuanVideoStarGetV10OrderField = "view_cnt"
	LIKE_CNT_QianchuanQianchuanVideoStarGetV10OrderField                        QianchuanQianchuanVideoStarGetV10OrderField = "like_cnt"
	COMMENT_CNT_QianchuanQianchuanVideoStarGetV10OrderField                     QianchuanQianchuanVideoStarGetV10OrderField = "comment_cnt"
	SHARE_CNT_QianchuanQianchuanVideoStarGetV10OrderField                       QianchuanQianchuanVideoStarGetV10OrderField = "share_cnt"
)

// Ptr returns reference to qianchuan_qianchuan_video_star_get_v1.0_order_field value
func (v QianchuanQianchuanVideoStarGetV10OrderField) Ptr() *QianchuanQianchuanVideoStarGetV10OrderField {
	return &v
}
