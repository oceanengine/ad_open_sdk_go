/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeGetV2DataListStatus
type CreativeGetV2DataListStatus string

// List of creative_get_v2_data_list_status
const (
	CREATIVE_STATUS_ADVERTISER_BUDGET_EXCEED_CreativeGetV2DataListStatus      CreativeGetV2DataListStatus = "CREATIVE_STATUS_ADVERTISER_BUDGET_EXCEED"
	CREATIVE_STATUS_ADVERTISER_PRE_OFFLINE_BUDGET_CreativeGetV2DataListStatus CreativeGetV2DataListStatus = "CREATIVE_STATUS_ADVERTISER_PRE_OFFLINE_BUDGET"
	CREATIVE_STATUS_AD_AUDIT_CreativeGetV2DataListStatus                      CreativeGetV2DataListStatus = "CREATIVE_STATUS_AD_AUDIT"
	CREATIVE_STATUS_AD_AUDIT_DENY_CreativeGetV2DataListStatus                 CreativeGetV2DataListStatus = "CREATIVE_STATUS_AD_AUDIT_DENY"
	CREATIVE_STATUS_AD_DISABLE_CreativeGetV2DataListStatus                    CreativeGetV2DataListStatus = "CREATIVE_STATUS_AD_DISABLE"
	CREATIVE_STATUS_AD_EXTERNAL_URL_DISABLE_CreativeGetV2DataListStatus       CreativeGetV2DataListStatus = "CREATIVE_STATUS_AD_EXTERNAL_URL_DISABLE"
	CREATIVE_STATUS_AD_PRE_OFFLINE_BUDGET_CreativeGetV2DataListStatus         CreativeGetV2DataListStatus = "CREATIVE_STATUS_AD_PRE_OFFLINE_BUDGET"
	CREATIVE_STATUS_AD_REAUDIT_CreativeGetV2DataListStatus                    CreativeGetV2DataListStatus = "CREATIVE_STATUS_AD_REAUDIT"
	CREATIVE_STATUS_ALL_CreativeGetV2DataListStatus                           CreativeGetV2DataListStatus = "CREATIVE_STATUS_ALL"
	CREATIVE_STATUS_AUDIT_CreativeGetV2DataListStatus                         CreativeGetV2DataListStatus = "CREATIVE_STATUS_AUDIT"
	CREATIVE_STATUS_AUDIT_DENY_CreativeGetV2DataListStatus                    CreativeGetV2DataListStatus = "CREATIVE_STATUS_AUDIT_DENY"
	CREATIVE_STATUS_AUTHORIZE_CANCEL_CreativeGetV2DataListStatus              CreativeGetV2DataListStatus = "CREATIVE_STATUS_AUTHORIZE_CANCEL"
	CREATIVE_STATUS_AWEME_ACCOUNT_DISABLED_CreativeGetV2DataListStatus        CreativeGetV2DataListStatus = "CREATIVE_STATUS_AWEME_ACCOUNT_DISABLED"
	CREATIVE_STATUS_AWEME_ANCHOR_DISABLED_CreativeGetV2DataListStatus         CreativeGetV2DataListStatus = "CREATIVE_STATUS_AWEME_ANCHOR_DISABLED"
	CREATIVE_STATUS_AWEME_ITEM_DISABLED_CreativeGetV2DataListStatus           CreativeGetV2DataListStatus = "CREATIVE_STATUS_AWEME_ITEM_DISABLED"
	CREATIVE_STATUS_BALANCE_EXCEED_CreativeGetV2DataListStatus                CreativeGetV2DataListStatus = "CREATIVE_STATUS_BALANCE_EXCEED"
	CREATIVE_STATUS_BUDGET_EXCEED_CreativeGetV2DataListStatus                 CreativeGetV2DataListStatus = "CREATIVE_STATUS_BUDGET_EXCEED"
	CREATIVE_STATUS_CAMPAIGN_DISABLE_CreativeGetV2DataListStatus              CreativeGetV2DataListStatus = "CREATIVE_STATUS_CAMPAIGN_DISABLE"
	CREATIVE_STATUS_CAMPAIGN_EXCEED_CreativeGetV2DataListStatus               CreativeGetV2DataListStatus = "CREATIVE_STATUS_CAMPAIGN_EXCEED"
	CREATIVE_STATUS_CAN_NOT_EDIT_CreativeGetV2DataListStatus                  CreativeGetV2DataListStatus = "CREATIVE_STATUS_CAN_NOT_EDIT"
	CREATIVE_STATUS_CAN_NOT_LAUNCH_CreativeGetV2DataListStatus                CreativeGetV2DataListStatus = "CREATIVE_STATUS_CAN_NOT_LAUNCH"
	CREATIVE_STATUS_CREATE_CreativeGetV2DataListStatus                        CreativeGetV2DataListStatus = "CREATIVE_STATUS_CREATE"
	CREATIVE_STATUS_DATA_ERROR_CreativeGetV2DataListStatus                    CreativeGetV2DataListStatus = "CREATIVE_STATUS_DATA_ERROR"
	CREATIVE_STATUS_DELETE_CreativeGetV2DataListStatus                        CreativeGetV2DataListStatus = "CREATIVE_STATUS_DELETE"
	CREATIVE_STATUS_DELIVERY_OK_CreativeGetV2DataListStatus                   CreativeGetV2DataListStatus = "CREATIVE_STATUS_DELIVERY_OK"
	CREATIVE_STATUS_DISABLE_CreativeGetV2DataListStatus                       CreativeGetV2DataListStatus = "CREATIVE_STATUS_DISABLE"
	CREATIVE_STATUS_DONE_CreativeGetV2DataListStatus                          CreativeGetV2DataListStatus = "CREATIVE_STATUS_DONE"
	CREATIVE_STATUS_FROZEN_CreativeGetV2DataListStatus                        CreativeGetV2DataListStatus = "CREATIVE_STATUS_FROZEN"
	CREATIVE_STATUS_LIVE_ROOM_OFF_CreativeGetV2DataListStatus                 CreativeGetV2DataListStatus = "CREATIVE_STATUS_LIVE_ROOM_OFF"
	CREATIVE_STATUS_NOT_START_CreativeGetV2DataListStatus                     CreativeGetV2DataListStatus = "CREATIVE_STATUS_NOT_START"
	CREATIVE_STATUS_NO_DELIVERY_CreativeGetV2DataListStatus                   CreativeGetV2DataListStatus = "CREATIVE_STATUS_NO_DELIVERY"
	CREATIVE_STATUS_NO_SCHEDULE_CreativeGetV2DataListStatus                   CreativeGetV2DataListStatus = "CREATIVE_STATUS_NO_SCHEDULE"
	CREATIVE_STATUS_PRODUCT_OFFLINE_CreativeGetV2DataListStatus               CreativeGetV2DataListStatus = "CREATIVE_STATUS_PRODUCT_OFFLINE"
	CREATIVE_STATUS_REAUDIT_CreativeGetV2DataListStatus                       CreativeGetV2DataListStatus = "CREATIVE_STATUS_REAUDIT"
)

// Ptr returns reference to creative_get_v2_data_list_status value
func (v CreativeGetV2DataListStatus) Ptr() *CreativeGetV2DataListStatus {
	return &v
}
