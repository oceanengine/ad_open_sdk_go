/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCreateV2RequestBricksInnerBackground
type ToolsSiteCreateV2RequestBricksInnerBackground struct {
	//
	Color       *string                                                   `json:"color,omitempty"`
	Description *ToolsSiteCreateV2RequestBricksInnerBackgroundDescription `json:"description,omitempty"`
	//
	Image *string                                             `json:"image,omitempty"`
	Title *ToolsSiteCreateV2RequestBricksInnerBackgroundTitle `json:"title,omitempty"`
	//
	Type *string `json:"type,omitempty"`
}
