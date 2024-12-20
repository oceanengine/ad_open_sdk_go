/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPromotionListV30FilteringPromotionStatusSecond
type LocalPromotionListV30FilteringPromotionStatusSecond string

// List of local_promotion_list_v3.0_filtering_promotion_status_second
const (
	AUDIT_LocalPromotionListV30FilteringPromotionStatusSecond                    LocalPromotionListV30FilteringPromotionStatusSecond = "AUDIT"
	AUDIT_DENY_LocalPromotionListV30FilteringPromotionStatusSecond               LocalPromotionListV30FilteringPromotionStatusSecond = "AUDIT_DENY"
	AWEME_ACCOUNT_DISABLED_LocalPromotionListV30FilteringPromotionStatusSecond   LocalPromotionListV30FilteringPromotionStatusSecond = "AWEME_ACCOUNT_DISABLED"
	BALANCE_OFFLINE_BUDGET_LocalPromotionListV30FilteringPromotionStatusSecond   LocalPromotionListV30FilteringPromotionStatusSecond = "BALANCE_OFFLINE_BUDGET"
	DISABLED_LocalPromotionListV30FilteringPromotionStatusSecond                 LocalPromotionListV30FilteringPromotionStatusSecond = "DISABLED"
	DISABLE_BY_QUOTA_LocalPromotionListV30FilteringPromotionStatusSecond         LocalPromotionListV30FilteringPromotionStatusSecond = "DISABLE_BY_QUOTA"
	LIVE_ROOM_OFF_LocalPromotionListV30FilteringPromotionStatusSecond            LocalPromotionListV30FilteringPromotionStatusSecond = "LIVE_ROOM_OFF"
	NO_SCHEDULE_LocalPromotionListV30FilteringPromotionStatusSecond              LocalPromotionListV30FilteringPromotionStatusSecond = "NO_SCHEDULE"
	OFFLINE_BALANCE_LocalPromotionListV30FilteringPromotionStatusSecond          LocalPromotionListV30FilteringPromotionStatusSecond = "OFFLINE_BALANCE"
	PRODUCT_OR_POI_OFFLINE_LocalPromotionListV30FilteringPromotionStatusSecond   LocalPromotionListV30FilteringPromotionStatusSecond = "PRODUCT_OR_POI_OFFLINE"
	PROJECT_DISABLED_LocalPromotionListV30FilteringPromotionStatusSecond         LocalPromotionListV30FilteringPromotionStatusSecond = "PROJECT_DISABLED"
	PROJECT_OFFLINE_BUDGET_LocalPromotionListV30FilteringPromotionStatusSecond   LocalPromotionListV30FilteringPromotionStatusSecond = "PROJECT_OFFLINE_BUDGET"
	PROMOTION_OFFLINE_BUDGET_LocalPromotionListV30FilteringPromotionStatusSecond LocalPromotionListV30FilteringPromotionStatusSecond = "PROMOTION_OFFLINE_BUDGET"
	REAUDIT_LocalPromotionListV30FilteringPromotionStatusSecond                  LocalPromotionListV30FilteringPromotionStatusSecond = "REAUDIT"
	ROI2_DISABLE_LocalPromotionListV30FilteringPromotionStatusSecond             LocalPromotionListV30FilteringPromotionStatusSecond = "ROI2_DISABLE"
	TIME_NO_REACH_LocalPromotionListV30FilteringPromotionStatusSecond            LocalPromotionListV30FilteringPromotionStatusSecond = "TIME_NO_REACH"
)

// Ptr returns reference to local_promotion_list_v3.0_filtering_promotion_status_second value
func (v LocalPromotionListV30FilteringPromotionStatusSecond) Ptr() *LocalPromotionListV30FilteringPromotionStatusSecond {
	return &v
}
