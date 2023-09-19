/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInner struct for ToolsSiteTemplateGetV2ResponseDataListInnerBricksInner
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInner struct {
	Button *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButton `json:"button,omitempty"`
	Coupon *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerCoupon `json:"coupon,omitempty"`
	Form   *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerForm   `json:"form,omitempty"`
	// 组件在模板中的位置描述
	Index        string                                                              `json:"index"`
	Picture      *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPicture      `json:"picture,omitempty"`
	PictureGroup *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPictureGroup `json:"picture_group,omitempty"`
	Text         *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerText         `json:"text,omitempty"`
	Type         ToolsSiteTemplateGetV2DataListBricksType                            `json:"type"`
	Video        *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerVideo        `json:"video,omitempty"`
	WechatApplet *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerWechatApplet `json:"wechat_applet,omitempty"`
	WechatGame   *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerWechatGame   `json:"wechat_game,omitempty"`
}
