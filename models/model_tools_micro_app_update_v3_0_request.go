/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppUpdateV30Request struct for ToolsMicroAppUpdateV30Request
type ToolsMicroAppUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AppPage []*ToolsMicroAppUpdateV30RequestAppPageInner `json:"app_page"`
	//
	DiscountRate *int64 `json:"discount_rate,omitempty"`
	//
	HasDiscount *bool `json:"has_discount,omitempty"`
	//
	HasOnlineEarning *bool `json:"has_online_earning,omitempty"`
	//
	InstanceId          int64                                      `json:"instance_id"`
	MaxPaymentTierRange *ToolsMicroAppUpdateV30MaxPaymentTierRange `json:"max_payment_tier_range,omitempty"`
	MaxRechargeTier     *ToolsMicroAppUpdateV30MaxRechargeTier     `json:"max_recharge_tier,omitempty"`
	MembershipType      *ToolsMicroAppUpdateV30MembershipType      `json:"membership_type,omitempty"`
	MidPaymentTierRange *ToolsMicroAppUpdateV30MidPaymentTierRange `json:"mid_payment_tier_range,omitempty"`
	MinPaymentTierRange *ToolsMicroAppUpdateV30MinPaymentTierRange `json:"min_payment_tier_range,omitempty"`
	MinRechargeTier     *ToolsMicroAppUpdateV30MinRechargeTier     `json:"min_recharge_tier,omitempty"`
	PaymentForm         *ToolsMicroAppUpdateV30PaymentForm         `json:"payment_form,omitempty"`
	//
	PropName                *string                                        `json:"prop_name,omitempty"`
	RecommendedRechargeTier *ToolsMicroAppUpdateV30RecommendedRechargeTier `json:"recommended_recharge_tier,omitempty"`
	//
	Remark       *string                             `json:"remark,omitempty"`
	RevenueModel *ToolsMicroAppUpdateV30RevenueModel `json:"revenue_model,omitempty"`
	//
	TagInfo string `json:"tag_info"`
}
