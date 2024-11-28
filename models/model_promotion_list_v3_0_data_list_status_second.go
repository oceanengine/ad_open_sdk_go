/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30DataListStatusSecond
type PromotionListV30DataListStatusSecond string

// List of promotion_list_v3.0_data_list_status_second
const (
	APP_NOT_ALLOWED_PromotionListV30DataListStatusSecond             PromotionListV30DataListStatusSecond = "APP_NOT_ALLOWED"
	AUDIT_PromotionListV30DataListStatusSecond                       PromotionListV30DataListStatusSecond = "AUDIT"
	AUDIT_DENY_PromotionListV30DataListStatusSecond                  PromotionListV30DataListStatusSecond = "AUDIT_DENY"
	AWEME_ACCOUNT_DISABLED_PromotionListV30DataListStatusSecond      PromotionListV30DataListStatusSecond = "AWEME_ACCOUNT_DISABLED"
	AWEME_ACCOUNT_OPTIMIZABLE_PromotionListV30DataListStatusSecond   PromotionListV30DataListStatusSecond = "AWEME_ACCOUNT_OPTIMIZABLE"
	AWEME_ANCHOR_DISABLED_PromotionListV30DataListStatusSecond       PromotionListV30DataListStatusSecond = "AWEME_ANCHOR_DISABLED"
	AWEME_VIDEO_OPTIMIZABLE_PromotionListV30DataListStatusSecond     PromotionListV30DataListStatusSecond = "AWEME_VIDEO_OPTIMIZABLE"
	BALANCE_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond      PromotionListV30DataListStatusSecond = "BALANCE_OFFLINE_BUDGET"
	BUDGET_GROUP_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond PromotionListV30DataListStatusSecond = "BUDGET_GROUP_OFFLINE_BUDGET"
	DISABLED_PromotionListV30DataListStatusSecond                    PromotionListV30DataListStatusSecond = "DISABLED"
	DISABLE_BY_QUOTA_PromotionListV30DataListStatusSecond            PromotionListV30DataListStatusSecond = "DISABLE_BY_QUOTA"
	LIVE_ROOM_OFF_PromotionListV30DataListStatusSecond               PromotionListV30DataListStatusSecond = "LIVE_ROOM_OFF"
	NO_SCHEDULE_PromotionListV30DataListStatusSecond                 PromotionListV30DataListStatusSecond = "NO_SCHEDULE"
	OFFLINE_BALANCE_PromotionListV30DataListStatusSecond             PromotionListV30DataListStatusSecond = "OFFLINE_BALANCE"
	PRODUCT_OFFLINE_PromotionListV30DataListStatusSecond             PromotionListV30DataListStatusSecond = "PRODUCT_OFFLINE"
	PROJECT_DISABLED_PromotionListV30DataListStatusSecond            PromotionListV30DataListStatusSecond = "PROJECT_DISABLED"
	PROJECT_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond      PromotionListV30DataListStatusSecond = "PROJECT_OFFLINE_BUDGET"
	PROMOTION_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond    PromotionListV30DataListStatusSecond = "PROMOTION_OFFLINE_BUDGET"
	REAUDIT_PromotionListV30DataListStatusSecond                     PromotionListV30DataListStatusSecond = "REAUDIT"
	STAR_TASK_NOT_STARTED_PromotionListV30DataListStatusSecond       PromotionListV30DataListStatusSecond = "STAR_TASK_NOT_STARTED"
	TIME_NO_REACH_PromotionListV30DataListStatusSecond               PromotionListV30DataListStatusSecond = "TIME_NO_REACH"
)

// Ptr returns reference to promotion_list_v3.0_data_list_status_second value
func (v PromotionListV30DataListStatusSecond) Ptr() *PromotionListV30DataListStatusSecond {
	return &v
}
