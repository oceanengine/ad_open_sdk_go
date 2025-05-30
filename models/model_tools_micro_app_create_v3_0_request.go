/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppCreateV30Request struct for ToolsMicroAppCreateV30Request
type ToolsMicroAppCreateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AppId   string                               `json:"app_id"`
	AppPage ToolsMicroAppCreateV30RequestAppPage `json:"app_page"`
	//
	DiscountRate *int64 `json:"discount_rate,omitempty"`
	//
	HasDiscount *bool `json:"has_discount,omitempty"`
	//
	HasOnlineEarning    *bool                                      `json:"has_online_earning,omitempty"`
	MaxPaymentTierRange *ToolsMicroAppCreateV30MaxPaymentTierRange `json:"max_payment_tier_range,omitempty"`
	MaxRechargeTier     *ToolsMicroAppCreateV30MaxRechargeTier     `json:"max_recharge_tier,omitempty"`
	MembershipType      *ToolsMicroAppCreateV30MembershipType      `json:"membership_type,omitempty"`
	MidPaymentTierRange *ToolsMicroAppCreateV30MidPaymentTierRange `json:"mid_payment_tier_range,omitempty"`
	MinPaymentTierRange *ToolsMicroAppCreateV30MinPaymentTierRange `json:"min_payment_tier_range,omitempty"`
	MinRechargeTier     *ToolsMicroAppCreateV30MinRechargeTier     `json:"min_recharge_tier,omitempty"`
	PaymentForm         *ToolsMicroAppCreateV30PaymentForm         `json:"payment_form,omitempty"`
	//
	PropName                *string                                        `json:"prop_name,omitempty"`
	RecommendedRechargeTier *ToolsMicroAppCreateV30RecommendedRechargeTier `json:"recommended_recharge_tier,omitempty"`
	//
	Remark       string                              `json:"remark"`
	RevenueModel *ToolsMicroAppCreateV30RevenueModel `json:"revenue_model,omitempty"`
	//
	TagInfo string `json:"tag_info"`
}
