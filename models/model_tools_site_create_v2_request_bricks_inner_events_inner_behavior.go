/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCreateV2RequestBricksInnerEventsInnerBehavior
type ToolsSiteCreateV2RequestBricksInnerEventsInnerBehavior struct {
	AndroidLink *ToolsSiteCreateV2RequestBricksInnerEventsInnerBehaviorAndroidLink `json:"android_link,omitempty"`
	IosLink     *ToolsSiteCreateV2RequestBricksInnerEventsInnerBehaviorIosLink     `json:"ios_link,omitempty"`
	Link        *ToolsSiteCreateV2RequestBricksInnerEventsInnerBehaviorLink        `json:"link,omitempty"`
	//
	Name       *string                                                           `json:"name,omitempty"`
	SmartPhone *ToolsSiteCreateV2RequestBricksInnerEventsInnerBehaviorSmartPhone `json:"smart_phone,omitempty"`
}
