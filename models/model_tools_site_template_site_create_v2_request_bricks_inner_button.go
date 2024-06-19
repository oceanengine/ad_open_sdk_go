/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerButton 按钮组件描述
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerButton struct {
	AppointEvent *ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonAppointEvent `json:"appoint_event,omitempty"`
	// linkEvent自定义图片链接
	BgImageUrl    *string                                                             `json:"bg_image_url,omitempty"`
	DownloadEvent *ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonDownloadEvent `json:"download_event,omitempty"`
	EventType     ToolsSiteTemplateSiteCreateV2BricksButtonEventType                  `json:"event_type"`
	LinkEvent     *ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonLinkEvent     `json:"link_event,omitempty"`
	PhoneEvent    *ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonPhoneEvent    `json:"phone_event,omitempty"`
}
