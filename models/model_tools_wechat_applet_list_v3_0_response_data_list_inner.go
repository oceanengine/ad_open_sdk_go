/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletListV30ResponseDataListInner struct for ToolsWechatAppletListV30ResponseDataListInner
type ToolsWechatAppletListV30ResponseDataListInner struct {
	//
	AdvertiserId *int64                                       `json:"advertiser_id,omitempty"`
	AuditStatus  *ToolsWechatAppletListV30DataListAuditStatus `json:"audit_status,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	GuideText *string `json:"guide_text,omitempty"`
	//
	HeaderImageUrl *string `json:"header_image_url,omitempty"`
	//
	IconImageUrl *string `json:"icon_image_url,omitempty"`
	//
	ImagesHorizontalUrl []string `json:"images_horizontal_url,omitempty"`
	//
	ImagesVerticalUrl []string `json:"images_vertical_url,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	Introduction *string `json:"introduction,omitempty"`
	//
	Labels []string `json:"labels,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Path *string `json:"path,omitempty"`
	//
	Reason *string `json:"reason,omitempty"`
	//
	RemarkMessage *string `json:"remark_message,omitempty"`
	//
	UserName *string `json:"user_name,omitempty"`
}
