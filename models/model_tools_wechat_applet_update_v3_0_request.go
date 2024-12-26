/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletUpdateV30Request struct for ToolsWechatAppletUpdateV30Request
type ToolsWechatAppletUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	DiscountRate *int64 `json:"discount_rate,omitempty"`
	//
	GuideText *string `json:"guide_text,omitempty"`
	//
	HasDiscount *bool `json:"has_discount,omitempty"`
	//
	HasOnlineEarning *bool `json:"has_online_earning,omitempty"`
	//
	HeaderImageUrl *string `json:"header_image_url,omitempty"`
	//
	IconImageUrl *string `json:"icon_image_url,omitempty"`
	//
	ImagesHorizontalUrl []string `json:"images_horizontal_url,omitempty"`
	//
	ImagesVerticalUrl []string `json:"images_vertical_url,omitempty"`
	//
	InstanceId int64 `json:"instance_id"`
	//
	Introduction *string `json:"introduction,omitempty"`
	//
	Labels              []string                                       `json:"labels,omitempty"`
	MaxPaymentTierRange *ToolsWechatAppletUpdateV30MaxPaymentTierRange `json:"max_payment_tier_range,omitempty"`
	MaxRechargeTier     *ToolsWechatAppletUpdateV30MaxRechargeTier     `json:"max_recharge_tier,omitempty"`
	MembershipType      *ToolsWechatAppletUpdateV30MembershipType      `json:"membership_type,omitempty"`
	MidPaymentTierRange *ToolsWechatAppletUpdateV30MidPaymentTierRange `json:"mid_payment_tier_range,omitempty"`
	MinPaymentTierRange *ToolsWechatAppletUpdateV30MinPaymentTierRange `json:"min_payment_tier_range,omitempty"`
	MinRechargeTier     *ToolsWechatAppletUpdateV30MinRechargeTier     `json:"min_recharge_tier,omitempty"`
	PaymentForm         *ToolsWechatAppletUpdateV30PaymentForm         `json:"payment_form,omitempty"`
	//
	PropName                *string                                            `json:"prop_name,omitempty"`
	RecommendedRechargeTier *ToolsWechatAppletUpdateV30RecommendedRechargeTier `json:"recommended_recharge_tier,omitempty"`
	//
	RemarkMessage *string                                 `json:"remark_message,omitempty"`
	RevenueModel  *ToolsWechatAppletUpdateV30RevenueModel `json:"revenue_model,omitempty"`
	//
	TagInfo string `json:"tag_info"`
}
